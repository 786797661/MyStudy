/**
* @Author: Gosin
* @Date: 2021/4/23 21:52
 */

package main

import (
	"container/list"
	"fmt"
)

func main() {
	heights := []int{4, 2, 0, 3, 2, 5}
	fmt.Print(largestRectangleArea(heights))
}

func largestRectangleArea1(heights []int) int {
	max := 0
	for i := 0; i < len(heights); i++ {
		left := 0
		right := 0
		for j := i; j < len(heights); j++ {
			if heights[i] > heights[j] {
				left = heights[i] * (j - i)
				break
			} else {
				if j == len(heights)-1 {
					left = heights[i] * (j - i + 1)
					break
				}
			}

		}
		for j := i - 1; j >= 0; j-- {
			if heights[i] > heights[j] {
				right = heights[i] * (i - j - 1)
				break
			} else {
				if j == 0 {
					right = heights[i] * (i - j)
					break
				}
			}

		}
		if max < right+left {
			max = right + left
		}
	}
	return max
}

//[]int{5, 4, 1, 2}

func largestRectangleArea(heights []int) int {

	max := 0
	stack := list.New()
	for i := 0; i < len(heights); i++ {
		temp := -1
		for stack.Len() > 0 && heights[i] < heights[stack.Back().Value.(int)] {
			if temp == -1 && stack.Len() != 1 {
				temp = stack.Back().Value.(int)
				if max < heights[temp] {
					max = heights[temp]
				}
				stack.Remove(stack.Back())
			} else {
				if stack.Len() == 1 {
					if max < heights[stack.Back().Value.(int)]*(stack.Back().Value.(int)+1) {
						max = heights[stack.Back().Value.(int)] * (stack.Back().Value.(int) + 1)
						stack.Remove(stack.Back())
					}
				} else {
					index := stack.Back().Value.(int)
					stack.Remove(stack.Back())
					if max < heights[index]*(temp-stack.Back().Value.(int)) {
						max = heights[index] * (temp - stack.Back().Value.(int))
					}
				}

				//temp = stack.Back().Value.(int)

			}

		}
		stack.PushBack(i)

	}

	temp := -1
	for stack.Len() > 0 {
		if temp == -1 && stack.Len() != 1 {
			temp = stack.Back().Value.(int)
			if max < heights[temp] {
				max = heights[temp]
			}
		} else {
			if stack.Len() == 1 {
				if max < heights[stack.Back().Value.(int)]*(len(heights)) {
					max = heights[stack.Back().Value.(int)] * (len(heights))
				}
			} else {
				index := stack.Back().Value.(int)
				stack.Remove(stack.Back())
				if max < heights[index]*(temp-stack.Back().Value.(int)) {
					max = heights[index] * (temp - stack.Back().Value.(int))
				}
				//temp = stack.Back().Value.(int)

			}

		}
		stack.Remove(stack.Back())
	}
	return max
}
