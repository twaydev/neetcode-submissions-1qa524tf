/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {

	if len(intervals) == 0 || len(intervals) == 1 {
		return true
	}
	// sort ascending
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})
	// iterative way
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].end > intervals[i+1].start {
			return false
		}
	}

	return true
}
