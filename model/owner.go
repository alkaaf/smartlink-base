package model

import "time"

type Owner struct {
	Idowner             string     `json:"idowner" gorm:"primarykey;index"`
	Nama                string     `json:"nama"`
	Email               string     `json:"email"`
	Password            string     `json:"-"`
	Kota                string     `json:"kota"`
	Telp                string     `json:"telp"`
	Salt                string     `json:"salt"`
	RememberToken       string     `json:"remember_token"`
	VerifiedEmail       int        `json:"verified_email"`
	VerifiedTelp        int        `json:"verified_telp"`
	KodeVerifikasiTelp  string     `json:"kode_verifikasi_telp"`
	LevelAfiliasi       int        `json:"level_afiliasi"`
	KodeAfiliasi        string     `json:"kode_afiliasi"`
	kodeAfiliator       string     `json:"kode_afiliator"`
	PersentaseAfiliator int        `json:"persentase_afiliator"`
	SaldoAfiliator      float32    `json:"saldo_afiliator"`
	TutorNumber         int        `json:"tutor_number"`
	DataGenerated       int        `json:"data_generated"`
	ExpiryTrialDate     *time.Time `json:"expiry_trial_date"`
	Active              int        `json:"active"`
	LastLogin           *time.Time `json:"last_login"`
	EmailReminders      string     `json:"email_reminders"`
	Versi               string     `json:"versi"`
	DateReset           string     `json:"date_reset"`
	SmsPremium          int        `json:"sms_premium"`
	UploadToken         string     `json:"upload_token"`
	LastKoinNotif       int        `json:"last_koin_notif"`
	Idfotoprofil        string     `json:"idfotoprofil"`
	BaseModel
}
