package model

type Transaksi struct {
	IdTransaksi        string       `json:"idtransaksi" gorm:"column:idtransaksi;primary_key"`
	CustomerIdCustomer string       `json:"customer_idcustomer" gorm:"column:customer_idcustomer"`
	Grandtotal         float32      `json:"grandtotal" gorm:"column:grandtotal"`
	DetLayanan         []Detlayanan `json:"detlayanan" gorm:"foreignkey:transaksi_idtransaksi;association_foreignkey:idtransaksi"`
	BaseModel
}
