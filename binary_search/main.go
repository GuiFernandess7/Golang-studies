package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 13, 16}

    if !checkIfSorted(arr) {
        quickSort(arr, 0, len(arr)-1)
    }

    fmt.Println(arr)
}

func checkIfSorted(nums []int) bool {
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] > nums[i+1] {
            return false
        }
    }
    return true
}

func quickSort(nums []int, low, high int) {
    if low < high {
        pi := partition(nums, low, high)

        quickSort(nums, low, pi-1)
        quickSort(nums, pi+1, high)
    }
}

func partition(nums []int, low, high int) int {
    pivot := nums[high]
    i := low - 1

    for j := low; j < high; j++ {
        if nums[j] <= pivot {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    nums[i+1], nums[high] = nums[high], nums[i+1]
    return i + 1
}
