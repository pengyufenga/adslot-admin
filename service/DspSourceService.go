package service

import (
	"adslot-admin/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//获取dsp配置信息详情
func GetDspSource(c *gin.Context) {
	var dspSource model.DspSource

	//获取传参
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	dspSource.DspId = id

	source, err := dspSource.GetDspSource()
	if err != nil {
		log.Printf("获取dspSource详情失败,err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("获取dspSource详情失败", err))
	} else {
		c.JSON(http.StatusOK, model.SuccessResult(source))
	}
}

//获取所有
func GetAllDspSource(c *gin.Context) {
	var dspSource model.DspSource
	var pageParam model.PageParam
	//获取传参
	err := c.ShouldBindJSON(&pageParam)
	if err == nil {
		total, err := dspSource.CountOfDspSource()
		if err == nil {
			pageParam.Total = total
			dspSources, err := dspSource.GetAllDspSource(pageParam.PageSize, pageParam.PageNum)
			if err == nil {
				data := make(map[string]interface{})
				data["dspSources"] = dspSources
				data["pageParam"] = pageParam
				c.JSON(http.StatusOK, model.SuccessResult(data))
			} else {
				c.JSON(http.StatusBadRequest, model.FailResult("获取dspSources列表出错", err))
			}
		} else {
			c.JSON(http.StatusBadRequest, model.FailResult("获取dspSources条数出错", err))
		}
	} else {
		c.JSON(http.StatusBadRequest, model.FailResult("请求参数问题", err))
	}

}

//新增
func AddDspSource(c *gin.Context) {
	var dspSources model.DspSource
	//获取传参
	err := c.ShouldBindJSON(&dspSources)
	if err != nil {
		log.Printf("新增dspSource失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		log.Printf("新增dspSource:[%+v]", dspSources)
		err := dspSources.AddDspSource()
		if err != nil {
			log.Printf("新增dspSource失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("新增dspSource失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(dspSources))
		}
	}
}

//更新dsp配置信息
func UpdateDspSource(c *gin.Context) {
	var dspSource model.DspSource
	//获取参数
	err := c.ShouldBindJSON(&dspSource)
	if err != nil {
		log.Printf("更新dspSource失败,参数绑定异常，err=[%s]", err.Error())
		c.JSON(http.StatusBadRequest, model.FailResult("参数绑定异常", err))
	} else {
		log.Printf("更新dspSource:[%+v]", dspSource)
		err := dspSource.UpdateDspSource()
		if err != nil {
			log.Printf("更新dspSource失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("更新dspSource失败", err))
		} else {
			c.JSON(http.StatusOK, model.SuccessResult(dspSource))
		}
	}

}

//删除dsp配置信息
func DeleteDspSource(c *gin.Context) {
	var dspSource model.DspSource
	//获取参数
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	dspSource.DspId = id
	//查找该条记录是否存在
	_, err := dspSource.GetDspSource()
	if err != nil {
		log.Printf("删除dspSource，该dspSource不存在，dspId:[%v]", id)
		c.JSON(http.StatusBadRequest, model.FailResult("未找到该dspSource", err))
	} else {
		//存在则删除
		err := dspSource.DeleteDspSource()
		if err != nil {
			log.Printf("删除dspSource,失败，err=[%s]", err.Error())
			c.JSON(http.StatusBadRequest, model.FailResult("删除dspSource失败", err))
		} else {
			log.Printf("删除dspSource成功,AdslotId:[%v]", id)
			c.JSON(http.StatusOK, model.SuccessResult(id))
		}
	}
}
