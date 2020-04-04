package interactors

import (
	"net/http"
	"strconv"

	"github.com/cakirmuha/auction-bid-tracker/model"
	"github.com/cakirmuha/auction-bid-tracker/service"
	"github.com/labstack/echo/v4"
)

func getAllBidsByItem(c echo.Context) error {
	cc := c.(*ServerContext)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id < 1 {
		return model.NewApiError(model.ApiErrorBadRequest, "Invalid item id", nil)
	}

	bids, err := cc.DB().GetAllBidsByItemID(uint32(id))
	if err != nil {
		return model.NewApiError(model.ApiErrorNotAcceptable, err.Error(), nil)
	}

	return service.ApiResponse(c, http.StatusOK, bids)
}

func getCurrentWinningBidByItem(c echo.Context) error {
	cc := c.(*ServerContext)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id < 1 {
		return model.NewApiError(model.ApiErrorBadRequest, "Invalid item id", nil)
	}

	bid, err := cc.DB().GetCurrentWinningBidByItemID(uint32(id))
	if err != nil {
		return model.NewApiError(model.ApiErrorNotAcceptable, err.Error(), nil)
	}

	return service.ApiResponse(c, http.StatusOK, bid)
}
