func minPartitions(n string) int {
	max := -1
	for _, b := range n {
		if int(b-'0') > max {
			max = int(b - '0')
		}
	}
	return max
}