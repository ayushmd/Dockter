setup: debian_docker install_go protoc_install rpc_code

rpc_code:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	./rpc/workerrpc/worker.proto

	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	./rpc/echo/echo.proto
	
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	.\rpc\masterrpc\master.proto
	
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	.\rpc\builderrpc\builder.proto


protoc_install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	export PATH="$PATH:$(go env GOPATH)/bin"

debian_docker:
	sudo apt  install docker.io
	sudo groupadd docker
	sudo usermod -aG docker $USER
	newgrp docker

install_go:
	wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
	tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
	export PATH=$PATH:/usr/local/go/bin

docker_clean:
	docker kill $(docker ps -aq)
	docker rm -vf $(docker ps -aq)
	docker rmi -f $(docker images -aq)
