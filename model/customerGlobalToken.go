package model

import "time"

type CustomerGlobalToken struct {
	Id         int64     `json:"id" gorm:"primary_key;index;notNull;autoIncrement"`
	IdCustomer string    `json:"country_code" gorm:"index"`
	Token      string    `json:"token" gorm:"index"`
	ExpiredAt  time.Time `json:"expired_at"`
	RemoteIp   string    `json:"remote_ip"`
	UserAgent  string    `json:"user_agent"`
	BaseModel
}
