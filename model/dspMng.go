package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DspMng struct {
	//dsp_mng_id
	Id int64 `json:"id" gorm:"column:id"`
	//状态：0-生效，1-无效
	Status int `json:"status" gorm:"column:status"`
	//dspID,dsp_source表主键id
	DspId int64 `json:"dsp_id" gorm:"column:dsp_id"`
	//dsp媒体id
	DspMediaId string `json:"dspmedia_id" gorm:"column:dspmedia_id"`
	//dsp媒体名称
	DspMediaName string `json:"dspmedia_name" gorm:"column:dspmedia_name"`
	//dsp广告位id
	DspSlotId string `json:"dspslot_id" gorm:"column:dspslot_id"`
	//dsp广告位名称
	DspSlotName string `json:"dspslot_name" gorm:"column:dspslot_name"`
	//dsp渠道ID
	DspOpt string `json:"dspopt" gorm:"column:dspopt"`
	//广告前端缓存时间 0-不缓存，单位：秒
	CacheTime int `json:"cache_time" gorm:"column:cache_time"`
	//每日请求次数限制
	Quota int `json:"quota" gorm:"column:quota"`
	//banner 1 插屏 2 开屏 3 信息流 4
	SlotType int `json:"slot_type" gorm:"column:slot_type"`
	//dsp包名
	Package string `json:"package" gorm:"column:package"`
	//android 1 ios 2 wphone 3
	OsType int `json:"os_type" gorm:"column:os_type"`
	//TODO　描述　ｔｅｘｔ
	Comment string `json:"comment" gorm:"column:comment"`
	//创建时间
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	//最后更新时间
	ModifyTime time.Time `json:"modify_time" gorm:"column:modify_time"`
	//请求广告时的参数，0透传广告`请求携带的
	Width int `json:"width" gorm:"column:width"`
	//请求广告时的参数，0透传广告请求携带的
	Height int `json:"height" gorm:"column:height"`
}

func (DspMng) TableName() string {
	return "dsp_mng"
}

//增
func (d *DspMng) AddDspMng() error {
	err := DB.Create(d).Error
	return err
}

//删
func (d *DspMng) DeleteDspMng() error {
	err := DB.Delete(d).Error
	return err
}

//改
func (d *DspMng) UpdateDspMng() error {
	err := DB.Save(d).Error
	return err
}

//查
func (d *DspMng) GetDspMng() (DspMng, error) {
	err := DB.First(&d, d.Id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return *d, err
	}
	return *d, nil
}

//查所有
func (d *DspMng) GetAllDspMng(pageSize int, pageNum int) ([]DspMng, error) {
	var (
		dspMngList []DspMng
		err        error
	)
	if pageSize > 0 && pageNum > 0 {
		err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("id asc").Find(&dspMngList).Error
	} else {
		err = DB.Find(&dspMngList).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return dspMngList, nil
}

//total
func (d *DspMng) CountOfDspMng() (int, error) {
	var (
		list  []DspMng
		total int
	)
	err := DB.Find(&list).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
