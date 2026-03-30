func minEatingSpeed(piles []int, h int) int {// special case
	if len(piles) == 0 {
		return 0
	}

	// get maxPile
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	l, r := 1, maxPile
	res := r
	for l <= r {
		mid := l + (r-l)/2

		// chạy tính thời gian với mid
		hours := 0
		for _, pile := range piles {
			// hours += (pile + mid - 1) / mid
			tmpHour := pile / mid
			if pile%mid != 0 {
				tmpHour++
			}
			hours += tmpHour
			if hours > h {
				break
			}
		}
		if hours <= h {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
		// fmt.Println("res:", res, "mid:", mid, "hours:", hours, "l-r", l, r)
	}
	return res
}
