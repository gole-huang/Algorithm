package Algo

type MyCircularDeque struct {
	queue       []int
	head, tail  int
	empty, full bool
}

func mcdConstructor(k int) MyCircularDeque {
	mcd := new(MyCircularDeque)
	mcd.queue = make([]int, k)
	mcd.head, mcd.tail = 0, 0
	mcd.full, mcd.empty = false, true
	return *mcd
}

func (mcd *MyCircularDeque) InsertFront(value int) bool {
	if mcd.full {
		return false
	}
	mcd.head = (mcd.head - 1 + len(mcd.queue)) % len(mcd.queue)
	mcd.queue[mcd.head] = value
	mcd.empty = false
	if mcd.head == mcd.tail {
		mcd.full = true
	}
	return true
}

func (mcd *MyCircularDeque) InsertLast(value int) bool {
	if mcd.full {
		return false
	}
	mcd.queue[mcd.tail] = value
	mcd.empty = false
	mcd.tail = (mcd.tail + 1) % len(mcd.queue)
	if mcd.tail == mcd.head {
		mcd.full = true
	}
	return true
}

func (mcd *MyCircularDeque) DeleteFront() bool {
	if mcd.empty {
		return false
	}
	mcd.head = (mcd.head + 1) % len(mcd.queue)
	if mcd.head == mcd.tail {
		mcd.empty = true
	}
	mcd.full = false
	return true
}

func (mcd *MyCircularDeque) DeleteLast() bool {
	if mcd.empty {
		return false
	}
	mcd.tail = (mcd.tail - 1 + len(mcd.queue)) % len(mcd.queue)
	if mcd.tail == mcd.head {
		mcd.empty = true
	}
	mcd.full = false
	return true
}

func (mcd *MyCircularDeque) GetFront() int {
	if mcd.empty {
		return -1
	} else {
		return mcd.queue[mcd.head]
	}
}

func (mcd *MyCircularDeque) GetRear() int {
	if mcd.empty {
		return -1
	} else {
		return mcd.queue[(mcd.tail-1+len(mcd.queue))%len(mcd.queue)]
	}
}

func (mcd *MyCircularDeque) IsEmpty() bool {
	return mcd.empty
}

func (mcd *MyCircularDeque) IsFull() bool {
	return mcd.full
}
