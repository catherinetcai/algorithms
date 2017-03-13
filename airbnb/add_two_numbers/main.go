package main

import (
	"container/list"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// This is mostly just to test
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
	spew.Dump(res)
}

/*
Algorithm/pseudocode:
 The inputs are reversed, so if input is (2 -> 4 -> 3) and (5 -> 6 -> 4), then add the ones place (2 + 5). If this value is greater than 10, you'd have to carry that value over to the next, which you have to do with 4+6 = 10. Since it's 10, you'd leave the 0 and carry the 1 over to the next sum (4 + 3) + 1 = 8. So the answer is 807, which you'd reverse and return 7 -> 0 -> 8.
*/
func addTwoNumbers(l1, l2 *list.List) *list.List {
	var sum int

	sumList := list.New()

	// Gets the heads of the two linked lists
	current1 := l1.Front()
	current2 := l2.Front()

	for current1 != nil || current2 != nil {
		// We want to divide the sum by 10 because, for example, if the value carried over from the previous iteration is 10, then we'll want to add 1 to the sum of this next node
		sum /= 10

		if current1 != nil {
			sum += current1.Value.(int)
			current1 = current1.Next()
		}
		if current2 != nil {
			sum += current2.Value.(int)
			current2 = current2.Next()
		}
		// We want the modulo
		sumList.PushBack(sum % 10)
	}
	return sumList
}
