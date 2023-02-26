# Variables
APP_NAME=gosls
DIST_DIR=dist

# Targets
.PHONY: build test release clean

build:
	@echo "Building $(APP_NAME)"
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux .
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME)-windows.exe .
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)-darwin .

test:
	@echo "Running tests"
	go test -v ./...

release: clean test build
	@echo "Releasing $(APP_NAME)"
	mkdir -p $(DIST_DIR)
	mv bin/$(APP_NAME)-windows.exe bin/$(APP_NAME).exe
	7z a bin/$(APP_NAME).zip bin/$(APP_NAME).exe tpl/*
	mv bin/$(APP_NAME).exe bin/$(APP_NAME)-windows.exe
	mv bin/$(APP_NAME).zip $(DIST_DIR)
	@echo "Done"

clean:
	@echo "Cleaning up"
	go clean
	rm -f $(APP_NAME)
	rm -rf $(DIST_DIR)
