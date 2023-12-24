grpc-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative dictpb/dict.proto


attack:
	go run ./ddos/main.go