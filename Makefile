.PHONY: all

PROJ_PATH := ${CURDIR}
DOCKER_PATH := ${PROJ_PATH}/docker

ENR_PARSER=enrollments-parser

BASIC_IMAGE=default
IMAGE_POSTFIX=-image

build-app:
	go build -o .bin/${ENR_PARSER} cmd/${ENR_PARSER}/main.go

build-docker:
	docker build -t ${BASIC_IMAGE} -f ${DOCKER_PATH}/builder.Dockerfile .
	docker build -t ${ENR_PARSER}${IMAGE_POSTFIX} -f ${DOCKER_PATH}/enrollments_parser.Dockerfile .

app-setup-and-up: build-docker app-up

app-up:
	docker-compose up

all: app-setup-and-up

app-enrollments-parser-bash:
	docker-compose run --rm --no-deps --name temp_enrl_parser_service ${ENR_PARSER} bash
