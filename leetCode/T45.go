package leetCode

import "math"

/*
*  @author liqiqiorz
*  @data 2021/6/17 16:43
 */

func jump(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	dp[size-1] = 0
	for i := size - 2; i >= 0; i-- {
		dp[i] = 1<<31 - 1
		for x := 1; x <= nums[i]; x++ {
			if i+x < size {
				dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i+x])))
			}
		}
	}
	return dp[0]
}
