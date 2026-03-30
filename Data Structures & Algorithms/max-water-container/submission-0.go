func maxArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights)-1; i++ {
		for j := i + 1; j < len(heights); j++ {
			// calculate area
			// min() to get X
			// j-i to get Y
			area := (j - i) * min(heights[i], heights[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}