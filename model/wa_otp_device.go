package model

import "time"

type WaOtpDevice struct {
	Iddevice  string     `json:"iddevice" gorm:"primary_key"`
	Number    string     `json:"number"`
	Token     string     `json:"token"`
	LastServe *time.Time `json:"last_serve"`
	BaseModel
}
