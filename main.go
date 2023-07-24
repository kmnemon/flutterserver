package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Pizza struct {
	Id          int     `json:"id"`
	PizzaName   string  `json:"pizzaName"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	pizza1 := &Pizza{
		Id:          1,
		PizzaName:   "sweat",
		Description: "it's lovely pizza",
		Price:       32.0,
		ImageUrl:    "http://xxx.xxx",
	}

	pizza2 := &Pizza{
		Id:          2,
		PizzaName:   "saulty pizza",
		Description: "it's a little saulty one",
		Price:       23.0,
		ImageUrl:    "http://yyy.yyy",
	}

	pizzas := []Pizza{
		*pizza1, *pizza2,
	}

	json, _ := json.Marshal(pizzas)
	writer.Write(json)
}

func pizzaHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var pizza Pizza
	if err := json.Unmarshal([]byte(request.), &pizza); err != nil{
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}


}

func main() {
	http.HandleFunc("/pizzas", handler)
	http.HandleFunc("/pizza", pizzaHandler)
	http.ListenAndServe(":8080", nil)
}
