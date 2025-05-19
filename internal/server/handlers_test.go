package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestServer_AddTask(t *testing.T) {
	srv := NewServer()

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/api/tasks", srv.AddTask)
	body := []byte(`{"tool": "upgrader", "task": "healthchecks", "status": "completed", "duration": 120}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	tests := []struct {
		name             string
		body             []byte
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Shoud_Succeed",
			body:             []byte(`{"tool": "upgrader", "task": "healthchecks", "status": "completed", "duration": 120}`),
			expectedStatus:   http.StatusCreated,
			expectedResponse: `{"message": "Created"}`,
		},
		{
			name:             "Shoud_Fail_Invalid_Status",
			body:             []byte(`{"tool": "upgrader", "task": "healthchecks", "status": "wowzers", "duration": 1}`),
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error": "Invalid JSON"}`,
		},
		{
			name:             "Shoud_Fail_Invalid_Fields",
			body:             []byte(`{"toolzers": "upgrader", "taskzers": "healthchecks", "statuszers": "completed", "durationzers": 100500}`),
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error": "Invalid JSON"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedResponse, w.Body.String())
		})
	}
}
