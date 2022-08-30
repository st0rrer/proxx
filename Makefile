APP_NAME:=proxx
IMAGE_TAG=$(if $(TAG),$(TAG),latest)
IMAGE_NAME=$(if $(IMAGE),$(IMAGE),$(APP_NAME))

docker_build:
	@DOCKER_BUILDKIT=1 docker build . -t $(IMAGE_NAME):$(IMAGE_TAG)

docker_run: docker_build
	docker run -ti $(IMAGE_NAME):$(IMAGE_TAG) /app/proxx

build:
	go build -o ./bin/$(APP_NAME) cmd/*.go

clean:
	rm ./bin/$(APP_NAME)

run: build
	./bin/$(APP_NAME)

