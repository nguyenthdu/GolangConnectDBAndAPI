package storage

import (
	"GolangDatabases/demo-todo-list/common"
	"GolangDatabases/demo-todo-list/modules/item/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.ToDoItem, error) {
	var data model.ToDoItem
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		//se co 2 truong hop loi: 1 la khong tim thay, 2 la loi db
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound // nhung cho khac tuong tu
		}

		return nil, common.ErrorDB(err)
	}
	return &data, nil

}
