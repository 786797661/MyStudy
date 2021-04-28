/**
* @Author: Gosin
* @Date: 2021/4/22 21:28
 */

package main

import (
	"container/list"
	"fmt"
)

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Print(dailyTemperatures(temperatures))
}

func dailyTemperatures(T []int) []int {
	stack := list.New()
	stack.PushBack(0)
	day := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for stack.Len() > 0 && T[i] > T[stack.Back().Value.(int)] {
			day[stack.Back().Value.(int)] = i - stack.Back().Value.(int)
			stack.Remove(stack.Back())
		}
		stack.PushBack(i)
	}
	return day
}
