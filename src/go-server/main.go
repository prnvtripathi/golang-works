package main //just like including the header file in C/C++

import (
	"fmt"
	"log"
	"net/http"
)

//just like importing various packages in python/js

func helloHandler(req http.ResponseWriter, res *http.Request) {
	if res.URL.Path != "/hello" {
		http.Error(req, "404 not found.", http.StatusNotFound)
		return
	} 
	//this is the error handling, if the path is not '/hello' then it will show 404 error

	if res.Method != "GET" {
		http.Error(req, "Method is not supported.", http.StatusNotFound)
		return
	} 
	//this is the error handling, if the method is not 'GET' then it will show 404 error

	fmt.Fprintf(req, "Hello!") 
	//this is the response to the request
}


// formHandler function - handles the request and response of '/form' endpoint
func formHandler(req http.ResponseWriter, res *http.Request) {
	if err := res.ParseForm(); err != nil {
		fmt.Fprintf(req, "ParseForm() err: %v", err)
		return
	}
	//this is the error handling, if the form is not parsed then it will show the error

	fmt.Fprintf(req, "POST request successful\n") //this is the response to the request

	name := res.FormValue("name") //this is the name of the input field in the form
	email := res.FormValue("email") //this is the email of the input field in the form

	//showing the name and email in the console
	fmt.Fprintf(req, "Name = %s\n", name) 
	fmt.Fprintf(req, "Email = %s\n", email)
}


// main function
func main() {
	fileServer := http.FileServer(http.Dir("./static")) 
	//serving static files, it is here to look out for the static folder and serving html files

	http.Handle("/", fileServer) //this is the default handler - root handler
	http.HandleFunc("/form", formHandler) //this is the handler for '/form' endpoint
	http.HandleFunc("/hello", helloHandler) //this is the handler for '/hello' endpoint

	fmt.Println("Server is listening at port 8080") //just like print statement in python/js
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} //this is the server listening at port 8080, and showing the error if any
}
