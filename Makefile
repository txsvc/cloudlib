.PHONY: all
all: test test_coverage
	golangci-lint run > lint.txt

.PHONY: test
test:
	#go test
	cd settings && go test
	cd logger && go test
	cd rest && go test
	cd helpers && go test
	cd provider && go test
	cd observer && go test
	cd observer/provider && go test
	cd storage && go test
	cd storage/provider && go test
	
	
.PHONY: test_coverage
test_coverage:
	go test `go list ./... | grep -v 'hack\|deprecated'` -coverprofile=coverage.txt -covermode=atomic
