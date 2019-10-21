package inset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

// has
func (set *IntSet) Has(x int) bool {
	word, bit := x/64, x%64
	return word < len(set.words) && set.words[word]&(1<<bit) != 0
}

// add
func (set *IntSet) Add(x int) {
	word, bit := x/64, x%64
	for word >= len(set.words) {
		// allocate a new uint64 element
		set.words = append(set.words, 0)
	}
	// mapping x to words[word] and set bit
	set.words[word] |= 1 << bit
}

// add all
func (set *IntSet) AddALL(elements ...int) {
	for _, v := range elements {
		set.Add(v)
	}
}

// union
func (set *IntSet) UnionWith(anotherSet *IntSet) {
	for i, word := range anotherSet.words {
		if i < len(set.words) {
			set.words[i] |= word
		} else {
			set.words = append(set.words, word)
		}
	}
}

// intersect
func (set *IntSet) IntersectWith(anotherSet *IntSet) {
	for i, word := range anotherSet.words {
		if i < len(set.words) {
			set.words[i] &= word
		}
	}
}

// difference
func (set *IntSet) DifferenceWith(anotherSet *IntSet) {
	for i, word := range anotherSet.words {
		if i < len(set.words) {
			set.words[i] &= word
		}
	}
}

// string
func (set *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range set.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", i*64+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
