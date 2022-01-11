package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("bismillah walhamdulillah")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("listening on port 8000...")
	http.ListenAndServe(":8000", nil)

}
