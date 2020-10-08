package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type GslbServiceGroup struct {
	ServiceGroupName GslbServiceGroupInstance `json:"service-group-instance,omitempty"`
}

type GslbServiceGroupInstance struct {
	DependencySite      int                               `json:"dependency-site,omitempty"`
	Disable             int                               `json:"disable,omitempty"`
	DisableSite         []GslbServiceGroupDisableSiteList `json:"disable-site-list,omitempty"`
	MemberName          []GslbServiceGroupMember          `json:"member,omitempty"`
	PersistentAgingTime int                               `json:"persistent-aging-time,omitempty"`
	PersistentIpv6Mask  int                               `json:"persistent-ipv6-mask,omitempty"`
	PersistentMask      string                            `json:"persistent-mask,omitempty"`
	PersistentSite      int                               `json:"persistent-site,omitempty"`
	ServiceGroupName    string                            `json:"service-group-name,omitempty"`
	UUID                string                            `json:"uuid,omitempty"`
	UserTag             string                            `json:"user-tag,omitempty"`
}

type GslbServiceGroupDisableSiteList struct {
	DisableSite string `json:"disable-site,omitempty"`
}

type GslbServiceGroupMember struct {
	MemberName string `json:"member-name,omitempty"`
}

func PostGslbServiceGroup(id string, inst GslbServiceGroup, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGslbServiceGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/gslb/service-group", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbServiceGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetGslbServiceGroup(id string, name string, host string) (*GslbServiceGroup, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGslbServiceGroup")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/gslb/service-group/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbServiceGroup
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

func PutGslbServiceGroup(id string, name string, inst GslbServiceGroup, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutGslbServiceGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/gslb/service-group/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbServiceGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteGslbServiceGroup(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteGslbServiceGroup")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/gslb/service-group/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbServiceGroup
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
