package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFIFOUniqueList(t *testing.T) {
	elements := []int{0, 1, 2, 3, 4, 5}
	list := NewFIFOUniqueList(10, func(left, right int) bool { return left == right }, elements...)
	assert.Equal(t, list.FirstElem.Value, elements[len(elements)-1])
	assert.Equal(t, list.LastElem.Value, elements[0])
	assert.Equal(t, list.Size, len(elements))
}

func TestFIFOUniqueList_Push(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}
	assert.Equal(t, listSize-1, list.FirstElem.Value)
	assert.Equal(t, 0, list.LastElem.Value)
	assert.Equal(t, listSize, list.Size)

	for i := 0; i < listSize; i++ {
		list.Push(i)
	}
	assert.Equal(t, listSize-1, list.FirstElem.Value)
	assert.Equal(t, 0, list.LastElem.Value)
	assert.Equal(t, listSize, list.Size)

	shift := 10
	for i := 0; i < (listSize + shift); i++ {
		list.Push(i)
	}
	assert.Equal(t, (listSize+shift)-1, list.FirstElem.Value)
	assert.Equal(t, shift, list.LastElem.Value)
	assert.Equal(t, listSize, list.Size)
}

func TestFIFOUniqueList_Remove(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}
	list.Remove(listSize - 1)
	assert.Equal(t, listSize-2, list.FirstElem.Value)
	assert.Equal(t, listSize-1, list.Size)

	list.Remove(0)
	assert.Equal(t, 1, list.LastElem.Value)
	assert.Equal(t, listSize-2, list.Size)
}

func TestFIFOUniqueList_RemoveOne(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}

	list.RemoveOne(func(prev, curr, next int) bool {
		return curr%2 == 0 && curr > 5
	})
	assert.Equal(t, listSize-1, list.Size)
}

func TestFIFOUniqueList_RemoveMany(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}

	list.RemoveMany(func(prev, curr, next int) bool {
		return curr%2 == 0
	})
	assert.Equal(t, listSize/2, list.Size)
}

func TestFIFOUniqueList_Pop(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}

	list.Pop()
	assert.Equal(t, 1, list.LastElem.Value)
	assert.Equal(t, listSize-1, list.Size)
}

func TestFIFOUniqueList_FindOne(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}

	actual, isFound := list.FindOne(func(prev, curr, next int) bool {
		return curr%2 == 0 && curr > 5
	})
	assert.True(t, isFound)
	assert.NotEqual(t, 6, actual)

	_, isFound = list.FindOne(func(prev, curr, next int) bool {
		return curr == 10
	})
	assert.False(t, isFound)
}

func TestFIFOUniqueList_FindMany(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}

	actual := list.FindMany(func(prev, curr, next int) bool {
		return curr%2 == 0
	})
	expected := []int{8, 6, 4, 2, 0}
	assert.Equal(t, len(expected), len(actual))
	for i := 0; i < len(actual); i++ {
		assert.Equal(t, expected[i], actual[i])
	}
}
