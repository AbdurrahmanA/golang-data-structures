package linked_list

import "testing"

func setup() *List {
	return &List{}
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
	list := setup()
	cases := []struct {
		value int
	}{
		{1},
		{2},
		{3},
		{4},
	}

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
