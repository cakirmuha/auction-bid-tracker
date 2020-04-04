package database

import (
	"fmt"

	"github.com/cakirmuha/auction-bid-tracker/model"
	"github.com/cakirmuha/auction-bid-tracker/util"
)

func (db *DB) SaveUserBidOnItem(bid model.Bid) error {
	val, ok := db.cache.itemBidCache.Load(bid.ItemID)
	if !ok {
		// New singly linked list
		list := util.LinkedList{
			Size: 0,
		}

		// Add new node
		n := &util.Node{
			Value: bid,
		}

		list.Append(n)
		db.cache.itemBidCache.Store(bid.ItemID, list)
	} else {
		list := val.(util.LinkedList)
		leadingBid := list.Head.Value
		if bid.Amount > leadingBid.Amount {
			// Add new node
			n := &util.Node{
				Value: bid,
			}
			list.Append(n)
			db.cache.itemBidCache.Store(bid.ItemID, list)
		} else {
			return fmt.Errorf("bid amount(%v) must be greater than the highest(%v)", bid.Amount, leadingBid.Amount)
		}
	}
	return nil
}
