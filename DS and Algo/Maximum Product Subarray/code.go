package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxInArr(values []int) int {
    val := values[0]
    for _, v := range values {
        if (v > val) {
            val = v
        }
    }

    return val
}

func maxProduct(nums []int) int {
    m := maxInArr(nums)
    currMax, currMin := 1, 1

    for _, num := range nums {
        if num == 0 {
            currMax, currMin = 1, 1
            continue
        }

        tmp := currMax * num
        currMax = max(tmp, max(currMin * num, num))
        currMin = min(tmp, min(currMin * num, num))

        m = max(m, currMax)
    }

    return m
}

func main() {
    arr := []int{2,3,-2,4} // example
    ans := maxProduct(arr)
    fmt.Printf("%v", ans)
}