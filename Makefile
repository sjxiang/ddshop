.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# go install github.com/google/wire/cmd/wire@latest


.PHONY: api
# generate api proto
api:
	protoc --proto_path=./user_srv/pb \
 	       --go_out=paths=source_relative:./user_srv/pb \
 	       --go-grpc_out=paths=source_relative:./user_srv/pb \
	       user.proto
