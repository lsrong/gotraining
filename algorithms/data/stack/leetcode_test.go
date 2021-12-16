package stack

import "testing"

func TestIsValidSymbol(t *testing.T) {
	// 1. s = "()"， true ;
	// 2. s = "()[]{}"， true ;
	// 3. s = "(]"， false ;
	// 4. s = "([)]"， false ;
	// 5. s = "{[]}"， true ;
	demos := map[string]bool{
		"()":         true,
		"()[]{}":     true,
		"(]":         false,
		"([)]":       false,
		"{[]}":       true,
		"((({[]}":    false,
		"((({[]})))": true,
		")]}":        false,
	}
	for s, res := range demos {
		check := IsValidSymbol(s)
		if check != res {
			t.Fatalf("IsValidSymbol() check [%s] Got %t, Expected %t", s, check, res)
		}
	}
}
