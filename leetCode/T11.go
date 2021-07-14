package leetCode

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var ret int
	for i < j {
		var area int
		if height[i] > height[j] {
			area = (j - i) * height[j]
		} else {
			area = (j - i) * height[i]
		}
		if area > ret {
			ret = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}
	return ret
}
