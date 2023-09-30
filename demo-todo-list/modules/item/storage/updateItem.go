package storage

import (
	"GolangDatabases/demo-todo-list/modules/item/model"
	"context"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ToDoItemUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err

	}
	return nil

}
