package levenshtein

import (
	"testing"
)

func BenchmarkDistance(b *testing.B) {
	text := "TAGCTATCACGACCGCGGTCGATTTGCCCGATAGCTATCACGACCGCGGTCGATTTGCCC"
	word := "AGGCTATCACCTGACCTCCAGGCCGATGCCCAGGCTATCACCTGACCTCCAGGCCGATGC"
	text = text[:5]
	word = word[:5]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Distance(text, word)
	}
}

func BenchmarkMyerDistReg(b *testing.B) {
	text := "TAGCTATCACGACCGCGGTCGATTTGCCCGATAGCTATCACGACCGCGGTCGATTTGCCC"
	word := "AGGCTATCACCTGACCTCCAGGCCGATGCCCAGGCTATCACCTGACCTCCAGGCCGATGC"
	text = text[:5]
	word = word[:5]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MyerDistReg(text, word)
	}
}
