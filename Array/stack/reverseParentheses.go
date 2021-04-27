package stack

func reverseParentheses(s string) string {
     var f []int // 记录（前面有多少字符
     var res []byte // 返回最后的结果
     for i := 0; i< len(s); i++ {
     	m := s[i]
     	if m == '(' {
     		f = append(f, len(res))
		} else if m == ')' {
			l, r := f[len(f) -1], len(res) -1
			for l < r {
				res[l], res[r] = res[r], res[l]
				l++
				r--
			}
			f = f[:len(f)-1]
		} else {
			res = append(res, m)
		}
	 }
	 return string(res)
}
