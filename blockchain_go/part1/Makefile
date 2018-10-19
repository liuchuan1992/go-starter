#定义变量
BINARY := blockchain

#执行命令
all: build run

build:
	@echo "==> Go build"
	@go build -o $(BINARY) -v

run:
	@echo "==> Running"
	@./$(BINARY)


.PHONY: build run