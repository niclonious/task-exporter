//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ../../openapi.yaml
package api

import (
	"fmt"
	"log/slog"
	"task-exporter/internal/config"
	"task-exporter/internal/server"

	"github.com/gin-gonic/gin"
)

type API struct {
	router *gin.Engine
	config config.Config
}

func New() (api *API, err error) {
	api = &API{}
	server := server.NewServer()
	api.config, err = config.LoadConfig("./") //TODO make it configurable via app arguments
	if err != nil {
		slog.Error("Failed to read config", slog.String("error", err.Error()))
		return nil, err
	}

	if api.config.Env == config.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}
	api.router = gin.New()
	api.router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	api.router.SetTrustedProxies(nil)

	RegisterHandlers(api.router, server)

	return api, nil
}

func (a *API) Run() {
	a.router.Run(":" + fmt.Sprintf("%d", a.config.Port))
}
