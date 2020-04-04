package interactors

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/cakirmuha/auction-bid-tracker/cmd/bid.tracker/assets"
	"github.com/cakirmuha/auction-bid-tracker/service"
)

func (s *ServerContext) SetupHandlers() {
	const handlerPrefix = "/api/v1"

	api, endSetupFunc := s.Context.SetupHandlers(handlerPrefix, nil)

	var (
		jsonBody = service.EnsureContentTypeFunc(echo.MIMEApplicationJSON)
	)

	{
		fs := http.FileServer(assets.Assets)
		api.GET("/assets/*", echo.WrapHandler(http.StripPrefix(handlerPrefix+"/assets/", fs)), service.NoCacheMiddleware).Name = "GetStaticAssets"
	}
	{
		g := api.Group("/user")
		g.POST("/bid", createUserBid, jsonBody).Name = "createUserBid"
		g.GET("/:id/items", getAllItemsByUser).Name = "getAllItemsByUser"
	}
	{
		g := api.Group("/item")
		g.GET("/:id/bids", getAllBidsByItem).Name = "getAllBidsByItem"
		g.GET("/:id/bids/winning", getCurrentWinningBidByItem).Name = "getCurrentWinningBidByItem"
	}
	endSetupFunc()
}
