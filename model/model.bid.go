package model

type Bid struct {
	UserID   uint32 `json:"user_id"`
	UserName string `json:"user_name,omitempty"`
	ItemID   uint32 `json:"item_id"`
	ItemName string `json:"item_name,omitempty"`
	Amount   uint64 `json:"amount"`
	BidTime  int64  `json:"bid_time,omitempty"`
}

type BidNode struct {
	Value Bid
	Next  *BidNode
}

// Singly linked list
type BidLinkedList struct {
	Head *BidNode
	Size uint32
}

// Appends node n to list s
func (s *BidLinkedList) Prepend(n *BidNode) {
	if s.Head == nil {
		s.Head = n
	} else {
		n.Next = s.Head
		s.Head = n
	}

	s.Size++
}

// Has user bid on item
func (s *BidLinkedList) HasUserBidOnItem(userID uint32) bool {
	temp := s.Head
	for temp != nil {
		if temp.Value.UserID == userID {
			return true
		}
		temp = temp.Next
	}

	return false
}

func (s *BidLinkedList) LinkedList2Slice() (slice []Bid) {
	if s.Head == nil {
		return
	}

	head := s.Head
	for head.Next != nil {
		slice = append(slice, head.Value)
		head = head.Next
	}
	slice = append(slice, head.Value)
	return
}
