class Solution:
    def repeatedCharacter(self, s: str) -> str:
        cur = set()
        for i in s:
            if i in cur:
                return i
            cur.add(i)
