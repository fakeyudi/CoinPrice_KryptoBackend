package main

import (
    "fmt"
    "CoinPrice_KryptoBackendTask/router"
	"CoinPrice_KryptoBackendTask/middleware"
    "log"
    "net/http"
)

func main() {
    r := router.Router()
    // fs := http.FileServer(http.Dir("build"))
    // http.Handle("/", fs)
	go middleware.AllAlerts()
    fmt.Println("Starting server on the port 3000...")

    log.Fatal(http.ListenAndServe(":3000", r))
}