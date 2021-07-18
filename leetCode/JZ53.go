package leetCode
func search(nums []int,target int)int{
	l:=helper(nums,0,len(nums)-1,target)
	r:=helper(nums,l,len(nums)-1,target+1)
	return r-l
}

func helper(nums []int, l,r,target int)int{
	for l<=r{
		mid:=(l+r)/2
		if nums[mid] < target {
			l=mid+1
		}else {
			r=mid-1
		}
	}
	return l
}