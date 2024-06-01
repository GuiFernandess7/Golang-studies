package main

import "fmt"

func main(){
	var nums []int = []int{8, 3, 5, 2, 1, 6, 9, 4, 7}

	for i := 0; i < len(nums) - 1; i++ {
		for j := 0; j < len(nums) - i - 1; j++ {
			if nums[j] > nums[j+1]{
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}

	fmt.Println(nums)
}