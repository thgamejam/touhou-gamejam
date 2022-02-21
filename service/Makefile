GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find "./app/*/internal" -name "*.proto")
API_PROTO_FILES=$(shell find "./api" -name "*.proto" -type f ! -name "error_reason.proto")
#API_PROTO_FILES=$(shell find "./api" -name "*.proto")
ERROR_PROTO_FILES=$(shell find "./api" -name "error_reason.proto")

PKG_CONF_PROTO_FILES=$(shell find "./pkg/conf" -name *.proto)

WIRE_FILES=$(shell find "./app/*/cmd/service" -name "wire.go")

# 判断Docker的环境
DOCKER_EXE="$(shell command -v docker.exe)"
DOCKER=""
ifeq ($(DOCKER_EXE), $(DOCKER))
	DOCKER="$(shell command -v docker)"
else
	DOCKER=$(DOCKER_EXE)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.6.1
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	bash ./script/dev_environment.sh

.PHONY: app
# 创建一个微服务
app:
	bash ./script/create_app.sh
	make all;

.PHONY: service
# 生成 Service 代码
service:
	bash ./script/generate_service.sh

.PHONY: wire
# 依赖注入
wire:
	@for file in $(WIRE_FILES) ; do \
		path=$$( dirname $$file ) ; \
		wire $$path ; \
	done

.PHONY: error
# generate errors code
error:
	protoc --proto_path=. \
               --proto_path=./third_party \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(ERROR_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	for file in $(INTERNAL_PROTO_FILES) ; do \
		protoc --proto_path=. \
                --proto_path=./third_party \
                --go_out=paths=source_relative:. \
                $$file ; \
	done
	protoc --proto_path=. \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:. \
	       $(PKG_CONF_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
 	       --validate_out=paths=source_relative,lang=go:. \
 	       --openapiv2_out . \
           --openapiv2_opt logtostderr=true \
           --openapiv2_opt json_names_for_fields=false \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# 代码生成
generate:
	go get github.com/google/wire/cmd/wire
	go install github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# 生成所有代码
all:
	make api;
	make error;
	make config;
	make wire;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
