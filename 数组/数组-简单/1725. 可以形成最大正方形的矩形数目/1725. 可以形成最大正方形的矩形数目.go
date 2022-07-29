func countGoodRectangles(rectangles [][]int) (ans int) {
	max := 0
	for _,v := range rectangles{
		cur := min(v[0],v[1])
		if cur > max {
			max = cur
			ans = 1
		} else if cur == max {
			ans ++
		}
	}
	return
}

func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}