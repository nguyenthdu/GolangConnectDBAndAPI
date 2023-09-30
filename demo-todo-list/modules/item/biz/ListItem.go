package biz

import (
	"GolangDatabases/demo-todo-list/common"
	"GolangDatabases/demo-todo-list/modules/item/model"
	"context"
)

// khong can chi dinh cu the, lam tuong minh
// dung o dau khai bao o do
// auto implement
// truyen vao condition
type ListItemStorage interface {
	ListItem(ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]model.ToDoItem, error)
}

type listItemBiz struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *listItemBiz {
	return &listItemBiz{store: store}
}

func (biz *listItemBiz) ListItem(
	ctx context.Context,
	filter *model.Filter, // dung con tro de tranh copy du lieu
	paging *common.Paging) ([]model.ToDoItem, error) {
	data, err := biz.store.ListItem(ctx, filter, paging)
	//neu muon mapping doi tuong co the truyen them vao morekyes vi du neu muon lay them thong tin user "User
	if err != nil {
		return nil, err
	}
	return data, nil

}
