package main

import (
	"container/list"
	"testing"
)

// This is going to implement the test case we were looking for
func TestAddTwoNumbers(t *testing.T) {
	l1, l2 := list.New(), list.New()
	elems1 := []int{2, 4, 3}
	elems2 := []int{5, 6, 4}
	for i := range elems1 {
		l1.PushBack(elems1[i])
	}
	for i := range elems2 {
		l2.PushBack(elems2[i])
	}
	res := addTwoNumbers(l1, l2)
	expected := []int{7, 0, 8}
	current := res.Front()
	index := 0
	for current != nil {
		if current.Value.(int) != expected[index] {
			t.Errorf("Expected: %v, got: %v\n", expected[index], current.Value.(int))
		}
		current = current.Next()
		index++
	}
}
