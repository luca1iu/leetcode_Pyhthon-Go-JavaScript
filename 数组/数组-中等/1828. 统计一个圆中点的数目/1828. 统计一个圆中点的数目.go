func countPoints(points [][]int, queries [][]int) []int {
    ans:=[]int{}
    for i:=0;i<len(queries);i++{
        pq:=queries[i]
        tmp:=0
        for j:=0;j<len(points);j++{
            p0:=points[j]
            x:=p0[0]-pq[0]
            y:=p0[1]-pq[1]
            if x*x+y*y<=pq[2]*pq[2]{
                tmp++
            }
        }
        ans =append(ans, tmp)
    }
    return ans
}
