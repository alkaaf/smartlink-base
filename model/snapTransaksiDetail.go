package model

import (
	"time"
)

type SnapTransaksiDetail struct {
	Id               string       `json:"id" gorm:"primaryKey;index;notNull"`
	IdsnapTransaksi  string       `json:"idsnap_transaksi" gorm:"index"`
	IdsnapLayanan    string       `json:"idsnap_layanan" gorm:"index"`
	IdsnapMesin      *string      `json:"idsnap_mesin" gorm:"index"`
	IdcustomerGlobal *string      `json:"idcustomer_global" gorm:"index"`
	Idkaryawan       *string      `json:"idkaryawan" gorm:"index"`
	UsedAt           *time.Time   `json:"used_at"`
	Status           int          `json:"status"` // 1 belum pakai
	Qty              float64      `json:"qty"`
	Harga            float64      `json:"harga"`
	Durasi           int64        `json:"durasi"`
	Tipe             int          `json:"tipe"`
	IdsnapTag        string       `json:"idsnap_tag"`
	Idoutlet         string       `json:"idoutlet" gorm:"index"`
	SnapTag          *SnapTag     `json:"snap_tag"`
	SnapLayanan      *SnapLayanan `json:"snap_layanan" gorm:"foreignKey:idsnap_layanan;association_foreignkey:id"`
	SnapMesin        *SnapMesin   `json:"snap_mesin" gorm:"foreignKey:idsnap_mesin;association_foreignkey:id"`
	Misc             interface{}  `json:"misc" gorm:"-"`
	BaseModel
}

//
//func (d *SnapTransaksiDetail) getSlotStatus() {
//	module := pwa.GetModuleByIdMesin(*d.IdsnapMesin)
//	module.InfoSlot(module.SingleSlot.Id)
//	d.Misc = module
//}
