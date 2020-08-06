package vthunder

//vThunder resource Virtual Server

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVirtualServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualServerCreate,
		Update: resourceVirtualServerUpdate,
		Read:   resourceVirtualServerRead,
		Delete: resourceVirtualServerDelete,

		Schema: map[string]*schema.Schema{
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"arp_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_logging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"stats_data_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_policy_acl_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_acl_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_acl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extended_stats": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_file_inspection": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"use_default_if_no_server": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_tcp_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"use_alternate_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_http_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"when_down_protocol2": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"serv_sel_fail": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"support_http2": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"pool_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_persist_destination_ip_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_udp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alt_protocol1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"alt_protocol2": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_tcp_proxy_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_policy_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv6_acl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"alternate_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_server_ssl_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"when_down": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"pool": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"scaleout_device_group": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"trunk_fwd": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"message_switching": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_udp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_pool": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_persist_ssl_sid_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"no_auto_up_on_aflex": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_persist_ssl_sid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_persist_ssl_sid_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reset": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_external_service_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"trunk_rev": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ipinip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_http_policy_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_respmod_icap": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_diameter_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"scaleout_bucket_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_persist_source_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_virtual_port_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"snat_on_vip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"redirect_to_https": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_cache_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"range": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_smtp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"aflex_scripts": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aflex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"aflex_shared": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ha_conn_mirror": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_ftp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"on_syn": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alternate_port_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_connection_reuse_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_external_service_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"persist_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"force_routing_mode": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_dynamic_service_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_reqmod_icap": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"no_logging": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_dns_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"clientip_sticky_nat": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gslb_enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_client_ssl_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_persist_cookie_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"use_cgnv6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"conn_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"req_fail": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"secs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"use_rcv_hop_for_resp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_persist_cookie_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_tcp_proxy_client": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_persist_source_ip_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_cache": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_dynamic_service": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_fix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"auto": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_http": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_virtual_port": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_persist_destination_ip_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ip_map_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"waf_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"acl_id_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl_id_src_nat_pool": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_id_src_nat_pool_shared": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_id_seq_num_shared": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_id_seq_num": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"shared_partition_pool_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"precedence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"view": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"def_selection_if_pref_failed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_persist_destination_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_dns": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_dblb": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"acl_name_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"shared_partition_pool_name": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_name_seq_num_shared": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_name_src_nat_pool_shared": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_name_src_nat_pool": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_name_seq_num": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"shared_partition_tcp_proxy_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_imap_pop3": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"no_dest_nat": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_persist_cookie": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"reset_on_server_selection_fail": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"rtp_sip_call_id_match": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_udp_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"expand": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"skip_rev_hash": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"enable_playerid_check": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_dynamic_service_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_dns_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_virtual_port_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_external_service": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_scaleout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_http_policy_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_sip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_diameter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_tcp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_client_ssl_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"stats_data_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_smpp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_connection_reuse_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_ssli": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"rate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_connection_reuse": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"syn_cookie": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"optimization_level": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_tcp_proxy_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_persist_source_ip_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_policy_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_tcp_proxy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_http_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_diameter_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_server_ssl_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"eth_fwd": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_http_policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"sampling_enable": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"template_cache_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_client_ssl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_server_ssh": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"auth_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aaa_policy": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"template_client_ssh": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_server_ssl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"l7_hardware_assist": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"eth_rev": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_tcp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"port_translation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"redistribute_route_map": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"acl_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ha_dynamic": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_if_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"netmask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"migrate_vip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_data_cpu": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"target_floating_ipv4": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"target_floating_template_policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"finish_migration": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cancel_migration": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"disable_vip_adv": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vrid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_policy_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_policy_acl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_policy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"enable_disable_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"redistribution_flagged": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_virtual_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_policy_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ethernet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"acl_id_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"acl_name_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"acl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_scaleout": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func validateName() {
	return
}

func resourceVirtualServerCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating virtual server   (Inside resourceVirtualServerCreate    " + name)
		v := dataToVs(name, d)
		logger.Println("[INFO] received formatted data from method data to vs --" + v.Name.Name + ",--" + v.Name.UUID)
		d.SetId(name)
		go_vthunder.PostVS(client.Token, v, client.Host)

		return resourceVirtualServerRead(d, meta)
	}
	return nil
}

func resourceVirtualServerRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading virtual server (Inside resourceVirtualServerRead)")

	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()

		logger.Println("[INFO] Fetching virtual server Read" + name)

		vs, err := go_vthunder.GetVS(client.Token, name, client.Host)

		if vs == nil {
			logger.Println("[INFO] No virtual server found " + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceVirtualServerUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying virtual server   (Inside resourceVirtualServerUpdate    " + name)
		v := dataToVs(name, d)
		logger.Println("[INFO] received formatted data from method data to vs --" + v.Name.Name + ",--" + v.Name.UUID)
		d.SetId(name)
		go_vthunder.PutVS(client.Token, name, v, client.Host)

		return resourceVirtualServerRead(d, meta)
	}
	return nil
}

func resourceVirtualServerDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()
	name := d.Id()
	logger.Println("[INFO] Deleting virtual server (Inside resourceVirtualServerDelete) " + name)

	if client.Host != "" {
		err := go_vthunder.DeleteVS(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete Virtual Server  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//Utility method to instantiate Virtual Server Structure
func dataToVs(name string, d *schema.ResourceData) go_vthunder.VirtalServerInstanceMain {
	var v go_vthunder.VirtalServerInstanceMain

	var vsMain go_vthunder.VirtualServerMain
	vsMain.StatsDataAction = d.Get("stats_data_action").(string)
	vsMain.Ipv6ACLShared = d.Get("ipv6_acl_shared").(string)
	vsMain.ACLName = d.Get("acl_name").(string)
	vsMain.EnableDisableAction = d.Get("enable_disable_action").(string)
	vsMain.HaDynamic = d.Get("ha_dynamic").(int)
	vsMain.RedistributeRouteMap = d.Get("redistribute_route_map").(string)
	vsMain.ACLNameShared = d.Get("acl_name_shared").(string)
	vsMain.IPAddress = d.Get("ip_address").(string)
	vsMain.UseIfIP = d.Get("use_if_ip").(int)
	vsMain.UUID = d.Get("uuid").(string)
	vsMain.Vrid = d.Get("vrid").(int)
	vsMain.DisableVipAdv = d.Get("disable_vip_adv").(int)
	vsMain.TemplateVirtualServer = d.Get("template_virtual_server").(string)
	vsMain.ArpDisable = d.Get("arp_disable").(int)
	vsMain.Description = d.Get("description").(string)
	vsMain.RedistributionFlagged = d.Get("redistribution_flagged").(int)
	vsMain.Netmask = d.Get("netmask").(string)
	vsMain.ACLID = d.Get("acl_id").(int)
	vsMain.Ipv6ACL = d.Get("ipv6_acl").(string)
	vsMain.TemplateLogging = d.Get("template_logging").(string)
	vsMain.ExtendedStats = d.Get("extended_stats").(int)
	vsMain.Name = d.Get("name").(string)
	vsMain.TemplateScaleout = d.Get("template_scaleout").(string)
	vsMain.TemplatePolicy = d.Get("template_policy").(string)
	vsMain.UserTag = d.Get("user_tag").(string)
	vsMain.TemplatePolicyShared = d.Get("template_policy_shared").(string)
	vsMain.Ethernet = d.Get("ethernet").(int)
	vsMain.SharedPartitionPolicyTemplate = d.Get("shared_partition_policy_template").(int)
	vsMain.ACLIDShared = d.Get("acl_id_shared").(int)

	portListCount := d.Get("port_list.#").(int)
	vsMain.Protocol = make([]go_vthunder.PortList, 0, portListCount)
	for i := 0; i < portListCount; i++ {
		var pl go_vthunder.PortList
		prefix := fmt.Sprintf("port_list.%d", i)
		pl.HaConnMirror = d.Get(prefix + "ha_conn_mirror").(int)
		pl.OnSyn = d.Get(prefix + "on_syn").(int)
		pl.Name = d.Get(prefix + ".name").(string)
		pl.OptimizationLevel = d.Get(prefix + ".optimization_level").(string)
		pl.Protocol = d.Get(prefix + ".protocol").(string)
		pl.StatsDataAction = d.Get(prefix + ".stats_data_action").(string)
		pl.PersistType = d.Get(prefix + ".persist_type").(string)
		pl.DefSelectionIfPrefFailed = d.Get(prefix + ".def_selection_if_pref_failed").(string)
		pl.Action = d.Get(prefix + ".action").(string)
		pl.AltProtocol2 = d.Get(prefix + ".alt_protocol2").(string)
		pl.AltProtocol1 = d.Get(prefix + ".alt_protocol1").(string)
		pl.Precedence = d.Get(prefix + ".precedence").(int)
		pl.PortNumber = d.Get(prefix + ".port_number").(int)
		pl.ServiceGroup = d.Get(prefix + ".service_group").(string)
		pl.Auto = d.Get(prefix + ".auto").(int)

		aclNameListCount := d.Get(prefix + ".acl_name_list.#").(int)
		pl.ACLName = make([]go_vthunder.ACLNameList, aclNameListCount, aclNameListCount)
		for x := 0; x < aclNameListCount; x++ {
			var a go_vthunder.ACLNameList
			mapEntity(d.Get(fmt.Sprintf("%s.acl_name_list.%d", prefix, x)).(map[string]interface{}), &a)
			pl.ACLName[x] = a
		}

		vsMain.Protocol = append(vsMain.Protocol, pl)
	}

	v.Name = vsMain

	return v
}
