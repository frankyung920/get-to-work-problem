.PHONY: unit-test lint init-mocks

unit-test:
	go test -race -tags=unit -covermode=atomic -coverprofile=coverage.out ./... 
	go tool cover -html=coverage.out -o ./coverage.html
