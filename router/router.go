package router

import (
	"adslot-admin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//媒体模块路由
	r.GET("/media/detail/:id", service.GetMedia)
	r.GET("/media/list",service.GetAllMedia)
	r.POST("/media/add",service.InsertMedia)
	r.PUT("/media/update",service.UpdateMedia)
	r.DELETE("/media/delete/:id",service.DeleteMedia)

	//广告位模块路由
	r.GET("/adslot/detail/:id",service.GetAdslot)
	r.POST("/adslot/list",service.GetAllAdslot)
	r.POST("/adslot/add",service.AddAdslot)

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello")
	})
	return r
}
