package model

import (
	"GolangDatabases/demo-todo-list/common"
	"errors"
)

const EntityName = "Item"

type ToDoItem struct {
	common.SQLModule             //embed struct
	Title            string      `json:"tile"`
	Description      string      `json:"description"`
	Status           *ItemStatus `json:"status"`
	//them tuy chon ,omitempty bo di cac gia tri null: 0, string null, flase, * null
}

// Tao table

// Khi tao thi khong can phai truyen het du lieu vao nen khai bao 1 func con
type ToDoItemCreate struct {
	Id          int         `json:"-" gorm:"column:id;"`
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

// update item
type ToDoItemUpdate struct {
	// dung string thi mac dinh neu du lieu truyen vao rong thi gorm se khong thay doi du lieu
	// neu muon thay doi du lieu khi truyen rong thi phai khi bao *string
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

// gioi han lai so item co the hien thi trong 1 trang

// create item
func (ToDoItemCreate) TableName() string {
	return "todo_items" // return ve ten bang chua du lieu trong csdl

}

// get a item
func (ToDoItem) TableName() string {
	return "todo_items"
}

// update a tiem
func (ToDoItemUpdate) TableName() string {
	return "todo_items"

}

// create ErrTitleIsBlank
var (
	ErrTitleIsBlank = errors.New("Title is blank")

	ErrItemDeleted = errors.New("Item is deleted")
	//Cach toi nhat nen khai tung loi cu the trong model
	ErrorItemNotFound = common.NewCustomError(errors.New("Item not found"), "Item not found", "ItemNotFound")
)
