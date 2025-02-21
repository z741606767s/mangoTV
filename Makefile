# 项目名称
APP_NAME := mangoTV

# 源文件目录
SRC_DIR := ./app

# 输出目录
BUILD_DIR := ./build

# 默认目标（编译）
.PHONY: all
all: build

# 编译目标
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

# 清理目标
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# 运行目标
.PHONY: run
run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

# 跨平台编译 - Linux
.PHONY: build-linux
build-linux:
	@echo "Building $(APP_NAME) for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux $(SRC_DIR)

# 跨平台编译 - Linux 沙箱
.PHONY: build-linux-sandbox
build-linux-sandbox:
	@echo "Building $(APP_NAME) for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux-sandbox $(SRC_DIR)

# 跨平台编译 - Windows
.PHONY: build-windows
build-windows:
	@echo "Building $(APP_NAME) for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME).exe $(SRC_DIR)

# 跨平台编译 - Mac M1 (ARM64)
.PHONY: build-mac-m1
build-mac-m1:
	@echo "Building $(APP_NAME) for Mac M1..."
	@GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)-mac-m1 $(SRC_DIR)

# 跨平台编译 - Mac Intel (AMD64)
.PHONY: build-mac-intel
build-mac-intel:
	@echo "Building $(APP_NAME) for Mac Intel..."
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-mac-intel $(SRC_DIR)

# 帮助目标
.PHONY: help
help:
	@echo "Makefile 使用说明:"
	@echo "  make          编译项目"
	@echo "  make build    编译项目"
	@echo "  make clean    清理生成的文件"
	@echo "  make run      编译并运行项目"
	@echo "  make build-linux   为 Linux 编译"
	@echo "  make build-windows 为 Windows 编译"
