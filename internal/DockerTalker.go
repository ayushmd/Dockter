package internal

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ayush18023/Load_balancer_Fyp/internal/auth"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	registery "github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/builder/remotecontext/urlutil"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/go-connections/nat"
	"github.com/go-git/go-git/v5"
)

type Dockter struct {
	cli *client.Client
}

type Reader struct {
	in io.ReadCloser // Stream to read from
}

// type DockerFile struct {
// 	RuntimeEnv string
// 	BuildCmd   string
// 	StartCmd   string
// 	EnvVars    map[string]string
// }

func NewReader(in io.ReadCloser) *Reader {
	return &Reader{
		in: in,
	}
}

func (d *Dockter) Init() {
	var err error
	d.cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
}

func NewDockter() *Dockter {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	return &Dockter{
		cli: cli,
	}
}

func (d *Dockter) Close() {
	d.cli.Close()
}

func (d *Dockter) BuildNewImage(imageName string, imagePath string) {
	ctx := context.Background()
	dockerBuildContext, err := archive.TarWithOptions(imagePath, &archive.TarOptions{})
	if err != nil {
		panic(err)
	}
	defer dockerBuildContext.Close()

	buildOptions := types.ImageBuildOptions{
		Tags:        []string{imageName + ":latest"},
		Remove:      true,
		ForceRemove: true,
		NoCache:     true,
		Dockerfile:  "Dockerfile",
		Context:     dockerBuildContext,
	}
	imageBuildResponse, err := d.cli.ImageBuild(ctx, dockerBuildContext, buildOptions)
	if err != nil {
		panic(err)
	}
	defer imageBuildResponse.Body.Close()

	// Print the build output
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		panic(err)
	}
}

func (d *Dockter) BuildImage(options BuildOptions) {
	var (
		err      error
		buildCtx io.ReadCloser
		// dockerfileCtx io.ReadCloser
		// contextDir    string
		// tempDir       string
		relDockerFile string
		// progBuff      io.Writer
		// buildBuff     io.Writer
		// remote        string
	)
	if urlutil.IsGitURL(options.Context) {
		_, err = git.PlainClone(options.Label, false, &git.CloneOptions{
			URL:      options.Context,
			Progress: os.Stdout,
		})
		if err != nil {
			if err != git.ErrRepositoryAlreadyExists {
				log.Fatal(err)
			}
		}
		defer os.RemoveAll(options.Label)
	}
	if !options.UseDefaultDockerFile {
		relDockerFile = filepath.Join(options.Label, "Dockerfile")
		err = os.WriteFile(
			relDockerFile,
			[]byte(options.DockerfileContent),
			0644,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	// buildCtx, err = archive.TarWithOptions(options.Label, &archive.TarOptions{})
	buildCtx, err = os.Open("./ayush")
	// relDockerFile, err = filepath.Abs(relDockerFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// relDockerFile = filepath.ToSlash(relDockerFile)
	if err != nil {
		log.Fatal(err)
	}
	var p []byte
	buildCtx.Read(p)
	fmt.Println(string(p))
	defer buildCtx.Close()
	buildOptions := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{options.Label + ":latest"},
	}
	buildResponse, err := d.cli.ImageBuild(
		context.Background(),
		buildCtx,
		buildOptions,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer buildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Dockter) CreateImage(imageName string, DockerfileCtx string, reader io.Reader) {
	buildOptions := types.ImageBuildOptions{
		Dockerfile: DockerfileCtx,
		Tags:       []string{imageName + ":latest"},
	}

	buildResponse, err := d.cli.ImageBuild(
		context.Background(),
		reader,
		buildOptions,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer buildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Dockter) ListContainers() []types.Container {
	containers, err := d.cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}
	return containers
}

func (d *Dockter) InspectContainer(containerID string) (types.ContainerJSON, error) {
	return d.cli.ContainerInspect(context.Background(), containerID)
}

func (d *Dockter) StatsContainer(containerID string) (types.ContainerStats, error) {
	return d.cli.ContainerStats(context.Background(), containerID, false)
}

func (d *Dockter) ContainerMetrics(containerID string) (ContainerMetricInfo, error) {
	containerInfo, err := d.cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		return ContainerMetricInfo{}, err
	}
	defer containerInfo.Body.Close()
	data, err := io.ReadAll(containerInfo.Body)
	if err != nil {
		return ContainerMetricInfo{}, err
	}
	var containerMetrics ContainerMetricInfo
	if err = json.Unmarshal(data, &containerMetrics); err != nil {
		return ContainerMetricInfo{}, err
	}
	return containerMetrics, nil
}

func GetBasedMetrics(containerID string) (*ContainerBasedMetric, error) {
	d := NewDockter()
	defer d.Close()
	metrics, err := d.ContainerMetrics(containerID)
	if err != nil {
		return nil, err
	}
	f := filters.NewArgs(filters.KeyValuePair{Key: "id", Value: containerID})
	containers, err := d.cli.ContainerList(context.Background(), container.ListOptions{
		Size:    true,
		Filters: f,
	})
	if err != nil {
		return nil, err
	}
	//Below Calculation from Docker api reference
	cpuDelta := metrics.Cpu_stats.Cpu_usage.Total_usage - metrics.Precpu_stats.Cpu_usage.Total_usage
	systemCpuDelta := metrics.Cpu_stats.System_cpu_usage - metrics.Precpu_stats.System_cpu_usage
	numberCpus := metrics.Cpu_stats.Online_cpus
	var diskuse int64
	if len(containers) > 0 {
		diskuse = containers[0].SizeRootFs
	}
	basedMetric := &ContainerBasedMetric{
		MemUsage:   metrics.Memory_stats.Usage,
		CpuPercent: (cpuDelta / systemCpuDelta) * int64(numberCpus) * 100,
		DiskUsage:  diskuse,
	}
	return basedMetric, nil
}

func (d *Dockter) RunContainer(imageName string, containerName string, ports []string) (string, error) {
	expports, portBindings, err := nat.ParsePortSpecs(ports)
	if err != nil {
		log.Fatal("Error in ports")
	}

	cont, err := d.cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image:        imageName,
			ExposedPorts: expports,
			//may require Cmd:
		},
		&container.HostConfig{
			PortBindings: portBindings,
		}, nil, nil,
		containerName,
	)

	if err != nil {
		return "", err
	}

	if err := d.cli.ContainerStart(context.Background(), cont.ID, container.StartOptions{}); err != nil {
		return "", err
	}
	// hostConfig := &container.HostConfig{
	// 	PortBindings: nat.PortMap{
	// 		"8080/tcp": []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: "8080"}},
	// 	},
	// }

	return cont.ID, nil
}

func (d *Dockter) TrashContainer(containerID string) error {
	if err := d.cli.ContainerKill(context.Background(), containerID, "SIGKILL"); err != nil {
		return err
	}
	removeOptions := container.RemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}
	if err := d.cli.ContainerRemove(context.Background(), containerID, removeOptions); err != nil {
		return err
	}
	return nil
}

func (d *Dockter) ClearImages(imageName string) error {
	_, err := d.cli.ImageRemove(context.Background(), imageName, types.ImageRemoveOptions{})
	if err != nil {
		return err
	}
	// _, err = d.cli.ImageRemove(context.Background(), imageName+":latest", types.ImageRemoveOptions{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return nil
}

func (d *Dockter) ExecuteCommand(containerID string, cmd []string) []string {
	execConfig := types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
		Tty:          false,
	}

	execID, err := d.cli.ContainerExecCreate(context.Background(), containerID, execConfig)
	if err != nil {
		panic(err)
	}

	response, err := d.cli.ContainerExecAttach(context.Background(), execID.ID, types.ExecStartCheck{})
	if err != nil {
		panic(err)
	}
	defer response.Close()

	scanner := bufio.NewScanner(response.Reader)
	var outputLines []string
	for scanner.Scan() {
		outputLines = append(outputLines, scanner.Text())
	}
	return outputLines
}

func encodeAuthToBase64(authConfig registery.AuthConfig) (string, error) {
	authJSON, err := json.Marshal(authConfig)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(authJSON), nil
}

func (d *Dockter) PushToRegistry(imageName, reponame string) string {
	authConfig := registery.AuthConfig{
		Username:      auth.GetKey("DOCKER_USER"),
		Password:      auth.GetKey("DOCKER_PAT"),
		ServerAddress: "https://index.docker.io/v1/",
	}

	encodedAuth, err := encodeAuthToBase64(authConfig)
	if err != nil {
		panic(err)
	}

	tag := reponame + ":" + imageName

	// Tag the image
	if err := d.cli.ImageTag(context.Background(), imageName, tag); err != nil {
		log.Fatal(err)
	}

	// Push the image to the registry
	pushResponse, err := d.cli.ImagePush(context.Background(), tag, types.ImagePushOptions{
		RegistryAuth: encodedAuth,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer pushResponse.Close()

	// Log push output
	_, err = io.Copy(os.Stdout, pushResponse)
	if err != nil {
		log.Fatal(err)
	}
	return tag
}

func (d *Dockter) PullFromRegistery(repoimageName string) error {
	out, err := d.cli.ImagePull(context.Background(), repoimageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
	return nil
}

// func (d *Dockter) Inspect() {
// 	containerjson, err := d.cli.ContainerInspect(context.Background(), "")
// }

func FindFirstPort(outputLines []string) string {
	var runningPorts []string
	for _, line := range outputLines {
		if strings.Contains(line, "LISTEN") {
			fields := strings.Fields(line)
			if len(fields) >= 4 {
				splits := strings.Split(fields[3], ":")
				runningPorts = append(runningPorts, splits[len(splits)-1])
			}
		}
	}
	if len(runningPorts) > 0 {
		return runningPorts[0]
	}

	return ""
}
