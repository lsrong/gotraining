package search

import "testing"

var testFound []Result

func BenchmarkCMMatch(b *testing.B) {
	var result []Result
	var err error

	cm := NewCM()

	for i := 0; i < b.N; i++ {
		result, err = cm.match("https://c.m.163.com/nc/article/headline/T1348647853363/0-100.html", "人民")
		if err != nil {
			b.FailNow()
		}
	}
	testFound = result
}
