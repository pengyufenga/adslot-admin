package model

import "github.com/jinzhu/gorm"

//dsp_source结构体
type DspSource struct {
	DspId      int    `json:"dsp_id"  gorm:"column:dsp_id"`     //dsp_id
	DspType    int    `json:"dsp_type"  gorm:"column:dsp_type"` //dsp类型：1-sdk，2-api
	DspName    string `json:"dsp_name"  gorm:"column:dsp_name"`
	Comment    string `json:"comment"  gorm:"column:comment"`           //TODO 描述 text
	ReqUrl     string `json:"req_url"  gorm:"column:req_url"`           //广告请求的地址
	ReqTestUrl string `json:"req_test_url"  gorm:"column:req_test_url"` //测试地址
	TimeOut    int    `json:"timeout"  gorm:"column:timeout"`           //超时时间：单位：毫秒
	AdLogo     string `json:"ad_logo"  gorm:"column:ad_logo"`           //广告logo
}

//type DspSource struct {
//	DspId      int64  `gorm:"column:dsp_id"`   //dsp_id
//	DspType    int    `gorm:"column:dsp_type"` //dsp类型：1-sdk，2-api
//	DspName    string `gorm:"column:dsp_name"`
//	Comment    string `gorm:"column:comment"`      //TODO 描述 text
//	ReqUrl     string `gorm:"column:req_url"`      //广告请求的地址
//	ReqTestUrl string `gorm:"column:req_test_url"` //测试地址
//	TimeOut    int    `gorm:"column:timeout"`      //超时时间：单位：毫秒
//	AdLogo     string `gorm:"column:ad_logo"`      //广告logo
//}

func (DspSource) TableName() string {
	return "dsp_source"
}

//增
func (d *DspSource) AddDspSource() error {
	err := DB.Create(d).Error
	return err
}

//删
func (d *DspSource) DeleteDspSource() error {
	err := DB.Delete(d).Error
	return err
}

//改
func (d *DspSource) UpdateDspSource() error {
	err := DB.Save(d).Error
	return err
}

//查
func (d *DspSource) GetDspSource() (DspSource, error) {
	err := DB.Where("dsp_id = ?", d.DspId).First(&d).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return *d, err
	}
	return *d, nil
}

//查所有
func (d *DspSource) GetAllDspSource(pageSize int, pageNum int) ([]DspSource, error) {
	var (
		dspSourceList []DspSource
		err           error
	)
	if pageSize > 0 && pageNum > 0 {
		err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("id asc").Find(&dspSourceList).Error
	} else {
		err = DB.Find(&dspSourceList).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return dspSourceList, nil
}

//total
func (d *DspSource) CountOfDspSource() (int, error) {
	var (
		list  []DspSource
		total int
	)
	err := DB.Find(&list).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
