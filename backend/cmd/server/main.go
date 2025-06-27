package main


import (
    "log"

    "github.com/garyhopkins/adminex/api"
)

func main() {
    log.Println("Starting server")

    a := api.New()
    log.Fatalln(a.ListenAndServe("0.0.0.0:8888"))
}
