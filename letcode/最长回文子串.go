/**
* @Author: Gosin
* @Date: 2021/4/21 21:35
 */

package main

//
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
//	fmt.Println(longestPalindrome("cbbd"))
//}

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	//对角线上的类型都为true
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	//定义初始位置
	index := 0
	length := 1
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > length {
				index = i
				length = j - i + 1
			}
		}
	}
	content := s[index : length+index]

	return content
}
