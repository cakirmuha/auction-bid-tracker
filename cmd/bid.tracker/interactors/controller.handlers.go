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
		_ = service.EnsureContentTypeFunc(echo.MIMEApplicationJSON)
	)

	{
		fs := http.FileServer(assets.Assets)
		api.GET("/assets/*", echo.WrapHandler(http.StripPrefix(handlerPrefix+"/assets/", fs)), service.NoCacheMiddleware).Name = "GetStaticAssets"
	}
	endSetupFunc()
}
