package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BaseResponse struct {
	Msg string `json:"msg"`
}

type Module struct {
	Id   string `json:"id"`
	Msg  string `json:"msg"`
	Hook *Hook  `json:"hook"`
	Slot []Slot `json:"slot"`
	Data []Data `json:"data"`
	BaseResponse
}

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Hook struct {
	Time   string `json:"time"`
	Event  string `json:"event"`
	SlotId int    `json:"slot_id"`
}
type Slot struct {
	SlotId   int     `json:"slot_id"`
	TimeLeft int64   `json:"time_left"`
	Status   string  `json:"status"`
	Power    float64 `json:"power"`
	Current  float64 `json:"current"`
}

type SnapModule struct {
	Id             string           `json:"id" gorm:"primary_key;index;not null"`
	Idowner        string           `json:"idowner" gorm:"index"`
	Idoutlet       string           `json:"idoutlet" gorm:"index"`
	Url            string           `json:"url"`
	Port           int              `json:"port"`
	Path           string           `json:"path"`
	File           string           `json:"file"`
	SnapModuleSlot []SnapModuleSlot `json:"snap_module_slot"`
	SingleSlot     *SnapModuleSlot  `json:"single_slot" gorm:"foreignKey:idsnap_module;associative_foreignKey:id"`
	BaseModel
}

func (d *SnapModule) GetAddress() string {
	return fmt.Sprintf("%s:%d/%s%s", d.Url, d.Port, d.Path, d.File)
	//return "http://alldatamonitoring.com:30006/laundry/test.php"
}

func (d *SnapModule) On(slot int, duration int64) (int, *BaseResponse) {
	//sReq := fmt.Sprintf("%s/on/%d/%d", d.GetAddress(), slot, duration)
	sReq := fmt.Sprintf("%s?power=on&slot=%d&durasi=%d", d.GetAddress(), slot, duration)
	r, err := http.Get(sReq)
	if err != nil {
		return 500, &BaseResponse{Msg: err.Error()}
	} else {
		var response BaseResponse
		code := parseResponse(r, &response)
		return code, &response
	}
}

func (d *SnapModule) Off(slot int) (int, *BaseResponse) {
	sReq := fmt.Sprintf("%s?power=off&slot=%d", d.GetAddress(), slot)
	r, err := http.Get(sReq)
	if err != nil {
		return 500, &BaseResponse{Msg: err.Error()}
	} else {
		var response BaseResponse
		code := parseResponse(r, &response)
		return code, &response
	}
}

func (d *SnapModule) InfoSlot(id int) (int, *Slot) {
	sReq := fmt.Sprintf("%s?log=info", d.GetAddress())
	r, err := http.Get(sReq)
	if err != nil {
		log.Println(err)
		return 500, nil
	} else {
		var response Module
		code := parseResponse(r, &response)
		for _, v := range response.Slot {
			if v.SlotId == id {
				return code, &v
			}
		}
		return 404, nil
	}
}

func (d *SnapModule) Info() (int, *Module) {
	sReq := fmt.Sprintf("%s?log=info", d.GetAddress())
	r, err := http.Get(sReq)
	if err != nil {
		log.Println(err)
		return 500, nil
	} else {
		var response Module
		code := parseResponse(r, &response)
		return code, &response
	}
}

func (d *SnapModule) StatusSlot(slot int) string {
	// unknown, on, off
	code, module := d.Info()
	log.Println(code)
	if code != 200 {
		return "unknown"
	}
	for _, v := range module.Slot {
		if v.SlotId == slot {
			return v.Status
		}
	}
	return "unknown"
}

func parseResponse(r *http.Response, model interface{}) int {
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return 500
	}
	err = json.Unmarshal(s, model)
	if err != nil {
		return 500
	}
	return r.StatusCode
}
