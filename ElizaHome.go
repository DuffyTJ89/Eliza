//Eliza home page
// Author : Thomas Duffy
 
 package main
 
 import (
 	"fmt"
 	"net/http"
     "log"
 )
 
 //function that prints the name to the web page
 func printName(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "Welcome to Eliza")
 }//end printName
 
 func main() {
 	//handle requests by calling printName
     http.HandleFunc("/", printName)
 	//start webserver and serve on port 8080
     log.Println("Listening....")
     http.ListenAndServe(":8080", nil)
 }