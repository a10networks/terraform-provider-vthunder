package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type GslbProtocolLimitDiff struct {
	ArdtResponse GslbProtocolLimitInstance `json:"limit-instance,omitempty"`
}

type GslbProtocolLimitInstance struct {
	ArdtQuery    int    `json:"ardt-query,omitempty"`
	ArdtResponse int    `json:"ardt-response,omitempty"`
	ArdtSession  int    `json:"ardt-session,omitempty"`
	ConnResponse int    `json:"conn-response,omitempty"`
	Message      int    `json:"message,omitempty"`
	Response     int    `json:"response,omitempty"`
	UUID         string `json:"uuid,omitempty"`
}

func PostGslbProtocolLimit(id string, inst GslbProtocolLimitDiff, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGslbProtocolLimit")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/gslb/protocol/limit", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbProtocolLimitDiff
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetGslbProtocolLimit(id string, host string) (*GslbProtocolLimitDiff, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGslbProtocolLimit")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/gslb/protocol/limit/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbProtocolLimitDiff
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
