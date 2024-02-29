# Command for compilation .proto files
generate:
	protoc -I api api/sso_v1/sso.proto --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative