package service

import (
	"adslot-admin/model"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

//获取sdk-channel详情
func GetSdkChannel(c *gin.Context) {
	var sdkChannel model.SdkChannel

	//获取传参
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	sdkChannel.Id = id

	mng, err := sdkChannel.GetSdkChannel()
	if err != nil {
		log.Printf("获取sdkChannel详情失败,err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("获取sdkChannel详情失败", err))
	} else {
		c.JSON(http.StatusOK, model.SuccessResult(mng))
	}
}

//获取所有
func GetAllSdkChannel(c *gin.Context) {
	var sdkChannel model.SdkChannel
	var pageParam model.PageParam
	//获取传参
	err := c.ShouldBindJSON(&pageParam)
	if err == nil {
		total, err := sdkChannel.CountOfSdkChannel()
		if err == nil {
			pageParam.Total = total
			sdkChannels, err := sdkChannel.GetAllSdkChannel(pageParam.PageSize, pageParam.PageNum)
			if err == nil {
				data := make(map[string]interface{})
				data["sdkChannels"] = sdkChannels
				data["pageParam"] = pageParam
				c.JSON(http.StatusOK, model.SuccessResult(data))
			} else {
				c.JSON(http.StatusBadRequest, model.FailResult("获取sdkChannel列表出错", err))
			}
		} else {
			c.JSON(http.StatusBadRequest, model.FailResult("获取sdkChannel条数出错", err))
		}
	} else {
		c.JSON(http.StatusBadRequest, model.FailResult("请求参数问题", err))
	}

}

//新增
func AddSdkChannel(c *gin.Context) {
	//当前时间
	now := time.Now()

	var sdkChannel model.SdkChannel
	//获取传参
	err := c.ShouldBindJSON(&sdkChannel)
	if err != nil {
		log.Printf("新增sdkChannel失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		sdkChannel.CreateTime = now
		log.Printf("新增sdkChannel:[%+v]", sdkChannel)
		err := sdkChannel.AddSdkChannel()
		if err != nil {
			log.Printf("新增sdkChannel失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("新增sdkChannel失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(sdkChannel))
		}
	}
}

//更新dsp配置信息
func UpdateSdkChannel(c *gin.Context) {
	var sdkChannel model.SdkChannel
	//获取参数
	err := c.ShouldBindJSON(&sdkChannel)
	if err != nil {
		log.Printf("更新sdkChannel失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		log.Printf("更新sdkChannel:[%+v]", sdkChannel)
		err := sdkChannel.UpdateSdkChannel()
		if err != nil {
			log.Printf("更新sdkChannel失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("更新sdkChannel失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(sdkChannel))
		}
	}

}

//删除dsp配置信息
func DeleteSdkChannel(c *gin.Context) {
	var sdkChannel model.SdkChannel
	//获取参数
	param := c.Param("id")
	id := com.StrTo(param).MustInt()
	sdkChannel.Id = id
	//查找该条记录是否存在
	_, err := sdkChannel.GetSdkChannel()
	if err != nil {
		log.Printf("删除sdkChannel，该sdkChannel不存在，dspmngId:[%v]", id)
		c.JSON(http.StatusBadRequest, model.FailResult("未找到该sdkChannel", err))
	} else {
		//存在则删除
		err := sdkChannel.DeleteSdkChannel()
		if err != nil {
			log.Printf("删除sdkChannel,失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("删除sdkChannel失败", err))
		} else {
			log.Printf("删除sdkChannel成功,AdslotId:[%v]", id)
			c.JSON(http.StatusOK, model.SuccessResult(id))
		}
	}
}
