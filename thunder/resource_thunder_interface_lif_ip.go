package thunder

//Thunder resource InterfaceLifIp

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceInterfaceLifIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceLifIpCreate,
		Update: resourceInterfaceLifIpUpdate,
		Read:   resourceInterfaceLifIpRead,
		Delete: resourceInterfaceLifIpDelete,
		Schema: map[string]*schema.Schema{
			"dhcp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv4_netmask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"allow_promiscuous_vip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cache_spoofing_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"inside": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"outside": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"generate_membership_query": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"query_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_resp_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"router": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"isis": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"rip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"string": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"mode": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mode": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"key_chain": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key_chain": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
								},
							},
						},
						"send_packet": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"receive_packet": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"send_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"send": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"version": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"receive_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"receive": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"version": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"split_horizon_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": {
										Type:        schema.TypeString,
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
					},
				},
			},
			"ospf": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ospf_global": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"authentication_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"bfd_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bfd": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"disable": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"database_filter_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"database_filter": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"out": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"dead_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"disable": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"message_digest_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"message_digest_key": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"md5_value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"mtu": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"mtu_ignore": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"network": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"broadcast": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"non_broadcast": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"point_to_point": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"point_to_multipoint": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"p2mp_nbma": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"retransmit_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"transmit_delay": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ospf_ip_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_addr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"authentication": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"authentication_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"database_filter": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"out": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dead_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"message_digest_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"message_digest_key": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"md5_value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"mtu_ignore": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"retransmit_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"transmit_delay": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"ifname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceInterfaceLifIpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceLifIp (Inside resourceInterfaceLifIpCreate) ")
		name1 := d.Get("ifname").(string)
		data := dataToInterfaceLifIp(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceLifIp --")
		d.SetId(name1)
		go_thunder.PostInterfaceLifIp(client.Token, name1, data, client.Host)

		return resourceInterfaceLifIpRead(d, meta)

	}
	return nil
}

func resourceInterfaceLifIpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceLifIp (Inside resourceInterfaceLifIpRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetInterfaceLifIp(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceInterfaceLifIpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceInterfaceLifIpRead(d, meta)
}

func resourceInterfaceLifIpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceInterfaceLifIpRead(d, meta)
}
func dataToInterfaceLifIp(d *schema.ResourceData) go_thunder.InterfaceLifIp {
	var vc go_thunder.InterfaceLifIp
	var c go_thunder.InterfaceLifIPInstance
	c.Dhcp = d.Get("dhcp").(int)

	AddressListCount := d.Get("address_list.#").(int)
	c.Ipv4Address = make([]go_thunder.InterfaceLifAddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj1 go_thunder.InterfaceLifAddressList
		prefix1 := fmt.Sprintf("address_list.%d.", i)
		obj1.Ipv4Address = d.Get(prefix1 + "ipv4_address").(string)
		obj1.Ipv4Netmask = d.Get(prefix1 + "ipv4_netmask").(string)
		c.Ipv4Address = append(c.Ipv4Address, obj1)
	}

	c.AllowPromiscuousVip = d.Get("allow_promiscuous_vip").(int)
	c.CacheSpoofingPort = d.Get("cache_spoofing_port").(int)
	c.Inside = d.Get("inside").(int)
	c.Outside = d.Get("outside").(int)
	c.GenerateMembershipQuery = d.Get("generate_membership_query").(int)
	c.QueryInterval = d.Get("query_interval").(int)
	c.MaxRespTime = d.Get("max_resp_time").(int)

	var obj2 go_thunder.InterfaceLifRouter
	prefix2 := "router.0."

	var obj2_1 go_thunder.InterfaceLifIsis
	prefix2_1 := prefix2 + "isis.0."
	obj2_1.Tag = d.Get(prefix2_1 + "tag").(string)

	obj2.Tag = obj2_1

	c.Isis = obj2

	var obj3 go_thunder.InterfaceLifRip
	prefix3 := "rip.0."

	var obj3_1 go_thunder.InterfaceLifAuthentication
	prefix3_1 := prefix3 + "authentication.0."

	var obj3_1_1 go_thunder.InterfaceLifStr
	prefix3_1_1 := prefix3_1 + "str.0."
	obj3_1_1.String = d.Get(prefix3_1_1 + "string").(string)

	obj3_1.String = obj3_1_1

	var obj3_1_2 go_thunder.InterfaceLifMode
	prefix3_1_2 := prefix3_1 + "mode.0."
	obj3_1_2.Mode = d.Get(prefix3_1_2 + "mode").(string)

	obj3_1.Mode = obj3_1_2

	var obj3_1_3 go_thunder.InterfaceLifKeyChain
	prefix3_1_3 := prefix3_1 + "key_chain.0."
	obj3_1_3.KeyChain = d.Get(prefix3_1_3 + "key_chain").(string)

	obj3_1.KeyChain = obj3_1_3

	obj3.Str = obj3_1

	obj3.SendPacket = d.Get(prefix3 + "send_packet").(int)
	obj3.ReceivePacket = d.Get(prefix3 + "receive_packet").(int)

	var obj3_2 go_thunder.InterfaceLifSendCfg
	prefix3_2 := prefix3 + "send_cfg.0."
	obj3_2.Send = d.Get(prefix3_2 + "send").(int)
	obj3_2.Version = d.Get(prefix3_2 + "version").(string)

	obj3.Send = obj3_2

	var obj3_3 go_thunder.InterfaceLifReceiveCfg
	prefix3_3 := prefix3 + "receive_cfg.0."
	obj3_3.Receive = d.Get(prefix3_3 + "receive").(int)
	obj3_3.Version = d.Get(prefix3_3 + "version").(string)

	obj3.Receive = obj3_3

	var obj3_4 go_thunder.InterfaceLifSplitHorizonCfg
	prefix3_4 := prefix3 + "split_horizon_cfg.0."
	obj3_4.State = d.Get(prefix3_4 + "state").(string)

	obj3.State = obj3_4

	c.Authentication = obj3

	var obj4 go_thunder.InterfaceLifOspf
	prefix4 := "ospf.0."

	var obj4_1 go_thunder.InterfaceLifOspfGlobal
	prefix4_1 := prefix4 + "ospf_global.0."

	var obj4_1_1 go_thunder.InterfaceLifAuthenticationCfg
	prefix4_1_1 := prefix4_1 + "authentication_cfg.0."
	obj4_1_1.Authentication = d.Get(prefix4_1_1 + "authentication").(int)
	obj4_1_1.Value = d.Get(prefix4_1_1 + "value").(string)

	obj4_1.Authentication = obj4_1_1

	obj4_1.AuthenticationKey = d.Get(prefix4_1 + "authentication_key").(string)

	var obj4_1_2 go_thunder.InterfaceLifBfdCfg
	prefix4_1_2 := prefix4_1 + "bfd_cfg.0."
	obj4_1_2.Bfd = d.Get(prefix4_1_2 + "bfd").(int)
	obj4_1_2.Disable = d.Get(prefix4_1_2 + "disable").(int)

	obj4_1.Bfd = obj4_1_2

	obj4_1.Cost = d.Get(prefix4_1 + "cost").(int)

	var obj4_1_3 go_thunder.InterfaceLifDatabaseFilterCfg
	prefix4_1_3 := prefix4_1 + "database_filter_cfg.0."
	obj4_1_3.DatabaseFilter = d.Get(prefix4_1_3 + "database_filter").(string)
	obj4_1_3.Out = d.Get(prefix4_1_3 + "out").(int)

	obj4_1.DatabaseFilter = obj4_1_3

	obj4_1.DeadInterval = d.Get(prefix4_1 + "dead_interval").(int)
	obj4_1.Disable = d.Get(prefix4_1 + "disable").(string)
	obj4_1.HelloInterval = d.Get(prefix4_1 + "hello_interval").(int)

	MessageDigestCfgCount := d.Get(prefix4_1 + "message_digest_cfg.#").(int)
	obj4_1.MessageDigestKey = make([]go_thunder.InterfaceLifMessageDigestCfg, 0, MessageDigestCfgCount)

	for i := 0; i < MessageDigestCfgCount; i++ {
		var obj4_1_4 go_thunder.InterfaceLifMessageDigestCfg
		prefix4_1_4 := fmt.Sprintf(prefix4_1+"message_digest_cfg.%d.", i)
		obj4_1_4.MessageDigestKey = d.Get(prefix4_1_4 + "message_digest_key").(int)
		obj4_1_4.Md5Value = d.Get(prefix4_1_4 + "md5_value").(string)
		obj4_1.MessageDigestKey = append(obj4_1.MessageDigestKey, obj4_1_4)
	}

	obj4_1.Mtu = d.Get(prefix4_1 + "mtu").(int)
	obj4_1.MtuIgnore = d.Get(prefix4_1 + "mtu_ignore").(int)

	var obj4_1_5 go_thunder.InterfaceLifNetwork
	prefix4_1_5 := prefix4_1 + "network.0."
	obj4_1_5.Broadcast = d.Get(prefix4_1_5 + "broadcast").(int)
	obj4_1_5.NonBroadcast = d.Get(prefix4_1_5 + "non_broadcast").(int)
	obj4_1_5.PointToPoint = d.Get(prefix4_1_5 + "point_to_point").(int)
	obj4_1_5.PointToMultipoint = d.Get(prefix4_1_5 + "point_to_multipoint").(int)
	obj4_1_5.P2MpNbma = d.Get(prefix4_1_5 + "p2mp_nbma").(int)

	obj4_1.Broadcast = obj4_1_5

	obj4_1.Priority = d.Get(prefix4_1 + "priority").(int)
	obj4_1.RetransmitInterval = d.Get(prefix4_1 + "retransmit_interval").(int)
	obj4_1.TransmitDelay = d.Get(prefix4_1 + "transmit_delay").(int)

	obj4.AuthenticationCfg = obj4_1

	OspfIpListCount := d.Get(prefix4 + "ospf_ip_list.#").(int)
	obj4.IPAddr = make([]go_thunder.InterfaceLifOspfIPList, 0, OspfIpListCount)

	for i := 0; i < OspfIpListCount; i++ {
		var obj4_2 go_thunder.InterfaceLifOspfIPList
		prefix4_2 := fmt.Sprintf(prefix4+"ospf_ip_list.%d.", i)
		obj4_2.IPAddr = d.Get(prefix4_2 + "ip_addr").(string)
		obj4_2.Authentication = d.Get(prefix4_2 + "authentication").(int)
		obj4_2.Value = d.Get(prefix4_2 + "value").(string)
		obj4_2.AuthenticationKey = d.Get(prefix4_2 + "authentication_key").(string)
		obj4_2.Cost = d.Get(prefix4_2 + "cost").(int)
		obj4_2.DatabaseFilter = d.Get(prefix4_2 + "database_filter").(string)
		obj4_2.Out = d.Get(prefix4_2 + "out").(int)
		obj4_2.DeadInterval = d.Get(prefix4_2 + "dead_interval").(int)
		obj4_2.HelloInterval = d.Get(prefix4_2 + "hello_interval").(int)

		MessageDigestCfgCount := d.Get(prefix4_2 + "message_digest_cfg.#").(int)
		obj4_2.MessageDigestKey = make([]go_thunder.InterfaceLifMessageDigestCfg, 0, MessageDigestCfgCount)

		for i := 0; i < MessageDigestCfgCount; i++ {
			var obj4_2_1 go_thunder.InterfaceLifMessageDigestCfg
			prefix4_2_1 := fmt.Sprintf(prefix4_2+"message_digest_cfg.%d.", i)
			obj4_2_1.MessageDigestKey = d.Get(prefix4_2_1 + "message_digest_key").(int)
			obj4_2_1.Md5Value = d.Get(prefix4_2_1 + "md5_value").(string)
			obj4_2.MessageDigestKey = append(obj4_2.MessageDigestKey, obj4_2_1)
		}

		obj4_2.MtuIgnore = d.Get(prefix4_2 + "mtu_ignore").(int)
		obj4_2.Priority = d.Get(prefix4_2 + "priority").(int)
		obj4_2.RetransmitInterval = d.Get(prefix4_2 + "retransmit_interval").(int)
		obj4_2.TransmitDelay = d.Get(prefix4_2 + "transmit_delay").(int)
		obj4.IPAddr = append(obj4.IPAddr, obj4_2)
	}

	c.OspfGlobal = obj4

	vc.Dhcp = c
	return vc
}
