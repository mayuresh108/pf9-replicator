all:
	#env GOOS=linux GOARCH=amd64 go build -race -a -gcflags="all=-N -l" -o grpcClient/grpcClient.bin grpcClient/main.go grpcClient/data.go grpcClient/types.go grpcClient/helper.go
	env GOOS=linux GOARCH=amd64 go build -a -gcflags="all=-N -l" -o grpcClient/grpcClient.bin grpcClient/main.go grpcClient/data.go grpcClient/types.go grpcClient/helper.go
	cp grpcClient/grpcClient.bin ./replicator/
	#env GOOS=linux GOARCH=amd64 go build -race -a -gcflags="all=-N -l" -o replicator/replicator.bin replicator/main.go replicator/grpcServer.go
	env GOOS=linux GOARCH=amd64 go build -a -gcflags="all=-N -l" -o replicator/replicator.bin replicator/main.go replicator/grpcServer.go

replicator:
	go build -race -a -gcflags="all=-N -l" -o replicator.bin main.go grpcServer.go

grpcClient:
	go build -race -a -gcflags="all=-N -l" -o grpcClient/grpcClient.bin grpcClient/main.go grpcClient/data.go grpcClient/types.go grpcClient/helper.go
