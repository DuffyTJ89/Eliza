//Eliza home page
// Author : Thomas Duffy
 
 package main
 
 import (
 	"net/http"
     "log"
 )
 
 //function that prints the name to the web page
 func server(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "ElizaChat.html")
 }//end printName
 
 func main() {
 	//handle requests by calling printName
	 http.HandleFunc("/", server)
	 
 	//start webserver and serve on port 8080
     log.Println("Listening....")
     http.ListenAndServe(":8080", nil)
 }