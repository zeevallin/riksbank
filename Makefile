packr:
	@packr

gen: packr
	@go run gen/gen.go ./api/series
	@cd api/swea && wsdl2go -p="swea" -yolo < sweaWS.wsdl > swea.go

install:
	@dep ensure -v -update

run: packr
	@go run test/main.go