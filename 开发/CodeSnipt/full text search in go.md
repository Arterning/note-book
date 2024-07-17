

define the core data struture

```go
// index is an inverted index. It maps tokens to document IDs.
// key is text token, value is doc list
type index map[string][]int
```

then we can define the search mehod
if we have "i love you" text to be search
the work flow is :

- analye text into string[], ["i", "love", "you"]
- find the doc ids of token "i" C1
- find the doc ids of token "love" C2
- find the doc ids of token "you" C3
- get the intersection of [C1, C2, C3]


```go
// search queries the index for the given text.
func (idx index) search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				//intersection 取交集
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


analize return []string, for example:
input is :` i love you`
output is `[i,love,you]`

```go
// analyze analyzes the text and returns a slice of tokens.
func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}

```