package app

import (
	"fmt"
	"log/slog"
	"task-exporter/internal/api"
	"task-exporter/internal/config"
	"task-exporter/internal/prom"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricName = "task_duration"
)

type App struct {
	router *gin.Engine
	config config.Config
}

func New() (app *App, err error) {
	app = &App{}
	server := &Server{}
	app.config, err = config.LoadConfig("./")
	if err != nil {
		slog.Error("Failed to read config", slog.String("error", err.Error()))
		return nil, err
	}
	server.prometheus = *prom.NewPrometheusServer()
	server.prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricName,
			Help: "A gauge of task execution durations in seconds.",
		},
		[]string{
			"tool",
			"task",
			"status",
		},
	)
	app.router = gin.New()
	app.router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	app.router.SetTrustedProxies(nil)

	api.RegisterHandlers(app.router, server)

	return app, nil
}

func (a *App) Run() {
	a.router.Run(":" + fmt.Sprintf("%d", a.config.Port))
}
