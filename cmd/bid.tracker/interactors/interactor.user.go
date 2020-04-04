package interactors

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cakirmuha/auction-bid-tracker/model"
	"github.com/cakirmuha/auction-bid-tracker/service"
	"github.com/labstack/echo/v4"
)

func createUserBid(c echo.Context) error {
	cc := c.(*ServerContext)

	r := model.Bid{}
	if err := c.Bind(&r); err != nil {
		return model.NewApiError(model.ApiErrorBadRequest, "bind", err)
	}

	// If Item Available
	if _, err := cc.DB().GetItemNameByID(r.ItemID); err != nil {
		return model.NewApiError(model.ApiErrorNotFound, err.Error(), nil)
	}

	// If User Available
	if _, err := cc.DB().GetUserNameByID(r.UserID); err != nil {
		return model.NewApiError(model.ApiErrorNotFound, err.Error(), nil)
	}

	r.BidTime = time.Now().Unix()
	if err := cc.DB().SaveUserBidOnItem(r); err != nil {
		return model.NewApiError(model.ApiErrorNotAcceptable, err.Error(), nil)
	}

	return service.ApiResponse(c, http.StatusCreated, nil)
}

func getAllItemsByUser(c echo.Context) error {
	cc := c.(*ServerContext)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id < 1 {
		return model.NewApiError(model.ApiErrorBadRequest, "Invalid user id", nil)
	}

	// If User Available
	if _, err := cc.DB().GetUserNameByID(uint32(id)); err != nil {
		return model.NewApiError(model.ApiErrorNotFound, err.Error(), nil)
	}

	items, err := cc.DB().GetAllItemsByUserID(uint32(id))
	if err != nil {
		return model.NewApiError(model.ApiErrorNotFound, err.Error(), nil)
	}

	return service.ApiResponse(c, http.StatusOK, items)
}
