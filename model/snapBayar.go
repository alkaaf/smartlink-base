package model

import "time"

type SnapBayar struct {
	Id               string     `json:"id" gorm:"primarykey"`
	IdsnapTransaksi  string     `json:"idsnap_transaksi" gorm:"index"`
	Bayar            int64      `json:"bayar"`
	Kembalian        int64      `json:"kembalian"`
	JenisBayar       int        `json:"jenis_bayar"`
	Idoutlet         string     `json:"idoutlet"`
	Idowner          string     `json:"idowner"`
	Keterangan       string     `json:"keterangan"`
	IdakunDebit      string     `json:"idakun_debit"`
	TransferedAt     *time.Time `json:"transfered_at"`
	TransferedAtLong int64      `json:"transfered_at_long"`
	CreatedAtLong    int64      `json:"created_at_long"`
	BaseModel
}
