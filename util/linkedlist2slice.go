package util

import "github.com/cakirmuha/auction-bid-tracker/model"

func LinkedList2Slice(list LinkedList) (slice []model.Bid) {
	if list.Head == nil {
		return
	}

	head := list.Head
	for head.Next != nil {
		slice = append(slice, head.Value)
		head = head.Next
	}
	slice = append(slice, head.Value)
	return
}
