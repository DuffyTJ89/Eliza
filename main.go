package main

//func chat
//adapted from https://stackoverflow.com/questions/23282311/parse-input-from-html-form-in-golang


import (
	"fmt"
	"log"
	"net/http"

	"./eliza"
)


func chat(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "web/ElizaChat.html")

	//Call to ParseForm makes form fields available.
	err := r.ParseForm()
	if err != nil {
		// Handle error here via logging and then return
	}
	//take user input from the form on ElizaChat
	userQuestion := r.PostFormValue("userInput")
	fmt.Println(eliza.Ask((userQuestion))) //id ElizaAnswer
}

func main() {

	// test question
	//q1 := "hi my name is tom"
	//fmt.Println(eliza.Ask(q1))

	//handle requests by calling chat function
	http.HandleFunc("/", chat)

	//start webserver and serve on port 8080
	log.Println("Listening....")
	http.ListenAndServe(":8080", nil)

}
