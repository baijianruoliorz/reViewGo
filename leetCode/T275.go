package leetCode

func hIndex(citations []int)int{
	n:=len(citations)
	idx:=func()int{
		low,high:=0,n-1
		for low<=high{
			mid:=(low+high)>>1
			if n-mid>citations[mid]{
				low=mid+1
			}else {
				high=mid-1
			}
		}
		return low
	}()
	if citations[idx]==0{
		return 0
	}
	return n-idx
}
