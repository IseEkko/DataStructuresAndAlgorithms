package Stack

import "errors"

type Stack []interface{}

//get lenth
func (stack Stack) Len() int {
	return len(stack)
}

//get cap
func (stack Stack) Cap() int {
	return cap(stack)
}

//add element
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

//get value for stack top
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index,len is 0")
	}
	return stack[len(stack)-1], nil
}

//get value change slice ,kick top value
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("stack out of index,len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(*stack)-1]
	return value, nil
}

//verify this stack whether nil
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
