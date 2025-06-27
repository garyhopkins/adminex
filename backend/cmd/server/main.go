package main

import (
	"log"

	"github.com/garyhopkins/adminex/api"
	"github.com/garyhopkins/adminex/app"
)

func main() {
	log.Println("Starting server")
	app.GetApp()

	a := api.New()
	log.Fatalln(a.ListenAndServe("0.0.0.0:8888"))
}
