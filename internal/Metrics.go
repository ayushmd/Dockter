package internal

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func HealthMetrics() (float64, float64, float64, error) {
	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0.0, 0.0, 0.0, err
	}
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0.0, 0.0, 0.0, err
	}
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return 0.0, 0.0, 0.0, err
	}
	return cpuUsage[0] * 100, memInfo.UsedPercent, diskUsage.UsedPercent, nil
}

func HealthMetricsBased() (*ContainerBasedMetric, error) {
	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}
	return &ContainerBasedMetric{
		CpuPercent:       int64(cpuUsage[0]),
		MemUsage:         int32(memInfo.Used),
		TotalMem:         int32(memInfo.Available),
		MemUsedPercent:   memInfo.UsedPercent,
		DiskUsage:        int64(diskUsage.Used),
		TotalDisk:        int64(diskUsage.Total),
		DiskUsagePercent: diskUsage.UsedPercent,
	}, nil
}
