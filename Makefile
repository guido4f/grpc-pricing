#DST_DIR=./generated


generate-source:
	- mkdir -p gen
	- rm -rf gen/*.go
	docker run  -u $(id -u ${USER}):$(id -g ${USER}) \
				-v $(PWD)/proto:/defs:z \
 			 	-v $(PWD)/gen:/go:z \
 			namely/protoc-all \
 			 	-d byhiras/pricing/ \
 			 	-o /go  \
 			 	-l go

build: generate-source
	go build -o server server/main.go