
func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([][2]float64, n)
	for i := range n {
		cars[i] = [2]float64{float64(position[i]), float64(speed[i])}
		// fmt.Println(cars, cars[i], position[i], speed[i], i)
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})
	fleets := []float64{}
	for i := range n {
		timeReachTarget := (float64(target) - cars[i][0]) / cars[i][1]
		// fmt.Printf("i=%d, fleets=%v, timeReachTarget=%v\n", i, fleets, timeReachTarget)

		// mặc định thêm xe hiện tại vào fleets
		fleets = append(fleets, timeReachTarget)

		// nếu xe hiện tại chạy nhanh hơn xe phía trước (thời gian tới đích nhỏ hơn)
		// thì hợp nhất 2 xe bằng cách xóa xe phía trước khỏi fleets
		for len(fleets) >= 2 && fleets[len(fleets)-1] <= fleets[len(fleets)-2] {
			fleets = fleets[:len(fleets)-1]
		}
	}
	return len(fleets)
}