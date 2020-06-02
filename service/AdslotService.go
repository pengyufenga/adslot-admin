package service

import (
	"adslot-admin/model"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//获取广告位信息
func GetAdslot(c *gin.Context) {
	var adslot model.Adslot
	//获取传参
	queryId := c.Param("id")
	adslot.Id = com.StrTo(queryId).MustInt64()

	detail, err := adslot.GetAdslot()
	if err != nil {
		log.Printf("获取Adslot详情失败,err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.Result{http.StatusBadRequest, "获取Adslot详情失败", err.Error()})
	} else {
		result := model.Result{
			Code:    http.StatusOK,
			Message: "请求成功",
			Data:    detail,
		}
		c.JSON(http.StatusOK, result)
	}
}

//获取广告位信息列表
func GetAllAdslot(c *gin.Context) {
	var adslotDao model.Adslot
	var pageParam model.PageParam
	//获取传参
	err := c.ShouldBindJSON(&pageParam)
	if err == nil {
		total, err := adslotDao.CountOfAdslot()
		if err == nil {
			pageParam.Total = total
			adslots, err := adslotDao.GetAllAdslot(pageParam.PageSize, pageParam.PageNum)
			if err == nil {
				data := make(map[string]interface{})
				data["adslots"] = adslots
				data["pageParam"] = pageParam
				c.JSON(http.StatusOK, model.SuccessResult(data))
			} else {
				c.JSON(http.StatusBadRequest, model.FailResult("获取adslot列表出错", err))
			}
		} else {
			c.JSON(http.StatusBadRequest, model.FailResult("获取adslot条数出错", err))
		}
	} else {
		c.JSON(http.StatusBadRequest, model.FailResult("请求参数问题", err))
	}
}

//新增广告位信息
func AddAdslot(c *gin.Context)  {
	//当前时间
	now := time.Now()

	var adslotDao model.Adslot
	adslotDao.CreateTime = now
	adslotDao.ModifyTime = now
	//获取传参
	err := c.ShouldBindJSON(&adslotDao)
	if err != nil{
		c.JSON(http.StatusBadRequest,model.FailResult("参数绑定异常",err))
	}else {

	}
}
