package app

import (
	"net/http"
	"task-exporter/internal/api"
	"task-exporter/internal/prom"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type Server struct {
	prometheus prom.PrometheusServer
}

// Add a new Task
// (POST /api/tasks)
func (a *Server) AddTask(c *gin.Context) {
	var task api.AddTaskJSONRequestBody
	if err := c.ShouldBindJSON(&task); err != nil || !api.ValidateTask(task) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	labels := prometheus.Labels{
		"tool":   task.Tool,
		"task":   task.Task,
		"status": string(task.Status),
	}
	a.prometheus.GaugeVecs[metricName].With(labels).Set(float64(task.Duration))
	c.JSON(http.StatusCreated, gin.H{"message": "Created"})
}

// Get prometheus metrics
// (GET /metrics)
func (a *Server) GetPrometheusMetrics(c *gin.Context) {
	a.prometheus.Handler(c)
}
