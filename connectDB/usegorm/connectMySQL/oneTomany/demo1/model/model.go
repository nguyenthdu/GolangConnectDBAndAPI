package model

type Foo struct {
	Id   string `gorm:"primary_key"`
	Name string
	Bars []Bar `gorm:"foreignKey:FooId"`
}
type Bar struct {
	Id    string `gorm:"primary_key"`
	Name  string
	FooId string `gorm:"column:foo_id"`
	Foo   Foo    `gorm:"foreignKey:FooId"`
}

// doi ten bang
func (f *Foo) TableName() string {
	return "foo"

}
func (b *Bar) TableName() string {
	return "bar"

}

type Tabler interface {
	TableName() string
}
