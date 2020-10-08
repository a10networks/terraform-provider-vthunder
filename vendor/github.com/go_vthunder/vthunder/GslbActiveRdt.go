package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type GslbActiveRdt struct {
	Domain GslbActiveRdtInstance `json:"active-rdt-instance,omitempty"`
}

type GslbActiveRdtInstance struct {
	Domain   string `json:"domain,omitempty"`
	Icmp     int    `json:"icmp,omitempty"`
	Interval int    `json:"interval,omitempty"`
	Port     int    `json:"port,omitempty"`
	Retry    int    `json:"retry,omitempty"`
	Sleep    int    `json:"sleep,omitempty"`
	Timeout  int    `json:"timeout,omitempty"`
	Track    int    `json:"track,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}

func PostGslbActiveRdt(id string, inst GslbActiveRdt, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGslbActiveRdt")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/gslb/active-rdt", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbActiveRdt
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetGslbActiveRdt(id string, host string) (*GslbActiveRdt, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGslbActiveRdt")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/gslb/active-rdt/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbActiveRdt
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}

}
