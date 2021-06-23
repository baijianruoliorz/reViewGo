package leetCode

import "strings"

/*
*  @author liqiqiorz
*  @data 2021/6/23 16:50
 */
func reverseWords(s string) string {
	res := strings.Fields(s)
	length := len(res)
	for i := 0; i < length/2; i++ {
		res[i], res[length-1-i] = res[length-1-i], res[i]
	}
	return strings.Join(res, " ")
}
