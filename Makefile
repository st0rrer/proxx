APP_NAME:=proxx
IMAGE_TAG=$(if $(TAG),$(TAG),latest)
IMAGE_NAME=$(if $(IMAGE),$(IMAGE),$(APP_NAME))

docker_build:
	@DOCKER_BUILDKIT=1 docker build . -t $(IMAGE_NAME):$(IMAGE_TAG)

clean:
	rm -f ./bin/$(APP_NAME)

build: clean
	go build -o ./bin/$(APP_NAME) cmd/*.go
