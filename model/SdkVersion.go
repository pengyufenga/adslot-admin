package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//sdk_version 结构体
type SdkVersion struct {
	Id         int       `json:"id" gorm:"column:id"`
	Ver        string    `json:"ver" gorm:"column:ver"`         //版本号
	Context    string    `json:"context" gorm:"column:context"` // 版本描述
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

func (SdkVersion) TableName() string {
	return "sdk_version"
}

//增
func (d *SdkVersion) AddSdkVersion() error {
	err := DB.Create(d).Error
	return err
}

//删
func (d *SdkVersion) DeleteSdkVersion() error {
	err := DB.Delete(d).Error
	return err
}

//改
func (d *SdkVersion) UpdateSdkVersion() error {
	err := DB.Save(d).Error
	return err
}

//查
func (d *SdkVersion) GetSdkVersion() (SdkVersion, error) {
	err := DB.First(&d, d.Id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return *d, err
	}
	return *d, nil
}

//查所有
func (d *SdkVersion) GetAllSdkVersion(pageSize int, pageNum int) ([]SdkVersion, error) {
	var (
		list []SdkVersion
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
func (d *SdkVersion) CountOfSdkVersion() (int, error) {
	var (
		list  []SdkVersion
		total int
	)
	err := DB.Find(&list).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
