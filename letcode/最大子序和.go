/**
* @Author: Gosin
* @Date: 2021/4/28 23:04
 */

package main

func main() {
	dp := []int{5, 4, -1, 7, 8}
	maxSubArray(dp)
}
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	max := 0
	if len(nums) > 0 {
		max = nums[0]
		dp[0] = nums[0]

		for index, v := range nums {
			if index == 0 {
				continue
			}
			if dp[index-1] > 0 {
				dp[index] = dp[index-1] + v
			} else {
				dp[index] = v
			}
			if max < dp[index] {
				max = dp[index]
			}
		}
	}
	return max

}
