/**
* @Author: Gosin
* @Date: 2021/4/21 22:50
 */

package main

import (
	"strings"
)

//func main() {
//	//year := []string{"你", "好", "世", "界", "A", "B", "C", "D", "E"}
//	//s0 := year[0:1]
//	//s1 := year[3:5]
//	//s2 := year[0:len(year)]
//	//fmt.Println(s0)
//	//fmt.Println(s1)
//	//fmt.Println(s2)
//	////t.Log(s0, len(s0), cap(s0))
//	////t.Log(s1, len(s1), cap(s1))
//	////t.Log(s2, len(s2), cap(s2))
//	////对S2进行扩容发现 S2指针已经指向了另一个地址。
//	////s2 = append(s2, "F")
//	////fmt.Printf("%p %p", &year, &s2)
//	fmt.Println(lengthOfLongestSubstring("au"))
//}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var count = 1
	for i := 0; i < len(s); i++ {
		var temp = string(s[i])
		for j := i + 1; j < len(s); j++ {
			index := strings.Index(temp, string(s[j]))
			if index != -1 {
				if len(temp) > count {
					count = len(temp)
				}
				break
			} else {
				temp += string(s[j])
			}
			if j == len(s)-1 {
				if len(temp) > count {
					count = len(temp)
				}
			}
		}
	}
	return count
}
