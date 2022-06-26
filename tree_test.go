package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeTree(name string) *Node {
	root := NewNode(name, 0)
	root.AddChild(NewNode("A", 1))
	root.AddChild(NewNode("B", 2))
	c := NewNode("C", 3)
	root.AddChild(c)
	return root
}

func makeDeepTree(name string) *Node {
	root := NewNode(name, 0)
	root.AddChild(NewNode("A", 1))
	root.AddChild(NewNode("B", 2))
	c := NewNode("C", 3)
	root.AddChild(c)
	c.AddChild(NewNode("C1", 31))
	c.AddChild(NewNode("C2", 32))
	c3 := NewNode("C3", 33)
	c.AddChild(c3)
	c3.AddChild(NewNode("C31", 331))
	c3.AddChild(NewNode("C32", 332))
	c3.AddChild(NewNode("C33", 333))
	return root
}

func TestTreeSimple(t *testing.T) {
	assert := assert.New(t)

	t0 := NewNode("t0", 0)
	assert.Equal(0, t0.Size())
	assert.Equal("t0", t0.Name())

	t1 := makeTree("t1")
	assert.Equal(3, t1.Size())
	assert.Equal("t1", t1.Name())
}

func TestTreeDeep(t *testing.T) {
	assert := assert.New(t)
	root := makeDeepTree("deep1")

	found := root.GetChild([]string{"C", "C3", "C33"})
	assert.Equal("C3", found.Parent().name)
	assert.Equal("C", found.Parent().Parent().name)
	assert.Equal("deep1", found.Parent().Parent().Parent().name)

	assert.Equal("deep1", found.Root().name)
	assert.Equal("deep1", found.Parent().Root().name)
	assert.Equal("deep1", found.Parent().Parent().Root().name)

	found = root.GetChild([]string{"C"})
	assert.Equal("deep1", found.Root().name)
}

func TestTreeChildParent(t *testing.T) {
	assert := assert.New(t)

	root := makeDeepTree("t1")
	assert.Equal(3, root.Size())

	found := root.GetChild([]string{"xyz"})
	assert.Equal("", found.name)
	assert.Nil(found.data)

	found = root.GetChild([]string{"C"})
	assert.Equal("C", found.name)
	assert.Equal(3, found.data)

	parent := found.Parent()
	assert.Equal(root.name, parent.name)
	assert.Equal(root.data, parent.data)

	found = root.GetChild([]string{"C", "C3"})
	assert.Equal("C3", found.name)
	assert.Equal(33, found.data)

	parent = found.Parent().Parent()
	assert.Equal(root.name, parent.name)
	assert.Equal(root.data, parent.data)
}
