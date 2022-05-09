package pkg

import (
	"testing"
)

func mockTestData() *Node {
	var project = &Node{
		ID:     "proj",
		Type:   "proj",
		Name:   "proj",
		Parent: root,
	}

	var iteration = &Node{
		ID:     "iter",
		Type:   "iter",
		Name:   "iter",
		Parent: project,
	}
	return iteration
}

func mockForPush() *Node {
	return &Node{
		ID:     "iter2",
		Type:   "iter2",
		Name:   "iter2",
		Parent: nil,
	}
}

func TestIntegrity(t *testing.T) {
	var l = mockTestData()
	n := l.Pop()
	if n.ID != "proj" {
		t.Fail()
	}
	n2 := n.Push(mockForPush())
	if n2.ID != "iter2" {
		t.Fail()
	}
	s := n2.LineAge()
	if s != "~/proj/iter2" {
		t.Fail()
	}
}

func TestLineAge(t *testing.T) {
	s := mockTestData().LineAge()
	if s != "~/proj/iter" {
		t.Fail()
	}
}
