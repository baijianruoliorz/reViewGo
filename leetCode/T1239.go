package leetCode

/*
*  @author liqiqiorz
*  @data 2021/6/19 1:09
 */
func maxLength(arr []string) int {
	var backtrace func(int, int) int
	backtrace = func(idx, bit int) (rs int) {
		if idx == len(arr) {
			return
		}
		tmp := bit
		for _, v := range arr[idx] {
			dif := v - 'a'
			if bit&(1<<dif) > 0 {
				return backtrace(idx+1, tmp)
			}
			bit |= (1 << dif)
		}
		tb, rs := backtrace(idx+1, bit), backtrace(idx+1, tmp)
		if cur := tb + len(arr[idx]); cur > rs {
			return cur
		}
		return
	}
	return backtrace(0, 0)
}
