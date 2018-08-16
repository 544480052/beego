package main

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/",home)
	http.ListenAndServe(":9090",nil)
}

func home(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"hello world!!!!")
}