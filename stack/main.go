package stack

import "container/list"

type MinStack struct {
	stack *list.List
	minStack *list.List
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: list.New(),
		minStack:list.New(),
	}
}


func (this *MinStack) Push(x int)  {
	this.stack.PushBack(x)
	if this.minStack.Len() != 0{
		v := this.minStack.Back()
		if x < v.Value.(int) {
			this.minStack.PushBack(x)
		}
	}else{
		this.minStack.PushBack(x)
	}

}


func (this *MinStack) Pop()  {
	x := this.stack.Back()
	if this.minStack.Len() != 0 {
		v := this.minStack.Back()
		if x.Value.(int) == v.Value.(int) {
			this.minStack.Remove(v)
		}
	}

	this.stack.Remove(x)
}


func (this *MinStack) Top() int {
	x := this.stack.Back()
	return x.Value.(int)
}


func (this *MinStack) Min() int {
	if this.minStack.Len() != 0 {
		v := this.minStack.Back()
		return v.Value.(int)
	}else{
		return -1
	}

}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
