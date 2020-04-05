package store

import (
	"context"
	"fmt"
	"sync"

	"github.com/cakirmuha/auction-bid-tracker/model"
)

type DB struct {
	ctx context.Context

	cache *dbCache
}

type dbCache struct {
	itemBidCache sync.Map

	userCache map[uint32]model.User
	userMu    sync.RWMutex

	itemCache map[uint32]model.Item
	itemMu    sync.RWMutex
}

func New(ctx context.Context) (*DB, error) {
	c := dbCache{
		userCache: make(map[uint32]model.User),
		itemCache: make(map[uint32]model.Item),
	}

	// Set initial mock users and items
	c.userCache = CreateMockUsers()
	c.itemCache = CreateMockItems()

	return &DB{
		ctx:   ctx,
		cache: &c,
	}, nil
}

func CreateMockUsers() map[uint32]model.User {
	users := make(map[uint32]model.User)
	for i := 1; i <= 100; i++ {
		users[uint32(i)] = model.User{
			ID:   uint32(i),
			Name: fmt.Sprintf("User#%v", i),
		}
	}
	return users
}

func CreateMockItems() map[uint32]model.Item {
	items := make(map[uint32]model.Item)
	for i := 1; i <= 4; i++ {
		items[uint32(i)] = model.Item{
			ID:   uint32(i),
			Name: fmt.Sprintf("Item#%v", i),
		}
	}
	return items
}
