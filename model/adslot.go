package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Adslot struct {
	//广告位ID
	Id int64 `json:"id" gorm:"column:id"`
	//广告位名称
	Name string `json:"name" gorm:"column:name"`
	//每日请求数限制，0为不限制
	Quota int `json:"quota" gorm:"column:quota"`
	//所属媒体，media中的ID
	MediaId int64 `json:"media" gorm:"column:media"`
	//生效状态：0-生效，1-无效
	Status int `json:"status" gorm:"column:status"`
	//TODO 广告位配置json ,text
	Config string `json:"config" gorm:"column:config"`
	//创建时间
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	//更新时间
	ModifyTime time.Time `json:"modify_time" gorm:"column:modify_time"`
}

func (Adslot) TableName() string {
	return "adslot"
}

//增
func (a *Adslot) AddAdslot() error {
	err := DB.Create(a).Error
	return err
}

//删
func (a *Adslot) DeleteAdslot() error {
	err := DB.Delete(a).Error
	return err
}

//改
func (a *Adslot) UpdateAdslot() error {
	err := DB.Save(a).Error
	return err
}

//查
func (a *Adslot) GetAdslot() (Adslot, error) {
	err := DB.First(&a, a.Id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return *a, err
	}
	return *a, nil
}

//查全部
func (a *Adslot) GetAllAdslot(pageSize int, pageNum int) ([]Adslot, error) {
	var (
		adslotList []Adslot
		err        error
	)
	if pageSize > 0 && pageNum > 0 {
		err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("id asc").Find(&adslotList).Error
	} else {
		err = DB.Find(&adslotList).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return adslotList, nil
}

//数量
func (a *Adslot) CountOfAdslot() (int, error) {
	var (
		list  []Adslot
		total int
	)
	err := DB.Find(&list).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
