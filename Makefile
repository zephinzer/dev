CMD_ROOT=dev
DOCKER_NAMESPACE=zephinzer
DOCKER_IMAGE_NAME=dev
PROJECT_NAME=dev
GIT_COMMIT=$$(git rev-parse --verify HEAD)
GIT_TAG=$$(git describe --tag $$(git rev-list --tags --max-count=1))
TIMESTAMP=$(shell date +'%Y%m%d%H%M%S')

-include ./Makefile.properties

BIN_PATH=$(CMD_ROOT)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT}

deps:
	go mod vendor -v
	go mod tidy -v
setup_build:
	# this is required to create icons for use in the system tray
	go get github.com/cratonica/2goarray
	# this is required to compile manifest resources for windows
	go get github.com/akavel/rsrc
	# these are required to compile dev without cgo
	go install github.com/mattn/go-sqlite3
setup_build_linux:
	if ! apt-get --version >/dev/null; then \
		sudo apt-get install libgtk-3-dev libappindicator3-dev libwebkit2gtk-4.0-dev; \
	fi
run:
	go run -v -mod=vendor ./cmd/$(CMD_ROOT) ${args}
test:
	go test -v -mod=vendor ./... -cover -coverprofile c.out
prepare_icon:
	@if ! 2goarray -h >/dev/null; then \
		printf -- "\033[1m\033[31m⚠️ you need 2goarray in your path for this to work, ensure you've run 'make setup_build'\033[0m\n"; \
		exit 1; \
	fi
	2goarray SystrayIconDark constants < ./assets/icon/512-dark.png > ./internal/constants/icon_dark.go
	2goarray SystrayIconLight constants < ./assets/icon/512-light.png > ./internal/constants/icon_light.go
install_local:
	go install -v -mod=vendor ./cmd/$(CMD_ROOT)
build:
	go build -mod=vendor \
		-o ./bin/$(BIN_PATH) \
		./cmd/$(CMD_ROOT)
	rm -rf ./bin/$(CMD_ROOT)
	cd ./bin \
		&& ln -s ./$(BIN_PATH) ./$(CMD_ROOT)
build_production:
	go build -a -v \
		-ldflags "-X main.Commit=$(GIT_COMMIT) \
			-X main.Version=$(GIT_TAG) \
			-X main.Timestamp=$(TIMESTAMP) \
			-s -w" \
		-o ./bin/$(BIN_PATH) \
		./cmd/$(CMD_ROOT)
	sha256sum -b ./bin/$(BIN_PATH) \
		| cut -f 1 -d ' ' > ./bin/$(BIN_PATH).sha256
compress:
	ls -lah ./bin/$(BIN_PATH)
	upx -9 -v -o ./bin/.$(BIN_PATH) \
		./bin/$(BIN_PATH)
	upx -t ./bin/.$(BIN_PATH)
	rm -rf ./bin/$(BIN_PATH)
	mv ./bin/.$(BIN_PATH) \
		./bin/$(BIN_PATH)
	sha256sum -b ./bin/$(BIN_PATH) \
		| cut -f 1 -d ' ' > ./bin/$(BIN_PATH).sha256
	ls -lah ./bin/$(BIN_PATH)

image:
	docker build \
		--build-arg GIT_COMMIT_ID=$(GIT_COMMIT) \
		--build-arg GIT_TAG=$(GIT_TAG) \
		--build-arg BUILD_TIMESTAMP=$(TIMESTAMP) \
		--file ./deploy/Dockerfile \
		--tag $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest \
		.
test_image:
	container-structure-test test \
		--config ./deploy/Dockerfile.yaml \
		--image $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest
save:
	mkdir -p ./build
	docker save --output ./build/$(PROJECT_NAME).tar.gz $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest
load:
	docker load --input ./build/$(PROJECT_NAME).tar.gz
dockerhub:
	docker push $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest
	git fetch
	docker tag $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest \
		$(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):$$(git describe --tag $$(git rev-list --tags --max-count=1))
	docker push $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):$$(git describe --tag $$(git rev-list --tags --max-count=1))

see_ci:
	xdg-open https://gitlab.com/zephinzer/dev/pipelines

.ssh:
	mkdir -p ./.ssh
	ssh-keygen -t rsa -b 8192 -f ./.ssh/id_rsa -q -N ""
	cat ./.ssh/id_rsa | base64 -w 0 > ./.ssh/id_rsa.base64
