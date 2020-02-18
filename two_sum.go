package main

/*

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

Runtime: 28 ms, faster than 34.41% of Go online submissions for Two Sum.
Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Two Sum.
*/
import "fmt"

func main() {
	nums := []int{1, 2, 7, 11, 15}
	target := 9
	fmt.Printf("%v", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		gap := target - num
		for j, n := range nums {
			if n == gap && j != i {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
