package Array

import "fmt"

type queue struct {
	items           []int
	head, tail, cap int
}

func NewQueue(cap int) *queue {
	return &queue{
		items: make([]int, cap),
		head:  0,
		tail:  0,
		cap:   cap,
	}
}

func (q *queue) enQueue(item int) error {
	// 判断队列是否满, 牺牲一个空间
	if (q.tail+1)%q.cap == q.head {
		return fmt.Errorf("queque is over!")
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.cap
	return nil
}

func (q *queue) deQueue() (int, error) {
	//判断队列是否为空
	if q.tail == q.head {
		return 0, fmt.Errorf("queue is free")
	}
	ret := q.items[q.head]
	q.head = (q.head + 1) % q.cap
	return ret, nil
}

func (q *queue) isEmpty() bool {
	if q.tail == q.head {
		return true
	}
	return false
}
