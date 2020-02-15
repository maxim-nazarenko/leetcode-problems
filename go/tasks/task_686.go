package tasks

import (
	"strings"
)

// Given two strings A and B, find the minimum number of times A has to be repeated such that B is a substring of it. If no such solution, return -1.
// For example, with A = "abcd" and B = "cdabcdab".

// Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it; and B is not a substring of A repeated two times ("abcdabcd").

// Note:
// The length of A and B will be between 1 and 10000.

func repeatedStringMatch(A string, B string) int {
	if A == B || A == "" || B == "" {
		return 1
	}

	repeatCount := len(B)/len(A) + 1

	if len(B) < len(A) {
		repeatCount++
	}

	if len(B) > len(A) && len(B)%len(A) > 0 {
		repeatCount++
	}

	newA := strings.Repeat(A, repeatCount)
	index := strings.Index(newA, B)
	if index < 0 {
		return -1
	}

	if len(newA)-(index+len(B)) >= len(A) {
		repeatCount--
	}

	return repeatCount
}
