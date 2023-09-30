package biz

import (
	"GolangDatabases/demo-todo-list/modules/item/model"
	"context"
)

// khong can chi dinh cu the, lam tuong minh
// dung o dau khai bao o do
// auto implement
// truyen vao condition
type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.ToDoItem, error)
}

type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.ToDoItem, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil

}
