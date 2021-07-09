package thunder

//Thunder resource InterfaceLif

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceInterfaceLif() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceLifCreate,
		Update: resourceInterfaceLifUpdate,
		Read:   resourceInterfaceLifRead,
		Delete: resourceInterfaceLifDelete,
		Schema: map[string]*schema.Schema{
			"ifname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"access_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_id": {
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
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
			"ip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
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
									"ipv6_addr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"anycast": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"link_local": {
										Type:        schema.TypeInt,
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
					},
				},
			},
			"ipv6": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_addr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"anycast": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"link_local": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ipv6_enable": {
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
									"ripng": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rip": {
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
									"ospf": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"area_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"area_id_num": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"area_id_addr": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"tag": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"instance_id": {
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
											},
										},
									},
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
						"ospf": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"broadcast_type": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"p2mp_nbma": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"network_instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
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
									"cost_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"dead_interval_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dead_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"hello_interval_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hello_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"mtu_ignore_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mtu_ignore": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"neighbor_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"neighbor": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"neig_inst": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"neighbor_cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"neighbor_poll_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"neighbor_priority": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"priority_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"retransmit_interval_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"retransmit_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"transmit_delay_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"transmit_delay": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
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
								},
							},
						},
					},
				},
			},
			"bfd": {
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
									"key_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"echo": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"demand": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"interval_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"min_rx": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"multiplier": {
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
					},
				},
			},
			"isis": {
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
									"send_only_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"send_only": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"level": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"mode_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mode": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"level": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"key_chain_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key_chain": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"level": {
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
						"circuit_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"csnp_interval_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"csnp_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"padding": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hello_interval_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"hello_interval_minimal_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_interval_minimal": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"hello_multiplier_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_multiplier": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"lsp_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"mesh_group": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"blocked": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"metric_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metric": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"network": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"password_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"priority_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"retransmit_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"wide_metric_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"wide_metric": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
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
		},
	}
}

func resourceInterfaceLifCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceLif (Inside resourceInterfaceLifCreate) ")
		name1 := d.Get("ifname").(string)
		data := dataToInterfaceLif(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceLif --")
		d.SetId(name1)
		go_thunder.PostInterfaceLif(client.Token, data, client.Host)

		return resourceInterfaceLifRead(d, meta)

	}
	return nil
}

func resourceInterfaceLifRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceLif (Inside resourceInterfaceLifRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetInterfaceLif(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceInterfaceLifUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying InterfaceLif   (Inside resourceInterfaceLifUpdate) ")
		data := dataToInterfaceLif(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceLif ")
		go_thunder.PutInterfaceLif(client.Token, name1, data, client.Host)

		return resourceInterfaceLifRead(d, meta)

	}
	return nil
}

func resourceInterfaceLifDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceInterfaceLifDelete) " + name1)
		err := go_thunder.DeleteInterfaceLif(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToInterfaceLif(d *schema.ResourceData) go_thunder.InterfaceLif {
	var vc go_thunder.InterfaceLif
	var c go_thunder.InterfaceLifInstance
	c.Ifname = d.Get("ifname").(string)
	c.Mtu = d.Get("mtu").(int)
	c.Action = d.Get("action").(string)
	c.Name = d.Get("name").(string)

	var obj1 go_thunder.InterfaceAccessList
	prefix1 := "access_list.0."
	obj1.AclID = d.Get(prefix1 + "acl_id").(int)
	obj1.AclName = d.Get(prefix1 + "acl_name").(string)

	c.AclID = obj1

	c.UserTag = d.Get("user_tag").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.InterfaceSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj2 go_thunder.InterfaceSamplingEnable
		prefix2 := fmt.Sprintf("sampling_enable.%d.", i)
		obj2.Counters1 = d.Get(prefix2 + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj2)
	}

	var obj3 go_thunder.InterfaceIP
	prefix3 := "ip.0."
	obj3.Dhcp = d.Get(prefix3 + "dhcp").(int)

	AddressListCount := d.Get(prefix3 + "address_list.#").(int)
	obj3.Ipv6Addr = make([]go_thunder.InterfaceLifIPAddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj3_1 go_thunder.InterfaceLifIPAddressList
		prefix3_1 := fmt.Sprintf(prefix3+"address_list.%d.", i)
		obj3_1.Ipv6Addr = d.Get(prefix3_1 + "ipv6_addr").(string)
		obj3_1.Anycast = d.Get(prefix3_1 + "anycast").(int)
		obj3_1.LinkLocal = d.Get(prefix3_1 + "link_local").(int)
		obj3.Ipv6Addr = append(obj3.Ipv6Addr, obj3_1)
	}

	obj3.AllowPromiscuousVip = d.Get(prefix3 + "allow_promiscuous_vip").(int)
	obj3.CacheSpoofingPort = d.Get(prefix3 + "cache_spoofing_port").(int)
	obj3.Inside = d.Get(prefix3 + "inside").(int)
	obj3.Outside = d.Get(prefix3 + "outside").(int)
	obj3.GenerateMembershipQuery = d.Get(prefix3 + "generate_membership_query").(int)
	obj3.QueryInterval = d.Get(prefix3 + "query_interval").(int)
	obj3.MaxRespTime = d.Get(prefix3 + "max_resp_time").(int)

	var obj3_2 go_thunder.InterfaceLifIPRouter
	prefix3_2 := prefix3 + "router.0."

	var obj3_2_1 go_thunder.InterfaceLifIPRipIsis
	prefix3_2_1 := prefix3_2 + "isis.0."
	obj3_2_1.Tag = d.Get(prefix3_2_1 + "tag").(string)

	obj3_2.Tag = obj3_2_1

	obj3.Isis = obj3_2

	var obj3_3 go_thunder.InterfaceRip
	prefix3_3 := prefix3 + "rip.0."

	var obj3_3_1 go_thunder.InterfaceLifIpripAuthentication
	prefix3_3_1 := prefix3_3 + "authentication.0."

	var obj3_3_1_1 go_thunder.InterfaceStr
	prefix3_3_1_1 := prefix3_3_1 + "str.0."
	obj3_3_1_1.String = d.Get(prefix3_3_1_1 + "string").(string)

	obj3_3_1.String = obj3_3_1_1

	var obj3_3_1_2 go_thunder.InterfaceMode
	prefix3_3_1_2 := prefix3_3_1 + "mode.0."
	obj3_3_1_2.Mode = d.Get(prefix3_3_1_2 + "mode").(string)

	obj3_3_1.Mode = obj3_3_1_2

	var obj3_3_1_3 go_thunder.InterfaceKeyChain
	prefix3_3_1_3 := prefix3_3_1 + "key_chain.0."
	obj3_3_1_3.KeyChain = d.Get(prefix3_3_1_3 + "key_chain").(string)

	obj3_3_1.KeyChain = obj3_3_1_3

	obj3_3.Str = obj3_3_1

	obj3_3.SendPacket = d.Get(prefix3_3 + "send_packet").(int)
	obj3_3.ReceivePacket = d.Get(prefix3_3 + "receive_packet").(int)

	var obj3_3_2 go_thunder.InterfaceSendCfg
	prefix3_3_2 := prefix3_3 + "send_cfg.0."
	obj3_3_2.Send = d.Get(prefix3_3_2 + "send").(int)
	obj3_3_2.Version = d.Get(prefix3_3_2 + "version").(string)

	obj3_3.Send = obj3_3_2

	var obj3_3_3 go_thunder.InterfaceReceiveCfg
	prefix3_3_3 := prefix3_3 + "receive_cfg.0."
	obj3_3_3.Receive = d.Get(prefix3_3_3 + "receive").(int)
	obj3_3_3.Version = d.Get(prefix3_3_3 + "version").(string)

	obj3_3.Receive = obj3_3_3

	var obj3_3_4 go_thunder.InterfaceSplitHorizonCfg
	prefix3_3_4 := prefix3_3 + "split_horizon_cfg.0."
	obj3_3_4.State = d.Get(prefix3_3_4 + "state").(string)

	obj3_3.State = obj3_3_4

	obj3.Authentication = obj3_3

	var obj3_4 go_thunder.InterfaceLifIpOspf
	prefix3_4 := prefix3 + "ospf.0."

	var obj3_4_1 go_thunder.InterfaceOspfGlobal
	prefix3_4_1 := prefix3_4 + "ospf_global.0."

	var obj3_4_1_1 go_thunder.InterfaceAuthenticationCfg
	prefix3_4_1_1 := prefix3_4_1 + "authentication_cfg.0."
	obj3_4_1_1.Authentication = d.Get(prefix3_4_1_1 + "authentication").(int)
	obj3_4_1_1.Value = d.Get(prefix3_4_1_1 + "value").(string)

	obj3_4_1.Authentication = obj3_4_1_1

	obj3_4_1.AuthenticationKey = d.Get(prefix3_4_1 + "authentication_key").(string)

	var obj3_4_1_2 go_thunder.InterfaceBfdCfg
	prefix3_4_1_2 := prefix3_4_1 + "bfd_cfg.0."
	obj3_4_1_2.Bfd = d.Get(prefix3_4_1_2 + "bfd").(int)
	obj3_4_1_2.Disable = d.Get(prefix3_4_1_2 + "disable").(int)

	obj3_4_1.Bfd = obj3_4_1_2

	obj3_4_1.Cost = d.Get(prefix3_4_1 + "cost").(int)

	var obj3_4_1_3 go_thunder.InterfaceDatabaseFilterCfg
	prefix3_4_1_3 := prefix3_4_1 + "database_filter_cfg.0."
	obj3_4_1_3.DatabaseFilter = d.Get(prefix3_4_1_3 + "database_filter").(string)
	obj3_4_1_3.Out = d.Get(prefix3_4_1_3 + "out").(int)

	obj3_4_1.DatabaseFilter = obj3_4_1_3

	obj3_4_1.DeadInterval = d.Get(prefix3_4_1 + "dead_interval").(int)
	obj3_4_1.Disable = d.Get(prefix3_4_1 + "disable").(string)
	obj3_4_1.HelloInterval = d.Get(prefix3_4_1 + "hello_interval").(int)

	MessageDigestCfgCount := d.Get(prefix3_4_1 + "message_digest_cfg.#").(int)
	obj3_4_1.MessageDigestKey = make([]go_thunder.InterfaceMessageDigestCfg, 0, MessageDigestCfgCount)

	for i := 0; i < MessageDigestCfgCount; i++ {
		var obj3_4_1_4 go_thunder.InterfaceMessageDigestCfg
		prefix3_4_1_4 := fmt.Sprintf(prefix3_4_1+"message_digest_cfg.%d.", i)
		obj3_4_1_4.MessageDigestKey = d.Get(prefix3_4_1_4 + "message_digest_key").(int)
		obj3_4_1_4.Md5Value = d.Get(prefix3_4_1_4 + "md5_value").(string)
		obj3_4_1.MessageDigestKey = append(obj3_4_1.MessageDigestKey, obj3_4_1_4)
	}

	obj3_4_1.Mtu = d.Get(prefix3_4_1 + "mtu").(int)
	obj3_4_1.MtuIgnore = d.Get(prefix3_4_1 + "mtu_ignore").(int)

	var obj3_4_1_5 go_thunder.InterfaceNetwork
	prefix3_4_1_5 := prefix3_4_1 + "network.0."
	obj3_4_1_5.Broadcast = d.Get(prefix3_4_1_5 + "broadcast").(int)
	obj3_4_1_5.NonBroadcast = d.Get(prefix3_4_1_5 + "non_broadcast").(int)
	obj3_4_1_5.PointToPoint = d.Get(prefix3_4_1_5 + "point_to_point").(int)
	obj3_4_1_5.PointToMultipoint = d.Get(prefix3_4_1_5 + "point_to_multipoint").(int)
	obj3_4_1_5.P2MpNbma = d.Get(prefix3_4_1_5 + "p2mp_nbma").(int)

	obj3_4_1.Broadcast = obj3_4_1_5

	obj3_4_1.Priority = d.Get(prefix3_4_1 + "priority").(int)
	obj3_4_1.RetransmitInterval = d.Get(prefix3_4_1 + "retransmit_interval").(int)
	obj3_4_1.TransmitDelay = d.Get(prefix3_4_1 + "transmit_delay").(int)

	obj3_4.AuthenticationCfg = obj3_4_1

	OspfIpListCount := d.Get(prefix3_4 + "ospf_ip_list.#").(int)
	obj3_4.IPAddr = make([]go_thunder.InterfaceOspfIPList, 0, OspfIpListCount)

	for i := 0; i < OspfIpListCount; i++ {
		var obj3_4_2 go_thunder.InterfaceOspfIPList
		prefix3_4_2 := fmt.Sprintf(prefix3_4+"ospf_ip_list.%d.", i)
		obj3_4_2.IPAddr = d.Get(prefix3_4_2 + "ip_addr").(string)
		obj3_4_2.Authentication = d.Get(prefix3_4_2 + "authentication").(int)
		obj3_4_2.Value = d.Get(prefix3_4_2 + "value").(string)
		obj3_4_2.AuthenticationKey = d.Get(prefix3_4_2 + "authentication_key").(string)
		obj3_4_2.Cost = d.Get(prefix3_4_2 + "cost").(int)
		obj3_4_2.DatabaseFilter = d.Get(prefix3_4_2 + "database_filter").(string)
		obj3_4_2.Out = d.Get(prefix3_4_2 + "out").(int)
		obj3_4_2.DeadInterval = d.Get(prefix3_4_2 + "dead_interval").(int)
		obj3_4_2.HelloInterval = d.Get(prefix3_4_2 + "hello_interval").(int)

		MessageDigestCfgCount := d.Get(prefix3_4_2 + "message_digest_cfg.#").(int)
		obj3_4_2.MessageDigestKey = make([]go_thunder.InterfaceMessageDigestCfg, 0, MessageDigestCfgCount)

		for i := 0; i < MessageDigestCfgCount; i++ {
			var obj3_4_2_1 go_thunder.InterfaceMessageDigestCfg
			prefix3_4_2_1 := fmt.Sprintf(prefix3_4_2+"message_digest_cfg.%d.", i)
			obj3_4_2_1.MessageDigestKey = d.Get(prefix3_4_2_1 + "message_digest_key").(int)
			obj3_4_2_1.Md5Value = d.Get(prefix3_4_2_1 + "md5_value").(string)
			obj3_4_2.MessageDigestKey = append(obj3_4_2.MessageDigestKey, obj3_4_2_1)
		}

		obj3_4_2.MtuIgnore = d.Get(prefix3_4_2 + "mtu_ignore").(int)
		obj3_4_2.Priority = d.Get(prefix3_4_2 + "priority").(int)
		obj3_4_2.RetransmitInterval = d.Get(prefix3_4_2 + "retransmit_interval").(int)
		obj3_4_2.TransmitDelay = d.Get(prefix3_4_2 + "transmit_delay").(int)
		obj3_4.IPAddr = append(obj3_4.IPAddr, obj3_4_2)
	}

	obj3.OspfGlobal = obj3_4

	c.Dhcp = obj3

	var obj4 go_thunder.InterfaceIpv6
	prefix4 := "ipv6.0."

	AddressListCount = d.Get(prefix4 + "address_list.#").(int)
	obj4.Ipv6Addr = make([]go_thunder.InterfaceLifIPAddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj4_1 go_thunder.InterfaceLifIPAddressList
		prefix4_1 := fmt.Sprintf(prefix4+"address_list.%d.", i)
		obj4_1.Ipv6Addr = d.Get(prefix4_1 + "ipv6_addr").(string)
		obj4_1.Anycast = d.Get(prefix4_1 + "anycast").(int)
		obj4_1.LinkLocal = d.Get(prefix4_1 + "link_local").(int)
		obj4.Ipv6Addr = append(obj4.Ipv6Addr, obj4_1)
	}

	obj4.Ipv6Enable = d.Get(prefix4 + "ipv6_enable").(int)
	obj4.Inside = d.Get(prefix4 + "inside").(int)
	obj4.Outside = d.Get(prefix4 + "outside").(int)

	var obj4_2 go_thunder.InterfaceRouter
	prefix4_2 := prefix4 + "router.0."

	var obj4_2_1 go_thunder.InterfaceRipng
	prefix4_2_1 := prefix4_2 + "ripng.0."
	obj4_2_1.Rip = d.Get(prefix4_2_1 + "rip").(int)

	obj4_2.Rip = obj4_2_1

	var obj4_2_2 go_thunder.InterfaceLifIpv6RouterOspf
	prefix4_2_2 := prefix4_2 + "ospf.0."

	AreaListCount := d.Get(prefix4_2_2 + "area_list.#").(int)
	obj4_2_2.AreaIDNum = make([]go_thunder.InterfaceAreaList, 0, AreaListCount)

	for i := 0; i < AreaListCount; i++ {
		var obj4_2_2_1 go_thunder.InterfaceAreaList
		prefix4_2_2_1 := fmt.Sprintf(prefix4_2_2+"area_list.%d.", i)
		obj4_2_2_1.AreaIDNum = d.Get(prefix4_2_2_1 + "area_id_num").(int)
		obj4_2_2_1.AreaIDAddr = d.Get(prefix4_2_2_1 + "area_id_addr").(string)
		obj4_2_2_1.Tag = d.Get(prefix4_2_2_1 + "tag").(string)
		obj4_2_2_1.InstanceID = d.Get(prefix4_2_2_1 + "instance_id").(int)
		obj4_2_2.AreaIDNum = append(obj4_2_2.AreaIDNum, obj4_2_2_1)
	}

	obj4_2.AreaList = obj4_2_2

	var obj4_2_3 go_thunder.InterfaceLifIPRipIsis
	prefix4_2_3 := prefix4_2 + "isis.0."
	obj4_2_3.Tag = d.Get(prefix4_2_3 + "tag").(string)

	obj4_2.Tag = obj4_2_3

	obj4.Ripng = obj4_2

	var obj4_3 go_thunder.InterfaceOspf
	prefix4_3 := prefix4 + "ospf.0."

	NetworkListCount := d.Get(prefix4_3 + "network_list.#").(int)
	obj4_3.BroadcastType = make([]go_thunder.InterfaceNetworkList, 0, NetworkListCount)

	for i := 0; i < NetworkListCount; i++ {
		var obj4_3_1 go_thunder.InterfaceNetworkList
		prefix4_3_1 := fmt.Sprintf(prefix4_3+"network_list.%d.", i)
		obj4_3_1.BroadcastType = d.Get(prefix4_3_1 + "broadcast_type").(string)
		obj4_3_1.P2MpNbma = d.Get(prefix4_3_1 + "p2mp_nbma").(int)
		obj4_3_1.NetworkInstanceID = d.Get(prefix4_3_1 + "network_instance_id").(int)
		obj4_3.BroadcastType = append(obj4_3.BroadcastType, obj4_3_1)
	}

	obj4_3.Bfd = d.Get(prefix4_3 + "bfd").(int)
	obj4_3.Disable = d.Get(prefix4_3 + "disable").(int)

	CostCfgCount := d.Get(prefix4_3 + "cost_cfg.#").(int)
	obj4_3.Cost = make([]go_thunder.InterfaceCostCfg, 0, CostCfgCount)

	for i := 0; i < CostCfgCount; i++ {
		var obj4_3_2 go_thunder.InterfaceCostCfg
		prefix4_3_2 := fmt.Sprintf(prefix4_3+"cost_cfg.%d.", i)
		obj4_3_2.Cost = d.Get(prefix4_3_2 + "cost").(int)
		obj4_3_2.InstanceID = d.Get(prefix4_3_2 + "instance_id").(int)
		obj4_3.Cost = append(obj4_3.Cost, obj4_3_2)
	}

	DeadIntervalCfgCount := d.Get(prefix4_3 + "dead_interval_cfg.#").(int)
	obj4_3.DeadInterval = make([]go_thunder.InterfaceDeadIntervalCfg, 0, DeadIntervalCfgCount)

	for i := 0; i < DeadIntervalCfgCount; i++ {
		var obj4_3_3 go_thunder.InterfaceDeadIntervalCfg
		prefix4_3_3 := fmt.Sprintf(prefix4_3+"dead_interval_cfg.%d.", i)
		obj4_3_3.DeadInterval = d.Get(prefix4_3_3 + "dead_interval").(int)
		obj4_3_3.InstanceID = d.Get(prefix4_3_3 + "instance_id").(int)
		obj4_3.DeadInterval = append(obj4_3.DeadInterval, obj4_3_3)
	}

	HelloIntervalCfgCount := d.Get(prefix4_3 + "hello_interval_cfg.#").(int)
	obj4_3.HelloInterval = make([]go_thunder.InterfaceHelloIntervalCfg, 0, HelloIntervalCfgCount)

	for i := 0; i < HelloIntervalCfgCount; i++ {
		var obj4_3_4 go_thunder.InterfaceHelloIntervalCfg
		prefix4_3_4 := fmt.Sprintf(prefix4_3+"hello_interval_cfg.%d.", i)
		obj4_3_4.HelloInterval = d.Get(prefix4_3_4 + "hello_interval").(int)
		obj4_3_4.InstanceID = d.Get(prefix4_3_4 + "instance_id").(int)
		obj4_3.HelloInterval = append(obj4_3.HelloInterval, obj4_3_4)
	}

	MtuIgnoreCfgCount := d.Get(prefix4_3 + "mtu_ignore_cfg.#").(int)
	obj4_3.MtuIgnore = make([]go_thunder.InterfaceMtuIgnoreCfg, 0, MtuIgnoreCfgCount)

	for i := 0; i < MtuIgnoreCfgCount; i++ {
		var obj4_3_5 go_thunder.InterfaceMtuIgnoreCfg
		prefix4_3_5 := fmt.Sprintf(prefix4_3+"mtu_ignore_cfg.%d.", i)
		obj4_3_5.MtuIgnore = d.Get(prefix4_3_5 + "mtu_ignore").(int)
		obj4_3_5.InstanceID = d.Get(prefix4_3_5 + "instance_id").(int)
		obj4_3.MtuIgnore = append(obj4_3.MtuIgnore, obj4_3_5)
	}

	NeighborCfgCount := d.Get(prefix4_3 + "neighbor_cfg.#").(int)
	obj4_3.Neighbor = make([]go_thunder.InterfaceNeighborCfg, 0, NeighborCfgCount)

	for i := 0; i < NeighborCfgCount; i++ {
		var obj4_3_6 go_thunder.InterfaceNeighborCfg
		prefix4_3_6 := fmt.Sprintf(prefix4_3+"neighbor_cfg.%d.", i)
		obj4_3_6.Neighbor = d.Get(prefix4_3_6 + "neighbor").(string)
		obj4_3_6.NeigInst = d.Get(prefix4_3_6 + "neig_inst").(int)
		obj4_3_6.NeighborCost = d.Get(prefix4_3_6 + "neighbor_cost").(int)
		obj4_3_6.NeighborPollInterval = d.Get(prefix4_3_6 + "neighbor_poll_interval").(int)
		obj4_3_6.NeighborPriority = d.Get(prefix4_3_6 + "neighbor_priority").(int)
		obj4_3.Neighbor = append(obj4_3.Neighbor, obj4_3_6)
	}

	PriorityCfgCount := d.Get(prefix4_3 + "priority_cfg.#").(int)
	obj4_3.Priority = make([]go_thunder.InterfacePriorityCfg, 0, PriorityCfgCount)

	for i := 0; i < PriorityCfgCount; i++ {
		var obj4_3_7 go_thunder.InterfacePriorityCfg
		prefix4_3_7 := fmt.Sprintf(prefix4_3+"priority_cfg.%d.", i)
		obj4_3_7.Priority = d.Get(prefix4_3_7 + "priority").(int)
		obj4_3_7.InstanceID = d.Get(prefix4_3_7 + "instance_id").(int)
		obj4_3.Priority = append(obj4_3.Priority, obj4_3_7)
	}

	RetransmitIntervalCfgCount := d.Get(prefix4_3 + "retransmit_interval_cfg.#").(int)
	obj4_3.RetransmitInterval = make([]go_thunder.InterfaceRetransmitIntervalCfg, 0, RetransmitIntervalCfgCount)

	for i := 0; i < RetransmitIntervalCfgCount; i++ {
		var obj4_3_8 go_thunder.InterfaceRetransmitIntervalCfg
		prefix4_3_8 := fmt.Sprintf(prefix4_3+"retransmit_interval_cfg.%d.", i)
		obj4_3_8.RetransmitInterval = d.Get(prefix4_3_8 + "retransmit_interval").(int)
		obj4_3_8.InstanceID = d.Get(prefix4_3_8 + "instance_id").(int)
		obj4_3.RetransmitInterval = append(obj4_3.RetransmitInterval, obj4_3_8)
	}

	TransmitDelayCfgCount := d.Get(prefix4_3 + "transmit_delay_cfg.#").(int)
	obj4_3.TransmitDelay = make([]go_thunder.InterfaceTransmitDelayCfg, 0, TransmitDelayCfgCount)

	for i := 0; i < TransmitDelayCfgCount; i++ {
		var obj4_3_9 go_thunder.InterfaceTransmitDelayCfg
		prefix4_3_9 := fmt.Sprintf(prefix4_3+"transmit_delay_cfg.%d.", i)
		obj4_3_9.TransmitDelay = d.Get(prefix4_3_9 + "transmit_delay").(int)
		obj4_3_9.InstanceID = d.Get(prefix4_3_9 + "instance_id").(int)
		obj4_3.TransmitDelay = append(obj4_3.TransmitDelay, obj4_3_9)
	}

	obj4.NetworkList = obj4_3

	c.AddressList = obj4

	var obj5 go_thunder.InterfaceBfd
	prefix5 := "bfd.0."

	var obj5_1 go_thunder.InterfaceAuthentication
	prefix5_1 := prefix5 + "authentication.0."
	obj5_1.KeyID = d.Get(prefix5_1 + "key_id").(int)
	obj5_1.Method = d.Get(prefix5_1 + "method").(string)
	obj5_1.Password = d.Get(prefix5_1 + "password").(string)

	obj5.KeyID = obj5_1

	obj5.Echo = d.Get(prefix5 + "echo").(int)
	obj5.Demand = d.Get(prefix5 + "demand").(int)

	var obj5_2 go_thunder.InterfaceIntervalCfg
	prefix5_2 := prefix5 + "interval_cfg.0."
	obj5_2.Interval = d.Get(prefix5_2 + "interval").(int)
	obj5_2.MinRx = d.Get(prefix5_2 + "min_rx").(int)
	obj5_2.Multiplier = d.Get(prefix5_2 + "multiplier").(int)

	obj5.Interval = obj5_2

	c.Authentication = obj5

	var obj6 go_thunder.InterfaceIsis
	prefix6 := "isis.0."

	var obj6_1 go_thunder.InterfaceLifIsisAuthentication
	prefix6_1 := prefix6 + "authentication.0."

	SendOnlyListCount := d.Get(prefix6_1 + "send_only_list.#").(int)
	obj6_1.SendOnly = make([]go_thunder.InterfaceSendOnlyList, 0, SendOnlyListCount)

	for i := 0; i < SendOnlyListCount; i++ {
		var obj6_1_1 go_thunder.InterfaceSendOnlyList
		prefix6_1_1 := fmt.Sprintf(prefix6_1+"send_only_list.%d.", i)
		obj6_1_1.SendOnly = d.Get(prefix6_1_1 + "send_only").(int)
		obj6_1_1.Level = d.Get(prefix6_1_1 + "level").(string)
		obj6_1.SendOnly = append(obj6_1.SendOnly, obj6_1_1)
	}

	ModeListCount := d.Get(prefix6_1 + "mode_list.#").(int)
	obj6_1.Mode = make([]go_thunder.InterfaceModeList, 0, ModeListCount)

	for i := 0; i < ModeListCount; i++ {
		var obj6_1_2 go_thunder.InterfaceModeList
		prefix6_1_2 := fmt.Sprintf(prefix6_1+"mode_list.%d.", i)
		obj6_1_2.Mode = d.Get(prefix6_1_2 + "mode").(string)
		obj6_1_2.Level = d.Get(prefix6_1_2 + "level").(string)
		obj6_1.Mode = append(obj6_1.Mode, obj6_1_2)
	}

	KeyChainListCount := d.Get(prefix6_1 + "key_chain_list.#").(int)
	obj6_1.KeyChain = make([]go_thunder.InterfaceKeyChainList, 0, KeyChainListCount)

	for i := 0; i < KeyChainListCount; i++ {
		var obj6_1_3 go_thunder.InterfaceKeyChainList
		prefix6_1_3 := fmt.Sprintf(prefix6_1+"key_chain_list.%d.", i)
		obj6_1_3.KeyChain = d.Get(prefix6_1_3 + "key_chain").(string)
		obj6_1_3.Level = d.Get(prefix6_1_3 + "level").(string)
		obj6_1.KeyChain = append(obj6_1.KeyChain, obj6_1_3)
	}

	obj6.SendOnlyList = obj6_1

	var obj6_2 go_thunder.InterfaceBfdCfg
	prefix6_2 := prefix6 + "bfd_cfg.0."
	obj6_2.Bfd = d.Get(prefix6_2 + "bfd").(int)
	obj6_2.Disable = d.Get(prefix6_2 + "disable").(int)

	obj6.Bfd = obj6_2

	obj6.CircuitType = d.Get(prefix6 + "circuit_type").(string)

	CsnpIntervalListCount := d.Get(prefix6 + "csnp_interval_list.#").(int)
	obj6.CsnpInterval = make([]go_thunder.InterfaceCsnpIntervalList, 0, CsnpIntervalListCount)

	for i := 0; i < CsnpIntervalListCount; i++ {
		var obj6_3 go_thunder.InterfaceCsnpIntervalList
		prefix6_3 := fmt.Sprintf(prefix6+"csnp_interval_list.%d.", i)
		obj6_3.CsnpInterval = d.Get(prefix6_3 + "csnp_interval").(int)
		obj6_3.Level = d.Get(prefix6_3 + "level").(string)
		obj6.CsnpInterval = append(obj6.CsnpInterval, obj6_3)
	}

	obj6.Padding = d.Get(prefix6 + "padding").(int)

	HelloIntervalListCount := d.Get(prefix6 + "hello_interval_list.#").(int)
	obj6.HelloInterval = make([]go_thunder.InterfaceHelloIntervalList, 0, HelloIntervalListCount)

	for i := 0; i < HelloIntervalListCount; i++ {
		var obj6_4 go_thunder.InterfaceHelloIntervalList
		prefix6_4 := fmt.Sprintf(prefix6+"hello_interval_list.%d.", i)
		obj6_4.HelloInterval = d.Get(prefix6_4 + "hello_interval").(int)
		obj6_4.Level = d.Get(prefix6_4 + "level").(string)
		obj6.HelloInterval = append(obj6.HelloInterval, obj6_4)
	}

	HelloIntervalMinimalListCount := d.Get(prefix6 + "hello_interval_minimal_list.#").(int)
	obj6.HelloIntervalMinimal = make([]go_thunder.InterfaceHelloIntervalMinimalList, 0, HelloIntervalMinimalListCount)

	for i := 0; i < HelloIntervalMinimalListCount; i++ {
		var obj6_5 go_thunder.InterfaceHelloIntervalMinimalList
		prefix6_5 := fmt.Sprintf(prefix6+"hello_interval_minimal_list.%d.", i)
		obj6_5.HelloIntervalMinimal = d.Get(prefix6_5 + "hello_interval_minimal").(int)
		obj6_5.Level = d.Get(prefix6_5 + "level").(string)
		obj6.HelloIntervalMinimal = append(obj6.HelloIntervalMinimal, obj6_5)
	}

	HelloMultiplierListCount := d.Get(prefix6 + "hello_multiplier_list.#").(int)
	obj6.HelloMultiplier = make([]go_thunder.InterfaceHelloMultiplierList, 0, HelloMultiplierListCount)

	for i := 0; i < HelloMultiplierListCount; i++ {
		var obj6_6 go_thunder.InterfaceHelloMultiplierList
		prefix6_6 := fmt.Sprintf(prefix6+"hello_multiplier_list.%d.", i)
		obj6_6.HelloMultiplier = d.Get(prefix6_6 + "hello_multiplier").(int)
		obj6_6.Level = d.Get(prefix6_6 + "level").(string)
		obj6.HelloMultiplier = append(obj6.HelloMultiplier, obj6_6)
	}

	obj6.LspInterval = d.Get(prefix6 + "lsp_interval").(int)

	var obj6_7 go_thunder.InterfaceMeshGroup
	prefix6_7 := prefix6 + "mesh_group.0."
	obj6_7.Value = d.Get(prefix6_7 + "value").(int)
	obj6_7.Blocked = d.Get(prefix6_7 + "blocked").(int)

	obj6.Value = obj6_7

	MetricListCount := d.Get(prefix6 + "metric_list.#").(int)
	obj6.Metric = make([]go_thunder.InterfaceMetricList, 0, MetricListCount)

	for i := 0; i < MetricListCount; i++ {
		var obj6_8 go_thunder.InterfaceMetricList
		prefix6_8 := fmt.Sprintf(prefix6+"metric_list.%d.", i)
		obj6_8.Metric = d.Get(prefix6_8 + "metric").(int)
		obj6_8.Level = d.Get(prefix6_8 + "level").(string)
		obj6.Metric = append(obj6.Metric, obj6_8)
	}

	obj6.Network = d.Get(prefix6 + "network").(string)

	PasswordListCount := d.Get(prefix6 + "password_list.#").(int)
	obj6.Password = make([]go_thunder.InterfacePasswordList, 0, PasswordListCount)

	for i := 0; i < PasswordListCount; i++ {
		var obj6_9 go_thunder.InterfacePasswordList
		prefix6_9 := fmt.Sprintf(prefix6+"password_list.%d.", i)
		obj6_9.Password = d.Get(prefix6_9 + "password").(string)
		obj6_9.Level = d.Get(prefix6_9 + "level").(string)
		obj6.Password = append(obj6.Password, obj6_9)
	}

	PriorityListCount := d.Get(prefix6 + "priority_list.#").(int)
	obj6.Priority = make([]go_thunder.InterfacePriorityList, 0, PriorityListCount)

	for i := 0; i < PriorityListCount; i++ {
		var obj6_10 go_thunder.InterfacePriorityList
		prefix6_10 := fmt.Sprintf(prefix6+"priority_list.%d.", i)
		obj6_10.Priority = d.Get(prefix6_10 + "priority").(int)
		obj6_10.Level = d.Get(prefix6_10 + "level").(string)
		obj6.Priority = append(obj6.Priority, obj6_10)
	}

	obj6.RetransmitInterval = d.Get(prefix6 + "retransmit_interval").(int)

	WideMetricListCount := d.Get(prefix6 + "wide_metric_list.#").(int)
	obj6.WideMetric = make([]go_thunder.InterfaceWideMetricList, 0, WideMetricListCount)

	for i := 0; i < WideMetricListCount; i++ {
		var obj6_11 go_thunder.InterfaceWideMetricList
		prefix6_11 := fmt.Sprintf(prefix6+"wide_metric_list.%d.", i)
		obj6_11.WideMetric = d.Get(prefix6_11 + "wide_metric").(int)
		obj6_11.Level = d.Get(prefix6_11 + "level").(string)
		obj6.WideMetric = append(obj6.WideMetric, obj6_11)
	}

	c.SendOnlyList = obj6

	vc.Ifname = c
	return vc
}
