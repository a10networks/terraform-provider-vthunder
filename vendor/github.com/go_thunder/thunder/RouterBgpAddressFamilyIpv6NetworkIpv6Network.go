package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpAddressFamilyIpv6NetworkIpv6Network struct {
	NetworkIpv6 RouterBgpAddressFamilyIpv6NetworkIpv6NetworkInstance `json:"ipv6-network,omitempty"`
}

type RouterBgpAddressFamilyIpv6NetworkIpv6NetworkInstance struct {
	Backdoor    int    `json:"backdoor,omitempty"`
	CommValue   string `json:"comm-value,omitempty"`
	Description string `json:"description,omitempty"`
	NetworkIpv6 string `json:"network-ipv6,omitempty"`
	RouteMap    string `json:"route-map,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}

func PostRouterBgpAddressFamilyIpv6NetworkIpv6Network(id string, name1 string, inst RouterBgpAddressFamilyIpv6NetworkIpv6Network, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpAddressFamilyIpv6NetworkIpv6Network")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/network/ipv6-network", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NetworkIpv6Network
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			check_api_status("PostRouterBgpAddressFamilyIpv6NetworkIpv6Network", data)

		}
	}

}

func GetRouterBgpAddressFamilyIpv6NetworkIpv6Network(id string, name1 string, name2 string, host string) (*RouterBgpAddressFamilyIpv6NetworkIpv6Network, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpAddressFamilyIpv6NetworkIpv6Network")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/network/ipv6-network/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NetworkIpv6Network
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			check_api_status("GetRouterBgpAddressFamilyIpv6NetworkIpv6Network", data)
			return &m, nil
		}
	}

}

func PutRouterBgpAddressFamilyIpv6NetworkIpv6Network(id string, name1 string, name2 string, inst RouterBgpAddressFamilyIpv6NetworkIpv6Network, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterBgpAddressFamilyIpv6NetworkIpv6Network")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/network/ipv6-network/"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NetworkIpv6Network
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			check_api_status("PutRouterBgpAddressFamilyIpv6NetworkIpv6Network", data)

		}
	}

}

func DeleteRouterBgpAddressFamilyIpv6NetworkIpv6Network(id string, name1 string, name2 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteRouterBgpAddressFamilyIpv6NetworkIpv6Network")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/network/ipv6-network/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NetworkIpv6Network
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			check_api_status("DeleteRouterBgpAddressFamilyIpv6NetworkIpv6Network", data)

		}
	}
	return nil
}
