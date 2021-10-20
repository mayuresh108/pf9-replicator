all:
	go build -race -a -gcflags="all=-N -l" -o grpcClient/grpcClient.bin grpcClient/main.go grpcClient/data.go grpcClient/types.go grpcClient/helper.go
	cp grpcClient/grpcClient.bin replicator
	go build -race -a -gcflags="all=-N -l" -o replicator/replicator.bin replicator/main.go replicator/grpcServer.go

replicator:
	go build -race -a -gcflags="all=-N -l" -o replicator.bin main.go grpcServer.go

grpcClient:
	go build -race -a -gcflags="all=-N -l" -o grpcClient/grpcClient.bin grpcClient/main.go grpcClient/data.go grpcClient/types.go grpcClient/helper.go
