package levenshtein

import (
	"testing"
	"fmt"
)

func TestMyerDist(t *testing.T) {
	// text := "ANNUAL"
	// word := "ANNEALING"
	text := "TAGCTATCACGACCGCGGTCGATTTGCCCGA"
	word := "AGGCTATCACCTGACCTCCAGGCCGATGCCC"
	MyerDist(text, word)
}

func TestMyerDistDiag(t *testing.T) {
	text := "TAGCTATCACGACCGCGGTCGATTT"
	word := "AGGCTATCACCTGACCTCCAGGCCG"
	// text := "ANNUAL"
	// word := "ANNEAL"
	for i := range MyerDistDiag(word, text, 11) {
		fmt.Println(i)
	}
}
