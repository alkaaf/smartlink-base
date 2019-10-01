package statis

type Provinsi struct {
	Id       string `json:"id" gorm:"primaryKey;index"`
	Nama     string `json:"nama"`
	NegaraId int    `json:"negara_id"`
}
