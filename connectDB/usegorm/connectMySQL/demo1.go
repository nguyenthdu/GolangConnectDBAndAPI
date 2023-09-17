package main

import (
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Student struct {
	Id       string `gorm:"primary_key"`
	FullName string
	Email    string
	Phone    string
	CardId   string
}

// TODO: generate id auto
// go get github.com/google/uuid
//
//	func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
//		//s.Id = uuid.New().String()
//		return nil
//
// }
//func NewID() (id string) {
//	id, _ = gonanoid.ID(10)
//	return id
//}

// Doi ten bang
/*type Tabler interface {
	TabelName() string
}

func (s*Student) TableName() string {
	return "Student"
}*/
//Custtom arrument
type custom struct {
	Name  string
	Email string
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:password@tcp(127.0.0.1:3306)/demoGolang1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Discard, //log cac cau lenh sql

	})
	if err != nil {
		panic("failed to connect database") //ket noi that bai
	}
	fmt.Println("Connect to database successfully")
	//Migrate the schema
	db.AutoMigrate(&Student{}) //tao bang
	//TODO: Create
	db.Create(&Student{Id: "1", FullName: "Nguyen Van A", Email: "email1", Phone: "phone1", CardId: "card1"})
	db.Create(&Student{Id: "2", FullName: "Nguyen Van B", Email: "email2", Phone: "phone2", CardId: "card2"})
	db.Create(&Student{Id: "3", FullName: "Nguyen Van C", Email: "email3", Phone: "phone3", CardId: "card3"})
	db.Create(&Student{Id: "4", FullName: "Nguyen Van D", Email: "email4", Phone: "phone4", CardId: "card4"})
	db.Create(&Student{Id: "5", FullName: "Nguyen Van E", Email: "email5", Phone: "phone5", CardId: "card5"})
	db.Create(&Student{Id: "6", FullName: "Nguyen Van F", Email: "email6", Phone: "phone6", CardId: "card6"})
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return
	}
	////Tao doi tuong dua trne model
	std := Student{Id: "7", FullName: "Nguyen Van G", Email: "email7", Phone: "phone7", CardId: "card7"}
	tx.Create(&std)
	//Tao doi tượng với các truường chỉ định
	tx.Select("Id", "FullName", "Email", "Phone", "CardId").Create(&Student{Id: "9", FullName: "Nguyen Van G", Email: "email7", Phone: "phone7", CardId: "card7"})
	// Tạo đối tượng với các thông tin ngoại trừ các thông tin có trong Omit
	tx.Omit("Email", "Phone").Create(&Student{Id: "100", FullName: "Nguyen Van G", Email: "email7", Phone: "phone7", CardId: "card7"})
	//Omit các trường chỉ định
	//Tao nhieu doi tuong cung luc
	db.Model(&Student{}).Create(map[string]interface{}{
		"Id":       "8",
		"FullName": "Nguyen Van H",
		"Email":    "email8",
		"Phone":    "phone8",
		"CardId":   "card8",
	})
	tx.Rollback()
	//TODO: Select Object
	//lấy bảng ghi đầu tiên sau khi sắp xếp theo khóa chính
	var student Student
	fmt.Println("Lay ban ghi dau tien sau khi sap xep theo khoa chinh")
	db.First(&student)
	fmt.Println(student)
	//Lấy bảng ghi đầu tiên
	fmt.Println("Lay ban ghi dau tien")
	db.Take(&student)
	fmt.Println(student)
	//Tìm kiếm theo khóa chính
	fmt.Println("Tim kiem theo khoa chinh")
	db.First(&student, "Id = ?", "2")
	db.First(&student, "3")
	fmt.Println(student)
	//Tìm kiếm theo điều kiện
	fmt.Println("Tim kiem theo dieu kien")
	db.First(&student, "FullName = ?", "Nguyen Van A")
	fmt.Println(student)
	//Tìm kiếm theo danh sách khóa chính
	fmt.Println("Tim kiem theo danh sach khoa chinh")
	db.Find(&student, []string{"1", "2", "3"})
	//Lấy tất cả các bản ghi
	fmt.Println("Lay tat ca cac ban ghi")
	var students []Student
	db.Find(&students)
	fmt.Println(students)
	//Condition
	//db.Where("name = ?", "jinzhu").Find(&users)
	//TODO: Update
	//Khoi tao transaction
	data := db.Begin()
	var student1 Student
	//Lay ra danh sach ban ghi
	data.First(&student1)
	//thay doi thong tin cua doi tuong
	student1.FullName = "Nguyen Van A"
	student1.Email = "email1"
	//update lai thong tin
	data.Save(&student1)
	//Update 1 cot voi ten collum va gia tri moi
	data.Model(&student1).Update("FullName", "Nguyen Van B")
	//su dung condition
	data.Model(&student1).Where("FullName = ?", "Nguyen Van A").Update("FullName", "Nguyen Van C")
	data.Rollback()
	//TODO: Delete
	data1 := db.Begin()
	if err := data1.Error; err != nil {
		return
	}
	//delete
	student2 := Student{Id: "1"}
	data1.Delete(&student2)

	//delete theo dieu kien
	data1.Where("FullName = ?", "Nguyen Van A").Delete(&Student{})
	//xoa theo khoa chinh
	data1.Delete(&Student{}, "Id = ?", "2")
	//xoa theo danh sach khoa chinh
	data1.Delete(&Student{}, []string{"3", "4", "5"})
	data1.Rollback()
	//TODO: Viet raw query
	data2 := db.Begin()
	if err := data2.Error; err != nil {
		return
	}
	var student3 Student
	// lay danh sach studnet theo id
	data2.Raw("SELECT * FROM student WHERE Id = ?", "1").Scan(&student3)
	//Lay danh sach
	var students4 []Student
	data2.Raw("SELECT * FROM student").Scan(&students4)
	//Lay danh sach 2 nguoi dau tien
	data2.Raw("SELECT * FROM student LIMIT 2").Scan(&students4)
	//Update name cua user co id bang 1
	data2.Exec("UPDATE student SET FullName = ? WHERE Id = ?", "Nguyen Van A", "1")
	data2.Rollback()
	//TODO: Name arrument
	//data.Where("FullName = ?", "Nguyen Van A").Find(&Student{})
	data2.Where("FullName = :FullName", map[string]interface{}{"FullName": "Nguyen Van A"}).Find(&Student{})
	data2.Where("FullName = ?", "Nguyen Van A").Or("FullName = ?", "Nguyen Van B").Find(&Student{})
	data2.Where("FullName <> ?", "Nguyen Van A").Find(&Student{})
	data2.Where("FullName LIKE ?", "%Nguyen%").Find(&Student{})
	data2.Where("FullName IN (?)", []string{"Nguyen Van A", "Nguyen Van B"}).Find(&Student{})
	data2.Where("FullName  = @name", sql.Named("name", "Nguyen Van A")).Find(&Student{})
	//Multiple Name Argument
	data2.Where("FullName = @name1 OR FullName = @name2", sql.Named("name1", "Nguyen Van A"), sql.Named("name2", "Nguyen Van B")).Find(&Student{})
	data2.Where("FullName = @name1 OR FullName = @name2", map[string]interface{}{"name1": "Nguyen Van A", "name2": "Nguyen Van B"}).Find(&Student{})
	data2.Where("FullName = @name1 OR FullName = @name2", map[string]interface{}{"name1": "Nguyen Van A"}).Or("FullName = @name2", map[string]interface{}{"name2": "Nguyen Van B"}).Find(&Student{})
	data2.Where("FullName = @name1 OR FullName = @name2", map[string]interface{}{"name1": "Nguyen Van A", "name2": "Nguyen Van B"}).Find(&Student{})
	data2.Where(("FullName = @name1 and Email = @email"), sql.Named("name", "bob"), sql.Named("email", "thanh@gmail.com")).Find(&Student{})
	//Custom arrument
	data2.Raw("SELECT * FROM student WHERE FullName = ? AND Email = ?", custom{Name: "Nguyen Van A", Email: "email1"}).Scan(&Student{})
	data2.Raw("SELECT * FROM student WHERE FullName = :name AND Email = :email", map[string]interface{}{"name": "Nguyen Van A", "email": "email1"}).Scan(&Student{})
	data2.Raw("SELECT * FROM student WHERE FullName = :name AND Email = :email", custom{Name: "Nguyen Van A", Email: "email1"}).Scan(&Student{})
	data2.Raw("SELECT * FROM student WHERE FullName = @Eame and email = @Email", custom{Name: "Nguyen Van A", Email: "email1"}).Find(&Student{})
}

// Viet hook Update
/*
BeforeSave
BeforeUpdate
AfterUpdate
AfterSave
*/
func (s *Student) BeforeUpdate(tx *gorm.DB) (err error) {
	if s.Email == "" {
		s.Email = gofakeit.Email()
	}
	return
}

// Viet hook Create
func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	if s.Email == "" {
		s.Email = gofakeit.Email()
	}
	return
}

// viet hook delete
/*
BeforeDelete
AfterDelete
*/
func (s *Student) BeforeDelete(tx *gorm.DB) (err error) {
	if s.Id == "1" {
		return fmt.Errorf("Khong the xoa ban ghi nay")
	}
	return
}
