package builder

import (
	"log"
	"strings"
	"testing"
)

func TestBuildRawWithoutEnv(t *testing.T) {
	// b := Builder{}
	// Name := "ayush"
	// GitLink := "https://github.com/johnpapa/node-hello.git"
	// Branch := "master"
	// BuildCmd := "npm i"
	// RunCmd := "npm start"
	// RuntimeEnv := "Node"
	// _, port, err := b.BuildRaw(
	// 	Name,
	// 	GitLink,
	// 	Branch,
	// 	BuildCmd,
	// 	RunCmd,
	// 	RuntimeEnv,
	// 	nil,
	// )
	// if port != "3000" || err != nil {
	// 	t.Fatalf(`BuildRaw = %q, %v, want 3000, nil`, port, err)
	// }
}

func TestGoBuildImage(t *testing.T) {
	b := Builder{}
	Name := "Amd"
	BuildCmd := "CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping"
	StartCmd := "/docker-gs-ping"
	RuntimeEnv := "Go"
	EnvVars := map[string]string{
		"PORT": "8000",
	}
	outputImage := b.BuildDockerByLang(Name, BuildCmd, StartCmd, RuntimeEnv, EnvVars)
	expectedImage := `
FROM golang:1.22
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
ENV PORT=8000

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
CMD ["/docker-gs-ping"]
`
	log.Println(strings.TrimSpace(outputImage), strings.TrimSpace(expectedImage))
	if strings.TrimSpace(outputImage) != strings.TrimSpace(expectedImage) {
		t.Fatalf(`Go image not matched`)
	}
}
