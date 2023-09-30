package storage

import (
	"GolangDatabases/demo-todo-list/common"
	"GolangDatabases/demo-todo-list/modules/item/model"
	"context"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.ToDoItemCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		//return err // day la loi cua thu vien
		// "error": "Error 1054 (42S22): Unknown column 'titles' in 'field list'"
		//quang loi goc ra
		return common.ErrorDB(err)
	}
	return nil

}
