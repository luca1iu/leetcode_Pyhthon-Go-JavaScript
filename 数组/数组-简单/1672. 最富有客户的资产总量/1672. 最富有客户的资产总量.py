class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        res = []
        for item in accounts:
            res.append(sum(item))
        return max(res)


class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        res = [sum(i) for i in accounts]
        return max(res)
