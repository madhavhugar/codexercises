package main

import "testing"

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		target int
		count int
	}{
		{
			name: "sorted array length == 3",
			nums: []int{1,2,3},
			target: 4,
			count: 7,
		},
		{
			name: "array length == 1",
			nums: []int{9},
			target: 3,
			count: 0,
		},
		{
			name: "array length == 1",
			nums: []int{2,1,3},
			target: 35,
			count: 1132436852,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func (t *testing.T) {
			got := combinationSum(test.nums, test.target)

			if got != test.count {
				t.Errorf("got != want; %d != %d", got, test.count)
			}
		})
	}
}