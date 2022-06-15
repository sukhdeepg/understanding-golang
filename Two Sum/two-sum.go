package main

import "fmt"

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

	for i, n := range nums {
		var req = target - n

		if val, ok := m[req]; ok {
			return []int{val, i}
		}

		m[n] = i
	}

	return []int{0, 0}
}

func main() {
	arr := []int{2,7,11,15} // example
	ans := twoSum(arr, 9)
	fmt.Printf("%v", ans)
}