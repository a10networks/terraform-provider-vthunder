package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type GslbProtocol struct {
	UUID GslbProtocolInstance `json:"protocol-instance,omitempty"`
}

type GslbProtocolInstance struct {
	AutoDetect                  int                      `json:"auto-detect,omitempty"`
	Type                        []GslbProtocolEnableList `json:"enable-list,omitempty"`
	ArdtResponse                GslbProtocolLimit        `json:"limit,omitempty"`
	MsgFormatAcos2X             int                      `json:"msg-format-acos-2x,omitempty"`
	PingSite                    string                   `json:"ping-site,omitempty"`
	StatusInterval              int                      `json:"status-interval,omitempty"`
	UUID                        string                   `json:"uuid,omitempty"`
	UseMgmtPort                 int                      `json:"use-mgmt-port,omitempty"`
	UseMgmtPortForAllPartitions int                      `json:"use-mgmt-port-for-all-partitions,omitempty"`
}

type GslbProtocolEnableList struct {
	Type string `json:"type,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

type GslbProtocolLimit struct {
	ArdtQuery    int    `json:"ardt-query,omitempty"`
	ArdtResponse int    `json:"ardt-response,omitempty"`
	ArdtSession  int    `json:"ardt-session,omitempty"`
	ConnResponse int    `json:"conn-response,omitempty"`
	Message      int    `json:"message,omitempty"`
	Response     int    `json:"response,omitempty"`
	UUID         string `json:"uuid,omitempty"`
}

func PostGslbProtocol(id string, inst GslbProtocol, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGslbProtocol")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/gslb/protocol", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbProtocol
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetGslbProtocol(id string, host string) (*GslbProtocol, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGslbProtocol")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/gslb/protocol/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbProtocol
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
