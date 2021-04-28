/**
* @Author: Gosin
* @Date: 2021/4/22 23:46
 */

package main

func countSquares(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	maxside := 0
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = matrix[i][j]
		}
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i != 0 && j != 0 && dp[i][j] != 0 {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}
			maxside += dp[i][j]
		}

	}
	return maxside
}
