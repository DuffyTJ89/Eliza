package main

import (
	"fmt"
	"./eliza"
	//"net/http"
	//"log"
)
/*
 //function that prints the name to the web page
 func server(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "ElizaChat.html")
 }//end server
*/


func main() {

	q1 := "hi my name is tom" // test question

	fmt.Println(eliza.Ask(q1))

	
}

/*
//handle requests by calling printName
	 http.HandleFunc("/", server)
	 
 	//start webserver and serve on port 8080
     log.Println("Listening....")
     http.ListenAndServe(":8080", nil)
*/
