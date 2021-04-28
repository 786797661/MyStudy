/**
* @Author: Gosin
* @Date: 2021/4/23 0:35
 */

package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Print(isValid(")"))
}

func isValid(s string) bool {
	stack := list.New()
	bytes := []byte(s)
	for _, v := range bytes {
		if v == '(' || v == '[' || v == '{' {
			stack.PushBack(v)
		} else {
			if stack.Len() > 0 {
				if v == ')' {
					if stack.Back().Value.(byte) == '(' {
						stack.Remove(stack.Back())
						continue
					} else {
						return false
					}
				} else if v == ']' {
					if stack.Back().Value.(byte) == '[' {
						stack.Remove(stack.Back())
						continue
					} else {
						return false
					}
				} else if v == '}' {
					if stack.Back().Value.(byte) == '{' {
						stack.Remove(stack.Back())
						continue
					} else {
						return false
					}
				} else {
					return false
				}
			} else {
				return false
			}

		}
	}

	if stack.Len() > 0 {
		return false
	} else {
		return true
	}
	//return true
}
