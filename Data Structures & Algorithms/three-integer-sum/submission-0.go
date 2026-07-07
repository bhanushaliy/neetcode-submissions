func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    for i := 0; i < len(nums); i++{
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        j := i+1
        k := len(nums)-1
        sum := -nums[i]
        for j < k{
            if nums[j]+nums[k] < sum{
                j++
            }else if nums[j]+nums[k] > sum{
                k--
            }else{
                res = append(res, []int{nums[i], nums[j], nums[k]})
                j++
                k--
                for j < k && nums[j] == nums[j-1]{
                    j++
                }
                for j < k && nums[k] == nums[k+1]{
                    k--
                }
            }
        }
    }
    return res
}
