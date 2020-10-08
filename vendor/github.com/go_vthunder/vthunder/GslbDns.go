package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type GslbDns struct {
	Action GslbDNSInstance `json:"dns-instance,omitempty"`
}

type GslbDNSInstance struct {
	Action    string                  `json:"action,omitempty"`
	Logging   string                  `json:"logging,omitempty"`
	Counters1 []GslbDnsSamplingEnable `json:"sampling-enable,omitempty"`
	Template  string                  `json:"template,omitempty"`
	UUID      string                  `json:"uuid,omitempty"`
}

type GslbDnsSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostGslbDns(id string, inst GslbDns, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGslbDns")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/gslb/dns", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbDns
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetGslbDns(id string, host string) (*GslbDns, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGslbDns")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/gslb/dns/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbDns
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
