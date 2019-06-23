package main

import "testing"

func TestMySum(t *testing.T){

	type test struct {
		data []int
		answer int
	}

	tests := []test{
		test {
			[]int{21, 21},
			42,
		},
		test {
			[]int{3, 4, 5},
			12,
		},
		test {
			[]int{1, 2},
			3,
		},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("expected ",v.answer, " got", x)
		}
	}

	x := mySum(2, 3)

	if x != 5 {
		t.Error("Expected ", 5, " got ", x)
	}
}

