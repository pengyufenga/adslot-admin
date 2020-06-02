package service

import (
	"adslot-admin/model"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Code    int
	Message string
	Data    interface{}
}

var  FailResult Result = Result{Code:2,Message:"错误",Data:nil}

//获取媒体信息
func GetMedia(c *gin.Context) {
	var m model.Media
	id := com.StrTo(c.Param("id")).MustInt64()
	m.Id = id
	media, err := m.GetMediaById()
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{2, "参数错误", err.Error()})
		return
	}
	result := Result{
		Code:    http.StatusOK,
		Message: "请求成功",
		Data:    media,
	}
	c.JSON(http.StatusOK, result)
}

//获取媒体列表
func GetAllMedia(c *gin.Context) {
	var mDao model.Media

	mediaList,err := mDao.GetAllMedia()
	if err != nil {
		c.JSON(http.StatusNoContent, Result{http.StatusBadRequest, "请求失败", err.Error()})
		return
	}

	result := Result{
		Code:    http.StatusOK,
		Message: "请求成功",
		Data:    mediaList,
	}

	c.JSON(http.StatusOK, result)
}

//新增媒体信息
func InsertMedia(c *gin.Context)  {
	now := time.Now()
	var meida model.Media
	err:= c.ShouldBindJSON(&meida)
	meida.CreateTime = now
	meida.ModifyTime = now
	if err != nil {
		log.Printf("新增media绑定参数问题，err:[%s]",err.Error())
		c.JSON(http.StatusBadRequest,FailResult)
	}else {
		log.Println("新增meidia信息:{}",&meida)
		err := meida.AddMedia()
		if err != nil{
			log.Printf("新增Media失败，err=[%s]",err.Error())
			c.JSON(http.StatusBadRequest,Result{http.StatusBadRequest,"新增Media失败",err.Error()})
		}else {
			c.JSON(http.StatusOK,Result{http.StatusOK,"新增Media成功",meida})
		}
	}
}

//更新媒体信息
func UpdateMedia(c *gin.Context)  {
	now := time.Now()

	var media model.Media
	err := c.ShouldBindJSON(&media)
	media.ModifyTime = now
	if err != nil {
		log.Printf("更新media绑定参数问题，err=[%s]",err.Error())
		c.JSON(http.StatusBadRequest,FailResult)
	}else {
		log.Printf("更新media信息:[%v]",media)
		err := media.UpdateMedia()
		if err != nil {
			log.Printf("更新media信息失败,ID=[%v]",media.Id)
			c.JSON(http.StatusBadRequest,FailResult)
		}else {
			log.Printf("更新media成功")
			c.JSON(http.StatusOK,Result{http.StatusOK,"更新media成功",media})
		}
	}
}

//删除媒体信息
func DeleteMedia(c *gin.Context)  {
	var m model.Media
	id := com.StrTo(c.Param("id")).MustInt64()
	m.Id = id
	_, err := m.GetMediaById()
	if  err != nil{
		log.Printf("ID为【%v】的媒体信息不存在",id)
		c.JSON(http.StatusBadRequest,Result{http.StatusBadRequest,"数据不存在",err.Error()})
	}else {
		m.DeleteMedia()
		c.JSON(http.StatusOK,Result{http.StatusOK,"删除成功",nil})
	}
}
