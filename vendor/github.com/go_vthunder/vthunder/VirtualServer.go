package go_vthunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type ACLNameList struct {
	ACLNameSrcNatPoolShared string `json:"acl-name-src-nat-pool-shared,omitempty"`
	SharedPartitionPoolName int    `json:"shared-partition-pool-name,omitempty"`
	ACLNameSeqNumShared     int    `json:"acl-name-seq-num-shared,omitempty"`
	ACLName                 string `json:"acl-name,omitempty"`
	ACLNameSrcNatPool       string `json:"acl-name-src-nat-pool,omitempty"`
	ACLNameSeqNum           int    `json:"acl-name-seq-num,omitempty"`
}

type ACLIDList struct {
	ACLIDSeqNum           int    `json:"acl-id-seq-num,omitempty"`
	ACLIDSrcNatPool       string `json:"acl-id-src-nat-pool,omitempty"`
	ACLIDSeqNumShared     int    `json:"acl-id-seq-num-shared,omitempty"`
	ACLID                 int    `json:"acl-id,omitempty"`
	SharedPartitionPoolID int    `json:"shared-partition-pool-id,omitempty"`
	ACLIDSrcNatPoolShared string `json:"acl-id-src-nat-pool-shared,omitempty"`
}

type AflexScripts struct {
	Aflex       string `json:"aflex,omitempty"`
	AflexShared string `json:"aflex-shared,omitempty"`
}

type AuthCfg struct {
	AaaPolicy string `json:"aaa-policy"`
}

type PortList struct {
	HaConnMirror                                int              `json:"ha-conn-mirror,omitempty"`
	Protocol                                    string           `json:"protocol,omitempty"`
	Precedence                                  int              `json:"precedence,omitempty"`
	PortTranslation                             int              `json:"port-translation,omitempty"`
	IPMapList                                   string           `json:"ip-map-list,omitempty"`
	TemplateReqmodIcap                          string           `json:"template-reqmod-icap,omitempty"`
	ACLName                                     []ACLNameList    `json:"acl-name-list,omitempty"`
	StatsDataAction                             string           `json:"stats-data-action,omitempty"`
	UseDefaultIfNoServer                        int              `json:"use-default-if-no-server,omitempty"`
	TemplateConnectionReuse                     string           `json:"template-connection-reuse,omitempty"`
	UUID                                        string           `json:"uuid,omitempty"`
	TemplateTCPShared                           string           `json:"template-tcp-shared,omitempty"`
	TemplateTCP                                 string           `json:"template-tcp,omitempty"`
	TemplatePersistCookie                       string           `json:"template-persist-cookie,omitempty"`
	SharedPartitionDynamicServiceTemplate       int              `json:"shared-partition-dynamic-service-template,omitempty"`
	SharedPartitionConnectionReuseTemplate      int              `json:"shared-partition-connection-reuse-template,omitempty"`
	WhenDown                                    int              `json:"when-down,omitempty"`
	TemplateClientSslShared                     string           `json:"template-client-ssl-shared,omitempty"`
	SharedPartitionPersistDestinationIPTemplate int              `json:"shared-partition-persist-destination-ip-template,omitempty"`
	SharedPartitionExternalServiceTemplate      int              `json:"shared-partition-external-service-template,omitempty"`
	PersistType                                 string           `json:"persist-type,omitempty"`
	SharedPartitionHTTPPolicyTemplate           int              `json:"shared-partition-http-policy-template,omitempty"`
	UseRcvHopForResp                            int              `json:"use-rcv-hop-for-resp,omitempty"`
	ScaleoutBucketCount                         int              `json:"scaleout-bucket-count,omitempty"`
	OptimizationLevel                           string           `json:"optimization-level,omitempty"`
	ReqFail                                     int              `json:"req-fail,omitempty"`
	NoDestNat                                   int              `json:"no-dest-nat,omitempty"`
	Name                                        string           `json:"name,omitempty"`
	TemplateSmpp                                string           `json:"template-smpp,omitempty"`
	UserTag                                     string           `json:"user-tag,omitempty"`
	TemplateDiameter                            string           `json:"template-diameter,omitempty"`
	samplingEnable                              []SamplingEnable `json:"sampling-enable,omitempty"`
	TemplateSsli                                string           `json:"template-ssli,omitempty"`
	SharedPartitionPersistCookieTemplate        int              `json:"shared-partition-persist-cookie-template,omitempty"`
	SharedPartitionPolicyTemplate               int              `json:"shared-partition-policy-template,omitempty"`
	TemplatePolicy                              string           `json:"template-policy,omitempty"`
	NoLogging                                   int              `json:"no-logging,omitempty"`
	ResetOnServerSelectionFail                  int              `json:"reset-on-server-selection-fail,omitempty"`
	WafTemplate                                 string           `json:"waf-template,omitempty"`
	Ipinip                                      int              `json:"ipinip,omitempty"`
	NoAutoUpOnAflex                             int              `json:"no-auto-up-on-aflex,omitempty"`
	Rate                                        int              `json:"rate,omitempty"`
	GslbEnable                                  int              `json:"gslb-enable,omitempty"`
	TemplateDNSShared                           string           `json:"template-dns-shared,omitempty"`
	TemplatePersistSslSid                       string           `json:"template-persist-ssl-sid,omitempty"`
	TemplateDNS                                 string           `json:"template-dns,omitempty"`
	SharedPartitionDNSTemplate                  int              `json:"shared-partition-dns-template,omitempty"`
	TemplateSip                                 string           `json:"template-sip,omitempty"`
	TemplateDblb                                string           `json:"template-dblb,omitempty"`
	SharedPartitionServerSslTemplate            int              `json:"shared-partition-server-ssl-template,omitempty"`
	TemplateClientSsl                           string           `json:"template-client-ssl,omitempty"`
	SupportHTTP2                                int              `json:"support-http2,omitempty"`
	TemplateClientSSH                           string           `json:"template-client-ssh,omitempty"`
	SharedPartitionTCPProxyTemplate             int              `json:"shared-partition-tcp-proxy-template,omitempty"`
	EnablePlayeridCheck                         int              `json:"enable-playerid-check,omitempty"`
	ServiceGroup                                string           `json:"service-group,omitempty"`
	SharedPartitionPersistSslSidTemplate        int              `json:"shared-partition-persist-ssl-sid-template,omitempty"`
	DefSelectionIfPrefFailed                    string           `json:"def-selection-if-pref-failed,omitempty"`
	SharedPartitionUDP                          int              `json:"shared-partition-udp,omitempty"`
	SynCookie                                   int              `json:"syn-cookie,omitempty"`
	AlternatePort                               int              `json:"alternate-port,omitempty"`
	AlternatePortNumber                         int              `json:"alternate-port-number,omitempty"`
	TemplatePersistSourceIPShared               string           `json:"template-persist-source-ip-shared,omitempty"`
	TemplateCache                               string           `json:"template-cache,omitempty"`
	TemplatePersistCookieShared                 string           `json:"template-persist-cookie-shared,omitempty"`
	RtpSipCallIDMatch                           int              `json:"rtp-sip-call-id-match,omitempty"`
	TemplateFileInspection                      string           `json:"template-file-inspection,omitempty"`
	TemplateFtp                                 string           `json:"template-ftp,omitempty"`
	ServSelFail                                 int              `json:"serv-sel-fail,omitempty"`
	TemplateUDP                                 string           `json:"template-udp,omitempty"`
	TemplateVirtualPortShared                   string           `json:"template-virtual-port-shared,omitempty"`
	Action                                      string           `json:"action,omitempty"`
	TemplateHTTP                                string           `json:"template-http,omitempty"`
	View                                        int              `json:"view,omitempty"`
	TemplatePersistSourceIP                     string           `json:"template-persist-source-ip,omitempty"`
	TemplateDynamicService                      string           `json:"template-dynamic-service,omitempty"`
	SharedPartitionVirtualPortTemplate          int              `json:"shared-partition-virtual-port-template,omitempty"`
	UseCgnv6                                    int              `json:"use-cgnv6,omitempty"`
	TemplatePersistDestinationIP                string           `json:"template-persist-destination-ip,omitempty"`
	TemplateVirtualPort                         string           `json:"template-virtual-port,omitempty"`
	ConnLimit                                   int              `json:"conn-limit,omitempty"`
	TrunkFwd                                    int              `json:"trunk-fwd,omitempty"`
	TemplateUDPShared                           string           `json:"template-udp-shared,omitempty"`
	TemplateHTTPPolicyShared                    string           `json:"template-http-policy-shared,omitempty"`
	Pool                                        string           `json:"pool,omitempty"`
	SnatOnVip                                   int              `json:"snat-on-vip,omitempty"`
	TemplateConnectionReuseShared               string           `json:"template-connection-reuse-shared,omitempty"`
	SharedPartitionTCP                          int              `json:"shared-partition-tcp,omitempty"`
	aclidList                                   []ACLIDList      `json:"acl-id-list,omitempty"`
	SharedPartitionHTTPTemplate                 int              `json:"shared-partition-http-template,omitempty"`
	TemplateExternalService                     string           `json:"template-external-service,omitempty"`
	OnSyn                                       int              `json:"on-syn,omitempty"`
	TemplatePersistSslSidShared                 string           `json:"template-persist-ssl-sid-shared,omitempty"`
	ForceRoutingMode                            int              `json:"force-routing-mode,omitempty"`
	TemplateHTTPPolicy                          string           `json:"template-http-policy,omitempty"`
	TemplatePolicyShared                        string           `json:"template-policy-shared,omitempty"`
	TemplateScaleout                            string           `json:"template-scaleout,omitempty"`
	WhenDownProtocol2                           int              `json:"when-down-protocol2,omitempty"`
	TemplateFix                                 string           `json:"template-fix,omitempty"`
	TemplateSMTP                                string           `json:"template-smtp,omitempty"`
	RedirectToHTTPS                             int              `json:"redirect-to-https,omitempty"`
	AltProtocol2                                string           `json:"alt-protocol2,omitempty"`
	AltProtocol1                                string           `json:"alt-protocol1,omitempty"`
	MessageSwitching                            int              `json:"message-switching,omitempty"`
	TemplateImapPop3                            string           `json:"template-imap-pop3,omitempty"`
	ScaleoutDeviceGroup                         int              `json:"scaleout-device-group,omitempty"`
	SharedPartitionPersistSourceIPTemplate      int              `json:"shared-partition-persist-source-ip-template,omitempty"`
	L7HardwareAssist                            int              `json:"l7-hardware-assist,omitempty"`
	TemplateTCPProxyShared                      string           `json:"template-tcp-proxy-shared,omitempty"`
	SharedPartitionCacheTemplate                int              `json:"shared-partition-cache-template,omitempty"`
	UseAlternatePort                            int              `json:"use-alternate-port,omitempty"`
	TemplateTCPProxyServer                      string           `json:"template-tcp-proxy-server,omitempty"`
	TrunkRev                                    int              `json:"trunk-rev,omitempty"`
	EthFwd                                      int              `json:"eth-fwd,omitempty"`
	PoolShared                                  string           `json:"pool-shared,omitempty"`
	TemplateRespmodIcap                         string           `json:"template-respmod-icap,omitempty"`
	Range                                       int              `json:"range,omitempty"`
	Reset                                       int              `json:"reset,omitempty"`
	TemplateExternalServiceShared               string           `json:"template-external-service-shared,omitempty"`
	Auto                                        int              `json:"auto,omitempty"`
	TemplateDynamicServiceShared                string           `json:"template-dynamic-service-shared,omitempty"`
	TemplateServerSSH                           string           `json:"template-server-ssh,omitempty"`
	aflexScripts                                []AflexScripts   `json:"aflex-scripts,omitempty"`
	TemplateHTTPShared                          string           `json:"template-http-shared,omitempty"`
	TemplateServerSsl                           string           `json:"template-server-ssl,omitempty"`
	SharedPartitionDiameterTemplate             int              `json:"shared-partition-diameter-template,omitempty"`
	TemplateServerSslShared                     string           `json:"template-server-ssl-shared,omitempty"`
	TemplatePersistDestinationIPShared          string           `json:"template-persist-destination-ip-shared,omitempty"`
	TemplateCacheShared                         string           `json:"template-cache-shared,omitempty"`
	PortNumber                                  int              `json:"port-number,omitempty"`
	TemplateTCPProxyClient                      string           `json:"template-tcp-proxy-client,omitempty"`
	SharedPartitionPool                         int              `json:"shared-partition-pool,omitempty"`
	TemplateTCPProxy                            string           `json:"template-tcp-proxy,omitempty"`
	ExtendedStats                               int              `json:"extended-stats,omitempty"`
	SharedPartitionClientSslTemplate            int              `json:"shared-partition-client-ssl-template,omitempty"`
	Expand                                      int              `json:"expand,omitempty"`
	SkipRevHash                                 int              `json:"skip-rev-hash,omitempty"`
	TemplateDiameterShared                      string           `json:"template-diameter-shared,omitempty"`
	ClientipStickyNat                           int              `json:"clientip-sticky-nat,omitempty"`
	Secs                                        int              `json:"secs,omitempty"`
	authCfg                                     AuthCfg          `json:"auth-cfg,omitempty"`
	EthRev                                      int              `json:"eth-rev,omitempty"`
}

type MigrateVip struct {
	TargetDataCPU      int    `json:"target-data-cpu,omitempty"`
	UUID               string `json:"uuid,omitempty"`
	FinishMigration    int    `json:"finish-migration,omitempty"`
	TargetFloatingIpv6 string `json:"target-floating-ipv6,omitempty"`
	TargetFloatingIpv4 string `json:"target-floating-ipv4,omitempty"`
	CancelMigration    int    `json:"cancel-migration,omitempty"`
}

type VirtualServerMain struct {
	Protocol                      []PortList `json:"port-list,omitempty"`
	StatsDataAction               string     `json:"stats-data-action,omitempty"`
	Ipv6ACLShared                 string     `json:"ipv6-acl-shared,omitempty"`
	ACLName                       string     `json:"acl-name,omitempty"`
	EnableDisableAction           string     `json:"enable-disable-action,omitempty"`
	HaDynamic                     int        `json:"ha-dynamic,omitempty"`
	RedistributeRouteMap          string     `json:"redistribute-route-map,omitempty"`
	ACLNameShared                 string     `json:"acl-name-shared,omitempty"`
	IPAddress                     string     `json:"ip-address,omitempty"`
	migrateVip                    MigrateVip `json:"migrate-vip,omitempty"`
	UseIfIP                       int        `json:"use-if-ip,omitempty"`
	UUID                          string     `json:"uuid,omitempty"`
	Vrid                          int        `json:"vrid,omitempty"`
	DisableVipAdv                 int        `json:"disable-vip-adv,omitempty"`
	TemplateVirtualServer         string     `json:"template-virtual-server,omitempty"`
	ArpDisable                    int        `json:"arp-disable,omitempty"`
	Description                   string     `json:"description,omitempty"`
	RedistributionFlagged         int        `json:"redistribution-flagged,omitempty"`
	Netmask                       string     `json:"netmask,omitempty"`
	ACLID                         int        `json:"acl-id,omitempty"`
	Ipv6ACL                       string     `json:"ipv6-acl,omitempty"`
	TemplateLogging               string     `json:"template-logging,omitempty"`
	ExtendedStats                 int        `json:"extended-stats,omitempty"`
	Name                          string     `json:"name,omitempty"`
	TemplateScaleout              string     `json:"template-scaleout,omitempty"`
	TemplatePolicy                string     `json:"template-policy,omitempty"`
	UserTag                       string     `json:"user-tag,omitempty"`
	TemplatePolicyShared          string     `json:"template-policy-shared,omitempty"`
	Ipv6Address                   string     `json:"ipv6-address,omitempty"`
	Ethernet                      int        `json:"ethernet,omitempty"`
	SharedPartitionPolicyTemplate int        `json:"shared-partition-policy-template,omitempty"`
	ACLIDShared                   int        `json:"acl-id-shared,omitempty"`
}

type VirtalServerInstanceMain struct {
	Name VirtualServerMain `json:"virtual-server,omitempty"`
}

func GetVS(id string, name string, host string) (*VirtalServerInstanceMain, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/virtual-server/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtalServerInstanceMain
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

func PostVS(id string, vs VirtalServerInstanceMain, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	logger.Println("[INFO] received V  --" + vs.Name.Name + ",--" + vs.Name.UUID)

	payloadBytes, err := json.Marshal(vs)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/virtual-server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtalServerInstanceMain
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func PutVS(id string, name string, vs VirtalServerInstanceMain, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	logger.Println("[INFO] received V  --" + vs.Name.Name + ",--" + vs.Name.UUID)

	payloadBytes, err := json.Marshal(vs)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/virtual-server/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtalServerInstanceMain
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func DeleteVS(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/virtual-server/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtalServerInstanceMain
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
