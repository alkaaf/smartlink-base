package helper

import (
	"encoding/json"
	"github.com/alkaaf/smartlink-base/model/template"
	"math/rand"
	"strconv"
)

func SendSms(p template.SMSPayload) {
	data, _ := json.Marshal(p)
	nc := NatsConn()
	_ = nc.Publish("sms.send", data)
}

func OtpGen() string {
	return strconv.Itoa(RangeIn(100000, 999999))
}

func RangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
