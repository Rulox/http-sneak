build:
	CGO_ENABLED=0 GOOS=linux go build -o http-sneak
	docker build -t rulox/http-sneak:latest .

run:
	go run main.go