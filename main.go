package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("OK"))
	})
	err := http.ListenAndServe(":5000",nil)
	if err!= nil{
		log.Fatalln("couldn't start the server")
	}
}
