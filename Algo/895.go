package algo

type stk struct {
	val int
	pos []int
}

type FreqStack struct {
	stack  []stk
	length int
}

func Constructor() FreqStack {
	stack := FreqStack{make([]stk, 0), 0}
	return stack
}

func (this *FreqStack) Push(val int) {
	this.length++
	exist := false
	for k, v := range this.stack {
		if v.val == val {
			exist = true
			this.stack[k].pos = append(this.stack[k].pos, this.length)
			break
		}
	}
	if !exist {
		this.stack = append(this.stack, stk{val, []int{this.length}})
	}
}

func (this *FreqStack) Pop() int {
	key, most, last := -1, 0, -1
	for k, v := range this.stack {
		if most < len(v.pos) {
			key = k
			most = len(v.pos)
			last = v.pos[most-1]
		} else if most == len(v.pos) && last < v.pos[most-1] {
			key = k
			last = v.pos[most-1]
		}
	}
	res := this.stack[key].val
	if len(this.stack[key].pos) == 1 {
		this.stack = append(this.stack[:key], this.stack[key+1:]...)
	} else {
		this.stack[key].pos = this.stack[key].pos[:most-1]
	}
	return res
}
