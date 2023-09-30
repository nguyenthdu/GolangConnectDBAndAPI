package biz

import (
	"GolangDatabases/demo-todo-list/common"
	"GolangDatabases/demo-todo-list/modules/item/model"
	"context"
	"strings"
)

// khong can chi dinh cu the, lam tuong minh
// dung o dau khai bao o do
// auto implement
type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.ToDoItemCreate) error
}
type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.ToDoItemCreate) error {
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTitleIsBlank
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		//return err
		//format lai loi
		return common.ErrorCannotCreate(model.EntityName, err)
	}
	return nil

}
