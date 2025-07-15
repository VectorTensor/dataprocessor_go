
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(GOPATH)/bin/protoc-gen-go-grpc

PROTO_DIR := protos
OUT_DIR := src/codegen

PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

.PHONY: all clean proto

all: proto

proto:
	@echo "Generating Go code from .proto files..."
	@mkdir -p $(OUT_DIR)
	@protoc \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		-I $(PROTO_DIR) \
		$(PROTO_FILES)

clean:
	@echo "Cleaning generated code..."
	@rm -f $(OUT_DIR)/*.pb.go $(OUT_DIR)/*_grpc.pb.go
