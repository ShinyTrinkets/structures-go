package structs

import (
	"reflect"
)

type Node struct {
	name     string
	parent   *Node
	children []*Node
	data     interface{}
}

func NewNode(name string, data interface{}) *Node {
	return &Node{name: name, data: data}
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) NoChildren() int {
	return len(n.children)
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) Root() *Node {
	root := n.parent
	for {
		if root.parent == nil {
			return root
		}
		if root != nil {
			root = root.parent
		}
	}
}

func (n *Node) GetChild(path []string) *Node {
	found := &Node{}
	children := n.children
	for _, name := range path {
		current := &Node{}
		for _, c := range children {
			if c.name == name {
				current = c
				break
			}
		}
		if current.name == name {
			found = current
			children = current.children
		} else {
			return found
		}
	}
	return found
}

func (n *Node) SetName(name string) {
	n.name = name
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func (n *Node) AddChild(c *Node) bool {
	if reflect.DeepEqual(n, c) {
		// cannot add itself as child
		return false
	}
	for _, child := range n.children {
		if child.name == c.name {
			// cannot have 2 children with the same name
			return false
		}
	}
	c.parent = n
	n.children = append(n.children, c)
	return true
}

func (n *Node) DelChild(c *Node) bool {
	for i, child := range n.children {
		if reflect.DeepEqual(child, c) {
			n.children = append(n.children[:i], n.children[i+1:]...)
			for _, orphan := range c.children {
				orphan.parent = n
				n.children = append(n.children, orphan)
			}
			return true
		}
	}
	return false
}
