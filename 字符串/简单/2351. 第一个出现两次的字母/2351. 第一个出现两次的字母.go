func repeatedCharacter(s string) byte {
    myMap := make(map[rune] int)  
    var result byte
    for i, v := range s {
		if _, isMapContainsKey := myMap[v]; isMapContainsKey{
            result = byte(v)
            break       
		} else {
			myMap[v] = i
		}
	}
    return result
}