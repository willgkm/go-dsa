package main

type Stack struct {
	values []int8
}

func (stack *Stack) push(data int8) {
	stack.values = append(stack.values, data)
}

func (stack *Stack) pop() {
	if stack.isEmpty() {
		return
	}
	length := len(stack.values) - 1
	stack.values = stack.values[:length]
}

func (stack *Stack) top() int8 {
	length := len(stack.values)
	return stack.values[length-1]
}

func (stack *Stack) isEmpty() bool {
	return len(stack.values) == int(0)

}

func (stack *Stack) size() int8 {
	return int8(len(stack.values))

}

func main() {}
