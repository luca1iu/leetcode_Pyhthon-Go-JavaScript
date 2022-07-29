class Solution:
    def mostWordsFound(self, sentences: List[str]) -> int:
        res = []
        for i in sentences:
            res.append(len(i.split(' ')))
        return max(res)


