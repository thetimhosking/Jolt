package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {

		var myTime string
		myTime = time.Now().Format("02-01-2006")

		log.Println("Hello world")
		log.Println("Current time is ", myTime)
	})
	http.ListenAndServe(":9090", nil)

}
