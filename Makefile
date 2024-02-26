.PHONY: docker
docker:
	DOCKER_BUILDKIT=1 docker build --progress=plain -f ./build/Dockerfile .

.PHONY: inject
inject:
	wire gen ./inject

.PHONY: tidy
tidy:
	go mod tidy