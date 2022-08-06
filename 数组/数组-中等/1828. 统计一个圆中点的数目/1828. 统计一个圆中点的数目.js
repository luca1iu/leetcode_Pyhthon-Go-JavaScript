/**
 * @param {number[][]} points
 * @param {number[][]} queries
 * @return {number[]}
 */
 var countPoints = function(points, queries) {
    let res = new Array(queries.length).fill(0);
    for (let i=0; i<points.length; i++) {
        for (let j=0; j<queries.length; j++) {
            if ((points[i][0]-queries[j][0])**2+(points[i][1]-queries[j][1])**2 <= queries[j][2]**2) {
                res[j] ++
            }
        }
    }
    return res
};