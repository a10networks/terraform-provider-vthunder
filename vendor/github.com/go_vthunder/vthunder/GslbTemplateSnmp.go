package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type GslbTemplateSnmp struct {
	Username GslbTemplateSnmpInstance `json:"snmp-instance,omitempty"`
}

type GslbTemplateSnmpInstance struct {
	AuthKey          string `json:"auth-key,omitempty"`
	AuthProto        string `json:"auth-proto,omitempty"`
	Community        string `json:"community,omitempty"`
	ContextEngineID  string `json:"context-engine-id,omitempty"`
	ContextName      string `json:"context-name,omitempty"`
	Host             string `json:"host,omitempty"`
	Interface        int    `json:"interface,omitempty"`
	Interval         int    `json:"interval,omitempty"`
	Oid              string `json:"oid,omitempty"`
	Port             int    `json:"port,omitempty"`
	PrivKey          string `json:"priv-key,omitempty"`
	PrivProto        string `json:"priv-proto,omitempty"`
	SecurityEngineID string `json:"security-engine-id,omitempty"`
	SecurityLevel    string `json:"security-level,omitempty"`
	SnmpName         string `json:"snmp-name,omitempty"`
	UUID             string `json:"uuid,omitempty"`
	UserTag          string `json:"user-tag,omitempty"`
	Username         string `json:"username,omitempty"`
	Version          string `json:"version,omitempty"`
}

func PostGslbTemplateSnmp(id string, inst GslbTemplateSnmp, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGslbTemplateSnmp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/gslb/template/snmp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbTemplateSnmp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetGslbTemplateSnmp(id string, name string, host string) (*GslbTemplateSnmp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGslbTemplateSnmp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/gslb/template/snmp/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbTemplateSnmp
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

func PutGslbTemplateSnmp(id string, name string, inst GslbTemplateSnmp, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutGslbTemplateSnmp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/gslb/template/snmp/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbTemplateSnmp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteGslbTemplateSnmp(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteGslbTemplateSnmp")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/gslb/template/snmp/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbTemplateSnmp
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
