package model

type ResumeEmoney struct {
	CustomerIdcustomer string  `json:"customer_idcustomer"`
	Saldo              float64 `json:"saldo"`
	MasaAktif          int64   `json:"masa_aktif"`
}
