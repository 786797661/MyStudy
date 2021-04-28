/**
* @Author: Gosin
* @Date: 2021/4/22 21:55
 */

package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	//fmt.Print(reverseParentheses("(ed(et(oc))el)"))
	fmt.Print(strings.Contains("time", "me1"))
}

func reverseParentheses(s string) string {
	bytes := []byte(s)
	list := list.New()
	for i := 0; i < len(bytes); i++ {
		if (bytes[i]) == '(' {
			list.PushBack(i)
		}
		if (bytes[i]) == ')' {
			index := list.Back().Value.(int)
			chaneg(bytes, index+1, i-1)
			list.Remove(list.Back())
		}
	}
	var res strings.Builder
	for _, v := range bytes {
		if v != '(' && v != ')' {
			res.WriteByte(v)
		}
	}
	return res.String()

}

func chaneg(bytes []byte, i int, j int) {
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
}
