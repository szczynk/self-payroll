.PHONY: run
run:
	go run *.go

.PHONY: deps
deps:
	go mod tidy

.PHONY: lint
lint:
	@echo  -e 'Starting golangci-lint'
	golangci-lint run ./...
	@echo  -e '\nStarting staticcheck'
	staticcheck -checks U1000 ./...

.PHONY: mocks
mocks:
	mockery --dir=model --output=model/mocks --outpkg=mocks --all

.PHONY: test-coverage
test-coverage: 
	go test -v ./... -covermode=count -coverpkg=./... -coverprofile coverage/coverage.out -json > coverage/coverage.json
	go tool cover -html coverage/coverage.out -o coverage/coverage.html
	tparse -all -file coverage/coverage.json
	open coverage/coverage.html