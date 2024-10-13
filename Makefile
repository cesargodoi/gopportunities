.PHONY: default run build test docs clean

# Variables
APP_NAME=gopportunities
GOCMD=/usr/lib/go-1.22/bin/go

# Tasks
default: dcrun

run:
	$(GOCMD) run ./main.go
dcrun:
	swag init
	$(GOCMD) run ./main.go
build:
	$(GOCMD) build -o $(APP_NAME) main.go
test:
	$(GOCMD) test ./ ...
docs:
	swag init
clean:
	rm -f $(APP_NAME)
	rm -rf ./docs