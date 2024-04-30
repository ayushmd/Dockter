package internal

type BuildOptions struct {
	Label                string
	Context              string
	DockerfileName       string
	UseDefaultDockerFile bool
	GitBranch            string
	DockerfileContent    string
}

type ContainerMetricInfo struct {
	Name        string `json:"name"`
	ContainerId string `json:"id"`
	Cpu_stats   struct {
		Cpu_usage struct {
			Total_usage int64 `json:"total_usage"`
		} `json:"cpu_usage"`
		System_cpu_usage int64 `json:"system_cpu_usage"`
		Online_cpus      int32 `json:"online_cpus"`
	} `json:"cpu_stats"`
	Precpu_stats struct {
		Cpu_usage struct {
			Total_usage int64 `json:"total_usage"`
		} `json:"cpu_usage"`
		System_cpu_usage int64 `json:"system_cpu_usage"`
	} `json:"precpu_stats"`
	Memory_stats struct {
		Usage     int32 `json:"usage"`
		Max_usage int32 `json:"max_usage"`
		Stats     struct {
			Cache int32 `json:"cache"`
		} `json:"stats"`
		// Limit     int32 `json:"limit"`
	} `json:"memory_stats"`
}

type ContainerBasedMetric struct {
	MemUsage         int32   `json:"memUsage"`
	TotalMem         int32   `json:"totalMem,omitempty"`
	MemUsedPercent   float64 `json:"memUsedPercent,omitempty"`
	CpuPercent       int64   `json:"cpuPercent"`
	DiskUsage        int64   `json:"diskUsage"`
	TotalDisk        int64   `json:"totalDisk,omitempty"`
	DiskUsagePercent float64 `json:"diskUsagePercent,omitempty"`
}
