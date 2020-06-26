CMD_NAME=dev
DOCKER_IMAGE_NAMESPACE=zephinzer
DOCKER_IMAGE_NAME=dev
PROJECT_NAME=dev
GO_PACKAGE=$(shell cat go.mod | head -n 1 | cut -f 2 -d ' ')
GIT_COMMIT=$$(git rev-parse --verify HEAD)
GIT_TAG=$$(git describe --tag $$(git rev-list --tags --max-count=1))
TIMESTAMP=$(shell date +'%Y%m%d%H%M%S')

-include ./Makefile.properties

BIN_PATH=$(CMD_NAME)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT}

deps:
	go mod vendor -v
	go mod tidy -v
setup_build:
	# this is required to compile manifest resources for windows
	go get github.com/akavel/rsrc
	# these are required to compile dev without cgo
	go install github.com/mattn/go-sqlite3
run:
	go run -v -mod=vendor ./cmd/$(CMD_NAME) ${args}
test:
	go test -v -mod=vendor ./... -cover -coverprofile c.out
install:
	go install -v -mod=vendor ./cmd/$(CMD_NAME)
build:
	docker run -it \
		--env BIN_NAME=$(CMD_NAME) \
		--user $$(id -u) \
		--volume $$(pwd):/go/$(GO_PACKAGE) \
		--workdir /go/$(GO_PACKAGE) \
		usvc/ci:go-build entrypoint
build_static:
	docker run -it \
		--env BIN_NAME=$(CMD_NAME) \
		--env CGO_ENABLED=0 \
		--env STATIC=1 \
		--user $$(id -u) \
		--volume $$(pwd):/go/$(GO_PACKAGE) \
		--workdir /go/$(GO_PACKAGE) \
		usvc/ci:go-build entrypoint
build_production:
	docker run -it \
		--env BIN_NAME=$(CMD_NAME) \
		--env USE_CACHE=0 \
		--user $$(id -u) \
		--volume $$(pwd):/go/$(GO_PACKAGE) \
		--workdir /go/$(GO_PACKAGE) \
		usvc/ci:go-build-production entrypoint
build_static_production:
	docker run -it \
		--env BIN_NAME=$(CMD_NAME) \
		--env CGO_ENABLED=0 \
		--env STATIC=1 \
		--user $$(id -u) \
		--volume $$(pwd):/go/$(GO_PACKAGE) \
		--workdir /go/$(GO_PACKAGE) \
		usvc/ci:go-build-production entrypoint
checksum:
	sha256sum -b ./bin/$(BIN_PATH) | cut -f 1 -d ' ' > ./bin/$(BIN_PATH).sha256
	rm -rf ./bin/$(CMD_NAME)
	cd ./bin \
		&& ln -s ./$(BIN_PATH) ./$(CMD_NAME) \
		&& ln -s ./$(BIN_PATH).sha256 ./$(CMD_NAME).sha256
compress:
	ls -lah ./bin/$(BIN_PATH)
	upx -9 -v -o ./bin/.$(BIN_PATH) ./bin/$(BIN_PATH)
	upx -t ./bin/.$(BIN_PATH)
	rm -rf ./bin/$(BIN_PATH)
	mv ./bin/.$(BIN_PATH) ./bin/$(BIN_PATH)
	sha256sum -b ./bin/$(BIN_PATH) | cut -f 1 -d ' ' > ./bin/$(BIN_PATH).sha256
	ls -lah ./bin/$(BIN_PATH)
image:
	docker build \
		--build-arg GIT_COMMIT_ID=$(GIT_COMMIT) \
		--build-arg GIT_TAG=$(GIT_TAG) \
		--build-arg BUILD_TIMESTAMP=$(TIMESTAMP) \
		--file ./deploy/Dockerfile \
		--tag $(DOCKER_IMAGE_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest \
		.
test_image:
	container-structure-test test \
		--config ./deploy/Dockerfile.yaml \
		--image $(DOCKER_IMAGE_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest
start_docker:
	docker-compose --file ./deploy/docker/docker-compose.yml up --build -V
stop_docker:
	docker-compose --file ./deploy/docker/docker-compose.yml down --rm
init_k8s:
	if ! kind get clusters | grep $(PROJECT_NAME); then \
		kind create cluster \
			--config ./deploy/kind.config.yaml \
			--name $(PROJECT_NAME); \
	fi
start_k8s: image
	kubectl config use-context kind-$(DOCKER_IMAGE_NAME)
	kind load docker-image --name $(DOCKER_IMAGE_NAME) $(DOCKER_IMAGE_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest

denit_k8s:
	if kind get clusters | grep $(PROJECT_NAME); then \
		kind delete cluster --name $(PROJECT_NAME); \
	fi
dockerhub:
	docker push $(DOCKER_IMAGE_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest
	git fetch
	docker tag $(DOCKER_IMAGE_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest \
		$(DOCKER_IMAGE_NAMESPACE)/$(DOCKER_IMAGE_NAME):$$(git describe --tag $$(git rev-list --tags --max-count=1))
	docker push $(DOCKER_IMAGE_NAMESPACE)/$(DOCKER_IMAGE_NAME):$$(git describe --tag $$(git rev-list --tags --max-count=1))
see_ci:
	xdg-open https://gitlab.com/zephinzer/dev/pipelines
.ssh:
	mkdir -p ./.ssh
	ssh-keygen -t rsa -b 8192 -f ./.ssh/id_rsa -q -N ""
	cat ./.ssh/id_rsa | base64 -w 0 > ./.ssh/id_rsa.base64
