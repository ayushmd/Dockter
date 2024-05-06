package master

import (
	"net/http"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/master/templ"
)

func P1(w http.ResponseWriter, r *http.Request) {
	// Master_.Join("127.0.0.5:5000", "WORKER", 20.8, 40.2, 69.9)
	Master_.Join("127.0.0.5:6000", "WORKER", internal.ContainerBasedMetric{
		CpuPercent:       2,
		MemUsage:         231129088,
		TotalMem:         591691776,
		MemUsedPercent:   23.213184,
		DiskUsage:        4632453120,
		TotalDisk:        8510222336,
		DiskUsagePercent: 54.433983,
	})
	templ.RenderAndExecute("index.html", w, Master_.ServerPool)
}

func LoadUI(router *Router) {
	router.Get("/", P1)
}
