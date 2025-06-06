package stats

import (
	"time"

	"github.com/aandrku/portfolio-v2/pkg/model"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func Get() (model.Stats, error) {
	var stats model.Stats

	m, err := mem.VirtualMemory()
	if err != nil {
		return stats, err
	}
	stats.MemoryUsed = m.Used
	stats.MemoryTotal = m.Total

	c, err := cpu.Percent(time.Second, false)
	if err != nil {
		return stats, err
	}
	stats.CPUUsage = c

	hi, err := host.Info()
	if err != nil {
		return stats, err
	}
	stats.Uptime = hi.Uptime

	du, err := disk.Usage("/")
	if err != nil {
		return stats, err
	}
	stats.DiskUsed = du.Used
	stats.DiskTotal = du.Total

	return stats, nil
}
