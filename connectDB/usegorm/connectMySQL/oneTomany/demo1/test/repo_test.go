package test

import (
	"GolangDatabases/connectDB/usegorm/connectMySQL/oneTomany/demo1/repo"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetFooBar(t *testing.T) {
	foos, err := repo.GetFooBar()
	assert.Nil(t, err) //Neu co loi thi fail
	for _, foo := range foos {
		fmt.Println(foo.Name)
		for _, bar := range foo.Bars {
			fmt.Println(bar.Name)
		}
	}
	//Positivef: kiem tra xem mot gia tri co la so duong hay khong
	assert.Positivef(t, len(foos), "foos should be positive") //Neu khong co du lieu thi fail
	/*
			assert.Positivef(t *testing.T, value interface{}, format string, args ...interface{}) bool
		t *testing.T: Đối tượng kiểm tra, được sử dụng để ghi nhận kết quả của kiểm tra. Nó được đưa vào để báo cáo kết quả.

		value interface{}: Giá trị cần kiểm tra. Trong trường hợp này, nó là độ dài của slice foos.

		format string: Một chuỗi định dạng cho thông báo lỗi (nếu kiểm tra fail).

		args ...interface{}: Các đối số (nếu có) được sử dụng trong chuỗi định dạng.

		Phương thức assert.Positivef được sử dụng để kiểm tra xem một giá trị có là số dương hay không.
			Nếu giá trị là số dương, kiểm tra sẽ pass. N
			gược lại, nếu giá trị không là số dương, kiểm tra sẽ fail và thông báo lỗi sẽ được in ra.

	*/
}

// Test_GetFooById: kiem tra lay du lieu cua bang con trong bang cha
func Test_GetFooById(t *testing.T) {
	foo, err := repo.GetFooById("ox-01")
	assert.Nil(t, err)
	fmt.Println(foo.Name)
	for _, bar := range foo.Bars {
		fmt.Println(bar.Name)
	}

}

// Test GetBar by id
func Test_GetBarById(t *testing.T) {
	bar, err := repo.GetBarById("bar1")
	assert.Nil(t, err)
	fmt.Println(bar.FooId)
	fmt.Println("" + bar.Id)
	fmt.Println("" + bar.Name)

}

// test createData
func Test_CreateData(t *testing.T) {
	err := repo.CreateData()
	assert.Nil(t, err)

}
