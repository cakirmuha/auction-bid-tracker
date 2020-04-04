package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/cakirmuha/auction-bid-tracker/database"
)

type SomeContext interface {
	Echo() *echo.Echo
	DB() *database.DB
}

type Context struct {
	echo.Context

	buildInfo string

	ec     *echo.Echo
	logger echo.Logger

	db *database.DB
}

func (c *Context) BuildInfo() string {
	return c.buildInfo
}

func (c *Context) Echo() *echo.Echo {
	return c.ec
}

func (c *Context) DB() *database.DB {
	return c.db
}

func (c *Context) Log() echo.Logger {
	return c.logger
}

func New(ctx context.Context, eLogger echo.Logger, opts ...Option) (*Context, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	serviceOpts := options{}
	for _, opt := range opts {
		opt(&serviceOpts)
	}

	cc := &Context{
		ec:     serviceOpts.ec,
		logger: eLogger,
	}

	if serviceOpts.db {
		db, err := database.New(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "DB")
		}
		cc.db = db
	}

	return cc, nil
}
