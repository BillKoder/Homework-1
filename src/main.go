package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var sentMessage string

type numbers struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

var numberSlice []numbers

func main() {
	router := mux.NewRouter()

	var port string
	portPtr := flag.Int("port", 5309, "the port the router listens to")
	messagePtr := flag.String("message", "", "a message to display at /message")
	flag.Parse()

	if *portPtr <= 1024 || *portPtr >= 65534 {
		fmt.Println("Error: not a valid port. Enter between 1025 and 65533.")
		return
	}
	port = strconv.Itoa(*portPtr)
	port = `:` + port

	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/v1/joke", joke).Methods("GET")
	router.HandleFunc("/transform", transform).Methods("POST")

	if *messagePtr != "" {
		sentMessage = *messagePtr
		router.HandleFunc("/message", message).Methods("GET")

	}
	fmt.Println("router is listening to port:", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}
func joke(w http.ResponseWriter, r *http.Request) {
	var joke string
	switch rand.Intn(10) {
	case 0:
		joke = "Justice is a dish best served cold, if it were served warm it would be justwater."
	case 1:
		joke = "When does a joke become a dad joke? When it becomes apparent."
	case 2:
		joke = "My wife is really mad at the fact that I have no sense of direction. So I packed up my stuff and right."
	case 3:
		joke = "How do you make holy water? You boil the hell out of it."
	case 4:
		joke = "I bought some shoes from a drug dealer. I don't know what he laced them with, but I was tripping all day!"
	case 5:
		joke = "Did you know the first French fries weren't actually cooked in France? They were cooked in Greece."
	case 6:
		joke = "If a child refuses to sleep during nap time, are they guilty of resisting a rest?"
	case 7:
		joke = "I'm reading a book about anti-gravity. It's impossible to put down!"
	case 8:
		joke = "What is the least spoken language in the world? Sign language"
	case 9:
		joke = "When a dad drives past a graveyard: Did you know that's a popular cemetery? Yep, people are just dying to get in there!"
	}
	fmt.Fprint(w, joke)

}
func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, sentMessage)
}
func transform(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newNumbers numbers
	_ = json.NewDecoder(r.Body).Decode(&newNumbers)
	numberSlice = append(numberSlice, newNumbers)
	// fmt.Fprint(w, newNumbers.Sum)
	json.NewEncoder(w).Encode(newNumbers.Number1 + newNumbers.Number2)
}
