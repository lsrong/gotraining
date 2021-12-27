package main

import (
	"bytes"
	"testing"
)

var output bytes.Buffer
var in []byte
var find = []byte("elvis")
var repl = []byte("elvis")

func init() {
	in, _ = assembleStream(data)
}

func BenchmarkAlgoOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output.Reset()
		algoOne(in, find, repl, &output)
	}
}

func BenchmarkAlgoTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output.Reset()
		algoTwo(in, find, repl, &output)
	}
}
