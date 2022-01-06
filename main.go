package main

import (
	"fmt"
)

func main() {

	number := []int{2, 2, 7, 15}
	target := 9
	result := twoSum(number, target)
	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {

	//get the target

	//loop through the nums and keep reducing the target from the
	//[1][2] [1][3] [1][4]
	//[2][3] [2][4]
	//	[3][4]

	//	which pair match the target return

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums); j++ {
			if i == j {
				j = j + 1
			}

			sum := nums[i] + nums[j]
			reduce := target - sum
			if reduce == 0 {
				result := []int{i + 1, j + 1}
				return result
			}
		}

	}

	return []int{}

}
