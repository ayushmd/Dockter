package master

import (
	"net/http"

	"github.com/ayush18023/Load_balancer_Fyp/master/templ"
)

func P1(w http.ResponseWriter, r *http.Request) {
	// Master_.Join("127.0.0.5:5000", "WORKER", 20.8, 40.2, 69.9)
	// Master_.Join("127.0.0.5:6000", "WORKER", 80.8, 21.2, 99.9)
	templ.RenderAndExecute("index.html", w, Master_.ServerPool)
}

func LoadUI(router *Router) {
	router.Get("/", P1)
}
