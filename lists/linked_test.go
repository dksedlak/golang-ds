package lists_test

import (
	"testing"

	"github.com/dksedlak/golang-ds/lists"
	"github.com/stretchr/testify/assert"
)

func TestCreateLinkedList(t *testing.T) {
	linkedList := lists.NewLinkedList[string]()

	assert.NotNil(t, linkedList)
	assert.Zero(t, linkedList.Size())
	assert.True(t, linkedList.IsEmpty())
}

func TestAddLinkedList(t *testing.T) {
	linkedList := lists.NewLinkedList[uint]()

	linkedList.Add(10)
	linkedList.Add(20)
	linkedList.Add(30)

	assert.NotNil(t, linkedList)
	assert.Equal(t, linkedList.Size(), 3)
	assert.False(t, linkedList.IsEmpty())
}
