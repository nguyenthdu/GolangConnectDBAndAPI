package main

import (
	"GolangDatabases/demo-todo-list/middleware"
	ginItem "GolangDatabases/demo-todo-list/modules/item/transport/gin"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

//
//// vi status chi co 3 trang thai nen se khai bao enum
//type ItemStatus int
//
//const (
//	Doing ItemStatus = iota
//	Done
//	Deleted
//)
//
//// khai bao vien luu tru cac trang thai status
//var StatusTypes = [3]string{"Doing", "Done", "Deleted"}
//
//// do ItemsStatus dang la kieu int nen se chuyen ve kieu string
//// do la ItemStatus tu 0-2 va mang StatusTypes cung chay tu 0-2 nen co the return ve ngay lap tuc
//func (item *ItemStatus) String() string {
//	return StatusTypes[*item]
//
//}
//
//// chuyen kieu int ve ItemStatus
//func parseStrItemStatus(s string) (ItemStatus, error) {
//	// tim ItemStatus phu hop voi string
//	for i := range StatusTypes {
//		if StatusTypes[i] == s {
//			return ItemStatus(i), nil
//		}
//
//	}
//
//	return ItemStatus(0), errors.New("Invalid status string")
//}
//
//// vi item status dang la kieu int nen gorm se khong gan duoc
//// dung de cho du lieu tu ben duoi db len khoi voi du lieu ItemStatus
//func (item *ItemStatus) Scan(value interface{}) error {
//	//dua value ve mang byte
//	bytes, ok := value.([]byte)
//	if !ok {
//		return errors.New(fmt.Sprintf("Fail to scan data from sql", value))
//	}
//	v, err := parseStrItemStatus(string(bytes))
//	if err != nil {
//		return errors.New(fmt.Sprintf("Fail to scan data from sql ", value))
//	}
//	*item = v
//	return nil
//}
//
//// Lay du lieu tu item status truyen xuong du leun ben duoi
//func (item *ItemStatus) Value() (driver.Value, error) {
//	if item == nil {
//		return nil, nil
//	}
//	return item.String(), nil
//
//}
//
//// decode
//func (item *ItemStatus) UnmarshalJSON(data []byte) error {
//	//do khi khi chuyen thi se lay ca ky tu " cua "Doing" nen can loai bo no
//	str := strings.ReplaceAll(string(data), "\"", "")
//	itemValue, err := parseStrItemStatus(str)
//	if err != nil {
//		return nil
//	}
//	*item = itemValue
//	return nil
//
//}
//
//// de status co the in ra chuoi string can implement MarshalJSON
//// Chuyen ItemStatus sang chuoi string
//func (item *ItemStatus) MarshalJSON() ([]byte, error) {
//	if item == nil {
//		return nil, nil
//	}
//	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
//
//}
//
//type ToDoItem struct {
//	common.SQLModule             //embed struct
//	Title            string      `json:"tile"`
//	Description      string      `json:"description"`
//	Status           *ItemStatus `json:"status"`
//	//them tuy chon ,omitempty bo di cac gia tri null: 0, string null, flase, * null
//}
//
//// Tao table
//
//// Khi tao thi khong can phai truyen het du lieu vao nen khai bao 1 func con
//type ToDoItemCreate struct {
//	Id          int         `json:"-" gorm:"column:id;"`
//	Title       string      `json:"title" gorm:"column:title;"`
//	Description string      `json:"description" gorm:"column:description;"`
//	Status      *ItemStatus `json:"status" gorm:"column:status;"`
//}
//
//// update item
//type ToDoItemUpdate struct {
//	// dung string thi mac dinh neu du lieu truyen vao rong thi gorm se khong thay doi du lieu
//	// neu muon thay doi du lieu khi truyen rong thi phai khi bao *string
//	Title       string `json:"title" gorm:"column:title;"`
//	Description string `json:"description" gorm:"column:description;"`
//	Status      string `json:"status" gorm:"column:status;"`
//}
//
//// gioi han lai so item co the hien thi trong 1 trang
//
//// create item
//func (ToDoItemCreate) TableName() string {
//	return "todo_items" // return ve ten bang chua du lieu trong csdl
//
//}
//
//// get a item
//func (ToDoItem) TableName() string {
//	return "todo_items"
//}
//
//// update a tiem
//func (ToDoItemUpdate) TableName() string {
//	return "todo_items"
//
//}

func main() {

	dsn := "root:password@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(db)
	//now := time.Now().UTC()
	//item := ToDoItem{
	//	Id:           1,
	//	Title:        "Xin chao",
	//	Descriptions: "Demo string",
	//	Statuss:      "Doing",
	//	CreateAt:     &now,
	//	UpdateAt:     nil,
	//}
	////jsonData, err := json.Marshal(item)
	////if err != nil {
	////	panic(err)
	////}
	////fmt.Println(string(jsonData))
	//////conver json to byte
	////var item2 ToDoItem
	////jsonStr := "{\"id\":1,\"tile\":\"Xin chao\",\"descriptions\":\"Demo string\",\"status\":\"Doing\",\"create_at\":\"2023-09-23T09:09:39.1451823Z\",\"update_at\":null}\n"
	////if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
	////	panic(err)
	////}
	////fmt.Println(item2)
	r := gin.Default()
	//su dung middleware
	//cach 1: tac dong vao tat ca cac api
	r.Use(middleware.Recovery())
	//cach 2: tac dong vao 1 group api
	v1 := r.Group("v1/") //co the them nhieu middleware
	//v1 := r.Group("v1/", middleware.Recovery()) //co the them nhieu middleware
	{
		item := v1.Group("/items")
		{
			//Cach 3: tac dong vao tung api
			//Create item
			item.POST("", ginItem.CreateNewItem(db))
			//item.POST("", middleware.Recovery(), ginItem.CreateNewItem(db))
			//Get all item
			item.GET("", ginItem.ListItem(db))
			//Get a item by id
			item.GET("/:id", ginItem.GetItemById(db))
			//Update a item
			item.PUT("/:id", ginItem.UpdateAItem(db))
			//Delete a item
			item.DELETE("/:id", ginItem.DeleteAItem(db))
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println([]int{}[0]) // chuong trinh se bi loi va cashes o day=> tao ra panic
		// nghia la Recover mac dinh cua gin bat loi
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//func CreateItem(db *gorm.DB) func(*gin.Context) {
//	return func(c *gin.Context) {
//		var data ToDoItemCreate //Lay thong tin struct
//		//ham JSON se duoc hoi tai ShouldBnd neu co viet
//		if err := c.ShouldBind(&data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"message: ": err.Error(),
//			})
//			return
//		}
//		//insert in to database
//		//Create se goi ham value
//		if err := db.Create(&data).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		c.JSON(http.StatusOK, common.SingleSuccessRespone(data.Id))
//	}
//}
//func GetAItem(db *gorm.DB) func(*gin.Context) {
//	return func(c *gin.Context) {
//
//		var data ToDoItem
//		id, err := strconv.Atoi(c.Param("id"))
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error()})
//			return
//		}
//		//data.Id = id
//		//cach 2 tuong minh hon
//		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		//khi tim kiem 1 dong thi dung fist, nhieu dong thi dung find
//		//db.First(&data)
//		c.JSON(http.StatusOK, common.SingleSuccessRespone(data))
//	}
//}
//func UpdateAItem(db *gorm.DB) func(*gin.Context) {
//	return func(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		var dataItem ToDoItemUpdate
//
//		if err := c.ShouldBind(&dataItem); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//			return
//		}
//
//		c.JSON(http.StatusOK, common.SingleSuccessRespone("Update success"))
//	}
//
//}
//func DeleteAItem(db *gorm.DB) func(*gin.Context) {
//	return func(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		//if err := db.Table("todo_items").Where("id=?", id).Delete(nil).Error; err != nil {
//		//short delete
//		if err := db.Table("todo_items").Where("id=?", id).Updates(map[string]interface{}{
//			"status": "deleted",
//		}).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		c.JSON(http.StatusOK, common.SingleSuccessRespone("Delete success"))
//	}
//
//}
//func getAllItems(db *gorm.DB) func(ctx *gin.Context) {
//	return func(c *gin.Context) {
//		var page common.Paging
//		if err := c.ShouldBind(&page); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//
//			return
//		}
//
//		page.Process()
//		var result []ToDoItem
//		//Loai bo nhung truong co status la deleted nhu o phan deleted da lam
//		//db = db.Where("status <> ?", "deleted")
//		if err := db.Table("todo_items").Count(&page.Total).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//		}
//		//them order vao de sap xep ket qua tim kiem theo id
//		if err := db.Order("id desc").
//			Offset((page.Page - 1) * page.Limit). //offset lấy kết quả truy vấn từ số bảng ghi. Nếu Offset(10) có nghĩa là lấy từ bảng ghi số 10
//			Limit(page.Limit).                    //giới hạn số kết nối truy vấn. ví dụ 10 thì sẽ là 10 item được hiển thị
//			Find(&result).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//
//		c.JSON(http.StatusOK, common.NewSuccessRespone(result, page, nil))
//
//	}
//}
