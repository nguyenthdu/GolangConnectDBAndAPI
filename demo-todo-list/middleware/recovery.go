package middleware

import (
	"GolangDatabases/demo-todo-list/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery() func(*gin.Context) {
	return func(c *gin.Context) {
		//log.Println("recovery middleware")
		//cashes loi ra:
		defer func() {
			if r := recover(); r != nil {
				//c.JSON(500, gin.H{
				//	"message": "internal server error",
				//})
				//log.Println("recovery middleware", r)
				//su dung boc loi
				if err, ok := r.(error); ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrorInternal(err))
					//abort: dung lai chuong trinh
					//statusJSON: tra ve status va json theo body
				}
				//Neu muon xem lai loi  panic(r)
			}
		}()
		c.Next()
	}

}
