all:
	go build -o replicator/replicator.bin replicator/main.go replicator/grpcServer.go
	go build -o grpcClient/grpcClient.bin grpcClient/main.go grpcClient/data.go grpcClient/types.go grpcClient/helper.go

replicator:
	go build -o replicator.bin main.go grpcServer.go

grpcClient:
	go build -o grpcClient/grpcClient.bin grpcClient/main.go grpcClient/data.go grpcClient/types.go grpcClient/helper.go
