package main

import (
	"connect_mysql/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import driver MySQL dưới dạng import rỗng
)

func main() {
	userService := service.NewUserService()
	//TODO: insert
	//user := &model.User{
	//	Name:  "Nguyen Van C",
	//	Email: "c@gmail.com",
	//}
	//err := userService.CreateUser(user)
	//if err != nil {
	//	fmt.Println("Error creating user:", err)
	//	return
	//}
	//TODO: update
	//userUpdate := &model.User{
	//	ID:    5,
	//	Name:  "Nguyen Van D",
	//	Email: "Thanhdu@gmail.com",
	//}
	//errUpdate := userService.UpdateUser(userUpdate)
	//if errUpdate != nil {
	//	fmt.Println("Error updating user:", errUpdate)
	//	return
	//}
	//TODO: delete
	errDelete := userService.DeleteUser(5)
	if errDelete != nil {
		fmt.Println("Error deleting user:", errDelete)
		return
	}

	//TODO: Get all
	users, errGetAll := userService.GetAllUsers()
	if errGetAll != nil {
		fmt.Println("Error fetching users:", errGetAll)
		return
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	//db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/demogo1")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close() //khi ket thuc chuong trinh thi dong ket noi
	////kiem tra ket noi den mysql
	//err = db.Ping()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Connected to MySQL database.")
	////TODO: select dữ liệu
	//row, err := db.Query("SELECT * FROM users")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer row.Close()
	//for row.Next() {
	//	var id int
	//	var name string
	//	var email string
	//	err = row.Scan(&id, &name, &email)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(id, name, email)
	//}
	////TODO: insert dữ liệu
	////insertStmt, err := db.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	////if err != nil {
	////	log.Fatal(err)
	////}
	////defer insertStmt.Close()
	////result, err := insertStmt.Exec("Nguyen Van A", "a@gmail.com")
	////if err != nil {
	////	log.Fatal(err)
	////}
	////lastInsertID, err := result.LastInsertId()
	////if err != nil {
	////	log.Fatal(err)
	////}
	////fmt.Println("Last insert ID:", lastInsertID)
	////TODO: update dữ liệu
	////updateStmt, err := db.Prepare("UPDATE users SET name = ? WHERE id = ?")
	////if err != nil {
	////	log.Fatal(err)
	////}
	////_, err = updateStmt.Exec("Nguyen Van B", 4)
	////if err != nil {
	////	log.Fatal(err)
	////
	////}
	//fmt.Println("Update success")
	////TODO: delete dữ liệu
	//deleteStmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_, err = deleteStmt.Exec(4)
	//if err != nil {
	//	log.Fatal(err)
	//
	//}
	//fmt.Println("Delete success")
}
