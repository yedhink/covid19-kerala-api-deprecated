build:
	go mod vendor
	go build -mod=vendor -v -o bin/covid19keralaapi cmd/covid19keralaapi/main.go
	PORT=5000 bin/covid19keralaapi

run:
	# runs on port 8000 by default
	go run -mod=vendor cmd/covid19keralaapi/main.go
