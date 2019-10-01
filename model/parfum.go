package model

type Parfum struct {
	Idparfum     int    `json:"idparfum" gorm:"primary_key"`
	NamaParfum   string `json:"nama_parfum"`
	OwnerIdowner string `json:"owner_idowner"`
	BaseModel
}
