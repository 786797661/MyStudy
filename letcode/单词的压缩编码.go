/**
* @Author: Gosin
* @Date: 2021/4/24 0:23
 */

package main

import (
	"fmt"
)

func main() {
	//s := "tixme"
	//x := "me"
	//temp := s[len(x):len(s)]
	//temp := s[len(s)-len(x) : len(s)]
	fmt.Print(minimumLengthEncoding([]string{"time", "me", "bell"}))
	//fmt.Print(temp)
}

func minimumLengthEncoding(words []string) int {
	mapwords := make(map[string]bool)
	for _, v := range words {
		mapwords[v] = true
	}
	for _, v := range words {
		for i := 1; i < len(v); i++ {
			delete(mapwords, v[i:])
		}
	}
	var count = 0
	for k, _ := range mapwords {
		count += len(k) + 1
	}
	return count

}
func comper(s, x string) (res string, flag bool) {
	if len(x) > len(s) {
		return change(x, s)
	} else {
		//temp := s[len(s)-len(x) : len(s)]
		//		//if temp == x {
		//		//	return s, true
		//		//}

		return change(s, x)
	}
}

func change(x, s string) (string, bool) {
	temp := x[len(x)-len(s) : len(x)]
	if temp == s {
		return x, true
	} else {
		return x, false
	}
}

//func comper(s, x string) (res string, flag bool) {
//	res = s
//	queue := list.New()
//	for i := 0; i < len(s); i++ {
//		queue.PushBack(i)
//	}
//	tempqueue := list.New()
//	for i := 0; i < len(x); i++ {
//		tempqueue.PushBack(i)
//	}
//
//	flag = true
//	temp := -1
//	for queue.Len() > 0 && tempqueue.Len() > 0 {
//		new := x[tempqueue.Front().Value.(int)]
//		old := s[queue.Front().Value.(int)]
//		if new != old {
//			queue.Remove(queue.Front())
//			flag = false
//		} else {
//			if temp == -1 {
//				temp = queue.Front().Value.(int)
//				flag = true
//			} else {
//				if queue.Front().Value.(int)-temp != 1 {
//					flag = false
//				}
//			}
//			tempqueue.Remove(tempqueue.Front())
//			queue.Remove(queue.Front())
//		}
//	}
//	if !flag {
//		return res + x, flag
//	} else {
//		for tempqueue.Len() > 0 {
//			res += string(x[tempqueue.Front().Value.(int)])
//			tempqueue.Remove(tempqueue.Front())
//		}
//		return res, flag
//	}
//
//}
