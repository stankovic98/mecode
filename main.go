package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {

	spin(100 * time.Millisecond)
	const n = 45
	fibN := Fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	// http.HandleFunc("/displayname", displayName)
	// log.Fatal(http.ListenAndServe(":8000", nil))
}

func displayName(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	result := &Person{
		q.Get("name"),
		q.Get("surname"),
	}

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "Unable to encode json", http.StatusInternalServerError)
		log.Println(err)
	}
}
