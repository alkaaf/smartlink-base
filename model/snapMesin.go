package model

type SnapMesin struct {
	Id         string  `json:"id"	gorm:"index;primaryKey;notNull"`
	Idowner    string  `json:"idowner" gorm:"index"`
	Idoutlet   string  `json:"idoutlet" gorm:"index"`
	Nama       string  `json:"nama"`
	Tipe       int     `json:"tipe"`
	Load       float32 `json:"load"`
	Keterangan string  `json:"keterangan"`
	Mode       int     `json:"mode"`
	IdsnapTag  string  `json:"idsnap_tag"`
	BaseModel
}
