package model

type SnapTag struct {
	Id     string  `json:"id" gorm:"primaryKey"`
	Load   float32 `json:"load"`
	Jenis  int     `json:"jenis"`
	Statis int     `json:"statis"`
	BaseModel
}
