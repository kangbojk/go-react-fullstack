
# go-react-fullstack



### Code Structure
```
.
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── main.go
├── docker-compose.yaml
├── go.mod
├── go.sum
├── pkg
│   ├── ID
│   │   └── id.go
│   ├── entity
│   │   ├── account
│   │   │   ├── entity.go
│   │   │   ├── fixture.go
│   │   │   └── manager.go
│   │   └── tenant
│   │       ├── entity.go
│   │       └── fixture.go
│   ├── password
│   │   └── password.go
│   ├── server
│   │   ├── data
│   │   │   ├── account.go
│   │   │   └── tenant.go
│   │   ├── middleware
│   │   ├── router
│   │   │   ├── auth.go
│   │   │   ├── handler.go
│   │   │   ├── jwt.go
│   │   │   ├── router.go
│   │   │   ├── router_test.go
│   │   │   └── ws.go
│   │   └── server.go
│   ├── storage
│   │   ├── db
│   │   │   └── account_repo_pg.go
│   │   └── memory
│   │       ├── account_repo_mem.go
│   │       ├── account_repo_mem_test.go
│   │       └── tenant_repo_mem.go
│   └── usecase
│       ├── interface.go
│       ├── usecase.go
│       └── usecase_test.go
└── web
```

The design of this application follow clean/hexagonal architecture.

To learn more about this pattern, please refer to [Building an enterprise service in Go](https://youtu.be/twcDf_Y2gXY) and [How Do You Structure Your Go Apps](https://youtu.be/oL6JBUk6tj0).
