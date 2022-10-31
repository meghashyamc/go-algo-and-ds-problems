package stacksandqueues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStackOperations(t *testing.T) {

	assert := assert.New(t)

	minStackTest := &minstack{}
	assert.Equal(-1, minStackTest.getMin())
	assert.Equal(-1, minStackTest.top())
	minStackTest.push(100)
	assert.Equal(100, minStackTest.top())
	minStackTest.push(5)
	assert.Equal(5, minStackTest.top())
	minStackTest.push(10)
	assert.Equal(10, minStackTest.top())
	assert.Equal(5, minStackTest.getMin())
	minStackTest.push(3)
	minStackTest.push(2)
	assert.Equal(2, minStackTest.getMin())
	minStackTest.pop()
	assert.Equal(3, minStackTest.getMin())
	minStackTest.push(1)
	assert.Equal(1, minStackTest.getMin())

}
