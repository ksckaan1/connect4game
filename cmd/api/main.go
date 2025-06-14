package main

import (
	"log"

	"github.com/ksckaan1/connect4backend/internal/core/adapter/left/httpapi"
)

func main() {
	server := httpapi.New().WithPort(":3000")
	server.Mount(httpapi.NewExampleHttpAPI().WithPattern("/example"))
	log.Println("server started on port 3000")
	err := server.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
