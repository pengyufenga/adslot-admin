package service

import (
	"adslot-admin/model"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//获取sdkVersion信息详情
func GetSdkVersion(c *gin.Context) {
	var sdkVersion model.SdkVersion

	//获取传参
	param := c.Param("id")
	id := com.StrTo(param).MustInt()
	sdkVersion.Id = id

	mng, err := sdkVersion.GetSdkVersion()
	if err != nil {
		log.Printf("获取sdkVersion详情失败,err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("获取sdkVersion详情失败", err))
	} else {
		c.JSON(http.StatusOK, model.SuccessResult(mng))
	}
}

//获取所有
func GetAllSdkVersion(c *gin.Context) {
	var sdkVersion model.SdkVersion
	var pageParam model.PageParam
	//获取传参
	err := c.ShouldBindJSON(&pageParam)
	if err == nil {
		total, err := sdkVersion.CountOfSdkVersion()
		if err == nil {
			pageParam.Total = total
			list, err := sdkVersion.GetAllSdkVersion(pageParam.PageSize, pageParam.PageNum)
			if err == nil {
				data := make(map[string]interface{})
				data["sdkVersions"] = list
				data["pageParam"] = pageParam
				c.JSON(http.StatusOK, model.SuccessResult(data))
			} else {
				c.JSON(http.StatusBadRequest, model.FailResult("获取sdkVersion列表出错", err))
			}
		} else {
			c.JSON(http.StatusBadRequest, model.FailResult("获取sdkVersion条数出错", err))
		}
	} else {
		c.JSON(http.StatusBadRequest, model.FailResult("请求参数问题", err))
	}

}

//新增
func AddSdkVersion(c *gin.Context) {
	//当前时间
	now := time.Now()

	var sdkVersion model.SdkVersion
	//获取传参
	err := c.ShouldBindJSON(&sdkVersion)
	if err != nil {
		log.Printf("新增sdkVersion失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		sdkVersion.CreateTime = now
		log.Printf("新增sdkVersion:[%+v]", sdkVersion)
		err := sdkVersion.AddSdkVersion()
		if err != nil {
			log.Printf("新增sdkVersion失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("新增sdkVersion失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(sdkVersion))
		}
	}
}

//更新sdkJs信息
func UpdateSdkVersion(c *gin.Context) {
	var sdkVersion model.SdkVersion
	//获取参数
	err := c.ShouldBindJSON(&sdkVersion)
	if err != nil {
		log.Printf("更新sdkVersion失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		log.Printf("更新sdkVersion:[%+v]", sdkVersion)
		err := sdkVersion.UpdateSdkVersion()
		if err != nil {
			log.Printf("更新sdkVersion失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("更新sdkVersion失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(sdkVersion))
		}
	}

}

//删除sdkJs信息
func DeleteSdkVersion(c *gin.Context) {
	var sdkJs model.SdkJs
	//获取参数
	param := c.Param("id")
	id := com.StrTo(param).MustInt()
	sdkJs.Id = id
	//查找该条记录是否存在
	_, err := sdkJs.GetSdkJs()
	if err != nil {
		log.Printf("删除sdkJs，该sdkJs不存在，sdkJsId:[%v]", id)
		c.JSON(http.StatusBadRequest, model.FailResult("未找到该sdkJs", err))
	} else {
		//存在则删除
		err := sdkJs.DeleteSdkJs()
		if err != nil {
			log.Printf("删除sdkJs,失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("删除sdkJs失败", err))
		} else {
			log.Printf("删除sdkJs成功,SdkJsId:[%v]", id)
			c.JSON(http.StatusOK, model.SuccessResult(id))
		}
	}
}
