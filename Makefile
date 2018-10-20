packr:
	@packr

gen: packr
	@go run gen/gen.go ./api/series

install:
	@dep ensure -v -update

run: packr
	@go run test/main.go