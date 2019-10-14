package model

type SnapLayanan struct {
	Id               string             `json:"id" gorm:"primaryKey"`
	Nama             string             `json:"nama"`
	Idowner          string             `json:"idowner" gorm:"index"`
	Durasi           int64              `json:"durasi" gorm:"default:0"`
	Tipe             int                `json:"tipe"` // 1  cuci 2 kering
	SnapLayananHarga []SnapLayananHarga `json:"snap_layanan_harga" gorm:"foreignkey:idsnap_layanan;association_foreignkey:id"`
	BaseModel
}
