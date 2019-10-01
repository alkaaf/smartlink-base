package model

type Item struct {
	Iditem           string           `json:"iditem"`
	Nama             string           `json:"nama"`
	OwnerIdowner     string           `json:"owner_idowner"`
	DaftarItemOutlet DaftarItemOutlet `json:"daftar_item_outlet" gorm:"foreignkey:iditem;association_foreignkey:item_iditem"`
	BaseModel
}
