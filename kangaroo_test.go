package main

import "testing"

var kangarooTests = []struct {
	n        int
	test     KangarooTask
	expected bool
}{
	{1, KangarooTask{x1: 0, x2: 0, v1: 0, v2: 0}, true},
	{2, KangarooTask{x1: -100000000, x2: 2, v1: 1, v2: 2}, false},
	{3, KangarooTask{x1: 1, x2: 2, v1: 1, v2: 0}, true},
	{4, KangarooTask{x1: 0, x2: 2, v1: -1, v2: 1}, false},
	{5, KangarooTask{x1: -1, x2: 1, v1: 1, v2: -1}, true},
	{6, KangarooTask{x1: 1, x2: 3, v1: 5, v2: 1}, false},
	{7, KangarooTask{x1: 1, x2: 3, v1: 4, v2: 2}, true},
	{8, KangarooTask{x1: 0, x2: 5, v1: 2, v2: 3}, false},
	{9, KangarooTask{x1: 0, x2: 4, v1: 3, v2: 2}, true},
}

func TestKangaroo(t *testing.T) {
	for _, kt := range kangarooTests {
		actual := kt.test.ExistDecision()
		if actual != kt.expected {
			t.Errorf("Fib(%d): unexpected result", kt.n)
		}
	}
}
