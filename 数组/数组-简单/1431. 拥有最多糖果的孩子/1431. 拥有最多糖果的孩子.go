func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := 0
    for _, c := range candies {
        if c > max {
            max = c
        }
    }

    result := []bool{}
    for _, c := range candies {
        res := (c + extraCandies >= max)
        result = append(result, res)
    }

    return result
}
