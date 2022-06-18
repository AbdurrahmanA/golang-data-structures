package main

import "fmt"

/*
 Question: Take the first k elements from sequence and add them to end of sequence.
 https://youtu.be/IPSaG9RRc-k?list=PLUl4u3cNGP63EdVPNLG3ToM6LaEUuStEY&t=2209
*/
func shift_left(list *List, k int) {
	if k == 0 || list.Len == 0 {
		return
	}

	if uint(k) > list.Len {
		panic("k is greater than list length")
	}

	head := list.RemoveFromHead()
	list.AddToTail(head.Value)
	shift_left(list, k-1)
}

/*
 Question: Given a singly linked list, find the middle of the linked list. For example, if the given linked list is 1->2->3->4->5 then the output should be 3.
 If there are even nodes, then there would be two middle nodes, we need to print the second middle element. For example, if the given linked list is 1->2->3->4->5->6 then the output should be 4.
 https://www.geeksforgeeks.org/write-a-c-function-to-print-the-middle-of-the-linked-list/
*/

func find_middle(list *List) {
	if list.Len == 0 {
		return
	}

	fmt.Printf("%+v \n", list.Mid)
}
