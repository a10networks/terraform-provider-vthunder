package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type GslbSystem struct {
	GslbServiceIP GslbSystemInstance `json:"system-instance,omitempty"`
}

type GslbSystemInstance struct {
	AgeInterval             int                          `json:"age-interval,omitempty"`
	GeoLocationIana         int                          `json:"geo-location-iana,omitempty"`
	GslbGroup               int                          `json:"gslb-group,omitempty"`
	GeoLocationLoadFilename []GslbSystemGslbLoadFileList `json:"gslb-load-file-list,omitempty"`
	GslbServiceIP           int                          `json:"gslb-service-ip,omitempty"`
	GslbSite                int                          `json:"gslb-site,omitempty"`
	Hostname                int                          `json:"hostname,omitempty"`
	IPTTL                   int                          `json:"ip-ttl,omitempty"`
	Module                  int                          `json:"module,omitempty"`
	SlbDevice               int                          `json:"slb-device,omitempty"`
	SlbServer               int                          `json:"slb-server,omitempty"`
	SlbVirtualServer        int                          `json:"slb-virtual-server,omitempty"`
	TTL                     int                          `json:"ttl,omitempty"`
	UUID                    string                       `json:"uuid,omitempty"`
	Wait                    int                          `json:"wait,omitempty"`
}

type GslbSystemGslbLoadFileList struct {
	GeoLocationLoadFilename string `json:"geo-location-load-filename,omitempty"`
	TemplateName            string `json:"template-name,omitempty"`
}

func PostGslbSystem(id string, inst GslbSystem, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGslbSystem")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/gslb/system", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbSystem
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetGslbSystem(id string, host string) (*GslbSystem, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGslbSystem")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/gslb/system/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbSystem
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
