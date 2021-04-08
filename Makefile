.PHONY: all

PROJ_PATH := ${CURDIR}
DOCKER_PATH := ${PROJ_PATH}/docker

ENR_PARSER=enrollments-parser

BASIC_IMAGE=default

build-app:
	go build -o .bin/${ENR_PARSER} cmd/${ENR_PARSER}/main.go

build-docker:
	docker build -t ${BASIC_IMAGE} -f ${DOCKER_PATH}/builder.Dockerfile .
	docker build -t enrollments-parser-service -f ${DOCKER_PATH}/enrollments_parser.Dockerfile .

app-up: build-docker
	docker-compose up --build
