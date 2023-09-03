# go-clean-arch

- Run app: `go run cmd/ordersystem/main.go  cmd/ordersystem/wire_gen.go`
- Test gRPC: `evans -r repl` -> `package pb` -> `service OrderService` -> `call ListOrders`
- Generate graphql schema: `go run github.com/99designs/gqlgen generate`
- Generate gRPC Proto: `protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto`
