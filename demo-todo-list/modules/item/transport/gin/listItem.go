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

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var page common.Paging
		if err := c.ShouldBind(&page); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		page.Process()

		// them filter
		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		store := storage.NewSQLStore(db)
		business := biz.NewListItemBiz(store)
		result, err := business.ListItem(c.Request.Context(), &filter, &page)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessRespone(result, page, filter))

	}
}
