package leetCode

/*
*  @author liqiqiorz
*  @data 2021/6/23 16:40
 */
func hammingWeight(num uint32) int {
	ans := 0
	for num != 0 {
		num = num & (num - 1)
		ans++
	}
	return ans
}
