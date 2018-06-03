build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t email-service .

run:
	docker run -e MICRO_REGISTRY=mdns -e MICRO_BROKER=nats email-service
