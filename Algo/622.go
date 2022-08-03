package Algo

type MyCircularQueue struct {
	queue []int
	front int
	rear  int
	empty bool
}

func Constructor(k int) MyCircularQueue {
	obj := new(MyCircularQueue)
	obj.queue = make([]int, k)
	obj.front, obj.rear = 0, 0
	obj.empty = true
	return *obj
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.front == this.rear && !this.empty {
		return false
	} else {
		this.queue[this.rear] = value
		this.rear = (this.rear + 1) % len(this.queue)
		this.empty = false
		return true
	}
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.front == this.rear && this.empty {
		return false
	} else {
		this.front = (this.front + 1) % len(this.queue)
		this.empty = true
		return true
	}
}

func (this *MyCircularQueue) Front() int {
	if this.front == this.rear && this.empty {
		return -1
	} else {
		return this.queue[this.front]
	}
}

func (this *MyCircularQueue) Rear() int {
	if this.rear == this.front && this.empty {
		return -1
	} else {
		return this.queue[(this.rear+len(this.queue)-1)%len(this.queue)]
	}
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear && this.empty
}

func (this *MyCircularQueue) IsFull() bool {
	return this.front == this.rear && !this.empty
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
