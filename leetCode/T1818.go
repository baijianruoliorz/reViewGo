package leetCode

import "sort"

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
	//这种方式不可用,是引用,并非复制
	//rec := sort.IntSlice(nums1)
	rec := append(sort.IntSlice(nil), nums1...)
	rec.Sort()
	sum, max := 0, 0
	//max = |num1[i]-nums2[i]|-|nums1[j]-nums2[i]|
	for i, v := range nums2 {
		a := Abs(nums1[i] - v)
		sum += a
		//找到最接近v的值,保证max的至最大
		j := rec.Search(v)
		if j < n {
			max = Max(max, a-Abs(rec[j]-v))
		}
		if j > 0 {
			max = Max(max, a-Abs(rec[j-1]-v))
		}
	}
	return (sum - max) % (1e9 + 7)
}

func Max(args ...int) int {
	max := args[0]
	for _, item := range args {
		if item > max {
			max = item
		}
	}
	return max
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
