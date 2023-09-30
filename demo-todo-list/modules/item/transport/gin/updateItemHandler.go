package gin

import (
	"GolangDatabases/demo-todo-list/common"
	"GolangDatabases/demo-todo-list/modules/item/biz"
	"GolangDatabases/demo-todo-list/modules/item/model"
	"GolangDatabases/demo-todo-list/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateAItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ToDoItemUpdate
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// nhung loi co ShouldBind nen dung invalid
		if err := c.ShouldBind(&data); err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.JSON(http.StatusBadRequest, common.ErrorValidate(err))
			//TODO: tuong tu o cac cho khac
			return
		}
		store := storage.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(store)
		if err := business.UpdateItemById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true))

	}

}
