package mediamodel

import "fmt"

type MediaCreateDTO struct {
	Filename  string `json:"filename" gorm:"column:filename;"`
	CloudName string `json:"cloudName" gorm:"column:cloud_name;"`
	Size      int64  `json:"size" gorm:"column:size;"`
	Ext       string `json:"ext" gorm:"column:ext;"`
	Url       string `json:"url" gorm:"column:url;"`
}

func (m *MediaCreateDTO) Fulfill(domain string) {
	m.Url = fmt.Sprintf("%s/%s", domain, m.Filename)
}

func (MediaCreateDTO) TableName() string {
	return Media{}.TableName()
}

func (m *MediaCreateDTO) Validate() error {
	return nil
}
