package gin

import (
	"GolangDatabases/demo-todo-list/common"
	"GolangDatabases/demo-todo-list/modules/item/biz"
	"GolangDatabases/demo-todo-list/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetItemById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		store := storage.NewSQLStore(db)
		busines := biz.NewGetItemBiz(store)
		data, err := busines.GetItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		//khi tim kiem 1 dong thi dung fist, nhieu dong thi dung find
		//db.First(&data)
		c.JSON(http.StatusOK, common.SimpleSuccessRespone(data))
	}
}
