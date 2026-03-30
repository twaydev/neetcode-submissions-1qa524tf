/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
res := 0
	if len(intervals) == 0 {
		return res
	}

	// declare array with start/end and increase/decrease room
	s := [][]int{}
	for _, interval := range intervals {
		s = append(s, []int{interval.start, +1})
		s = append(s, []int{interval.end, -1})
	}
	// sort asc by s[i][0], s[i][1]
	sort.Slice(s, func(i, j int) bool {
		if s[i][0] == s[j][0] {
			return s[i][1] < s[j][1]
		}
		return s[i][0] < s[j][0]
	})
	// itegrate to compare values
	activeRooms := 0
	for _, r := range s {
		activeRooms += r[1]
		if activeRooms > res {
			res = activeRooms
		}
	}
	return res
}
