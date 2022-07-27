/**
 * @param {number[]} nums
 * @return {number[]}
 */
 var buildArray = function(nums) {
    const ans = [];
    for (i=0; i<nums.length; i++) {
        ans.push(nums[nums[i]])
    }
    return ans
};