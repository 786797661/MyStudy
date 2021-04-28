/**
* @Author: Gosin
* @Date: 2021/4/28 21:51
 */

package main

import "container/list"

func longestValidParentheses(s string) int {

	stack := list.New()
	lowerIndex := -1
	max := 0
	for index, v := range s {
		if v == '(' {
			stack.PushBack(index)
		} else {
			if stack.Len() > 0 {
				//lower := stack.Back().Value.(int)
				stack.Remove(stack.Back())
				if stack.Len() > 0 {
					if index-stack.Back().Value.(int) > max {
						max = index - stack.Back().Value.(int)
					}
				} else {
					if index-lowerIndex > max {
						max = index - lowerIndex
					}
				}

			} else {
				lowerIndex = index
			}

		}
	}
	return max
}
