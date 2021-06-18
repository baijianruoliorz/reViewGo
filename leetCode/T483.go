package leetCode

import (
	"math"
	"math/bits"
	"strconv"
)

/*
*  @author liqiqiorz
*  @data 2021/6/18 17:51
 */
func smallestGoodBase(n string) string {
	v, _ := strconv.Atoi(n)
	lenmax := bits.Len(uint(int(v)))
	for i := lenmax - 1; i > 1; i-- {
		k := int(math.Pow(float64(v), 1/float64(i)))
		sum, s := 1, 1
		for j := 0; j < i; j++ {
			s *= k
			sum += s
		}
		if sum == v {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(v)
}
