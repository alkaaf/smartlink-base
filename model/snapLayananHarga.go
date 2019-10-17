package model

type SnapLayananHarga struct {
	IdsnapLayanan string       `json:"idsnap_layanan" gorm:"index;primaryKey"`
	IdsnapTag     string       `json:"idsnap_tag" gorm:"index"`
	Harga         float64      `json:"harga"`
	Idowner       string       `json:"idowner" gorm:"index"`
	Idoutlet      string       `json:"idoutlet" gorm:"index"`
	Owner         *Owner       `json:"owner"`
	Outlet        *Outlet      `json:"outlet"`
	SnapLayanan   *SnapLayanan `json:"snap_layanan" gorm:"foreignkey:id;association_foreignkey:idsnap_layanan"`
	SnapTag       *SnapTag     `json:"snap_tag" gorm:"foreignkey:id;association_foreignkey:idsnap_tag"`
	BaseModel
}
