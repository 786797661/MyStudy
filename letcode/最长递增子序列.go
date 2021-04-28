/**
* @Author: Gosin
* @Date: 2021/4/29 0:29
 */

package main

//0,1,0,3,2,3

//1,2,1,3,3,4
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 0
	if len(nums) > 0 {
		dp[0] = 1
		max = 1
		for i := 1; i < len(nums); i++ {
			//dp[i]=1
			maxtemp := 1 //2
			for j := i - 1; j >= 0; j-- {
				if nums[i] > nums[j] {
					if dp[j]+1 > maxtemp {
						maxtemp = dp[j] + 1
					}
				}
			}
			dp[i] = maxtemp
			if max < dp[i] {
				max = dp[i]
			}
		}
	}
	return max

}
