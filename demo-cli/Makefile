build:
	#GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags netgo

	docker build -t laracom-demo-cli .

run:
	docker run --name demo-cli -e MICRO_REGISTRY=mdns laracom-demo-cli