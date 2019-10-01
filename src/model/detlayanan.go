package model

type Detlayanan struct {
	IdTransaksi string  `gorm:"column:transaksi_idtransaksi"`
	Id          string  `gorm:"column:id;primary_key"`
	JumlahBeli  float32 `gorm:"column:jumlah_beli"`
	BaseModel
}
