package model

type Superowner struct {
	Id       string `json:"id" gorm:"primary_key"`
	Telp     string `json:"telp"`
	Password string `json:"-"`
	BaseModel
}
