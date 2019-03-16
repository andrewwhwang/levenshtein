package levenshtein

import (
	"testing"
)

func BenchmarkDistance(b *testing.B) {
	// text := "TAGCTATCACGACCGCGGTCGATTTGCCCGATAGCTATCACGACCGCGGTCGATTTGCCC"
	// word := "AGGCTATCACCTGACCTCCAGGCCGATGCCCAGGCTATCACCTGACCTCCAGGCCGATGC"
	text := "TAGCTATCACGACCGCGGTCGATTTGCCCG"
	word := "AGGCTATCACCTGACCTCCAGGCCGATGCC"
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
        Distance(text, word)
    }
}

func BenchmarkMyerDistReg(b *testing.B) {
	// text := "TAGCTATCACGACCGCGGTCGATTTGCCCGATAGCTATCACGACCGCGGTCGATTTGCCC"
	// word := "AGGCTATCACCTGACCTCCAGGCCGATGCCCAGGCTATCACCTGACCTCCAGGCCGATGC"
	text := "TAGCTATCACGACCGCGGTCGATTTGCCCG"
	word := "AGGCTATCACCTGACCTCCAGGCCGATGCC"
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
        MyerDistReg(text, word)
    }
}
