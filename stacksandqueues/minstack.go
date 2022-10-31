package stacksandqueues

/* Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) – Push element x onto stack.
pop() – Removes the element on top of the stack.
top() – Get the top element.
getMin() – Retrieve the minimum element in the stack.
Note that all the operations have to be constant time operations.

Questions to ask the interviewer :

Q: What should getMin() do on empty stack?
A: In this case, return -1.

Q: What should pop do on empty stack?
A: In this case, nothing.

Q: What should top() do on empty stack?
A: In this case, return -1
NOTE : If you are using your own declared global variables, make sure to clear them out in the constructor. */

type normalStack []int

func (n *normalStack) push(i int) {
	*n = append(*n, i)
}

func (n *normalStack) pop() {
	if len(*n) == 0 {
		return
	}
	*n = (*n)[:len(*n)-1]
}

func (n *normalStack) peek() int {
	if len(*n) == 0 {
		return -1
	}
	return (*n)[len(*n)-1]
}

func (n *normalStack) isEmpty() bool {
	return len(*n) == 0
}

type minstack struct {
	elements *normalStack
	mins     *normalStack
}

func (s *minstack) push(i int) {

	s.initIfNeeded()
	s.elements.push(i)
	if s.mins.isEmpty() {
		s.mins.push(i)
		return
	}

	s.mins.push(min(i, s.mins.peek()))
}

func (s *minstack) pop() {
	s.initIfNeeded()

	s.elements.pop()
	s.mins.pop()
}

func (s *minstack) top() int {

	s.initIfNeeded()
	return s.elements.peek()
}

func (s *minstack) getMin() int {
	s.initIfNeeded()
	return s.mins.peek()

}

func (s *minstack) initIfNeeded() {

	if s.elements == nil {
		s.elements = new(normalStack)
	}
	if s.mins == nil {
		s.mins = new(normalStack)
	}
}

func min(x, y int) int {

	if x < y {
		return x
	}
	return y

}
