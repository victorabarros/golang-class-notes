package main

import (
	"testing"
)

func benchmarkPopCount(b *testing.B, size uint64) {
	for ii := 0; ii < b.N; ii++ {
		PopCount(size)
	}
}

func benchmarkPopCountV2(b *testing.B, size uint64) {
	for ii := 0; ii < b.N; ii++ {
		popCountV2(size)
	}
}

func BenchmarkPopCount10(b *testing.B)   { benchmarkPopCount(b, 10) }
func BenchmarkPopCount1000(b *testing.B) { benchmarkPopCount(b, 1000) }

func BenchmarkPopCountV210(b *testing.B)   { benchmarkPopCountV2(b, 10) }
func BenchmarkPopCountV21000(b *testing.B) { benchmarkPopCountV2(b, 1000) }
