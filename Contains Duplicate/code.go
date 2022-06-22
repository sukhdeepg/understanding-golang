package main

import "fmt"

func containsDuplicate(nums []int) bool {
    mp := make(map[int]bool)

    for _, num := range nums {
        if _, ok := mp[num]; ok {
            return true
        }
        mp[num] = true
    }

    return false
}

func main() {
    arr := []int{1,2,3,4} // example
    ans := containsDuplicate(arr)
    fmt.Printf("%v", ans)
}