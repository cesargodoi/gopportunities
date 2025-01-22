.PHONY: default run build test docs clean

# Variables
APP_NAME=gopportunities
GO=/usr/local/go/bin/go

# Tasks
default: dcrun

run:
	@echo "Running..."
	@$(GO) run .

tidy:
	@echo "Tidying..."
	@$(GO) mod tidy

dcrun:
	@echo "Swaggering and Running..."
	@swag init
	@$(GO) run .

build:
	@echo "Building..."
	@$(GO) build -o $(APP_NAME) main.go

test:
	@echo "Testing..."
	@$(GO) test ./ ...

docs:
	@echo "Swaggering..."
	@swag init

clean:
	@echo "Cleaning..."
	@rm -f $(APP_NAME)
	@rm -rf ./docs
	