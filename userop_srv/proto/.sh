protoc -I . address.proto --go_out=. --go-grpc_out=.
protoc -I . message.proto --go_out=. --go-grpc_out=.
protoc -I . fav.proto --go_out=. --go-grpc_out=.