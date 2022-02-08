PRINT_WORK_DIRECTORY=$(shell pwd)
PROTOC_PATH=$(shell which protoc)
NODE_MODULES_PATH=$(shell npm root -g)
PROTOC_GEN_TS_PATH=$(NODE_MODULES_PATH)/ts-protoc-gen/bin/protoc-gen-ts
VERSION=$(shell git describe --tags --always)

PROTO_API_PATH=$(PRINT_WORK_DIRECTORY)/api
SERVICE_PATH=$(PRINT_WORK_DIRECTORY)/service
WEB_PATH=$(PRINT_WORK_DIRECTORY)/web

JAVA_OUT_DIR=$(SERVICE_PATH)/$$project_name/src/main/java
TS_OUT_DIR=$(WEB_PATH)/src/api/$$project_name

API_PROTO_FILES=$(shell find $(PROTO_API_PATH) -name *.proto)

.PHONY: test
# 测试
test:
	@echo $(PROTOC_PATH)
	@echo $(NODE_MODULES_PATH)
	@echo $(PROTOC_GEN_TS_PATH)
	@echo $(VERSION)
	@echo ''
	@echo 'out dir:'
	@echo $(JAVA_OUT_DIR)
	@echo $(TS_OUT_DIR)
	@echo ''
	@echo 'API proto:'
	@echo $(PROTO_API_PATH)
	for file in $(API_PROTO_FILES) ; do \
		project_name=$$( basename $$( dirname $$file ) ) ; \
		echo 'project name: ' $$project_name ; \
		echo 'file name: ' $$file ; \
		java_out=$(JAVA_OUT_DIR) ; \
		echo 'java out: ' $$java_out ; \
		ts_out=$(TS_OUT_DIR) ; \
		echo 'ts out: ' $$ts_out ; \
	done

.PHONY: init
# 初始化环境
init:
	sudo npm install -g ts-protoc-gen



.PHONY: api
# 构建 proto api 文件
api:
	@echo '开始构建 proto api 文件'
	for file in $(API_PROTO_FILES) ; do \
		proto_path=$$( dirname $$file ) ; \
		project_name=$$( basename $$proto_path ) ; \
		java_out=$(JAVA_OUT_DIR) ; \
		ts_out=$(TS_OUT_DIR) ; \
		for dir in $$java_out $$ts_out ; do \
			if [ ! -d "$$dir" ] ; then \
				mkdir $$dir ; \
			fi ; \
		done ; \
		$(PROTOC_PATH) \
			--proto_path="$$proto_path" \
			--plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
			--java_out="$$java_out" \
			--js_out="import_style=commonjs,binary:$$ts_out" \
			--ts_out="$$ts_out" \
			$$file ; \
	done
	@echo 'api 文件构建完成'

.PHONY: build
# 构建
build:


.PHONY: all
# 生成所有
all:
	make api;

# 帮助信息
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