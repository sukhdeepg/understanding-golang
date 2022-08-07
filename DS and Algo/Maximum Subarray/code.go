package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxSubArray(nums []int) int {
    n, m := len(nums), nums[0]
    dp := make([]int, n)
    dp[0] = nums[0]

    for idx := 1; idx < n; idx++ {
        if dp[idx - 1] > 0 {
            dp[idx] = nums[idx] + dp[idx - 1]
        } else {
            dp[idx] = nums[idx]
        }

        m = max(m, dp[idx])
    }

    return m
}

func main() {
    arr := []int{5,4,-1,7,8} // example
    ans := maxSubArray(arr)
    fmt.Printf("%v", ans)
}