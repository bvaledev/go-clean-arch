# go-clean-arch

- Run app: `go run cmd/ordersystem/main.go  cmd/ordersystem/wire_gen.go`
- Test gRPC: `evans -r repl` -> `package pb` -> `service OrderService` -> `call ListOrders`
- Generate graphql schema: `go run github.com/99designs/gqlgen generate`
- Generate gRPC Proto: `protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto`

## Migration
- Access docker mysql: `mysql -uroot -proot`
- Select database: `use orders;`
- Create table:
```
CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id));
```
- Show created table: `SHOW TABLES;`


## GraphQL Query
```
query List {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}

mutation Create {
  createOrder(input: {id: "124", Price: 239, Tax: 23}) {
    id
    Price
    Tax
    FinalPrice
  }
}
```