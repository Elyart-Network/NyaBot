.PHONY: all build clean test check run docker help
BIN_FILE=nyabot
all: check build run
build:
	@go build -o "${BIN_FILE}"
clean:
	@go clean
test:
	@go test
check:
	@go fmt ./
	@go vet ./
run:
	./"${BIN_FILE}"
docker:
	@docker build -t whydesd/nyabot:test .
help:
	@echo "make 格式化go代码 并编译生成二进制文件"
	@echo "make build 编译go代码生成二进制文件"
	@echo "make clean 清理中间目标文件"
	@echo "make test 执行测试case"
	@echo "make check 格式化go代码"
	@echo "make run 直接运行程序"
	@echo "make docker 构建docker镜像"