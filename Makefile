gen:
	protoc --proto_path=protofiles protofiles/*.proto --go_out=plugins=grpc:pbfiles
clean:
	rm pbfiles/*.go
run:
	go run main.go
server-run:
	clear
	go run main.go
server-clean-and-run:
	clear
	rm pbfiles/*.go
	protoc --proto_path=protofiles protofiles/*.proto --go_out=plugins=grpc:pbfiles
	go run main.go
server-clean-and-run-with-log:
	clear
	rm pbfiles/*.go
	protoc --proto_path=protofiles protofiles/*.proto --go_out=plugins=grpc:pbfiles
	GRPC_GO_LOG_SEVERITY_LEVEL=info GRPC_GO_LOG_VERBOSITY_LEVEL=2 go run main.go
client-run:
	clear
	go run client/catalog_service_client.go
