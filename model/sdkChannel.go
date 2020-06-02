package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//sdk_chanel结构体
type SdkChannel struct {
	Id         int       `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`
	Context    string    `json:"context" gorm:"column:context"` //TODO
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

func (SdkChannel) TableName() string {
	return "sdk_channel"
}

//增
func (s *SdkChannel) AddSdkChannel() error {
	err := DB.Create(s).Error
	return err
}

//删
func (s *SdkChannel) DeleteSdkChannel() error {
	err := DB.Delete(s).Error
	return err
}

//改
func (s *SdkChannel) UpdateSdkChannel() error {
	err := DB.Save(s).Error
	return err
}

//查
func (s *SdkChannel) GetSdkChannel() (SdkChannel, error) {
	err := DB.First(&s, s.Id).Error
	if err != nil {
		return *s, err
	}
	return *s, nil
}

//查全部
func (s *SdkChannel) GetAllSdkChannel(pageSize, pageNum int) ([]SdkChannel, error) {
	var (
		list []SdkChannel
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

//条数
func (s *SdkChannel) CountOfSdkChannel() (int, error) {
	var (
		list  []SdkChannel
		total int
	)
	err := DB.Find(&list).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
