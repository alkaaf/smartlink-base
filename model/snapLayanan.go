package model

type SnapLayanan struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Nama     string `json:"nama"`
	Idoutlet string `json:"idoutlet" gorm:"index"`
	Idowner  string `json:"idowner" gorm:"index"`
	//Qty              float32           `json:"-" gorm:"default:0"` // TODO decide to hide or show
	//Harga            float32           `json:"harga" gorm:"default:0"`
	Durasi                int64              `json:"durasi" gorm:"default:0"`
	Tipe                  int                `json:"tipe"` // 1  cuci 2 kering
	SnapLayananMesin      []SnapLayananMesin `json:"snap_layanan_mesin" gorm:"foreignKey:idsnap_layanan;association_foreignkey:id"`
	SnapLayananMesinHarga *SnapLayananMesin  `json:"snap_layanan_mesin_harga" gorm:"foreignKey:idsnap_layanan;association_foreignkey:id"`
	BaseModel
}
