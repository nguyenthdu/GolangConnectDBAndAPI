package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	social := Social{Facebook: "https://facebook.com", Twitter: "https://twitter.com"}
	user := User{Name: "LanKa", Type: "Author", Age: 25, Social: social}
	// jsonData, err := json.Marshal(user)
	// Giá trị đầu vào của MarshalIndent chứa prefix(tiền tố) và indent(thụt lề)
	jsonData, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

}
