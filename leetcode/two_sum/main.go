package main

import "fmt"

func main() {
	fmt.Println(bruteTwoSums([]int{2, 7, 11, 15}, 9))
	fmt.Println(hashTwoSums([]int{2, 7, 11, 15}, 9))
}

func bruteTwoSums(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; i < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func hashTwoSums(nums []int, target int) []int {
	subMap := map[int]int{}
	for i := range nums {
		if _, ok := subMap[nums[i]]; ok {
			return []int{subMap[nums[i]], i}
		}
		subMap[target-nums[i]] = i
	}
	return nil
}
