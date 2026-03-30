func minEatingSpeed(piles []int, h int) int {
if len(piles) == 0 {
		return 0
	}

	sort.Ints(piles)
	if len(piles) == h {
		return piles[h-1]
	}
	if len(piles) == 1 {
		return (piles[0] + h - 1) / h
	}

	l, r := 1, piles[len(piles)-1]
	res := r
	for l <= r {
		mid := l + (r-l)/2 // 4 + (30-4)/2 = 17 per hour

		// chạy tính thời gian với mid
		hours := 0
		for _, pile := range piles {
			tmpHour := pile / mid
			if pile%mid != 0 {
				tmpHour++
			}
			hours += tmpHour
		}
		if hours <= h {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
