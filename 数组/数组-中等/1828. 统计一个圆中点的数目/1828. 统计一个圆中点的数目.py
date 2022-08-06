# 两点距离公式求解与半径的大小关系
class Solution:
    def countPoints(self, points: List[List[int]], queries: List[List[int]]) -> List[int]:
        res = [0]*len(queries)
        for i in range(len(points)):
            for j in range(len(queries)):
                if (points[i][0]-queries[j][0])**2 + (points[i][1]-queries[j][1])**2 <= queries[j][2]**2:
                    res[j] += 1
        return res


