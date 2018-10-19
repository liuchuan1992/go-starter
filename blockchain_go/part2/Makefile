#定义变量
BINARY := blockchain_part2

#执行命令
all: build run

build:
	@echo "==> Go build"
	@go build -o $(BINARY) -v

run:
	@echo "==> Running"
	@./$(BINARY)


.PHONY: build run