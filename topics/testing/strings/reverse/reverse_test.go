package reverse

import (
	"testing"
)

// 基础测试命令以及结果：
// $ go test -bench .
//BenchmarkBytes-8                42441384                28.35 ns/op
//BenchmarkRunes-8                 6411446               190.5 ns/op
//BenchmarkCombiningChars-8        2329846               514.5 ns/op

// $ go test -bench . -benchmem
//BenchmarkBytes-8                42792712                27.90 ns/op           32 B/op          1 allocs/op
//BenchmarkRunes-8                 6498922               185.2 ns/op           144 B/op          2 allocs/op
//BenchmarkCombiningChars-8        2352848               508.7 ns/op           144 B/op          2 allocs/op

// $ go test -bench . -benchmem -benchtime 5s
//BenchmarkBytes-8                207782274               28.09 ns/op           32 B/op          1 allocs/op
//BenchmarkRunes-8                32183559               186.6 ns/op           144 B/op          2 allocs/op
//BenchmarkCombiningChars-8       11679547               513.4 ns/op           144 B/op          2 allocs/op

type reverseTest struct {
	in  string
	out string
}

func TestBytes(t *testing.T) {
	tests := []reverseTest{
		{"asdf", "fdsa"},
		{"Go语言", "\x80\xa8譯\xe8oG"},
		{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
	}
	for _, test := range tests {
		r := Bytes(test.in)
		if r != test.out {
			t.Errorf("reverse Bytes(%q) = %q, Expected %q", test.in, r, test.out)
		}
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes("asdsdfghjkqwwwertyuizxcvbnm")
	}
}

func TestRunes(t *testing.T) {
	tests := []reverseTest{
		{"asdf", "fdsa"},
		{"Go语言", "言语oG"},
		{"as⃝df̅", "̅fd⃝sa"},
		{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
	}
	for _, test := range tests {
		r := Runes(test.in)
		if r != test.out {
			t.Errorf("reverse Runes(%q) = %q, Expected %q", test.in, r, test.out)
		}
	}
}

func BenchmarkRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Runes("asdsdfghjkqwwwertyuizxcvbnm")
	}
}

func TestCombiningChars(t *testing.T) {
	tests := []reverseTest{
		{"asdf", "fdsa"},
		{"Go语言", "言语oG"},
		{"as⃝df̅", "f̅ds⃝a"},
		{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
	}
	for _, test := range tests {
		r := CombiningChars(test.in)
		if r != test.out {
			t.Errorf("Reverse CombiningChars(%q) = %q, Expected %q", test.in, r, test.out)
		}
	}
}

func BenchmarkCombiningChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CombiningChars("asdsdfghjkqwwwertyuizxcvbnm")
	}
}
