package model

type Karyawan struct {
	Idkaryawan string `json:"idkaryawan" gorm:"index;primaryKey"`
	Nama       string `json:"nama"`
	BaseModel
}
