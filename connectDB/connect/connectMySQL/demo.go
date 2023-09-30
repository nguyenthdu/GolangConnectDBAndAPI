package connectMySQL

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Connect mysql")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/demoGolang1")
	if err != nil {
		panic(err.Error())

		defer db.Close()

	}
}
