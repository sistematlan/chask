package system

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

// Metrics contains server resource usage information.
type Metrics struct {
	Hostname    string  `json:"hostname"`
	OS          string  `json:"os"`
	Arch        string  `json:"arch"`
	CPUs        int     `json:"cpus"`
	DiskTotal   uint64  `json:"disk_total_bytes"`
	DiskUsed    uint64  `json:"disk_used_bytes"`
	DiskPercent float64 `json:"disk_percent"`
	Uptime      string  `json:"uptime"`
}

// GetMetrics collects current system resource metrics.
func GetMetrics() (*Metrics, error) {
	hostname, _ := os.Hostname()

	// Disk usage for /
	var stat syscall.Statfs_t
	if err := syscall.Statfs("/", &stat); err != nil {
		return nil, err
	}
	diskTotal := stat.Blocks * uint64(stat.Bsize)
	diskFree := stat.Bavail * uint64(stat.Bsize)
	diskUsed := diskTotal - diskFree
	diskPercent := float64(diskUsed) / float64(diskTotal) * 100

	// Uptime
	uptime := "unknown"
	if out, err := exec.Command("uptime", "-p").Output(); err == nil {
		uptime = strings.TrimSpace(string(out))
	}

	return &Metrics{
		Hostname:    hostname,
		OS:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		CPUs:        runtime.NumCPU(),
		DiskTotal:   diskTotal,
		DiskUsed:    diskUsed,
		DiskPercent: diskPercent,
		Uptime:      uptime,
	}, nil
}
