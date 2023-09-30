package biz

import (
	"GolangDatabases/demo-todo-list/common"
	"GolangDatabases/demo-todo-list/modules/item/model"
	"context"
)

// lay du lieu item ra truoc khi update
type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.ToDoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ToDoItemUpdate) error
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{store: store}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *model.ToDoItemUpdate) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	//se co 2 truong hop loi: 1 la khong tim thay, 2 la loi db
	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrorCannotGet(model.EntityName, err)
		}
		return common.ErrorCannotUpdate(model.EntityName, err)
	}
	if data.Status != nil && *data.Status == model.Deleted {
		//return model.ErrItemDeleted
		//boc loi
		return common.ErrorCannotUpdate(model.EntityName, model.ErrItemDeleted)
	}
	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		//return err

		return common.ErrorCannotUpdate(model.EntityName, err)

	}
	//Sau khi format lai loi phai thay doi storage, bo gin.H thay bang err
	return nil

}
