buildapp:
	@go build -o build/adamenblog-api ./cmd/adamenblog-api

run_bin:
	@./build/adamenblog-api

run:
	@go run cmd/adamenblog-api/main.go