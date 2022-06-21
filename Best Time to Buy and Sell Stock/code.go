package main

import (
    "fmt"
    "math"
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxProfit(prices []int) int {
    minval, maxprofit := math.MaxInt32, 0

    for _, price := range prices {
        if price < minval {
            minval = price
        } else {
            maxprofit = max(maxprofit, price - minval)
        }
    }

    return maxprofit
}

func main() {
    arr := []int{7,1,5,3,6,4} // example
    ans := maxProfit(arr)
    fmt.Printf("%v", ans)
}