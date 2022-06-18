package main

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
