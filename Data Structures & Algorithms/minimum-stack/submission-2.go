
type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   0,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)


	if len(this.stack) == 1 {
		this.min = val
	} else if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	curr := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	// set new min if pop current min
	if curr == this.min && len(this.stack) > 0 {
		newMin := this.stack[0]
		for _, v := range this.stack {
			if v < newMin {
				newMin = v
			}
		}
		this.min = newMin
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}