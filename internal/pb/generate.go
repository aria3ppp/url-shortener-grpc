package pb

//go:generate protoc --proto_path=../../proto --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative url-shortener.proto
