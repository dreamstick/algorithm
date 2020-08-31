package Array

func canVisitAllRooms(rooms [][]int) bool {
	var n = len(rooms)
	q := NewQueue(n * n)
	isOpen := make([]bool, n)
	isOpen[0] = true
	q.enQueue(0)
	for{
		if q.isEmpty() {
			break
		}
		item, _ := q.deQueue()
		for _, v := range rooms[item] {
		   if !isOpen[v] {
		   	   isOpen[v] = true
		   	   q.enQueue(v)
		   }
		}
	}
	for i := 0; i < n; i++ {
		if !isOpen[i] {
			return false
		}
	}
	return true
}
