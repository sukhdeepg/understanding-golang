package main

import "fmt"

func findMin(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }

    left, right := 0, n - 1

    if nums[right] > nums[0] {
        return nums[0]
    }

    for left <= right {
        mid := (left + (right - left)) / 2

        if nums[mid] > nums[mid + 1] {
            return nums[mid + 1]
        }

        if nums[mid] < nums[mid - 1] {
            return nums[mid]
        }

        if nums[mid] > nums[0] {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return -1
}

func main() {
    arr := []int{3,4,5,1,2} // example
    ans := findMin(arr)
    fmt.Printf("%v", ans)
}