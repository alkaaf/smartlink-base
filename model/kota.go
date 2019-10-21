package model

import "time"

type Kota struct {
	Id        string     `json:"id" gorm:"primarykey;index"`
	NegaraId  string     `json:"negara_id"`
	Nama      string     `json:"nama"`
	Status    string     `json:"status"`
	CreatedBy string     `json:"created_by"`
	UpdatedBy string     `json:"updated_by"`
	Type      int        `json:"type"`
	DeletedAt *time.Time `json:"deleted_at"`
	BaseModel
}
