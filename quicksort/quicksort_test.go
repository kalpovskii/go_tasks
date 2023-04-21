package main

import "testing"

func TestPartition(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{4, 5, 2, 3, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 4, 6, 8, 1}, []int{1, 2, 4, 6, 8}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{-1, 12, 1337, 0, 2, 23, 21}, []int{-1, 0, 2, 12, 21, 23, 1337}},
	}

	for _, testCase := range testCases {
		arr := quickSortStart(testCase.input)

		// Проверим отсортирован ли массив
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				t.Errorf("Массив не отсортирован. Expected %v but got %v", testCase.expected, arr)
			}
		}
	}
}
