.PHONY: all
all: build fmt vet lint test

APP=kafka_hystrix_client
GLIDE_NOVENDOR=$(shell glide novendor)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(shell glide novendor | grep -v "featuretests")

APP_EXECUTABLE="./out/$(APP)"

setup:
	glide install
	@echo "Kafka Hystrix Client is setup!! Run make test to run tests"

build-deps:
	glide install

update-deps:
	glide update

build: update-deps compile fmt vet lint

install:
	go install ./...

fmt:
	go fmt $(GLIDE_NOVENDOR)

vet:
	go vet $(GLIDE_NOVENDOR)

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

test:
	go test $(GLIDE_NOVENDOR)

proto-gen:
	protoc --go_out=plugins=grpc,import_path=proto:. proto/*.proto
