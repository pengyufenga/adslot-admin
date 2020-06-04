package service

import (
	"adslot-admin/model"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//获取sdkJs信息详情
func GetSdkJs(c *gin.Context) {
	var sdkJs model.SdkJs

	//获取传参
	param := c.Param("id")
	id := com.StrTo(param).MustInt()
	sdkJs.Id = id

	mng, err := sdkJs.GetSdkJs()
	if err != nil {
		log.Printf("获取sdkJs详情失败,err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("获取sdkJs详情失败", err))
	} else {
		c.JSON(http.StatusOK, model.SuccessResult(mng))
	}
}

//获取所有
func GetAllSdkJs(c *gin.Context) {
	var sdkJs model.SdkJs
	var pageParam model.PageParam
	//获取传参
	err := c.ShouldBindJSON(&pageParam)
	if err == nil {
		total, err := sdkJs.CountOfSdkJs()
		if err == nil {
			pageParam.Total = total
			list, err := sdkJs.GetAllSdkJs(pageParam.PageSize, pageParam.PageNum)
			if err == nil {
				data := make(map[string]interface{})
				data["sdkJss"] = list
				data["pageParam"] = pageParam
				c.JSON(http.StatusOK, model.SuccessResult(data))
			} else {
				c.JSON(http.StatusBadRequest, model.FailResult("获取sdkJs列表出错", err))
			}
		} else {
			c.JSON(http.StatusBadRequest, model.FailResult("获取sdkJs条数出错", err))
		}
	} else {
		c.JSON(http.StatusBadRequest, model.FailResult("请求参数问题", err))
	}

}

//新增
func AddSdkJs(c *gin.Context) {
	//当前时间
	now := time.Now()

	var sdkJs model.SdkJs
	//获取传参
	err := c.ShouldBindJSON(&sdkJs)
	if err != nil {
		log.Printf("新增sdkJs失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		sdkJs.CreateTime = now
		sdkJs.ModifyTime = now
		log.Printf("新增sdkJs:[%+v]", sdkJs)
		err := sdkJs.AddSdkJs()
		if err != nil {
			log.Printf("新增sdkJs失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("新增sdkJs失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(sdkJs))
		}
	}
}

//更新sdkJs信息
func UpdateSdkJs(c *gin.Context) {
	now := time.Now()

	var sdkJs model.SdkJs
	//获取参数
	err := c.ShouldBindJSON(&sdkJs)
	if err != nil {
		log.Printf("更新sdkJs失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		sdkJs.ModifyTime = now
		log.Printf("更新sdkJs:[%+v]", sdkJs)
		err := sdkJs.UpdateSdkJs()
		if err != nil {
			log.Printf("更新sdkJs失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("更新sdkJs失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(sdkJs))
		}
	}

}

//删除sdkJs信息
func DeleteSdkJs(c *gin.Context) {
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
