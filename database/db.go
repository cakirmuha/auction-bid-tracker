package database

import (
	"context"
)

type DB struct {
	ctx context.Context

	cache *dbCache
}

type dbCache struct {
}

func New(ctx context.Context) (*DB, error) {
	c := dbCache{}

	return &DB{
		ctx:   ctx,
		cache: &c,
	}, nil
}
