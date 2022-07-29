func mostWordsFound(sentences []string) (ans int) {
    for _, v := range sentences {
        cnt := strings.Count(v, " ") +1
        if cnt > ans {
            ans = cnt
        }
    }
    return ans
}