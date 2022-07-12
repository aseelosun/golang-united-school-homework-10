package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", func(resp http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(resp, "Hello, PARAM!")
	})
	router.HandleFunc("/data{PARAM}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		param := vars["PARAM"]
		fmt.Fprintf(w, "I got message:\n%s", param)
	})
	//router.HandleFunc("/headers{\"a\":\"2\", \"b\":\"3\"}", func(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)
	//	param := vars["PARAM"]
	//	fmt.Fprintf(w, "I got message:\n%s", param)
	//})
	//router.HandleFunc("/bad", f2).Methods("GET")
	//router.HandleFunc("/data{PARAM}", f3).Methods("POST")
	//router.HandleFunc("/headers", f4).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
