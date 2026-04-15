package api

import (
	"encoding/json"
	"os"
	"runtime"
	"time"

	"github.com/sistematlan/chask/daemon/internal/system"
)

func handleHealth(_ json.RawMessage) (interface{}, error) {
	hostname, _ := os.Hostname()
	return map[string]interface{}{
		"status":   "healthy",
		"hostname": hostname,
		"version":  "0.1.0-dev",
		"go":       runtime.Version(),
		"uptime":   time.Now().Format(time.RFC3339),
	}, nil
}

func handleMetrics(_ json.RawMessage) (interface{}, error) {
	metrics, err := system.GetMetrics()
	if err != nil {
		return nil, err
	}
	return metrics, nil
}
