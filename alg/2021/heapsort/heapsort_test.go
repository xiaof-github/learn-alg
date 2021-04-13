package test

import (
	"testing"
)

var a = []int{45,145,245,32,5,2,69,239,12,40}

func Benchmark_buildMaxHeap(b *testing.B) {
	for i:=0; i < b.N;i++{
		buildMaxHeap(bA)
	}
}

func Benchmark_heapsort(b *testing.B) {
	for i:=0; i < b.N; i++ {
		heapsort(bA)
	}
}