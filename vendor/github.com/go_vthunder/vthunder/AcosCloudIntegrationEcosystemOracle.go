package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type EcosystemOracle struct {
	UUID EcosystemOracleInstance `json:"oracle,omitempty"`
}

type EcosystemOracleInstance struct {
	ServiceLabel        string `json:"service-label,omitempty"`
	Ipv4Address         string `json:"ipv4-address,omitempty"`
	Ipv6Address         string `json:"ipv6-address,omitempty"`
	HostName            string `json:"host-name,omitempty"`
	Port                int    `json:"port,omitempty"`
	HealthCheckInterval string `json:"health-check-interval,omitempty"`
	Action              string `json:"action,omitempty"`
	CompartmentID       string `json:"compartment-id,omitempty"`
	TenancyID           string `json:"tenancy-id,omitempty"`
	UserID              string `json:"user-id,omitempty"`
	Fingerprint         string `json:"fingerprint,omitempty"`
	PrivateKeyPath      string `json:"private-key-path,omitempty"`
	UUID                string `json:"uuid,omitempty"`
}

func PostAcosCloudIntegrationEcosystemOracle(id string, inst EcosystemOracle, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostAcosCloudIntegrationEcosystemOracle")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/acos-cloud-integration/ecosystem/oracle", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EcosystemOracle
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetAcosCloudIntegrationEcosystemOracle(id string, name string, host string) (*EcosystemOracle, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetAcosCloudIntegrationEcosystemOracle")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/acos-cloud-integration/ecosystem/oracle", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EcosystemOracle
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
