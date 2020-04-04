package util

import (
	"github.com/cakirmuha/auction-bid-tracker/model"
)

type Node struct {
	Value model.Bid
	Next  *Node
}

// Singly linked list
type LinkedList struct {
	Head *Node
	Size uint32
}

// Appends node n to list s
func (s *LinkedList) Prepend(n *Node) {
	if s.Head == nil {
		s.Head = n
	} else {
		n.Next = s.Head
		s.Head = n
	}

	s.Size++
}

// Has user bid on item
func (s *LinkedList) HasUserBidOnItem(userID uint32) bool {
	temp := s.Head
	for temp != nil {
		if temp.Value.UserID == userID {
			return true
		}
		temp = temp.Next
	}

	return false
}
