package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type GslbServiceIp struct {
	HealthCheckDisable GslbServiceIPInstance `json:"service-ip-instance,omitempty"`
}

type GslbServiceIPInstance struct {
	Action                     string                        `json:"action,omitempty"`
	ExternalIP                 string                        `json:"external-ip,omitempty"`
	HealthCheck                string                        `json:"health-check,omitempty"`
	HealthCheckDisable         int                           `json:"health-check-disable,omitempty"`
	HealthCheckProtocolDisable int                           `json:"health-check-protocol-disable,omitempty"`
	IPAddress                  string                        `json:"ip-address,omitempty"`
	Ipv6                       string                        `json:"ipv6,omitempty"`
	Ipv6Address                string                        `json:"ipv6-address,omitempty"`
	NodeName                   string                        `json:"node-name,omitempty"`
	PortProto                  []GslbServiceIpPortList       `json:"port-list,omitempty"`
	Counters1                  []GslbServiceIpSamplingEnable `json:"sampling-enable,omitempty"`
	UUID                       string                        `json:"uuid,omitempty"`
	UserTag                    string                        `json:"user-tag,omitempty"`
}

type GslbServiceIpPortList struct {
	Action                     string                        `json:"action,omitempty"`
	FollowPortProtocol         string                        `json:"follow-port-protocol,omitempty"`
	HealthCheck                string                        `json:"health-check,omitempty"`
	HealthCheckDisable         int                           `json:"health-check-disable,omitempty"`
	HealthCheckFollowPort      int                           `json:"health-check-follow-port,omitempty"`
	HealthCheckProtocolDisable int                           `json:"health-check-protocol-disable,omitempty"`
	PortNum                    int                           `json:"port-num,omitempty"`
	PortProto                  string                        `json:"port-proto,omitempty"`
	Counters1                  []GslbServiceIpSamplingEnable `json:"sampling-enable,omitempty"`
	UUID                       string                        `json:"uuid,omitempty"`
	UserTag                    string                        `json:"user-tag,omitempty"`
}

type GslbServiceIpSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostGslbServiceIp(id string, inst GslbServiceIp, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGslbServiceIp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/gslb/service-ip", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbServiceIp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetGslbServiceIp(id string, name string, host string) (*GslbServiceIp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGslbServiceIp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/gslb/service-ip/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbServiceIp
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

func PutGslbServiceIp(id string, name string, inst GslbServiceIp, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutGslbServiceIp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/gslb/service-ip/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbServiceIp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteGslbServiceIp(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteGslbServiceIp")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/gslb/service-ip/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbServiceIp
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
