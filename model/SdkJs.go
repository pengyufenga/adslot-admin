package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//sdk_jsscript结构体
type SdkJs struct {
	Id         int       `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`     //js脚本名称
	Ver        string    `json:"ver" gorm:"column:ver"`       //js脚本的版本号
	Status     int       `json:"status" gorm:"column:status"` //0-生效，1-不生效
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	ModifyTime time.Time `json:"modify_time" gorm:"column:modify_time"`
}

func (SdkJs) TableName() string {
	return "sdk_jsscript"
}

//增
func (d *SdkJs) AddSdkJs() error {
	err := DB.Create(d).Error
	return err
}

//删
func (d *SdkJs) DeleteSdkJs() error {
	err := DB.Delete(d).Error
	return err
}

//改
func (d *SdkJs) UpdateSdkJs() error {
	err := DB.Save(d).Error
	return err
}

//查
func (d *SdkJs) GetSdkJs() (SdkJs, error) {
	err := DB.First(&d, d.Id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return *d, err
	}
	return *d, nil
}

//查所有
func (d *SdkJs) GetAllSdkJs(pageSize int, pageNum int) ([]SdkJs, error) {
	var (
		list []SdkJs
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
func (d *SdkJs) CountOfSdkJs() (int, error) {
	var (
		list  []SdkJs
		total int
	)
	err := DB.Find(&list).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
