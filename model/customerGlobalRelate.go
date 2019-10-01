package model

type CustomerGlobalRelate struct {
	IdcustomerGlobal string `json:"idcustomer_global" gorm:"index;primarykey"`
	Idowner          string `json:"idowner" gorm:"index;primarykey"`
	Idoutlet         string `json:"idoutlet" gorm:"index;primarykey"`
}
