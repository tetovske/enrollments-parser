.PHONY:

BUILD_NAME=enrollments-parser
APP_NAME=enrollments-parser-app

app-build-local:
	go build -o .bin/$(BUILD_NAME) cmd/$(BUILD_NAME)/main.go

app-up-local:
	.bin/$(BUILD_NAME) "./configs/config.yml"

app-run: app-build-local app-up-local

app-build:
	docker build -t $(BUILD_NAME) .

app-up:
	docker run --name=$(APP_NAME) $(BUILD_NAME)