package leetCode
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

len1 := len(nums1)
len2 := len(nums2)
lenSum := len1+len2
if lenSum == 0 {
return float64(0)
}
l, r := 0, 0
a := make([]int,0,lenSum)
for l <len1 && r < len2{
if nums1[l] < nums2[r]{
a = append(a,nums1[l])
l++
}else {
a = append(a, nums2[r])
r++
}
}
a = append(a, nums1[l:]...)
a = append(a, nums2[r:]...)
if lenSum%2 != 0 {
return float64(a[lenSum/2])
}else {
return (float64(a[lenSum/2-1]) + float64(a[lenSum/2]))/2
}
}