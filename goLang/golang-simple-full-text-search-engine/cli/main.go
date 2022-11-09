package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"fts/pkg/search"
	"fts/pkg/wikipedia/dump"
)

const (
	// The path to the Wikipedia dump file.
	dumpFile = "test/testdata/enwiki-latest-abstract1.xml"

	// The maximum number of documents to return for a search.
	maxDoc = 10
)

func main() {
	// Open the Wikipedia dump file.
	fmt.Printf("[%v] loading dump file %s\n", time.Now(), dumpFile)
	docs, err := dump.LoadDocument(dumpFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%v] load %d documents\n", time.Now(), len(docs))

	// Create a new index.
	fmt.Printf("[%v] creating index\n", time.Now())
	idx := search.New()
	for _, doc := range docs {
		idx.Add(doc)
	}
	fmt.Printf("[%v] index created\n", time.Now())

	// Search the word given by the user.
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Input a word to search:")
		word, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Search the top 10 documents for the given word.
		fmt.Printf("[%v] searching for %s\n", time.Now(), word)
		results := idx.Search(word, maxDoc)
		fmt.Printf("[%v] found %d results\n", time.Now(), len(results))

		// Print the results.
		for _, result := range results {
			fmt.Printf("[%v] %d. %s\n", time.Now(), result, docs[result].Text)
		}
	}
}
