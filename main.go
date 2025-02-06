package main

import (
	"fmt"
	"log"
	"net/http"
	"netflix-backend/router"
)

func main() {
	fmt.Println("Netflix Backend Watchlist")
	r := router.Router()

	fmt.Println("Server is getting started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
