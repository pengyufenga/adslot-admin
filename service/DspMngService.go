package service

import (
	"adslot-admin/model"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//获取dsp配置信息详情
func GetDspMng(c *gin.Context) {
	var dspMng model.DspMng

	//获取传参
	param := c.Param("id")
	id := com.StrTo(param).MustInt64()
	dspMng.Id = id

	mng, err := dspMng.GetDspMng()
	if err != nil {
		log.Printf("获取dspmng详情失败,err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("获取dspmng详情失败", err))
	} else {
		c.JSON(http.StatusOK, model.SuccessResult(mng))
	}
}

//获取所有
func GetAllDspMng(c *gin.Context) {
	var dspMngDao model.DspMng
	var pageParam model.PageParam
	//获取传参
	err := c.ShouldBindJSON(&pageParam)
	if err == nil {
		total, err := dspMngDao.CountOfDspMng()
		if err == nil {
			pageParam.Total = total
			dspmngs, err := dspMngDao.GetAllDspMng(pageParam.PageSize, pageParam.PageNum)
			if err == nil {
				data := make(map[string]interface{})
				data["dspmngs"] = dspmngs
				data["pageParam"] = pageParam
				c.JSON(http.StatusOK, model.SuccessResult(data))
			} else {
				c.JSON(http.StatusBadRequest, model.FailResult("获取dspmng列表出错", err))
			}
		} else {
			c.JSON(http.StatusBadRequest, model.FailResult("获取dspmng条数出错", err))
		}
	} else {
		c.JSON(http.StatusBadRequest, model.FailResult("请求参数问题", err))
	}

}

//新增
func AddDspMng(c *gin.Context) {
	//当前时间
	now := time.Now()

	var dspmng model.DspMng
	//获取传参
	err := c.ShouldBindJSON(&dspmng)
	if err != nil {
		log.Printf("新增dspmng失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		dspmng.CreateTime = now
		dspmng.ModifyTime = now
		log.Printf("新增dspmng:[%+v]", dspmng)
		err := dspmng.AddDspMng()
		if err != nil {
			log.Printf("新增dspmng失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("新增dspmng失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(dspmng))
		}
	}
}

//更新dsp配置信息
func UpdateDspMng(c *gin.Context) {
	now := time.Now()

	var dspmng model.DspMng
	//获取参数
	err := c.ShouldBindJSON(&dspmng)
	if err != nil {
		log.Printf("更新dspmng失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		dspmng.ModifyTime = now
		log.Printf("更新dspmng:[%+v]", dspmng)
		err := dspmng.UpdateDspMng()
		if err != nil {
			log.Printf("更新dspmng失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("更新dspmng失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(dspmng))
		}
	}

}

//删除dsp配置信息
func DeleteDspMng(c *gin.Context) {
	var dspmng model.DspMng
	//获取参数
	param := c.Param("id")
	id := com.StrTo(param).MustInt64()
	dspmng.Id = id
	//查找该条记录是否存在
	_, err := dspmng.GetDspMng()
	if err != nil {
		log.Printf("删除dspmng，该dspmng不存在，dspmngId:[%v]", id)
		c.JSON(http.StatusBadRequest, model.FailResult("未找到该dspmng", err))
	} else {
		//存在则删除
		err := dspmng.DeleteDspMng()
		if err != nil {
			log.Printf("删除dspmng,失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("删除dspmng失败", err))
		} else {
			log.Printf("删除dspmng成功,AdslotId:[%v]", id)
			c.JSON(http.StatusOK, model.SuccessResult(id))
		}
	}
}
