package main

import (
	"inventory-go/db"
	"inventory-go/routes"
	"log"
)

func main() {
	db.Init()

	e := routes.Init()

	log.Println("Listening on :8080...")
	e.Logger.Fatal(e.Start(":8080"))
}
