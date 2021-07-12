package main

import "fmt"

func combine(n int, k int) [][]int {
	var res [][]int
	var dfs func(n, level, start int, list []int)
	dfs = func(n, level, start int, list []int) {
		if level == k {
			// finish
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
		}
		for i:=start; i <= n; i++ {
			list = append(list, i)
			dfs(n, level+1, i+1, list)
			list = list[:len(list)-1]
		}

	}
	dfs(n, 0, 1, []int{})
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var dfs  func(candidates []int, target int, idx int, list []int)
	dfs = func(candidates []int, target int, idx int, list []int) {
	     	if idx == len(candidates) {
				return
			}
			if target < 0 {
				return
			}
			if target == 0 {
				tmp := make([]int, len(list))
				copy(tmp, list)
				res = append(res, tmp)
				return
			}
			// 不选
			dfs(candidates, target, idx+1, list)

	     	//选
	     	list = append(list, candidates[idx])
	     	dfs(candidates, target-candidates[idx], idx, list)
	     	list = list[:len(list)-1]
	}
	dfs(candidates, target, 0, []int{})
	return res
}

func main()  {
	fmt.Println(combine(4,2))

}
