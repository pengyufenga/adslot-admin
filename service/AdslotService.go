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
func AddAdslot(c *gin.Context) {
	//当前时间
	now := time.Now()

	var adslot model.Adslot
	//获取传参
	err := c.ShouldBindJSON(&adslot)
	if err != nil {
		log.Printf("新增Adslot失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		adslot.CreateTime = now
		adslot.ModifyTime = now
		log.Printf("新增Adslot:[%v]", adslot)
		err := adslot.AddAdslot()
		if err != nil {
			log.Printf("新增Adslot失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("新增Adslot失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(adslot))
		}
	}
}

//更新广告位信息
func UpdateAdslot(c *gin.Context) {
	now := time.Now()

	var adslot model.Adslot
	//获取参数
	err := c.ShouldBindJSON(&adslot)
	if err != nil {
		log.Printf("更新Adslot失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		adslot.ModifyTime = now
		log.Printf("更新Adslot:[%+v]", adslot)
		err := adslot.UpdateAdslot()
		if err != nil {
			log.Printf("更新Adslot失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("更新Adslot失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(adslot))
		}
	}

}

//删除广告位信息
func DeleteAdslot(c *gin.Context) {
	var adslot model.Adslot
	//获取参数
	param := c.Param("id")
	id := com.StrTo(param).MustInt64()
	adslot.Id = id
	//查找该条记录是否存在
	_, err := adslot.GetAdslot()
	if err != nil {
		log.Printf("删除Adslot，该Adslot不存在，AdslotId:[%v]", id)
		c.JSON(http.StatusBadRequest, model.FailResult("未找到该Adslot", err))
	} else {
		//存在则删除
		err := adslot.DeleteAdslot()
		if err != nil {
			log.Printf("删除Adslot,失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("删除Adslot失败", err))
		} else {
			log.Printf("删除Adslot成功,AdslotId:[%v]", id)
			c.JSON(http.StatusOK, model.SuccessResult(id))
		}
	}

}
