build:
	protoc -I. --micro_out=. --go_out=. proto/user/user.proto
	docker build -t laracom-user-service .
run:
	docker run --name user-service -e MICRO_REGISTRY=mdns laracom-user-service