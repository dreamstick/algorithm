package main
import  "sort"
func subsets(nums []int) [][]int {
	res := [][]int{}
	var dfs func(start int, list []int)
	dfs = func(start int, list []int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			list = append(list, i)
			dfs(i+1, list)
			list = list[:len(list)-1]
		}
	}
	dfs(0, []int{})
	return res
}

// [4,4,4,1,4]
func subsets2(nums []int) [][]int {
	res := [][]int{}
	//sort
	sort.Slice(nums, func(i, j int) bool {
		return  nums[i] < nums[j]
	})
	var dfs func(start int, list []int)
	dfs = func(start int, list []int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)
		for i := start; i < len(nums) ; i++ {
			if (i > start && nums[i-1] == nums[i]) {
				continue
			}
			list = append(list, nums[i])
			dfs(i+1, list)
			list = list[:len(list)-1]
		}
	}
	dfs(0, []int{})
	return res
}

