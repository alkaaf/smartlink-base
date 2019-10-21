package model

type Superowner struct {
	Id       string `json:"id" gorm:"primarykey;index"`
	Telp     string `json:"telp"`
	Password string `json:"-"`
	BaseModel
}
