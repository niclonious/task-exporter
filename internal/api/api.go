//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ../../openapi.yaml
package api

import "slices"

func ValidateTask(task Task) bool {
	return slices.Contains([]TaskStatus{Completed, Failed, Succeeded}, task.Status)
}
