package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter , r *http.Request) {
	fmt.Fprintf(w , "hello , world");
}

func goodbyeHandler(w http.ResponseWriter , r *http.Request) {
	fmt.Println(r);
	fmt.Println("----------------------------")
	fmt.Println(w);
	fmt.Fprintf(w , "Good bye see you again !!!");
}

// json handling

type Response struct {
	Message string `json:"message"`
}

func greetHandler(w http.ResponseWriter , r *http.Request) {
	response := Response { Message: "Hi Welcome to the api..."}
	json.NewEncoder(w).Encode(response);	
}


func main() {
	http.HandleFunc("/" , helloHandler);
	http.HandleFunc("/goodbye" , goodbyeHandler);
	http.HandleFunc("/greet" , greetHandler);
	http.ListenAndServe(":3000" , nil);
}