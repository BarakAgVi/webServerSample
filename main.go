package main

import (
	"net/http"
	"fmt"
)



	func printOnWeb(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"<h1>welcome to the site<h1>")
	}

	func main(){
			http.HandleFunc("/", printOnWeb)
			http.ListenAndServe(":6060",nil)
	}