# golang simple full text search engine

> original article: https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/

## Corpus
We are going to search a part of the abstract of English Wikipedia.
The latest dump is available at dumps.wikimedia.org. As of today, the file size after decompression is 913 MB.
The XML file contains over 600K documents.

Document example:

```
<title>Wikipedia: Kit-Cat Klock</title>
<url>https://en.wikipedia.org/wiki/Kit-Cat_Klock</url>
<abstract>The Kit-Cat Klock is an art deco novelty wall clock shaped like a grinning cat with cartoon eyes that swivel in time with its pendulum tail.</abstract>
```

### Loading documents

`pkg/wikipedia/dump/dump.go`

```go
type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocument(filename string) ([]Document, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i, _ := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
```

## Analyzing documents


`pkg/search/analyze.go`

```go
func Analyze(text string) []string {
    tokens := tokenize(text)
    tokens = lowercaseFilter(tokens)
    tokens = stopwordFilter(tokens)
    tokens = stemmerFilter(tokens)
    return tokens
}
```

### Tokenize

The tokenizer is the first step of text analysis. Its job is to convert text into a list of tokens.
Our implementation splits the text on a word boundary and removes punctuation marks:

`pkg/search/tokenize.go`

```go
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		// Split on any character that is not a letter or a number.
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
```

### lowercase

In order to make the search case-insensitive, the lowercase filter converts tokens to lower case.
cAt, Cat and caT are normalized to cat. Later, when we query the index,
we'll lower case the search terms as well. This will make the search term cAt match the text Cat.

`pkg/search/filters.go`

```go
func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}
```

### stopword

Almost any English text contains commonly used words like a, I, the or be. Such words are called stop words. We are going to remove them since almost any document would match the stop words.


`pkg/search/filters.go`

```go
var stopwords = map[string]struct{}{
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func stopwordFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}
```

### stemming

Because of the grammar rules, documents may include different forms of the same word.
Stemming reduces words into their base form.
For example, fishing, fished and fisher may be reduced to the base form (stem) fish.

Implementing a stemmer is a non-trivial task, We'll take one of the existing modules:


`pkg/search/filters.go`

```go
import snowballeng "github.com/kljensen/snowball/english"

func stemmerFilter(tokens []string) []string {
    r := make([]string, len(tokens))
    for i, token := range tokens {
        r[i] = snowballeng.Stem(token, false)
    }
    return r
}
```

## Indexing documents

Building the index consists of analyzing the documents and adding their IDs to the map:


`pkg/search/index.go`

```go
type Index map[string][]int

func (idx index) add(docs []document) {
    for _, doc := range docs {
        for _, token := range analyze(doc.Text) {
            ids := idx[token]
            if ids != nil && ids[len(ids)-1] == doc.ID {
                // Don't add same ID twice.
                continue
            }
            idx[token] = append(ids, doc.ID)
        }
    }
}

func example() {
    idx := make(Index)
    idx.add([]document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
    idx.add([]document{{ID: 2, Text: "donut is a donut"}})
    fmt.Println(idx)
}
```

## Querying the index

To query the index, we are going to apply the same tokenizer and filters we used for indexing:

`pkg/search/index.go`

```go
func (idx Index) Search(text string) [][]int {
    var r [][]int
    for _, token := range analyze(text) {
        if ids, ok := idx[token]; ok {
            r = append(r, ids)
        }
    }
    return r
}

func example() {
	idx.search("Small wild cat")
	// [[24, 173, 303, ...], [98, 173, 765, ...], [[24, 51, 173, ...]]
}
```

And finally, we can find all documents that mention cats. Searching 600K documents took less than a millisecond!

With the inverted index, the time complexity of the search query is linear to the number of search tokens.
In the example query above, other than analyzing the input text, search had to perform only three map lookups.

## Boolen Querying

The query from the previous section returned a disjoined list of documents for each token. What we normally expect
to find when we type small wild cat in a search box is a list of results that contain small, wild and cat at the same time.
The next step is to compute the set intersection between the lists. This way we'll get a list of documents matching all tokens.

Luckily, IDs in our inverted index are inserted in ascending order. Since the IDs are sorted,
it's possible to compute the intersection between two lists in linear time.
The intersection function iterates two lists simultaneously and collect IDs that exist in both:

```go
func intersection(a []int, b []int) []int {
    maxLen := len(a)
    if len(b) > maxLen {
        maxLen = len(b)
    }
    r := make([]int, 0, maxLen)
    var i, j int
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            i++
        } else if a[i] > b[j] {
            j++
        } else {
            r = append(r, a[i])
            i++
            j++
        }
    }
    return r
}
```

Updated search analyzes the given query text, lookups tokens and computes the set intersection between lists of IDs:

```go
func (idx index) Search(text string) []int {
    var r []int
    for _, token := range analyze(text) {
        if ids, ok := idx[token]; ok {
            if r == nil {
                r = ids
            } else {
                r = intersection(r, ids)
            }
        } else {
            // Token doesn't exist.
            return nil
        }
    }
    return r
}
```
