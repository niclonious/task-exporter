package api

// Defines values for TaskStatus.
const (
	Completed TaskStatus = "completed"
	Failed    TaskStatus = "failed"
	Succeeded TaskStatus = "succeeded"
)

// Task defines model for Task.
type Task struct {
	Duration int        `json:"duration" binding:"required"`
	Status   TaskStatus `json:"status" binding:"required"`
	Task     string     `json:"task" binding:"required"`
	Tool     string     `json:"tool" binding:"required"`
}

// TaskStatus defines model for Task.Status.
type TaskStatus string

// AddTaskJSONRequestBody defines body for AddTask for application/json ContentType.
type AddTaskJSONRequestBody = Task
