package template

type SMSPayload struct {
	Number string `json:"number"`
	Text   string `json:"text"`
}
