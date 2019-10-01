package model

type SnapLayananMesin struct {
	IdsnapMesin   string       `json:"idsnap_mesin" gorm:"index;primaryKey"`
	IdsnapLayanan string       `json:"idsnap_layanan" gorm:"index;primaryKey"`
	SnapLayanan   *SnapLayanan `json:"snap_layanan" gorm:"foreignKey:id;association_foreignKey:idsnap_layanan"`
	SnapMesin     *SnapMesin   `json:"snap_mesin" gorm:"foreignKey:id;association_foreignkey:idsnap_mesin"`
	Harga         float64      `json:"harga"`
	BaseModel
}
