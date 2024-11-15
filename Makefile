UID=$(shell id -u)
GID=$(shell id -g)

DOCKER_BUILDKIT=1
DOCKER_TAG:=$(shell pwd | md5sum | cut -f1 -d ' ')
DOCKER_RUN_FLAGS+=-v $(PWD):/workspace
DOCKER_RUN_FLAGS+=-it
DOCKER_RUN_FLAGS+=--rm
DOCKER_RUN_FLAGS+=--mount type=volume,source=gomod-pkg,destination=/go/pkg
DOCKER_RUN_FLAGS+=--mount type=volume,source=gobuild-cache,destination=/root/.cache/go-build
DOCKER_RUN_FLAGS+=--mount type=volume,source=golangci-cache,destination=/root/.cache/golangci-lint
DOCKER_RUN_FLAGS+=--user $(UID):$(GID)

setup-docker:
	docker buildx build --build-arg UID=$(UID) --build-arg GID=$(GID) --tag $(DOCKER_TAG) .

run build test lint format clean setup: setup-docker
	@docker run $(DOCKER_RUN_FLAGS) $(DOCKER_TAG) make -f Makefile.main $@

enter: setup-docker
	@docker run $(DOCKER_RUN_FLAGS) $(DOCKER_TAG) /bin/sh
