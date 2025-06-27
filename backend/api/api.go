package api

import (
	"fmt"
	"log"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/garyhopkins/adminex/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const ApiDomain = "localhost"
const Name = "Adminex"
const Version = "0.1"

type Api struct {
	humaApi huma.API
	echo    *echo.Echo
}

func New() *Api {
	// Create a new router & API.
	router := echo.New()
	group := router.Group("/api")

	baseUrl := fmt.Sprintf("http://%s:8888/api", ApiDomain)
	config := huma.DefaultConfig(Name, Version)
	config.Servers = []*huma.Server{
		{URL: baseUrl},
	}

	SetupEchoMiddleware(router)
	api := humaecho.NewWithGroup(router, group, config)

	a := &Api{
		humaApi: api,
		echo:    router,
	}

	SetupHandlers(api)

	return a
}

func SetupEchoMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
}

func SetupHandlers(a huma.API) {
	huma.Register(a, handlers.OpTestHandler, handlers.TestHandler)

}

func (a *Api) ListenAndServe(addr string) error {
	log.Printf("Listening on %s", addr)
	return a.echo.Start(addr)
}
