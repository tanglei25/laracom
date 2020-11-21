build:
	#protoc -I. --go_out=plugins=grpc:. proto/demo/demo.proto
	protoc -I. --go_out=plugins=micro:. proto/demo/demo.proto
	GOOS=linux GOARCH=amd64 go build -tags netgo
	docker build -t laracom-demo-service .
run:
	#docker run -d --name demo -p 9091:9999 -p 8081:8000 laracom-demo-service
	docker run -d --name demo -p 9091:9091 -e MICRO_SERVER_ADDRESS=:9091 -e MICRO_REGISTRY=mdns laracom-demo-service