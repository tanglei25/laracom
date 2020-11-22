build:
	GOOS=linux GOARCH=amd64 go build -tags netgo
	docker build -t laracom-demo-cli .

run:
	docker run laracom-demo-cli