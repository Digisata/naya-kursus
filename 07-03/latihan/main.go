package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Items struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

func main() {
	items := []Items{
		{
			Id: 1,
			Name: "Shampoo",
			Qty: 15,
		},
		{
			Id: 2,
			Name: "Snack",
			Qty: 150,
		},
		{
			Id: 3,
			Name: "Cigaret",
			Qty: 90,
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World again!")
	})

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		result, _ := json.Marshal(items)
		
		fmt.Fprint(w, string(result))
	})


	log.Fatal(http.ListenAndServe(":5000", nil))
}
