package main

import (
	"connect_mysql_gorm/database"
	"connect_mysql_gorm/model"
	"connect_mysql_gorm/router"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&model.User{})
	// Khởi tạo router và chạy server
	r := router.SetupRouter()
	r.Run(":8080")
	//userService := service.NewUserService()
	//TODO: Thêm một user mới để kiểm tra
	//newUser := &model.User{Name: "Thanh Du", Email: "thanhdu@example.com"}
	//err := userService.CreateUser(newUser)
	//if err != nil {
	//	fmt.Println("Error creating user:", err)
	//	return
	//}
	//fmt.Println("User created:", newUser)
	//TODO: Lấy user theo ID
	//user, err := userService.GetUserByID(4)
	//if err != nil {
	//	fmt.Println("Error fetching user:", err)
	//	return
	//}
	//fmt.Printf("Fetched user: ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	//TODO: Cập nhật user
	//user := &model.User{ID: 4, Name: "Thanh Du11111111111", Email: "aa"}
	//err := userService.UpdateUser(user)
	//if err != nil {
	//	fmt.Println("Error updating user:", err)
	//	return
	//}
	//fmt.Println("User updated:", user)
	//TODO: Xóa user
	//err := userService.DeleteUser(4)
	//if err != nil {
	//	fmt.Println("Error deleting user:", err)
	//	return
	//}
	//fmt.Println("User deleted")
	//TODO: get all
	//users, errGetAll := userService.GetAllUsers()
	//if errGetAll != nil {
	//	fmt.Println("Error fetching users:", errGetAll)
	//	return
	//}
	//fmt.Println("Users:")
	//for _, user := range users {
	//	fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	//}
}
