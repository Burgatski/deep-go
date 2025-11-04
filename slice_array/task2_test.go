package main

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v ./slice_array

type CircularQueue[T constraints.Signed] struct {
	values []T
	head   int
	tail   int
	count  int
}

func NewCircularQueue[T constraints.Signed](size int) CircularQueue[T] {
	if size <= 0 {
		panic("size must be > 0")
	}
	return CircularQueue[T]{
		values: make([]T, size),
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	q.values[q.tail] = value
	q.tail = (q.tail + 1) % len(q.values)
	q.count++

	return true
}

func (q *CircularQueue[T]) Pop() (T, bool) {
	var zero T
	if q.Empty() {
		return zero, false
	}
	v := q.values[q.head]
	q.values[q.head] = zero
	q.head = (q.head + 1) % len(q.values)
	q.count--

	return v, true
}

func (q *CircularQueue[T]) Front() (T, bool) {
	var zero T
	if q.Empty() {
		return zero, false
	}

	return q.values[q.head], true
}

func (q *CircularQueue[T]) Back() (T, bool) {
	var zero T
	if q.Empty() {
		return zero, false
	}

	return q.values[(q.tail-1+len(q.values))%len(q.values)], true
}

func (q *CircularQueue[T]) Empty() bool {
	return q.count <= 0
}

func (q *CircularQueue[T]) Full() bool {
	return q.count == len(q.values)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	v, ok := queue.Front()
	assert.False(t, ok)
	assert.Equal(t, 0, v)

	v, ok = queue.Back()
	assert.False(t, ok)
	assert.Equal(t, 0, v)

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	v, ok = queue.Front()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = queue.Back()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	v, ok = queue.Front()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = queue.Back()
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	_, ok = queue.Pop()
	assert.True(t, ok)
	_, ok = queue.Pop()
	assert.True(t, ok)
	_, ok = queue.Pop()
	assert.True(t, ok)
	_, ok = queue.Pop()
	assert.False(t, ok)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
