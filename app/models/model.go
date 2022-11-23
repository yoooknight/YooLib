package models

type BaseModel struct {
	ID         int64  `json:"id"`
	CreateDate string `json:"create_date" gorm:"default:NULL"`
	UpdateDate string `json:"update_date" gorm:"default:NULL"`
}
