package model

type Layanan struct {
	Idlayanan              string        `json:"idlayanan" gorm:"primary_key"`
	NamaLayanan            string        `json:"nama_layanan"`
	OwnerIdowner           string        `json:"owner_idowner"`
	Jumlah                 float32       `json:"jumlah"`
	DurasiPenyelesaian     int64         `json:"durasi_penyelesaian"`
	DataSatuanIddataSatuan int           `json:"data_satuan_iddata_satuan"`
	Satuan                 *DataSatuan   `json:"satuan" gorm:"foreignkey:data_satuan_iddata_satuan;association_foreignkey:iddata_satuan"`
	HargaLayanan           *HargaLayanan `json:"harga_layanan" gorm:"foreginkey:layanan_idlayanan;association_foreignkey:idlayanan"`
	BaseModel
}
