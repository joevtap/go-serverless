FUNCTION_DIR=$(notdir $(CURDIR))
FUNCTION_NAME=$(FUNCTION_DIR)_handler
AWS_REGION={{.awsRegion}}

.PHONY: build clean test deploy

build:
	GOOS=linux GOARCH=amd64 go build -o $(FUNCTION_NAME) .

clean:
	rm -f $(FUNCTION_NAME)

test:
	go test -v

deploy:
	@echo "Deploying function $(FUNCTION_NAME) in directory $(FUNCTION_DIR)..."
	cd ../.. && serverless deploy function -f $(FUNCTION_NAME) --region $(AWS_REGION)
	@echo "Deployment complete."