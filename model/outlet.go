package model

type Outlet struct {
	Idoutlet string `json:"idoutlet" gorm:"primaryKey;index"`
	Nama     string `json:"nama"`
	Telp     string `json:"telp"`
	Alamat   string `json:"alamat"`
	BaseModel
}
