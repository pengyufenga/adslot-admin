package model

import (
	"time"
)

type Media struct {
	Id         int64     `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`
	User       int64     `json:"user" gorm:"column:user"`
	Type       int       `json:"type" gorm:"column:type"`
	OsType     int       `json:"os_type" gorm:"column:os_type"`
	SourceUrl  string    `json:"source_url" gorm:"column:source_url"`
	Comment    string    `json:"comment" gorm:"column:comment"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	ModifyTime time.Time `json:"modify_time" gorm:"column:modify_time"`
}

func (Media) TableName() string {
	return "media"
}

//获取所有媒体列表
func (m *Media) GetAllMedia() ([]Media, error) {
	var(
		mediaList []Media
		err error
	)
	err = DB.Find(&mediaList).Error
	if err != nil {
		return nil, err
	}
	return mediaList, nil
}

//根据Id获取媒体详情
func (m *Media) GetMediaById() (Media, error) {
	result := DB.First(&m, m.Id)
	return *m, result.Error
}

//新增媒体
func (m *Media) AddMedia() error{
	err:= DB.Create(m).Error
	return err
}

//更新媒体信息
func (m *Media) UpdateMedia() error{
	err := DB.Save(m).Error
	return err
}

//删除媒体信息
func (m *Media) DeleteMedia() {
	DB.Delete(m)
}
