ifeq ($(IMAGE), )
	IMAGE="123shang60"
endif

ifeq ($(VERSION), )
	VERSION="v0.0.0"
endif

run-server:
	@go run cmd/server.go

run-agent:
	@go run cmd/agent.go

gen-proto:
	@protoc --go_out=./pkg/register --go_opt=paths=source_relative \
		   --go-grpc_out=./pkg/register --go-grpc_opt=paths=source_relative \
    	   ./proto/register.proto

build-server:
	@go build -o server ./cmd/server.go

build-agent:
	@go build -o agent ./cmd/agent.go

docker-server:
	docker build -f Dockerfile-server -t $(IMAGE)/image-load-server:$(VERSION) .

docker-agent:
	docker build -f Dockerfile-agent -t $(IMAGE)/image-load-agent:$(VERSION) .

