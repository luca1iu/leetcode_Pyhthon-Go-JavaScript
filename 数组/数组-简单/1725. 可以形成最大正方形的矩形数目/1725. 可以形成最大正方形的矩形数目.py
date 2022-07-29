class Solution:
    def countGoodRectangles(self, rectangles: List[List[int]]) -> int:
        lens = [min(i) for i in rectangles]
        return lens.count(max(lens))