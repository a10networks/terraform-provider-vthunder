package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"util"
)

type VeBFD struct {
	UUID VeBFDInstance `json:"bfd,omitempty"`
}
type VeBFDIntervalCfg struct {
	Interval   int `json:"interval,omitempty"`
	MinRx      int `json:"min-rx,omitempty"`
	Multiplier int `json:"multiplier,omitempty"`
}
type VeBFDAuthentication struct {
	Encrypted string `json:"encrypted,omitempty"`
	Password  string `json:"password,omitempty"`
	Method    string `json:"method,omitempty"`
	KeyID     int    `json:"key-id,omitempty"`
}
type VeBFDInstance struct {
	Interval  VeBFDIntervalCfg    `json:"interval-cfg,omitempty"`
	Encrypted VeBFDAuthentication `json:"authentication,omitempty"`
	Echo      int                 `json:"echo,omitempty"`
	UUID      string              `json:"uuid,omitempty"`
	Demand    int                 `json:"demand,omitempty"`
}

func PostInterfaceVeBFD(id string, name int, inst VeBFD, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceVeBFD")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ve/"+strconv.Itoa(name)+"/bfd", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VeBFD
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostInterfaceVeBFD REQ RES..........................", m)
			check_api_status("PostInterfaceVeBFD", data)

		}
	}

}

func GetInterfaceVeBFD(id string, name string, host string) (*VeBFD, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceVeBFD")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ve/"+name+"/bfd/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VeBFD
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetInterfaceVeBFD REQ RES..........................", m)
			check_api_status("GetInterfaceVeBFD", data)
			return &m, nil
		}
	}

}
