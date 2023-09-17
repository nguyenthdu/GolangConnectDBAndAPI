package repo

import (
	"GolangDatabases/connectDB/usegorm/connectMySQL/oneTomany/demo1/model"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"math/rand"
	"time"
)

// thong so ket noi
var (
	host     string = "localhost"
	port     int    = 3306
	username string = "root"
	password string = "password"
	dbname   string = "demoGolang1"
)
var (
	DB     *gorm.DB   //ket noi database
	random *rand.Rand //doi tuong dung de tao random number
)

func init() {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	var err error
	DB, err = gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //log cac cau lenh sql
	})
	if err != nil {
		panic("failed to connect database") //ket noi that bai
	}

	//khoi dong engine random- so ngau nhien dua tren thoi gian hien tai
	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)

}

// Get Foo and Bar
func GetFooBar() (foos []model.Foo, err error) {
	//Preload: lay du lieu cua bang con trong bang cha
	if err := DB.Preload("Bars").Find(&foos).Error; err != nil {
		return nil, err
	}
	return foos, err
}

// Get Foo and Bar by id
func GetFooById(id string) (foo model.Foo, err error) {
	if err := DB.Preload("Bars").Find(&foo, "foo.id = ?", id).Error; err != nil {
		return model.Foo{}, err
	}
	return foo, nil

}

// Get info bar by id
func GetBarById(id string) (bar model.Bar, err error) {
	bar = model.Bar{
		Id: id,
	}
	if err = DB.Preload("Foo").Find(&bar).Error; err != nil {
		return model.Bar{}, err
	}
	return bar, nil

}

// Mockup data
// auto insert in database
func CreateData() (err error) {
	//create array save
	var foos []model.Foo
	//loop save object foo
	for i := 0; i < 5; i++ {
		foo_id := NewID() //ham random id
		foo := model.Foo{
			Id:   foo_id,
			Name: gofakeit.Animal(), //fake name
		}
		// voi moi doi tung foo -> tao 1 so doi tuong bar tuong ung
		for j := 0; j < 2+random.Intn(2); j++ {
			bar := model.Bar{
				Id:    NewID(),
				Name:  gofakeit.Animal(),
				FooId: foo_id,
			}
			foo.Bars = append(foo.Bars, bar)
		}
		foos = append(foos, foo)

	}
	// Insert database
	if err := DB.Create(&foos).Error; err != nil {
		return err
	}
	return nil

}
