package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Homepage enpoint hit")
}
func coba(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"coba")
}

func handleRequest(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/test",coba)
	log.Fatal(http.ListenAndServe(":1808",nil))
}

func main(){handleRequest()}


