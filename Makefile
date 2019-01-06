prepare:

install: prepare
	@go get -u github.com/gobuffalo/packr/packr
	@dep ensure -v -update
	
build: gen
	@go build -o build/riksbank cmd/riksbank/*.go

gen: prepare
	@packr

test: prepare
	@go test -cover ./...