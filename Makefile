IMAGE_NAME=localhost/scada-api
CONTAINER_NAME=scada-api
PORT=8080

up: build run logs

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run --rm -d -p $(PORT):$(PORT) --name $(CONTAINER_NAME) $(IMAGE_NAME)

logs:
	docker logs -f $(CONTAINER_NAME)

rebuild:
	docker build --no-cache -t $(IMAGE_NAME) .

stop:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

clean: stop
	sleep 3
	docker rmi $(IMAGE_NAME) || true

