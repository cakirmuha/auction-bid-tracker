package interactors

import "github.com/cakirmuha/auction-bid-tracker/service"

type ServerContext struct {
	service.Context // No ptr because we want a copy of echo.Context so that we can assign it per request to separate values
}
