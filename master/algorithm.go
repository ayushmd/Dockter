package master

import "github.com/ayush18023/Load_balancer_Fyp/internal"

func GetWeight(basedMetrics internal.ContainerBasedMetric) float64 {
	return 0.2*float64(basedMetrics.CpuPercent) + 0.0*basedMetrics.MemUsedPercent + 0.8*basedMetrics.DiskUsagePercent
}

func MasterPlanAlgo(ServerPool []*Backend, state string) *Backend {
	minWeight := float64(^uint(0) >> 1) // Initialize with a large positive value
	selectedi := 0
	for i, server := range ServerPool {
		var weight float64
		if server.IsAlive && server.State == state {
			initialWeight := GetWeight(server.Stats)
			if server.CurrentConnect > 0 {
				weight = initialWeight * float64(server.CurrentConnect)
			} else {
				weight = initialWeight
			}
			if weight < minWeight {
				minWeight = weight
				selectedi = i
			}
		}
	}
	return ServerPool[selectedi]
}

func MasterPlanAlgo2(ServerPool []*Backend, state string, taskStats internal.ContainerBasedMetric) *Backend {
	minWeight := float64(^uint(0) >> 1) // Initialize with a large positive value
	selectedi := 0
	for i, server := range ServerPool {
		var weight float64
		if server.IsAlive && server.State == state {
			afterMem := (server.Stats.MemUsage + taskStats.MemUsage) / server.Stats.TotalMem * 100
			afterDisk := server.Stats.DiskUsage + taskStats.DiskUsage
			if afterMem <= 80 && afterDisk < server.Stats.TotalDisk {
				initialWeight := GetWeight(server.Stats)
				if server.CurrentConnect > 0 {
					weight = initialWeight * float64(server.CurrentConnect)
				} else {
					weight = initialWeight
				}
				if weight < minWeight {
					minWeight = weight
					selectedi = i
				}
			}
		}
	}
	return ServerPool[selectedi]
}
