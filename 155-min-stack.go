package main

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

// Implement the MinStack class:

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.

type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	if len(ms.minStack) == 0 {
		ms.minStack = append(ms.minStack, val)
	} else {
		if min := ms.minStack[len(ms.minStack)-1]; val < min {
			ms.minStack = append(ms.minStack, val)
		} else {
			ms.minStack = append(ms.minStack, min)
		}
	}
	ms.stack = append(ms.stack, val)
}

func (ms *MinStack) Pop() {
	if sl := len(ms.stack); sl > 0 {
		ms.stack = ms.stack[:sl-1]
		ms.minStack = ms.minStack[:sl-1]
	}
}

func (ms *MinStack) Top() int {
	if sl := len(ms.stack); sl == 0 {
		return 0
	} else {
		return ms.stack[sl-1]
	}
}

func (ms *MinStack) GetMin() int {
	if msl := len(ms.minStack); msl == 0 {
		return 0
	} else {
		return ms.minStack[msl-1]
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
