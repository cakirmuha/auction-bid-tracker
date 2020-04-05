package store

import (
	"fmt"

	"github.com/cakirmuha/auction-bid-tracker/model"
)

func (db *DB) SaveUserBidOnItem(bid model.Bid) error {
	val, ok := db.cache.itemBidCache.Load(bid.ItemID)
	if !ok {
		// New singly linked list
		list := model.BidLinkedList{
			Size: 0,
		}

		// Add new node
		n := &model.BidNode{
			Value: bid,
		}

		list.Prepend(n)
		db.cache.itemBidCache.Store(bid.ItemID, list)
		return nil
	}

	list := val.(model.BidLinkedList)
	leadingBid := list.Head.Value
	if bid.Amount > leadingBid.Amount {
		// Add new node
		n := &model.BidNode{
			Value: bid,
		}
		list.Prepend(n)
		db.cache.itemBidCache.Store(bid.ItemID, list)
		return nil
	}

	return fmt.Errorf("bid amount(%v) must be greater than the highest(%v)", bid.Amount, leadingBid.Amount)
}
