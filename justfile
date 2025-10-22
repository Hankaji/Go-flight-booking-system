default:
    just --list

# Generate protobuf & gRPC Go code
gen-proto:
    protoc --go_out=. --go_opt=module=flight-booking-server \
           --go-grpc_out=. --go-grpc_opt=module=flight-booking-server \
           -I=proto \
           $(find proto -name "*.proto")
