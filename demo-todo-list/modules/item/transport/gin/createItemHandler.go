package gin

import (
	"GolangDatabases/demo-todo-list/common"
	"GolangDatabases/demo-todo-list/modules/item/biz"
	"GolangDatabases/demo-todo-list/modules/item/model"
	"GolangDatabases/demo-todo-list/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateNewItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.ToDoItemCreate
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)
		if err := business.CreateNewItem(ctx.Request.Context(), &data); err != nil {
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			//format lai loi
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(data.Id))
	}
}
