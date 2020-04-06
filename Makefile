build:
	go build -mod=vendor cmd/covid19keralaapi/main.go
	GIN_MODE=release PORT=5000 ./main

run:
	go run -mod=vendor cmd/covid19keralaapi/main.go
