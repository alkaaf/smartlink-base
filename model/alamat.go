package model

type Alamat struct {
	Idalamat           string `json:"idalamat" gorm:"primary_key"`
	CustomerIdcustomer string `json:"customer_idcustomer"`
	Lokasi             string `json:"lokasi"`
	Tetap              int    `json:"tetap"`
	BaseModel
}
