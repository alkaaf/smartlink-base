package model

type Owner struct {
	Idowner string `json:"idowner" gorm:"primarykey;index"`
	Nama    string `json:"nama"`
	BaseModel
}
