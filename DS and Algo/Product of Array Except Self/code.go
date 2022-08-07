package main

import "fmt"

func productExceptSelf(nums []int) []int {
    n, p := len(nums), 1
    var res []int

    for idx, _ := range nums {
        res = append(res, p)
        p *= nums[idx]
    }

    p = 1

    for idx := n - 1; idx >= 0; idx-- {
        res[idx] *= p
        p *= nums[idx]
    }

    return res
}

func main() {
    arr := []int{1,2,3,4} // example
    ans := productExceptSelf(arr)
    fmt.Printf("%v", ans)
}