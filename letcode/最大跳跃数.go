/**
* @Author: Gosin
* @Date: 2021/4/21 23:16
 */

package main

import "fmt"

func main() {

	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {

	var index = len(nums) - 1
	count := 0
	for index > 0 {
		for i := 0; i < len(nums); i++ {
			if i+nums[i] >= index {
				index = i
				count++
				break

			}
		}
	}

	return count
}
