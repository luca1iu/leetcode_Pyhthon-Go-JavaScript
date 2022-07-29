/**
 * @param {number[][]} rectangles
 * @return {number}
 */
 var countGoodRectangles = function(rectangles) {
    let max = 0, ans=0
    for (const r of rectangles) {
        const cur = Math.min(r[0],r[1])
        if (cur > max) {
            max = cur
            ans = 1
        } else if (cur==max){
            ans ++
        }
    }
    return ans
};