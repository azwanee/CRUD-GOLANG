package main

import (
	"crud-api/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("PORT_SERVER") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error read env file with err: %s", err)
		}
	}

	e := routes.Routes()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT_SERVER"))))
}
