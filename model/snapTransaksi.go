package model

type SnapTransaksi struct {
	Id                  string                `json:"id" gorm:"primaryKey;index;notNull"`
	Qty                 float64               `json:"qty"`
	Idoutlet            string                `json:"idoutlet" gorm:"index"`
	Idowner             string                `json:"idowner" gorm:"index" swagger:"ignore"`
	IdcustomerGlobal    *string               `json:"idcustomer_global" gorm:"index"`
	Idkaryawan          *string               `json:"idkaryawan" gorm:"index"`
	NamaCustomer        *string               `json:"nama_customer"`
	Jenis               int                   `json:"jenis"`        // 1 semi 2 full 3 drop
	StatusBayar         int                   `json:"status_bayar"` // 1 sukses 2 gagal 3 pending
	PersenDiskon        float64               `json:"persen_diskon"`
	Diskon              float64               `json:"diskon"`
	PersenPajak         float64               `json:"persen_pajak"`
	Pajak               float64               `json:"pajak"`
	Total               float64               `json:"total"`
	Grandtotal          float64               `json:"grandtotal"`
	Keterangan          *string               `json:"keterangan"`
	Lunas               int                   `json:"lunas" gorm:"default:0"`
	CreatedAtLong       int64                 `json:"created_at_long"`
	SnapTransaksiDetail []SnapTransaksiDetail `json:"snap_transaksi_detail" gorm:"foreignkey:idsnap_transaksi;association_foreignkey:id"`
	SnapBayar           []SnapBayar           `json:"snap_bayar" gorm:"foreignkey:idsnap_transaksi;association_foreignkey:id"`
	Outlet              *Outlet               `json:"outlet" gorm:"foreignkey:idoutlet;association_foreignkey:idoutlet"`
	Owner               *Owner                `json:"owner" gorm:"foreignkey:idowner;association_foreignkey:idowner"`
	Karyawan            *Karyawan             `json:"karyawan" gorm:"foreignkey:idkaryawan;association_foreignkey:idkaryawan"`
	BaseModel
}
