package model

type HargaLayanan struct {
	OutletIdoutlet   string  `json:"outlet_idoutlet"`
	Harga            float32 `json:"harga"`
	LayananIdlayanan string  `json:"layanan_idlayanan" gorm:"primary_key"`
	MinOrderReg      float32 `json:"min_order_reg"`
	MinOrderDepo     float32 `json:"min_order_depo"`
	BaseModel
}
