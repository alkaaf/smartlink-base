package model

import (
	"fmt"
	"time"
)

type CustomerGlobal struct {
	Id                   string                `json:"id" gorm:"primary_key;index;not null"`
	CountryCode          string                `json:"country_code" gorm:"index"`
	Email                string                `json:"email" gorm:"index"`
	Telp                 string                `json:"telp" gorm:"index"`
	Saldo                float64               `json:"saldo"`
	Idowner              string                `json:"idowner" gorm:"index"`
	Idoutlet             string                `json:"idoutlet" gorm:"index"`
	NamaAwal             string                `json:"nama_awal"`
	NamaAkhir            string                `json:"nama_akhir"`
	JenisKelamin         int                   `json:"jenis_kelamin"`
	Ttl                  time.Time             `json:"ttl"`
	Alamat               string                `json:"alamat"`
	Kota                 string                `json:"kota"`
	Idkota               int                   `json:"idkota" gorm:"index"`
	Provinsi             string                `json:"provinsi"`
	Idprovinsi           int                   `json:"idprovinsi"  gorm:"index"`
	Negara               string                `json:"negara"`
	Idnegara             int                   `json:"idnegara" gorm:"index"`
	Pass                 string                `json:"pass"`
	FullTelp             string                `json:"full_telp"`
	DataLengkap          int                   `json:"data_lengkap" gorm:"default:0"`
	Verified             int                   `json:"verified"`
	VerifiedAt           time.Time             `json:"verified_at"`
	LastLogin            time.Time             `json:"last_login"`
	Otp                  *string               `json:"-"`
	ExpiredOtp           time.Time             `json:"-"`
	Link                 *string               `json:"-"`
	ExpiredLink          time.Time             `json:"-"`
	OtpGlobal            *string               `json:"-"`
	ExpiredOtpGlobal     time.Time             `json:"-"`
	CustomerGlobalRelate *CustomerGlobalRelate `json:"customer_global_relate" gorm:"index;foreignKey:idcustomer_global;association_foreignkey:id"`
	BaseModel
}

func (c CustomerGlobal) SendableTelp() string {
	return fmt.Sprintf("+%s%s", c.CountryCode, c.Telp)
}

func (c *CustomerGlobal) AssignFullTelp() {
	c.FullTelp = c.CountryCode + c.Telp
}
