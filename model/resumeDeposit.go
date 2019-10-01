package model

type ResumeDeposit struct {
	CustomerIdcustomer       string        `json:"customer_idcustomer"`
	DataDepositIddataDeposit string        `json:"data_deposit_iddata_deposit"`
	Jumlah                   float64       `json:"jumlah"`
	MasaAktif                int64         `json:"masa_aktif"`
	Deposit                  *DataDeposit2 `json:"deposit" gorm:"foreignkey:iddata_deposit;association_foreignkey:data_deposit_iddata_deposit"`
}
