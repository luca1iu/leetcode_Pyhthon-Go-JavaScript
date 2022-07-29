class Solution:
    def decompressRLElist(self, nums: List[int]) -> List[int]:
        res = []
        for i in range(0,len(nums),2):
            num = nums[i]
            res.extend([nums[i+1]]*num)
        return res




