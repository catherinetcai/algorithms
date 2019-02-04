package main

import "container/list"

/*
Algorithm/pseudocode:
 The inputs are reversed, so if input is (2 -> 4 -> 3) and (5 -> 6 -> 4), then add the ones place (2 + 5). If this value is greater than 10, you'd have to carry that value over to the next, which you have to do with 4+6 = 10. Since it's 10, you'd leave the 0 and carry the 1 over to the next sum (4 + 3) + 1 = 8. So the answer is 807, which you'd reverse and return 7 -> 0 -> 8.
*/
func addTwoNumbers(l1, l2 *list.List) *list.List {
	var carry, sum int

	sumList := list.New()

	// Gets the heads of the two linked lists
	current1 := l1.Front()
	current2 := l2.Front()

	for current1 != nil || current2 != nil {
		sum = carry
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

		// We're going to carry over the tens
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
	}
	return sumList
}
