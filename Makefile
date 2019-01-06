prepare:

install: prepare
	@go get -t -v -u \
		github.com/gobuffalo/packr/packr \
		github.com/urfave/cli \
		github.com/davecgh/go-spew/spew
	
build: gen
	@go build -o build/riksbank cmd/riksbank/*.go

gen: prepare
	@packr

test: prepare
	@go test -cover ./...