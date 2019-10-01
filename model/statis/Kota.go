package statis

type Kota struct {
	Id         int    `json:"id" gorm:"primaryKey;index"`
	Nama       string `json:"nama"`
	ProvinsiId int    `json:"provinsi_id"`
	NegaraId   int    `json:"negara_id"`
}
