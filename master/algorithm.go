package master

func GetWeight(cpuUsage, memUsage, diskUsage float64) float64 {
	return 0.5*cpuUsage + 0.3*memUsage + 0.2*diskUsage
}

func MasterPlanAlgo(ServerPool []*Backend, state string) *Backend {
	minWeight := float64(^uint(0) >> 1) // Initialize with a large positive value
	selectedi := 0
	for i, server := range ServerPool {
		var weight float64
		if server.IsAlive && server.State == state {
			initialWeight := GetWeight(server.CpuUsage, server.MemUsage, server.DiskUsage)
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
