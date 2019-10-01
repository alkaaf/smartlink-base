package model

type SnapModuleSlot struct {
	Id           int    `json:"id" gorm:"primaryKey;index;notNull"`
	IdsnapModule string `json:"idsnap_module" gorm:"primaryKey;index"`
	Idmesin      string `json:"idmesin" gorm:"index"`
	Idowner      string `json:"idowner" gorm:"index"`
	BaseModel
}
