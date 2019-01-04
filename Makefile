prepare:

gen: prepare
	@packr

install: prepare
	@go get -u github.com/gobuffalo/packr/packr
	@dep ensure -v -update
	
test: gen
	@go test ./...