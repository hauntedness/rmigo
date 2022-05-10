package pkg

type Node struct {
	ID     string
	Type   string
	Name   string
	Parent *Node
}

func (n *Node) Pop() (parent *Node) {
	if n.Parent == nil {
		return n.Clear()
	}
	parent = n.Parent
	n.Parent = nil
	return
}

func (n *Node) Push(child *Node) *Node {
	child.Parent = n
	return child
}

func (n *Node) Clear() *Node {
	return root
}

func (n *Node) LineAge() string {
	if n.Parent == nil {
		return n.Name
	} else {
		return n.Parent.LineAge() + "/" + n.Name
	}
}
