package tree


func numTrees(n int) int {
	count := make([]int, 0)

	//init
	count[0] = 0
	count[1] = 1

	for i := 2; i <= n ; i++ {
		for j := 0; j < i; j++ {
			count[i] = count[j] * count[i-j]
		}
	}
    return count[n]
}
