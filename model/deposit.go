package model

type DataDeposit2 struct {
	IddataDeposit    string   `json:"iddata_deposit"`
	Nama             string   `json:"nama"`
	MasaAktif        int64    `json:"masa_aktif"`
	Jumlah           float64  `json:"jumlah"`
	LayananIdlayanan string   `json:"layanan_idlayanan"`
	Idowner          string   `json:"idowner"`
	HargaDeposit     float64  `json:"harga_deposit"`
	HargaLayanan     float64  `json:"harga_layanan"`
	Layanan          *Layanan `json:"layanan" gorm:"foreignkey:idlayanan;association_foreignkey:layanan_idlayanan"`
	BaseModel
}
