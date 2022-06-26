
.PHONY: test coverage clean

test:
	go test -v

coverage:
	go test -failfast -covermode=atomic -coverprofile=coverage.out

clean:
	go clean -x -i ./...
