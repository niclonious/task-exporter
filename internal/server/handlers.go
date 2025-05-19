package server

import (
	"net/http"
	"slices"
	"task-exporter/internal/prom"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricName = "task_duration"
)

type Server struct {
	prometheus prom.PrometheusServer
}

// NewServer initialises a Server structure that implements api.ServerInterface
func NewServer() *Server {
	server := &Server{}
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
	return server
}

// AddTask - handler for (POST /api/tasks)
func (a *Server) AddTask(c *gin.Context) {
	var task AddTaskJSONRequestBody
	if err := c.ShouldBindJSON(&task); err != nil || !validateTask(task) {
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

// GetPrometheusMetrics - handler for (GET /metrics)
func (a *Server) GetPrometheusMetrics(c *gin.Context) {
	a.prometheus.Handler(c)
}

// ValidateTask validates incoming task data
func validateTask(task Task) bool {
	return slices.Contains([]TaskStatus{Completed, Failed, Succeeded}, task.Status) && task.Duration > 0
}
