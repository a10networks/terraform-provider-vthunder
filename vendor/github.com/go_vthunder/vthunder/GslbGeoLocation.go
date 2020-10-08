package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type GslbGeoLocation struct {
	GeoLocnObjName GslbGeoLocationInstance `json:"geo-location-instance,omitempty"`
}

type GslbGeoLocationInstance struct {
	FirstIPAddress []GslbGeoLocationLocnMultipleAddresses `json:"geo-locn-multiple-addresses,omitempty"`
	GeoLocnObjName string                                 `json:"geo-locn-obj-name,omitempty"`
	UUID           string                                 `json:"uuid,omitempty"`
	UserTag        string                                 `json:"user-tag,omitempty"`
}

type GslbGeoLocationLocnMultipleAddresses struct {
	FirstIPAddress   string `json:"first-ip-address,omitempty"`
	FirstIpv6Address string `json:"first-ipv6-address,omitempty"`
	GeolIpv4Mask     string `json:"geol-ipv4-mask,omitempty"`
	GeolIpv6Mask     int    `json:"geol-ipv6-mask,omitempty"`
	IPAddr2          string `json:"ip-addr2,omitempty"`
	Ipv6Addr2        string `json:"ipv6-addr2,omitempty"`
}

func PostGslbGeoLocation(id string, inst GslbGeoLocation, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGslbGeoLocation")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/gslb/geo-location", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbGeoLocation
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetGslbGeoLocation(id string, name string, host string) (*GslbGeoLocation, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGslbGeoLocation")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/gslb/geo-location/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbGeoLocation
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

func PutGslbGeoLocation(id string, name string, inst GslbGeoLocation, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutGslbGeoLocation")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/gslb/geo-location/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbGeoLocation
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteGslbGeoLocation(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteGslbGeoLocation")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/gslb/geo-location/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GslbGeoLocation
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
