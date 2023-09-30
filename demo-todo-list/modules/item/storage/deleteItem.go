package storage

import (
	"GolangDatabases/demo-todo-list/modules/item/model"
	"context"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	// 1 cach khac
	deletedStatus := model.Deleted
	if err := s.db.Table("todo_items"). //(model.TodoItem{}.TableName).
						Where(cond).
						Updates(map[string]interface{}{
			//co the su dung cach nay
			//"status": "deleted",
			"status": deletedStatus.String(),
		}).Error; err != nil {
		return err
	}
	return nil
}
