prepare:

install: prepare
	@go get -u \
		github.com/gobuffalo/packr/packr \
		github.com/urfave/cli \
		github.com/davecgh/go-spew/spew \
		cloud.google.com/go/civil
	
build: gen
	@go build -o build/riksbank cmd/riksbank/*.go

gen: prepare
	@packr

test: prepare
	@go test -cover ./...