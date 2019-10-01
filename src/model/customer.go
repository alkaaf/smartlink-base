package model

type Customer struct {
	Idcustomer    string          `json:"idcustomer" gorm:"primary_key"`
	Nama          string          `json:"nama" "`
	OwnerIdowner  string          `json:"owner_idowner"`
	Telp          string          `json:"telp"`
	JenisKelamin  int             `json:"jenis_kelamin"`
	TanggalLahir  int64           `json:"tanggal_lahir"`
	Instansi      string          `json:"instansi"`
	Email         string          `json:"email"`
	Idoutlet      string          `json:"idoutlet"`
	CountryCode   string          `json:"country_code" gorm:"index"`
	Phone         string          `json:"phone" gorm:"index"`
	EmptyPhone    int             `json:"empty_phone"`
	Alamat        []Alamat        `json:"alamat" gorm:"foreginkey:customer_idcustomer;association_foreignkey:idcustomer"`
	ResumeEmoney  *ResumeEmoney   `json:"quota_emoney" gorm:"foreignkey:customer_idcustomer;association_foreignkey:idcustomer"`
	ResumeDeposit []ResumeDeposit `json:"quota_deposit" gorm:"foreignkey:customer_idcustomer;association_foreignkey:idcustomer"`
	BaseModel
}
