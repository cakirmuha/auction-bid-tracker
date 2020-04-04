package database

import (
	"fmt"

	"github.com/cakirmuha/auction-bid-tracker/model"
	"github.com/cakirmuha/auction-bid-tracker/util"
)

func (db *DB) GetItemNameByID(itemID uint32) (*string, error) {
	db.cache.userMu.RLock()
	user, ok := db.cache.userCache[itemID]
	db.cache.userMu.Lock()
	if ok {
		return &user.Name, nil
	}
	return nil, fmt.Errorf("item(%v) not available on the system", itemID)
}

func (db *DB) GetCurrentWinningBidByItemID(itemID uint32) (*model.Bid, error) {
	val, ok := db.cache.itemBidCache.Load(itemID)
	if !ok {
		return nil, fmt.Errorf("there is no winning bid for item(%v)", itemID)
	}

	itemBidList := val.(util.LinkedList)
	leadingBid := itemBidList.Head.Value

	// Set Item Name
	itemName, err := db.GetItemNameByID(itemID)
	if err != nil {
		return nil, err
	}
	leadingBid.ItemName = *itemName

	// Set User Name
	userName, err := db.GetUserNameByID(leadingBid.UserID)
	if err != nil {
		return nil, err
	}
	leadingBid.UserName = *userName

	return &leadingBid, nil
}

func (db *DB) GetAllBidsByItemID(itemID uint32) ([]model.Bid, error) {
	val, ok := db.cache.itemBidCache.Load(itemID)
	if !ok {
		return nil, fmt.Errorf("there is no bid for item(%v)", itemID)
	}

	var bids []model.Bid

	itemBidList := val.(util.LinkedList)
	slice := util.LinkedList2Slice(itemBidList)

	// Set Item Name
	itemName, err := db.GetItemNameByID(itemID)
	if err != nil {
		return nil, err
	}

	for _, s := range slice {
		s := s

		// Set item name
		s.ItemName = *itemName

		// Set User Name
		userName, err := db.GetUserNameByID(s.UserID)
		if err != nil {
			return nil, err
		}
		s.UserName = *userName

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

			// Set Item Name
			itemName, err2 := db.GetItemNameByID(itemID)
			if err2 != nil {
				err = err2
				return false
			}
			items = append(items, model.Item{
				ID:   itemID,
				Name: *itemName,
			})
		}
		return true
	})
	return items, nil
}
