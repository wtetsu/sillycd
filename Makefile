GOCMD=go
BINARY_NAME=sillycd
OUT=dist
CMD=cmd/sillycd/main.go

build:
	go build -o ${OUT}/$(BINARY_NAME) -v ${CMD}
all: build build-osx build-windows build-linux
build-osx:
	GOOS=darwin GOARCH=amd64 go build -o ${OUT}/osx/$(BINARY_NAME) -v ${CMD}
build-windows:
	GOOS=windows GOARCH=amd64 go build -o ${OUT}/windows/$(BINARY_NAME).exe -v ${CMD}
build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${OUT}/linux/$(BINARY_NAME) -v ${CMD}
test:
	go test -v ./...
clean:
	go clean ${CMD}
	rm -f ${OUT}/osx/$(BINARY_NAME)
	rm -f ${OUT}/windows/$(BINARY_NAME)
	rm -f ${OUT}/linux/$(BINARY_NAME)
