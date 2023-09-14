package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	user := User{Id: "ID001", Name: "LanKa", Password: "123465"}
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
	fmt.Println(string(jsonData))
}
