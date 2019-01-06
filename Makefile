prepare:

gen: prepare
	@packr

install: prepare
	@go get -u github.com/gobuffalo/packr/packr
	@dep ensure -v -update
	
build: gen
	@go build \
	-o build/riksbank

test: gen
	@go test -cover ./...