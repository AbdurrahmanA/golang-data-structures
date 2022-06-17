package main

import "testing"

var cases = []struct {
	value int
}{
	{1},
	{2},
	{3},
	{4},
}

func emptyList() *List {
	return &List{}
}

func fullList() *List {
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node1.Next = node2
	node2.Prev = node1

	return &List{Head: node1, Tail: node2, Len: 2}
}

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()
	if list.Len != 0 {
		t.Errorf("Expected list.Len to be 0, got %d", list.Len)
	}
	if list.Head != nil {
		t.Errorf("Expected list.Head to be nil, got %v", list.Head)
	}
	if list.Tail != nil {
		t.Errorf("Expected list.Tail to be nil, got %v", list.Tail)
	}
}

func TestAddToHead(t *testing.T) {
	list := emptyList()

	for i, c := range cases {
		temp := list.Head
		list.AddToHead(c.value)

		if list.Len != uint(i)+1 {
			t.Errorf("Expected list.Len to be %d, got %d", i, list.Len)
		}

		if list.Head.Value != c.value {
			t.Errorf("Expected list.Head.Value to be %d, got %d", c.value, list.Head.Value)
		}

		if list.Len == 1 {
			if list.Head != list.Tail {
				t.Errorf("Expected list.Head to be list.Tail, got %v and %v", list.Head, list.Tail)
			}
		} else {
			if list.Head.Next != temp {
				t.Errorf("Expected list.Head.Next to be %v, got %v", temp, list.Head.Next)
			}

			if temp.Prev != list.Head {
				t.Errorf("Expected temp.Prev to be list.Head, got %v", temp.Prev)
			}
		}
	}
}

func TestRemoveFromHead(t *testing.T) {
	list := fullList()

	oldHead := list.Head
	head := list.RemoveFromHead()

	if oldHead != head {
		t.Errorf("Expected head to be %v, got %v", oldHead, head)
	}

	emptyList := emptyList()
	if emptyList.RemoveFromHead() != nil {
		t.Errorf("Expected head to be nil, got %v", head)
	}

}

func TestAddToTail(t *testing.T) {
	list := emptyList()

	for i, c := range cases {
		temp := list.Tail
		list.AddToTail(c.value)

		if list.Len != uint(i)+1 {
			t.Errorf("Expected list.Len to be %d, got %d", i, list.Len)
		}

		if i == 0 {
			if temp != nil {
				t.Errorf("Expected temp to be nil, got %v", temp)
			}
		} else {
			if list.Tail.Prev != temp {
				t.Errorf("Expected list.Tail.Prev to be %v, got %v", temp, list.Tail.Prev)
			}
		}

	}
}

func TestRemoveFromTail(t *testing.T) {
	list := fullList()

	oldTail := list.Tail
	tail := list.RemoveFromTail()

	if oldTail != tail {
		t.Errorf("Expected tail to be %v, got %v", oldTail, tail)
	}

	emptyList := emptyList()

	if emptyList.RemoveFromTail() != nil {
		t.Errorf("Expected tail to be nil, got %v", tail)
	}

}
