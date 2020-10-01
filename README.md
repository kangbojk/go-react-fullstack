# github.com/kangbojk/go-react-fullstack


### Code Structure
```
.
├── Dockerfile
├── Makefile
├── cmd
│   └── main.go
├── config/
├── internal
│   ├── entity
│   │   ├── account
│   │   │   ├── entity.go
│   │   │   └── fixture.go
│   │   └── tenant
│   │       ├── entity.go
│   │       └── fixture.go
│   ├── server
│   │   ├── data/
│   │   ├── middleware/
│   │   ├── router
│   │   │   ├── handler.go
│   │   │   ├── router.go
│   │   │   └── router_test.go
│   │   └── server.go
│   ├── storage
│   │   ├── db
│   │   │   └── account_repo_pg.go
│   │   └── memory
│   │       ├── account_repo_mem.go
│   │       ├── account_repo_mem_test.go
│   │       └── tenant_repo_mem.go
│   └── usecase
│       ├── interface.go
│       ├── usecase.go
│       └── usecase_test.go
├── pkg
│   ├── ID
│   │   └── id.go
│   └── password
│       └── password.go
└── web/
```

The design of this application follow clean/hexagonal architecture.

To learn more about this pattern, please refer to [Building an enterprise service in Go](https://youtu.be/twcDf_Y2gXY) and [How Do You Structure Your Go Apps](https://youtu.be/oL6JBUk6tj0).