# GO - gRPC - ProtoBuffers




### Pre - Requirements

- go: -      https://go.dev/doc/install
- protoc: -  https://grpc.io/docs/protoc-installation/


### Instructions using Windows

```bash 
go version
go mod init pdaniel.com/go/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/protobuf
go get github.com/lib/pq
go get google.golang.org/grpc
go get google.golang.org/grpc/codes
go get google.golang.org/grpc/status

protoc --version
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative studentpb/student.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative testpb/test.proto

cd database
  docker build . -t pdaniel-grpc-db 
  docker run -p 54321:5432 -e POSTGRES_PASSWORD=postgres pdaniel-grpc-db

```
```bash 
go run server-student/main.go
```
Then use Postman,GRPC,  URL: localhost:5060, ServerReflectionInfo 