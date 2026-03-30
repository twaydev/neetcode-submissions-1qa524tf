
type TimeBasedItem struct {
	tValue string
	tTime  int
}
type TimeMap struct {
	store map[string][]TimeBasedItem
}

func Constructor() TimeMap {

	return TimeMap{
		store: make(map[string][]TimeBasedItem),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], TimeBasedItem{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	// fmt.Printf("key: %v - timestamp: %v - store: %v\n", key, timestamp, this.store)
	if items, exist := this.store[key]; exist {
		// we have items as an arrays asc by tTime
		l, r := 0, len(items)-1
		for l <= r {
			m := l + (r-l)/2
			// fmt.Printf("   %v - %v - %v\n", l, r, items[m])
			if items[m].tTime == timestamp {
				return items[m].tValue
			}
			if items[m].tTime >= timestamp { // move left
				r = m - 1
			} else {
				l = m + 1
			}
		}
		// fmt.Printf("   *%v - %v\n", l, r)
		if l > 0 && items[l-1].tTime <= timestamp {
			return items[l-1].tValue
		}
	}
	return ""
}