package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type GslbGroup struct {
	ConfigSave GslbGroupInstance `json:"group-instance,omitempty"`
}

type GslbGroupInstance struct {
	AutoMapLearn   int                    `json:"auto-map-learn,omitempty"`
	AutoMapPrimary int                    `json:"auto-map-primary,omitempty"`
	AutoMapSmart   int                    `json:"auto-map-smart,omitempty"`
	ConfigAnywhere int                    `json:"config-anywhere,omitempty"`
	ConfigMerge    int                    `json:"config-merge,omitempty"`
	ConfigSave     int                    `json:"config-save,omitempty"`
	DNSDiscover    int                    `json:"dns-discover,omitempty"`
	DataInterface  int                    `json:"data-interface,omitempty"`
	Enable         int                    `json:"enable,omitempty"`
	Learn          int                    `json:"learn,omitempty"`
	MgmtInterface  int                    `json:"mgmt-interface,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Primary        []GslbGroupPrimaryList `json:"primary-list,omitempty"`
	Priority       int                    `json:"priority,omitempty"`
	Standalone     int                    `json:"standalone,omitempty"`
	Suffix         string                 `json:"suffix,omitempty"`
	UUID           string                 `json:"uuid,omitempty"`
	UserTag        string                 `json:"user-tag,omitempty"`
}

type GslbGroupPrimaryList struct {
	Primary string `json:"primary,omitempty"`
}

func PostGslbGroup(id string, inst GslbGroup, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGslbGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/gslb/group", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetGslbGroup(id string, name string, host string) (*GslbGroup, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGslbGroup")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/gslb/group/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbGroup
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

func PutGslbGroup(id string, name string, inst GslbGroup, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutGslbGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/gslb/group/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteGslbGroup(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteGslbGroup")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/gslb/group/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
