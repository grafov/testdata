package main

import (
	"log"
	"math/rand"
	"time"
)

type testCase struct {
	input    []int
	expected []int
	n        int
}

func main() {
	testCases := []testCase{
		{
			input:    []int{1, 2, -3},
			expected: []int{1, 4, 9},
			n:        3,
		},
		{
			input:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81},
			n:        2,
		},
		{
			input:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81},
			n:        20,
		},
		{
			input:    []int{0, 0, 0, 0, 0, 5, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0, 25, 0, 0, 0, 0},
			n:        4,
		},
	}

	for _, tc := range testCases {
		actual := do(tc.n, tc.input)
		test(tc.expected, actual)
	}
}

// Must return a slice of elements in input squads.
func do(n int, input []int) []int {
	result := make([]int, len(input))

	ch := make(chan int, 10)

	for i := 0; i < n; i++ {
		go func() {
			val := <-ch
			result = append(result, pow(val))
		}()
	}

	for _, val := range input {
		ch <- val
	}

	return result
}

func pow(i int) int {
	r := rand.Int63n(100)
	time.Sleep(time.Duration(r) * time.Millisecond)

	return i * i
}

func test(expected, actual []int) {
	log.Println("========== Test ==========")
	log.Println("Expected", expected)
	log.Println("Actual", actual)

	for i := range expected {
		if expected[i] != actual[i] {
			log.Printf("Elem %d: failed!\n", i)
		} else {
			log.Printf("Elem %d: ok!\n", i)
		}
	}
}
