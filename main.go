package main

import (
	"ApiRest/routes"
	"ApiRest/services"
	"log"
)

func main() {
	services.InitDB()
	r := routes.InitRouter()
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
