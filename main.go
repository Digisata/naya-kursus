package main

import (
	"fmt"
	"reflect"
)

func main() {
	/* var map1 map[string]string
	map1 = map[string]string{} */
	map1 := map[string]interface{}{
		"name": "nopal",
		"age": 18,
		"is_married": false,
	}
	// map1 := make(map[string]interface{})
	map1["name"] = "hanip"
	map1["age"] = 20
	map1["is_married"] = true
	map1["address"] = "Jl Kaliurang"

	/* fmt.Println(map1["name"])
	fmt.Println(map1["age"])
	fmt.Println(map1["is_married"]) */

	for key, value := range map1 {
		fmt.Println(fmt.Sprintf("Key = %s, value = %v, data type = %v", key, value, reflect.TypeOf(value)))
	}

	fmt.Println()

	fmt.Println(map1)
	fmt.Println(len(map1))
	delete(map1, "name")
	fmt.Println(map1)
	fmt.Println(len(map1))

	fmt.Println(map[string]interface{}{
		"name": "nopal",
		"age": 18,
		"is_married": false,
	})
}
