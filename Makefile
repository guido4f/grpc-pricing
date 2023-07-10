#DST_DIR=./generated
.DEFAULT_GOAL := build

generate-source:
	- rm -rf gen/
	- mkdir -p gen
	docker run  -u $(id -u ${USER}):$(id -g ${USER}) -v $(PWD)/proto:/defs:z -v $(PWD)/gen:/go:z namely/protoc-all -d guido4f/fee/ 	-o /go  -l go

build: generate-source
	#go build -o server server/main.go
	docker build -f build/Dockerfile . -t fee-calculator-service
