package search

import (
	"fts/pkg/wikipedia/dump"
)

// Index is the implementation of a search index.
type Index map[string][]int

// New creates a new search index.
func New() Index {
	return make(Index)
}

// Add adds the given documents to the index
func (idx Index) Add(docs ...dump.Document) {
	for _, doc := range docs {
		for _, token := range Analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add same ID twice.
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

// Search returns the IDs of the documents that contain the given tokens.
func (idx Index) Search(text string, maxDoc int) []int {
	var r []int
	for _, token := range Analyze(text) {
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

	// cut off the result list if it's too long
	if maxDoc > 0 && len(r) > maxDoc {
		r = r[:maxDoc]
	}
	return r
}

// intersection returns the intersection of two lists.
// 求交集
func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		switch {
		case a[i] < b[j]:
			i++
		case a[i] > b[j]:
			j++
		default:
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}
