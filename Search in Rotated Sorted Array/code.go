package main

import "fmt"

func search(nums []int, target int) int {
    n := len(nums)
    l, r := 0, n - 1

    for l <= r {
        m := (l + r) / 2

        if target == nums[m] {
            return m
        }

        // left search space
        if nums[l] <= nums[m] {
            if target > nums[m] || target < nums[l] {
                l = m + 1
            } else {
                r = m - 1
            }
        } else { // right search space
            if target < nums[m] || target > nums[r] {
                r = m - 1
            } else {
                l = m + 1
            }
        }
    }

    return -1
}

func main() {
    arr := []int{4,5,6,7,0,1,2} // example
    ans := search(arr, 0)
    fmt.Printf("%v", ans)
}