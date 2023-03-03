package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone"`
}

// Nopal, S. Kom.

func main() {
	// customer := map[string]any{
	// 	"Id":    1,
	// 	"Name":  "John",
	// 	"Email": "john@example.com",
	// 	"Phone": "0812412453",
	// }

	// result, err := json.Marshal(customer)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(result))

	// customers := []map[string]any{
	// 	{
	// 		"id":    1,
	// 		"name":  "John",
	// 		"email": "john@example.com",
	// 		"phone": "0812412453",
	// 	},
	// 	{
	// 		"id":    2,
	// 		"name":  "Alex",
	// 		"email": "alex@example.com",
	// 		"phone": "08124123546",
	// 	},
	// 	{
	// 		"id":    3,
	// 		"name":  "Edwin",
	// 		"email": "edwin@example.com",
	// 		"phone": "0812412908",
	// 	},
	// }

	// result2, err2 := json.Marshal(customers)
	// if err2 != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(result2))

	// customer2 := Customer{
	// 	Id:    5,
	// 	Name:  "Zaenal",
	// 	Email: "zaenal@example.com",
	// 	Phone: "08734732323",
	// }

	// result3, err3 := json.Marshal(customer2)
	// if err3 != nil {
	// 	panic(err3)
	// }

	// fmt.Println(string(result3))

	/* jsonString := `
	[
		{
			"id": 5,
			"name": "Zaenal",
			"email": "zaenal@example.com",
			"phone": "08734732323"
		  },
		  {
			"id": 6,
			"name": "alex",
			"email": "alex@example.com",
			"phone": "087347323546"
		  },
		  {
			"id": 7,
			"name": "zendaya",
			"email": "zendaya@example.com",
			"phone": "08734732654"
		  }
	]
	`
	jsonByte := []byte(jsonString)

	// customer := Customer{}
	customer := map[string]any{}

	fmt.Println(customer)

	err := json.Unmarshal(jsonByte, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	// fmt.Printf("Id = %d, Name = %s, Email = %s, Phone = %s", customer.Id, customer.Name, customer.Email, customer.Phone)
	fmt.Printf("Id = %f, Name = %s, Email = %s, Phone = %s", customer["id"], customer["name"], customer["email"], customer["phone"]) */

	// jsonString := `
	// [
	// 	{
	// 	  "id": 5,
	// 	  "name": "Zaenal",
	// 	  "email": "zaenal@example.com",
	// 	  "phone": "08734732323"
	// 	},
	// 	{
	// 	  "id": 6,
	// 	  "name": "alex",
	// 	  "email": "alex@example.com",
	// 	  "phone": "087347323546"
	// 	},
	// 	{
	// 	  "id": 7,
	// 	  "name": "zendaya",
	// 	  "email": "zendaya@example.com",
	// 	  "phone": "08734732654"
	// 	}
	//   ]
	// `
	// jsonByte := []byte(jsonString)
	// customers := []Customer{}

	// err := json.Unmarshal(jsonByte, &customers)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers)

	// encoder := json.NewEncoder(os.Stdout)

	// err4 := encoder.Encode(customer2)
	// if err4 != nil {
	// 	panic(err4)
	// }

	reader, err5 := os.Open("07-01/latihan/anything.json")
	if err5 != nil {
		panic(err5)
	}

	decoder := json.NewDecoder(reader)

	anything := Object{}

	err5 = decoder.Decode(&anything)
	if err5 != nil {
		panic(err5)
	}

	fmt.Println(anything)
}

type Object struct {
	Data     []Data
	Included []Included
}

type Data struct {
	Type          string
	Id            string
	Attributes    Attribute
	Relationships Relationship
}

type Attribute struct {
	Title   string
	Body    string
	Created string
	Updated string
}

type Relationship struct {
	Author Author
}

type Author struct {
	Data AuthorData
}

type AuthorData struct {
	Id   string
	Type string
}

type Included struct {
	Type       string
	Id         string
	Attributes IncludedAttribute
}

type IncludedAttribute struct {
	Name   string
	Age    int
	Gender string
}
