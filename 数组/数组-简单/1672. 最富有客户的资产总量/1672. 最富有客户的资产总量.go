func maximumWealth(accounts [][]int) (ans int) {
    for _, account := range accounts {
        cur := 0
        for _, val := range account {
            cur += val
        }
        if cur > ans {
            ans = cur
        }
    }
    return
}

