package go_vthunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type VridInstance struct {
	UUID Vrid `json:"vrid,omitempty"`
}
type VlanCfg struct {
	Vlan         int `json:"vlan,omitempty"`
	Timeout      int `json:"timeout,omitempty"`
	PriorityCost int `json:"priority-cost,omitempty"`
}
type Ipv6DestinationCfg struct {
	Distance        int    `json:"distance,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	PriorityCost    int    `json:"priority-cost,omitempty"`
	Ipv6Destination string `json:"ipv6-destination,omitempty"`
	Gatewayv6       string `json:"gatewayv6,omitempty"`
}
type IPDestinationCfg struct {
	Distance      int    `json:"distance,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	Mask          string `json:"mask,omitempty"`
	PriorityCost  int    `json:"priority-cost,omitempty"`
	IPDestination string `json:"ip-destination,omitempty"`
	Gateway       string `json:"gateway,omitempty"`
}
type Route struct {
	Ipv6Destination []Ipv6DestinationCfg `json:"ipv6-destination-cfg,omitempty"`
	IPDestination   []IPDestinationCfg   `json:"ip-destination-cfg,omitempty"`
}
type BgpIpv4AddressCfg struct {
	BgpIpv4Address string `json:"bgp-ipv4-address,omitempty"`
	PriorityCost   int    `json:"priority-cost,omitempty"`
}
type BgpIpv6AddressCfg struct {
	BgpIpv6Address string `json:"bgp-ipv6-address,omitempty"`
	PriorityCost   int    `json:"priority-cost,omitempty"`
}
type Bgp struct {
	BgpIpv4Address []BgpIpv4AddressCfg `json:"bgp-ipv4-address-cfg,omitempty"`
	BgpIpv6Address []BgpIpv6AddressCfg `json:"bgp-ipv6-address-cfg,omitempty"`
}
type Interface struct {
	Ethernet     int `json:"ethernet,omitempty"`
	PriorityCost int `json:"priority-cost,omitempty"`
}
type Ipv4GatewayList struct {
	UUID         string `json:"uuid,omitempty"`
	IPAddress    string `json:"ip-address,omitempty"`
	PriorityCost int    `json:"priority-cost,omitempty"`
}
type Ipv6GatewayList struct {
	Ipv6Address  string `json:"ipv6-address,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	PriorityCost int    `json:"priority-cost,omitempty"`
}
type Gateway struct {
	IPAddress   []Ipv4GatewayList `json:"ipv4-gateway-list,omitempty"`
	Ipv6Address []Ipv6GatewayList `json:"ipv6-gateway-list,omitempty"`
}
type TrunkCfg struct {
	PriorityCost int `json:"priority-cost,omitempty"`
	Trunk        int `json:"trunk,omitempty"`
	PerPortPri   int `json:"per-port-pri,omitempty"`
}
type TrackingOptions struct {
	Vlan []VlanCfg `json:"vlan-cfg,omitempty"`
	UUID string    `json:"uuid,omitempty"`
	//IPDestination  Route       `json:"route,omitempty"`
	//BgpIpv4Address Bgp         `json:"bgp,omitempty"`
	Ethernet []Interface `json:"interface,omitempty"`
	//IPAddress      Gateway     `json:"gateway,omitempty"`
	Trunk []TrunkCfg `json:"trunk-cfg,omitempty"`
}
type BladeParameters struct {
	Priority               int             `json:"priority,omitempty"`
	FailOverPolicyTemplate string          `json:"fail-over-policy-template,omitempty"`
	UUID_BladeParams       string          `json:"uuid,omitempty"`
	UUID                   TrackingOptions `json:"tracking-options,omitempty"`
}
type PreemptMode struct {
	Threshold int `json:"threshold,omitempty"`
	Disable   int `json:"disable,omitempty"`
}
type Ipv6AddressPartCfg struct {
	Ethernet             int    `json:"ethernet,omitempty"`
	Ipv6AddressPartition string `json:"ipv6-address-partition,omitempty"`
	Ve                   int    `json:"ve,omitempty"`
	Trunk                int    `json:"trunk,omitempty"`
}
type IPAddressCfg struct {
	IPAddress string `json:"ip-address,omitempty"`
}
type IPAddressPartCfg struct {
	IPAddressPartition string `json:"ip-address-partition,omitempty"`
}
type Ipv6AddressCfg struct {
	Ipv6Address string `json:"ipv6-address,omitempty"`
	Ethernet    int    `json:"ethernet,omitempty"`
	Ve          int    `json:"ve,omitempty"`
	Trunk       int    `json:"trunk,omitempty"`
}
type FloatingIP struct {
	Ipv6AddressPartition []Ipv6AddressPartCfg `json:"ipv6-address-part-cfg,omitempty"`
	IPAddress            []IPAddressCfg       `json:"ip-address-cfg,omitempty"`
	IPAddressPartition   []IPAddressPartCfg   `json:"ip-address-part-cfg,omitempty"`
	Ipv6Address          []Ipv6AddressCfg     `json:"ipv6-address-cfg,omitempty"`
}
type Follow struct {
	VridLead string `json:"vrid-lead,omitempty"`
}
type Vrid struct {
	UUID_BladeParams BladeParameters `json:"blade-parameters,omitempty"`
	UUID             string          `json:"uuid,omitempty"`
	VridVal          *int            `json:"vrid-val,omitempty"`
	UserTag          string          `json:"user-tag,omitempty"`
	Threshold        PreemptMode     `json:"preempt-mode,omitempty"`
	Ipv6Address      FloatingIP      `json:"floating-ip,omitempty"`
	VridLead         Follow          `json:"follow,omitempty"`
}

func GetVrrpVrid(id string, name string, host string) (*VridInstance, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/vrrp-a/vrid/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VridInstance
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}
}

func PostVrrpVrid(id string, vc VridInstance, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/vrrp-a/vrid", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v VridInstance
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func PutVrrpVrid(id string, name string, vc VridInstance, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	logger.Println("[INFO] Inside putVrrpVrid")

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/vrrp-a/vrid/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VridInstance
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func DeleteVrrpVrid(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/vrrp-a/vrid/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VridInstance
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
		}
	}
	return nil
}
