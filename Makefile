start:
	go build cmd/covid19keralaapi/main.go
	GIN_MODE=release PORT=5000 ./main
