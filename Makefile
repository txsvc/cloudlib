.PHONY: all
all: test lint coverage

test:
	cd logger && go test
	cd helpers && go test
	cd observer && go test
	cd observer/provider && go test
	cd storage && go test
	cd storage/provider && go test
	
lint:
	golangci-lint run > lint.txt

coverage:
	go test `go list ./... | grep -v 'hack\|deprecated'` -coverprofile=coverage.txt -covermode=atomic
	go tool cover -func=coverage.txt

clean:
	rm -f coverage.txt lint.txt
	go clean
