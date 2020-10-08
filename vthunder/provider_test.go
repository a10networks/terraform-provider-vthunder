package vthunder

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var TEST_PARTITION = "Common"

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

var testProviders = map[string]terraform.ResourceProvider{
	"vthunder": Provider(),
}

func init() {

	testAccProvider = ProviderTest().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"vthunder": testAccProvider,
	}

}

func ProviderTest() terraform.ResourceProvider {
	return &schema.Provider{

		ResourcesMap: map[string]*schema.Resource{
			"vthunder_virtual_server":                    resourceVirtualServer(),
			"vthunder_service_group":                     resourceServiceGroup(),
			"vthunder_server":                            resourceServer(),
			"vthunder_ethernet":                          resourceEthernet(),
			"vthunder_rib_route":                         resourceRibRoute(),
			"vthunder_vrrp_vrid":                         resourceVrrpVrid(),
			"vthunder_glm":                               resourceGlm(),
			"vthunder_import":                            resourceImport(),
			"vthunder_vrrp_common":                       resourceVrrpCommon(),
			"vthunder_vrrp_peer_group":                   resourceVrrpPeerGroup(),
			"vthunder_vrrp_session_sync":                 resourceVrrpSessionSync(),
			"vthunder_dns_primary":                       resourceDnsPrimary(),
			"vthunder_harmony_controller_profile":        resourceHarmonyControllerProfile(),
			"vthunder_overlay_tunnel_options":            resourceOverlayTunnelOptions(),
			"vthunder_overlay_tunnel_vtep":               resourceOverlayTunnelVtep(),
			"vthunder_partition":                         resourceOverlayTunnelPartition(),
			"vthunder_slb_template_tcp":                  resourceSlbTemplateTcp(),
			"vthunder_slb_template_udp":                  resourceSlbTemplateUdp(),
			"vthunder_slb_template_smpp":                 resourceSlbTemplateSmpp(),
			"vthunder_slb_template_fix":                  resourceTemplateFix(),
			"vthunder_slb_template_ftp":                  resourceSlbTemplateFTP(),
			"vthunder_slb_template_mqtt":                 resourceSlbTemplateMqtt(),
			"vthunder_slb_template_http_policy":          resourceSlbTemplateHttpPolicy(),
			"vthunder_slb_template_ssli":                 resourceTemplateSSLI(),
			"vthunder_slb_template_cipher":               resourceTemplateCipher(),
			"vthunder_slb_template_imap_pop3":            resourceTemplateImap_POP3(),
			"vthunder_slb_template_dns":                  resourceTemplateDNS(),
			"vthunder_slb_template_connection_reuse":     resourceTemplateConnectionReuse(),
			"vthunder_slb_template_client_ssh":           resourceTemplateClientSsh(),
			"vthunder_slb_template_dblb":                 resourceTemplateDBLB(),
			"vthunder_slb_template_port":                 resourceTemplatePort(),
			"vthunder_slb_template_smtp":                 resourceSlbTemplateSMTP(),
			"vthunder_slb_template_tcp_proxy":            resourceSlbTemplateTcpProxy(),
			"vthunder_slb_template_server":               resourceSlbTemplateServer(),
			"vthunder_slb_template_virtual_port":         resourceSlbTemplateVirtualPort(),
			"vthunder_slb_template_diameter":             resourceSlbTemplateDiameter(),
			"vthunder_slb_template_cache":                resourceSlbTemplateCache(),
			"vthunder_slb_template_csv":                  resourceSlbTemplateCSV(),
			"vthunder_slb_template_reqmod_icap":          resourceSlbTemplateReqmodIcap(),
			"vthunder_slb_template_respmod_icap":         resourceSlbTemplateRespmodIcap(),
			"vthunder_slb_template_dynamic_service":      resourceSlbTemplateDynamicService(),
			"vthunder_slb_template_server_ssh":           resourceSlbTemplateServerSSH(),
			"vthunder_slb_template_policy":               resourceSlbTemplatePolicy(),
			"vthunder_slb_template_server_ssl":           resourceSlbTemplateServerSSL(),
			"vthunder_slb_template_client_ssl":           resourceSlbTemplateClientSSL(),
			"vthunder_slb_template_sip":                  resourceSlbTemplateSIP(),
			"vthunder_slb_template_http":                 resourceSlbTemplateHTTP(),
			"vthunder_slb_template_monitor":              resourceSlbTemplateMonitor(),
			"vthunder_slb_template_logging":              resourceSlbTemplateLogging(),
			"vthunder_slb_template_snmp":                 resourceSlbTemplateSNMP(),
			"vthunder_slb_template_external_service":     resourceSlbTemplateExternalService(),
			"vthunder_slb_template_virtual_server":       resourceSlbTemplateVirtualServer(),
			"vthunder_slb_aflow":                         resourceSlbAflow(),
			"vthunder_slb_common":                        resourceSlbCommon(),
			"vthunder_slb_common_conn_rate_limit_src_ip": resourceSlbCommonConnRateLimitSrcIP(),
			"vthunder_slb_connection_reuse":              resourceSlbConnectionReuse(),
			"vthunder_slb_crl_srcip":                     resourceSlbCrlSrcip(),
			"vthunder_slb_dns":                           resourceSlbDNS(),
			"vthunder_slb_sip":                           resourceSlbSip(),
			"vthunder_slb_smpp":                          resourceSlbSmpp(),
			"vthunder_slb_smtp":                          resourceSlbSmtp(),
			"vthunder_slb_switch":                        resourceSlbSwitch(),
			"vthunder_slb_spdy_proxy":                    resourceSlbSpdyProxy(),
			"vthunder_slb_svm_source_nat":                resourceSlbSvmSourceNat(),
			"vthunder_slb_sport_rate_limit":              resourceSlbSportRateLimit(),
			"vthunder_slb_transparent_acl_template":      resourceSlbTransparentAclTemplate(),
			"vthunder_slb_transparent_tcp_template":      resourceSlbTransperentTcpTemplate(),
			"vthunder_slb_dns_cache":                     resourceSlbDNSCache(),
			"vthunder_slb_dns_response_rate_limiting":    resourceSlbDNSResponseRateLimiting(),
			"vthunder_slb_fast_http_proxy":               resourceSlbFastHttpProxy(),
			"vthunder_slb_fix":                           resourceSlbFix(),
			"vthunder_slb_ftp_ctl":                       resourceSlbFTPCtl(),
			"vthunder_slb_ftp_data":                      resourceSlbFTPData(),
			"vthunder_slb_ftp_proxy":                     resourceSlbFTPProxy(),
			"vthunder_slb_generic_proxy":                 resourceSlbGenericProxy(),
			"vthunder_slb_passthrough":                   resourceSlbPassthrough(),
			"vthunder_slb_perf":                          resourceSlbPerf(),
			"vthunder_slb_persist":                       resourceSlbpersist(),
			"vthunder_slb_player_id_global":              resourceSlbPlayerIdGlobal(),
			"vthunder_slb_pop3_proxy":                    resourceSlbPop3Proxy(),
			"vthunder_slb_proxy":                         resourceSlbProxy(),
			"vthunder_slb_rate_limit_log":                resourceSlbRateLimitLog(),
			"vthunder_slb_rc_cache_global":               resourceSlbRcCacheGlobal(),
			"vthunder_slb_resource_usage":                resourceSlbResourceUsage(),
			"vthunder_slb_health_gateway":                resourceSlbHealthGateway(),
			"vthunder_slb_health_stat":                   resourceSlbHealthStat(),
			"vthunder_slb_http_proxy":                    resourceSlbHttpProxy(),
			"vthunder_slb_http2":                         resourceSlbHTTP2(),
			"vthunder_slb_hw_compress":                   resourceSlbHwCompress(),
			"vthunder_slb_ssl_cert_revoke":               resourceSlbSSLCertRevoke(),
			"vthunder_slb_ssl_forward_proxy":             resourceSlbSSLForwardProxy(),
			"vthunder_slb_icap_http":                     resourceSlbIcapHTTP(),
			"vthunder_slb_icap":                          resourceSlbIcap(),
			"vthunder_slb_imapproxy":                     resourceSlbImapproxy(),
			"vthunder_slb_l4":                            resourceSlL4(),
			"vthunder_slb_l7session":                     resourceSlL7session(),
			"vthunder_slb_mlb":                           resourceSlbMlb(),
			"vthunder_slb_mssql":                         resourceSlbMssql(),
			"vthunder_slb_mysql":                         resourceSlbMysql(),
			"vthunder_slb_ssl_expire_check":              resourceSlbSSLExpireCheck(),
			"vthunder_ip_dns_primary":                    resourceIpDnsPrimary(),
			"vthunder_ip_dns_secondary":                  resourceIpDnsSecondary(),
			"vthunder_ip_dns_suffix":                     resourceIpDnsSuffix(),
			"vthunder_ip_frag":                           resourceIpFrag(),
			"vthunder_ip_icmp":                           resourceIpIcmp(),
			"vthunder_ip_nat_alg_pptp":                   resourceIpNatAlgPptp(),
			"vthunder_ip_nat_global":                     resourceIpNatGlobal(),
			"vthunder_ip_nat_icmp":                       resourceIpNatIcmp(),
			"vthunder_ip_prefix_list":                    resourceIpPrefixList(),
			"vthunder_ip_reroute":                        resourceIpReroute(),
			"vthunder_ip_tcp":                            resourceIpTcp(),
			"vthunder_ipv6_frag":                         resourceIpv6Frag(),
			"vthunder_ipv6_icmpv6":                       resourceIpv6Icmpv6(),
			"vthunder_ipv6_nat_icmpv6":                   resourceIpv6NatIcmpv6(),
			"vthunder_ip_route_static_bfd":               resourceIPRouteStaticBfd(),
			"vthunder_ip_address":                        resourceIPAddress(),
			"vthunder_interface_management":              resourceInterfaceManagement(),
			"vthunder_interface_ethernet_bfd":            resourceInterfaceEthernetBFD(),
			"vthunder_interface_ethernet_lldp":           resourceInterfaceEthernetLLDP(),
			"vthunder_interface_ve_bfd":                  resourceInterfaceVeBFD(),
			"vthunder_interface_ethernet_trunk_group":    resourceInterfaceEthernetTrunkGroup(),
			"vthunder_interface_ethernet_ipv6":           resourceInterfaceEthernetIPv6(),
			"vthunder_interface_ve_ip":                   resourceInterfaceVeIP(),
			"vthunder_interface_ve_ipv6":                 resourceInterfaceVeIPv6(),
			"vthunder_interface_ethernet":                resourceInterfaceEthernet(),
			"vthunder_interface_ve":                      resourceInterfaceVE(),
			"vthunder_fw_alg_icmp":                       resourceFwAlgIcmp(),
			"vthunder_fw_active_rule_set":                resourceFwActiveRuleSet(),
			"vthunder_fw_alg_dns":                        resourceFwAlgDns(),
			"vthunder_fw_alg_ftp":                        resourceFwAlgFtp(),
			"vthunder_fw_alg_pptp":                       resourceFwAlgPptp(),
			"vthunder_fw_alg_rtsp":                       resourceFwAlgRtsp(),
			"vthunder_fw_alg_sip":                        resourceFwAlgSip(),
			"vthunder_fw_alg_tftp":                       resourceFwAlgTftp(),
			"vthunder_fw_apply_changes":                  resourceFwApplyChanges(),
			"vthunder_fw_app":                            resourceFwApp(),
			"vthunder_fw_global":                         resourceFwGlobal(),
			"vthunder_fw_gtp":                            resourceFwGtp(),
			"vthunder_fw_gtp_in_gtp_filtering":           resourceFwGtpInGtpFiltering(),
			"vthunder_fw_gtp_v0":                         resourceFwGtpV0(),
			"vthunder_fw_helper_sessions":                resourceFwHelperSessions(),
			"vthunder_fw_limit_entry":                    resourceFwLimitEntry(),
			"vthunder_fw_local_log":                      resourceFwLocalLog(),
			"vthunder_fw_logging":                        resourceFwLogging(),
			"vthunder_fw_radius_server":                  resourceFwRadiusServer(),
			"vthunder_fw_resource_usage":                 resourceFwResourceUsage(),
			"vthunder_fw_clear_session_filter":           resourceFwClearSessionFilter(),
			"vthunder_fw_full_cone_session":              resourceFwFullConeSession(),
			"vthunder_fw_server":                         resourceFwServer(),
			"vthunder_fw_service_group":                  resourceFwServiceGroup(),
			"vthunder_fw_status":                         resourceFwStatus(),
			"vthunder_fw_system_status":                  resourceFwSystemStatus(),
			"vthunder_fw_tap_monitor":                    resourceFwTapMonitor(),
			"vthunder_fw_tcp_mss_clamp":                  resourceFwTcpMssClamp(),
			"vthunder_fw_tcp_reset_on_error":             resourceFwTcpResetOnError(),
			"vthunder_fw_tcp_rst_close_immediate":        resourceFwTcpRstCloseImmediate(),
			"vthunder_fw_tcp_window_check":               resourceFwTcpWindowCheck(),
			"vthunder_fw_top_k_rules":                    resourceFwTopKRules(),
			"vthunder_fw_urpf":                           resourceFwUrpf(),
			"vthunder_fw_vrid":                           resourceFwVrid(),
			"vthunder_fw_template_logging":               resourceFwTemplateLogging(),
			"vthunder_gslb_active_rdt":                   resourceGslbActiveRdt(),
			"vthunder_gslb_dns":                          resourceGslbDns(),
			"vthunder_gslb_protocol":                     resourceGslbProtocol(),
			"vthunder_gslb_protocol_limit":               resourceGslbProtocolLimit(),
			"vthunder_gslb_system":                       resourceGslbSystem(),
			"vthunder_gslb_geo_location":                 resourceGslbGeoLocation(),
			"vthunder_gslb_group":                        resourceGslbGroup(),
			"vthunder_gslb_ip_list":                      resourceGslbIpList(),
			"vthunder_gslb_protocol_enable":              resourceGslbProtocolEnable(),
			"vthunder_gslb_service_group":                resourceGslbServiceGroup(),
			"vthunder_gslb_service_ip":                   resourceGslbServiceIp(),
			"vthunder_gslb_template_csv":                 resourceGslbTemplateCsv(),
			"vthunder_gslb_template_snmp":                resourceGslbTemplateSnmp(),
		},

		ConfigureFunc: providerConfigureTest,
	}
}

func providerConfigureTest(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Address:  os.Getenv("VTHUNDER_HOST"),     //hostname
		Username: os.Getenv("VTHUNDER_USER"),     //username
		Password: os.Getenv("VTHUNDER_PASSWORD"), //password
	}

	return config.Client()
}

func TestAccProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAcctPreCheck(t *testing.T) {
	if v := os.Getenv("VTHUNDER_HOST"); v == "" {
		t.Fatal("ADDRESS must be set for acceptance tests")
	}
	if v := os.Getenv("VTHUNDER_USER"); v == "" {
		t.Fatal("USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("VTHUNDER_PASSWORD"); v == "" {
		t.Fatal("PASSWORD must be set for acceptance tests")
	}
}
