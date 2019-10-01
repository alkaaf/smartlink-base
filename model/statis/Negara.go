package statis

type Negara struct {
	Id   int    `json:"id" gorm:"primaryKey;index"`
	Nama string `json:"nama"`
}
