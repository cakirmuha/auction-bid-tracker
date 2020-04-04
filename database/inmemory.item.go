package database

import (
	"fmt"

	"github.com/cakirmuha/auction-bid-tracker/model"
	"github.com/cakirmuha/auction-bid-tracker/util"
)

func (db *DB) GetCurrentWinningBidByItemID(itemID uint32) (*model.Bid, error) {
	val, ok := db.cache.itemBidCache.Load(itemID)
	if !ok {
		return nil, fmt.Errorf("there is no winning bid for item(%v)", itemID)
	}

	itemBidList := val.(util.LinkedList)
	leadingBid := itemBidList.Head.Value

	// Set item name
	db.cache.itemMu.RLock()
	item, ok := db.cache.itemCache[itemID]
	db.cache.itemMu.Lock()
	if ok {
		leadingBid.ItemName = item.Name
	} else {
		return nil, fmt.Errorf("item(%v) not available on the system", itemID)
	}

	// Set user name
	db.cache.userMu.RLock()
	user, ok := db.cache.userCache[leadingBid.UserID]
	db.cache.userMu.Lock()
	if ok {
		leadingBid.UserName = user.Name
	} else {
		return nil, fmt.Errorf("user(%v) not available on the system", itemID)
	}

	return &leadingBid, nil
}

func (db *DB) GetAllBidsByItemID(itemID uint32) ([]model.Bid, error) {
	val, ok := db.cache.itemBidCache.Load(itemID)
	if !ok {
		return nil, fmt.Errorf("there is no bid for item(%v)", itemID)
	}

	var (
		bids     []model.Bid
		itemName string
	)

	itemBidList := val.(util.LinkedList)
	slice := util.LinkedList2Slice(itemBidList)

	// Set item name
	db.cache.itemMu.RLock()
	item, ok := db.cache.itemCache[itemID]
	db.cache.itemMu.Lock()
	if ok {
		itemName = item.Name
	} else {
		return nil, fmt.Errorf("item(%v) not available on the system", itemID)
	}

	for _, s := range slice {
		s := s

		// Set item name
		s.ItemName = itemName

		// Set user name
		db.cache.userMu.RLock()
		user, ok := db.cache.userCache[s.UserID]
		db.cache.userMu.Lock()
		if ok {
			s.UserName = user.Name
		} else {
			return nil, fmt.Errorf("user(%v) not available on the system", itemID)
		}

		bids = append(bids, s)
	}

	return bids, nil
}

func (db *DB) GetAllItemsByUserID(userID uint32) ([]model.Item, error) {
	var (
		items []model.Item
		err   error
	)

	db.cache.itemBidCache.Range(func(k, v interface{}) bool {
		itemBidList := v.(util.LinkedList)
		itemID := k.(uint32)
		if itemBidList.HasUserBidOnItem(userID) {

			// Set item name
			db.cache.itemMu.RLock()
			item, ok := db.cache.itemCache[itemID]
			db.cache.itemMu.Lock()
			if ok {
				items = append(items, model.Item{
					ID:   itemID,
					Name: item.Name,
				})
			} else {
				err = fmt.Errorf("item(%v) not available on the system", itemID)
				return false
			}
		}
		return true
	})
	return items, nil
}
