package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerEnableTrapsRoutingBgp struct {
	BgpEstablishedNotification SnmpServerEnableTrapsRoutingBgpInstance `json:"bgp,omitempty"`
}

type SnmpServerEnableTrapsRoutingBgpInstance struct {
	BgpBackwardTransNotification int    `json:"bgpBackwardTransNotification,omitempty"`
	BgpEstablishedNotification   int    `json:"bgpEstablishedNotification,omitempty"`
	UUID                         string `json:"uuid,omitempty"`
}

func PostSnmpServerEnableTrapsRoutingBgp(id string, inst SnmpServerEnableTrapsRoutingBgp, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerEnableTrapsRoutingBgp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/enable/traps/routing/bgp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsRoutingBgp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			check_api_status("PostSnmpServerEnableTrapsRoutingBgp", data)

		}
	}

}

func GetSnmpServerEnableTrapsRoutingBgp(id string, host string) (*SnmpServerEnableTrapsRoutingBgp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerEnableTrapsRoutingBgp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/enable/traps/routing/bgp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsRoutingBgp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			check_api_status("GetSnmpServerEnableTrapsRoutingBgp", data)
			return &m, nil
		}
	}

}
