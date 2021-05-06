package leetCode

/*
*  @author liqiqiorz
*  @data 2021/4/26 21:39
 */
func shipWithinDays(weights []int, D int) int {
	sum := 0
	for _, v := range weights {
		sum += v
	}

	var check func(x int) bool
	check = func(x int) bool {
		count, temp := 0, 0
		for _, v := range weights {
			if v > x {
				return false
			} else if temp+v > x {
				count += 1
				temp = 0
			}
			temp += v
		}
		if temp != 0 {
			count += 1
		}
		return count <= D
	}
	l, r, mid := 0, sum, 0
	for l+1 < r {
		mid = 1 + (r-1)/2
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
