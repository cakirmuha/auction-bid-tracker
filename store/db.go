package store

import (
	"context"
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
	c.userCache = model.Users
	c.itemCache = model.Items

	return &DB{
		ctx:   ctx,
		cache: &c,
	}, nil
}
