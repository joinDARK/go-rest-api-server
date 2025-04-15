.PHONY: test
	go test -v -race -timeout 30s ./...

.PHONY: build

# Целевая директория для бинарника
BIN_DIR := ./bin
# Имя выходного бинарного файла
BINARY := $(BIN_DIR)/server

build: $(BINARY)

# Правило для сборки бинарника
$(BINARY): ./cmd/main.go
	@mkdir -p $(BIN_DIR)  # Создаем директорию bin, если она не существует
	go build -o $(BINARY) ./cmd/main.go

.DEFAULT_GOAL := build