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
	r.GET("/media/list", service.GetAllMedia)
	r.POST("/media/add", service.InsertMedia)
	r.PUT("/media/update", service.UpdateMedia)
	r.DELETE("/media/delete/:id", service.DeleteMedia)

	//广告位模块路由
	r.GET("/adslot/detail/:id", service.GetAdslot)
	r.POST("/adslot/list", service.GetAllAdslot)
	r.POST("/adslot/add", service.AddAdslot)
	r.PUT("/adslot/update", service.UpdateAdslot)
	r.DELETE("/adslot/delete/:id", service.DeleteAdslot)

	//dsp配置信息模块路由
	r.GET("/dspmng/detail/:id", service.GetDspMng)
	r.POST("/dspmng/list", service.GetAllDspMng)
	r.POST("/dspmng/add", service.AddDspMng)
	r.PUT("/dspmng/update", service.UpdateDspMng)
	r.DELETE("/dspmng/delete/:id", service.DeleteDspMng)

	//dsp_source模块路由
	r.GET("/dspsource/detail/:id", service.GetDspSource)
	r.POST("/dspsource/list", service.GetAllDspSource)
	r.POST("/dspsource/add", service.AddDspSource)
	r.PUT("/dspsource/update", service.UpdateDspSource)
	r.DELETE("/dspsource/delete/:id", service.DeleteDspSource)

	//sdk_channel模块路由
	r.GET("/sdk-channel/detail/:id", service.GetSdkChannel)
	r.POST("/sdk-channel/list", service.GetAllSdkChannel)
	r.POST("/sdk-channel/add", service.AddSdkChannel)
	r.PUT("/sdk-channel/update", service.UpdateSdkChannel)
	r.DELETE("/sdk-channel/delete/:id", service.DeleteSdkChannel)

	//sdk_js模块路由
	r.GET("/sdk-jsscript/detail/:id", service.GetSdkJs)
	r.POST("/sdk-jsscript/list", service.GetAllSdkJs)
	r.POST("/sdk-jsscript/add", service.AddSdkJs)
	r.PUT("/sdk-jsscript/update", service.UpdateSdkJs)
	r.DELETE("/sdk-jsscript/delete/:id", service.DeleteSdkJs)

	//sdk_version模块路由
	r.GET("/sdk-version/detail/:id", service.GetSdkVersion)
	r.POST("/sdk-version/list", service.GetAllSdkVersion)
	r.POST("/sdk-version/add", service.AddSdkVersion)
	r.PUT("/sdk-version/update", service.UpdateSdkVersion)
	r.DELETE("/sdk-version/delete/:id", service.DeleteSdkVersion)

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello")
	})
	return r
}
