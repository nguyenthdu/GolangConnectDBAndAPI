package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	jsonString := `{"title":"Learning JSON in Golang","author":"Lanka"}`
	var book Book
	err := json.Unmarshal([]byte(jsonString), &book)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", book)

}

/*
Đôi khi chúng ta phải làm việc với dữ liệu không có câu trúc.
Trong trường hợp này, chúng ta có thể sử dụng cách thay thế bằng map[string]interface{} là kiểu trả về của biển được giải nén
*/

/*
jsonString := `{"title":"Learning JSON in Golang","author":"Lanka"}`
var book map[string]interface{}
err := json.Unmarshal([]byte(jsonString), &book)
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", book)

*/
