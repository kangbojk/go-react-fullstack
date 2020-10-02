package main

import (
	"fmt"
	"log"

	"github.com/kangbojk/go-react-fullstack/internal/server"
	"github.com/kangbojk/go-react-fullstack/internal/storage/memory"
	"github.com/kangbojk/go-react-fullstack/internal/usecase"
)

func main() {

	// TODO: Conditional storage
	
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	// db, err := sql.Open("postgresql", dataSourceName)
	// if err != nil {
	// 		log.Fatal(err.Error())
	// }
	// log.Println("connecting to db ", dataSourceName)
	// defer db.Close()

	aRepo := in_memory.NewAccountRepoMem()
	tRepo := in_memory.NewTenantRepoMem()

	service := usecase.NewService(aRepo, tRepo)
	server := server.NewServer(service)

	// TODO: Read port from config
	fmt.Println("start server on port 8088")
	log.Fatal(server.ListenAndServe())
}