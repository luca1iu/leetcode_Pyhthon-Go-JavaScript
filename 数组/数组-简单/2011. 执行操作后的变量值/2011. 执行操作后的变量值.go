func finalValueAfterOperations(operations []string) (ans int) {
	for _, op := range operations {
		if op[1] == '+' {
			ans++
		} else {
			ans--
		}
	}
	return
}
