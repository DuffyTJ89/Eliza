package main

import (
	"fmt"
	"log"
	"net/http"

	"./eliza"
)


func chat(w http.ResponseWriter, r *http.Request) {

	//Call to ParseForm makes form fields available.
	err := r.ParseForm()
	if err != nil {
		// Handle error here via logging and then return
	}
	//take user input from the form on index.html
	userQuestion := r.URL.Query().Get("userInput") //get the users question from the url
	ElizaAnswer:=eliza.Ask(userQuestion) //pass the user question to the eliza file and get the answer
	fmt.Fprintf(w, ElizaAnswer) //eliza answer

}

func main() {

	dir:=http.Dir("./web") //all the files needed are in the web folder
	fileServer:=http.FileServer(dir)
	http.Handle("/", fileServer)

	//handle requests by calling chat function
	http.HandleFunc("/chat", chat)

	//start webserver and serve on port 8080
	log.Println("Listening....")
	http.ListenAndServe(":8080", nil)

}
