test:
	go test ./...

test-pprof:
	go test -cpuprofile main.pprof ./cmd/main_test.go

pprof:
	go tool pprof -http :8088 main.pprof
