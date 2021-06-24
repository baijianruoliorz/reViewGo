package leetCode

import "sort"

/*
*  @author liqiqiorz
*  @data 2021/6/24 12:11
 */
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	dfss(candidates, nil, target, 0, &res)
	return res
}
func dfss(candidates, nums []int, target, left int, res *[][]int) {
	if target == 0 { //结算
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := left; i < len(candidates); i++ { // left限定，形成分支
		if target < candidates[i] { //剪枝
			return
		}
		dfss(candidates, append(nums, candidates[i]), target-candidates[i], i, res) //分支
	}
}
