class Solution:
    def maxIncreaseKeepingSkyline(self, grid: List[List[int]]) -> int:
        n = len(grid)
        rowsMax, colsMax = [max(r) for r in grid], [max(c) for c in zip(*grid)]
        return sum(min(rowsMax[i], colsMax[j]) - grid[i][j] for i in range(n) for j in range(n))