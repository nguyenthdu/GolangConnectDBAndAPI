package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// vi status chi co 3 trang thai nen se khai bao enum
type ItemStatus int

const (
	Doing ItemStatus = iota
	Done
	Deleted
)

// khai bao vien luu tru cac trang thai status
var StatusTypes = [3]string{"Doing", "Done", "Deleted"}

// do ItemsStatus dang la kieu int nen se chuyen ve kieu string
// do la ItemStatus tu 0-2 va mang StatusTypes cung chay tu 0-2 nen co the return ve ngay lap tuc
func (item *ItemStatus) String() string {
	return StatusTypes[*item]

}

// chuyen kieu int ve ItemStatus
func parseStrItemStatus(s string) (ItemStatus, error) {
	// tim ItemStatus phu hop voi string
	for i := range StatusTypes {
		if StatusTypes[i] == s {
			return ItemStatus(i), nil
		}

	}

	return ItemStatus(0), errors.New("Invalid status string")
}

// vi item status dang la kieu int nen gorm se khong gan duoc
// dung de cho du lieu tu ben duoi db len khoi voi du lieu ItemStatus
func (item *ItemStatus) Scan(value interface{}) error {
	//dua value ve mang byte
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Fail to scan data from sql", value))
	}
	v, err := parseStrItemStatus(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintf("Fail to scan data from sql ", value))
	}
	*item = v
	return nil
}

// Lay du lieu tu item status truyen xuong du leun ben duoi
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil

}

// decode
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	//do khi khi chuyen thi se lay ca ky tu " cua "Doing" nen can loai bo no
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStrItemStatus(str)
	if err != nil {
		return nil
	}
	*item = itemValue
	return nil

}

// de status co the in ra chuoi string can implement MarshalJSON
// Chuyen ItemStatus sang chuoi string
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

// check status  == 'deleted'
func (item *ItemStatus) ItemStatusDeleted() bool {
	return *item == Deleted
}
