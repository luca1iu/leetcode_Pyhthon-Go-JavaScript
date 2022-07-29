func createTargetArray(nums []int, index []int) []int {
    target:=make([]int,len(index))
    for v,i:=range index{
        copy(target[i+1:], target[i:]) 
        target[i]=nums[v]
    }
    return target
}
