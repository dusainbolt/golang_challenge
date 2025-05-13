package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	fmt.Println("== 2. JSON Encoding (struct → JSON) ==")
	product := Product{Name: "Laptop", Price: 15000000}
	productJSON, _ := json.Marshal(product)
	fmt.Println("JSON Product:", string(productJSON))

	fmt.Println("\n== 3. JSON Decoding (JSON → struct) ==")
	userJSON := `{"name": "Du", "age": 27, "hobbies": ["Đọc sách", "Chơi game"]}`
	var user User
	_ = json.Unmarshal([]byte(userJSON), &user)
	fmt.Printf("Tên: %s\nTuổi: %d\nSở thích: %v\n", user.Name, user.Age, user.Hobbies)

	fmt.Println("\n== 4. Generic JSON with interface{} ==")
	jsonStr := `{"title": "Go Developer", "experience": 3, "skills": ["Go", "Docker", "Kubernetes"]}`
	var genericData map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &genericData)
	fmt.Println("Dữ liệu raw:", genericData)

	fmt.Println("\n== 5. Decoding Arbitrary Data ==")
	for key, val := range genericData {
		switch v := val.(type) {
		case string:
			fmt.Printf("%s (string): %s\n", key, v)
		case float64:
			fmt.Printf("%s (number): %.0f\n", key, v)
		case []interface{}:
			fmt.Printf("%s (array):\n", key)
			for i, skill := range v {
				fmt.Printf("  %d. %s\n", i+1, skill)
			}
		default:
			fmt.Printf("%s (unknown type)\n", key)
		}
	}

	fmt.Println("\n== 6. Reference Types (slice, map) ==")
	hobbyJSON := `{"name": "Du", "hobbies": ["Code", "Đọc sách"]}`
	var u2 User
	_ = json.Unmarshal([]byte(hobbyJSON), &u2)
	fmt.Printf("Tên: %s\n", u2.Name)
	for _, h := range u2.Hobbies {
		fmt.Println("Sở thích:", h)
	}
}
