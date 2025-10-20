default:
    just --list

# Generate protobuf & gRPC Go code
gen-proto:
    protoc --go_out=. \
           --go-grpc_out=. \
           proto/flight.proto
