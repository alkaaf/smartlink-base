package model

type DataSatuan struct {
	IddataSatuan int    `json:"iddata_satuan" gorm:"primary_key"`
	Nama         string `json:"nama"`
	JenisSatuan  int    `json:"jenis_satuan"`
	BaseModel
}
