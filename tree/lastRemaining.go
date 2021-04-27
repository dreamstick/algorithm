package tree

func lastRemaining(n int, m int) int {
	if n < 1 || m < 1 {
		return -1
	}

	// 模拟
	a := make([]int, n)
	for j := 0; j < n; j++ {
		a[j] = j
	}
	index := 0
	for n > 1 {
		index = (index + m -1) %n
		a = append(a[:index], a[index+1:]...)
		n--
	}
	return a[0]
}
