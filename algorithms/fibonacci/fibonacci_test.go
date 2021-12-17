package fibonacci

import "testing"

func testFib(t *testing.T, name string, f func(n int) []int) {
	fibNums := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	number := len(fibNums)
	fibFromRecs := f(number)

	if number != len(fibFromRecs) {
		t.Fatalf("NewFibViaRecursion len Got %d, Expoected %d", len(fibFromRecs), number)
	}
	for i, n := range fibNums {
		if n != fibFromRecs[i] {
			t.Fatalf("fibFromRecs index [%d] Got %d, Expoected %d", i, fibFromRecs[i], n)
		}
	}

	t.Logf("succeed to build fibonacci via [%s]: %v", name, fibFromRecs)
}

func TestNewFibViaRecursion(t *testing.T) {
	testFib(t, "NewFibViaRecursion", NewFibViaRecursion)
}

func TestNewNewFibViaDp(t *testing.T) {
	testFib(t, "NewFibViaDp", NewFibViaDp)
}
