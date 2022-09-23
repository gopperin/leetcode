package array

func MaxArea[T int | int64](height []T) T {
	var max T
	var start, end int
	max, start, end = 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = int(height[start])
			start++
		} else {
			high = int(height[end])
			end--
		}

		temp := width * high
		if temp > int(max) {

			var ti interface{} = &max
			switch v := ti.(type) {
			case *int:
				*v = temp
			case *int64:
				*v = int64(temp)
			}

		}
	}
	return max
}
