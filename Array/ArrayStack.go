package Array
var cap = 100
type ArrayStack struct {
	items []string
	cap int  //
 	n   int
}

//
func (a *ArrayStack) init() {
   a.cap = cap
   a.items = make([]string, a.cap)
   a.n = 0
}

//
func (a *ArrayStack) push(item string) bool{
	//
	if a.n == a.cap {
		return false
	}
	a.items[a.n] = item
	a.n++
	return true
}

//
func (a *ArrayStack) pop() (string, bool) {
	//
	if a.n <= 0 {
		return "", false
	}
	item :=  a.items[a.n - 1]
	a.n--
	return item, true
}