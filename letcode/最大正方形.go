/**
* @Author: Gosin
* @Date: 2021/4/22 22:23
 */

package main

import "strconv"

func main() {
	//[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
	//matrix := [][]byte{{1, 0, 1, 0, 0}}
}

func maximalSquare(matrix [][]byte) int {

	maxside := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(dp[i]); j++ {
			baseint := 0
			v, _ := strconv.ParseInt(string(matrix[i][j]), baseint, 10)
			dp[i][j] = int(v)
			if dp[i][j] == 1 {
				maxside = 1
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 || j == 0 {
				continue
			} else {
				if dp[i][j] == 1 {
					dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
					if dp[i][j] > maxside {
						maxside = dp[i][j]
					}
				}

			}
		}
	}

	return maxside * maxside
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
