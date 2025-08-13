APP_NAME=server
BIN_DIR=bin

#
# Так как сервер на чистом GO установка зависимостей не требуется
#

# Сборка с докера
docker-build:
	@echo "Docker build"
	@docker build -t task-api:latest .

# Запуск с докера
docker-run:
	@echo "Docker run"
	@docker run --rm -p 8080:8080 task-api:latest

# Запуск тестов
test:
	@echo "Running tests..."
	@go test ./... -v

# Сборка
build: clean
	@echo "Old files cleaned"
	@echo "Building $(APP_NAME)..."
	@go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/main.go
	@echo "Binary: $(BIN_DIR)/$(APP_NAME)"

# Очистка бинарников и тд
clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)
	@rm -f coverage.out

# Запуск
run: build
	@echo "Running $(APP_NAME)..."
	@./$(BIN_DIR)/$(APP_NAME)

.PHONY: build test clean rebuild run
