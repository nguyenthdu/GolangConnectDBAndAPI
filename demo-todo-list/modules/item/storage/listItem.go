package storage

import (
	"GolangDatabases/demo-todo-list/common"
	"GolangDatabases/demo-todo-list/modules/item/model"
	"context"
)

func (s *sqlStore) ListItem(ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]model.ToDoItem, error) {
	var result []model.ToDoItem
	//		//Loai bo nhung truong co status la deleted nhu o phan deleted da lam
	//db = db.Where("status <> ?", "deleted")
	db := s.db.Where("status <> ?", "Deleted")
	// kiem tra neu co filter thi moi dung
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}

	}
	if err := db.Table("todo_items").
		Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil

}
