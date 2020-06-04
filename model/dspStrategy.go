package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//dsp_strategy:流量配置
type DspStrategy struct {
	Id         int       `json:"id" gorm:"column:id"`         //策略ID = 媒体ID，或者广告位id
	Status     int       `json:"status" gorm:"column:status"` //状态0正常 1无效
	Config     string    `json:"config" gorm:"column:config"` //配置信息
	Comment    string    `json:"comment" gorm:"column:comment"`
	CreateTime time.Time `json:"create_time" gorm:"create_time"`
	ModifyTime time.Time `json:"modify_time" gorm:"modify_time"`
}

func (DspStrategy) TableName() string {
	return "dsp_strategy"
}

//增
func (d *DspStrategy) AddDspStrategy() error {
	err := DB.Create(d).Error
	return err
}

//删
func (d *DspStrategy) DeleteDspStrategy() error {
	err := DB.Delete(d).Error
	return err
}

//改
func (d *DspStrategy) UpdateDspStrategy() error {
	err := DB.Save(d).Error
	return err
}

//查
func (d *DspStrategy) GetDspStrategy() (DspStrategy, error) {
	err := DB.First(&d, d.Id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return *d, err
	}
	return *d, nil
}

//查所有
func (d *DspStrategy) GetAllDspStrategy(pageSize int, pageNum int) ([]DspStrategy, error) {
	var (
		list []DspStrategy
		err  error
	)
	if pageSize > 0 && pageNum > 0 {
		err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("id asc").Find(&list).Error
	} else {
		err = DB.Find(&list).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return list, nil
}

//total
func (d *DspStrategy) CountOfDspStrategy() (int, error) {
	var (
		list  []DspStrategy
		total int
	)
	err := DB.Find(&list).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
