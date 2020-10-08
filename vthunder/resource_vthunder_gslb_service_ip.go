package vthunder

//vThunder resource GslbServiceIp

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceGslbServiceIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbServiceIpCreate,
		Update: resourceGslbServiceIpUpdate,
		Read:   resourceGslbServiceIpRead,
		Delete: resourceGslbServiceIpDelete,
		Schema: map[string]*schema.Schema{
			"health_check_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"port_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_proto": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"health_check_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"follow_port_protocol": {
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
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"health_check_follow_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"health_check_protocol_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"health_check": {
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
			"external_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"node_name": {
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
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_protocol_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGslbServiceIpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbServiceIp (Inside resourceGslbServiceIpCreate) ")
		name := d.Get("node_name").(string)
		data := dataToGslbServiceIp(d)
		logger.Println("[INFO] received formatted data from method data to GslbServiceIp --")
		d.SetId(name)
		go_vthunder.PostGslbServiceIp(client.Token, data, client.Host)

		return resourceGslbServiceIpRead(d, meta)

	}
	return nil
}

func resourceGslbServiceIpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbServiceIp (Inside resourceGslbServiceIpRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbServiceIp(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbServiceIpUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying GslbServiceIp   (Inside resourceGslbServiceIpUpdate) ")
		name := d.Get("node_name").(string)
		data := dataToGslbServiceIp(d)
		logger.Println("[INFO] received formatted data from method data to GslbServiceIp ")
		d.SetId(name)
		go_vthunder.PutGslbServiceIp(client.Token, name, data, client.Host)

		return resourceGslbServiceIpRead(d, meta)

	}
	return nil
}

func resourceGslbServiceIpDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceGslbServiceIpDelete) " + name)
		err := go_vthunder.DeleteGslbServiceIp(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToGslbServiceIp(d *schema.ResourceData) go_vthunder.GslbServiceIp {
	var vc go_vthunder.GslbServiceIp
	var c go_vthunder.GslbServiceIPInstance
	c.HealthCheckDisable = d.Get("health_check_disable").(int)

	PortListCount := d.Get("port_list.#").(int)
	c.PortProto = make([]go_vthunder.GslbServiceIpPortList, 0, PortListCount)

	for i := 0; i < PortListCount; i++ {
		var obj1 go_vthunder.GslbServiceIpPortList
		prefix1 := fmt.Sprintf("port_list.%d.", i)
		obj1.PortProto = d.Get(prefix1 + "port_proto").(string)
		obj1.PortNum = d.Get(prefix1 + "port_num").(int)
		obj1.HealthCheckDisable = d.Get(prefix1 + "health_check_disable").(int)
		obj1.UserTag = d.Get(prefix1 + "user_tag").(string)
		obj1.FollowPortProtocol = d.Get(prefix1 + "follow_port_protocol").(string)

		SamplingEnableCount := d.Get("sampling_enable.#").(int)
		obj1.Counters1 = make([]go_vthunder.GslbServiceIpSamplingEnable, 0, SamplingEnableCount)

		for i := 0; i < SamplingEnableCount; i++ {
			var obj1_1 go_vthunder.GslbServiceIpSamplingEnable
			prefix2 := fmt.Sprintf(prefix1+"sampling_enable.%d.", i)
			obj1_1.Counters1 = d.Get(prefix2 + "counters1").(string)
			obj1.Counters1 = append(obj1.Counters1, obj1_1)
		}

		obj1.Action = d.Get(prefix1 + "action").(string)
		obj1.HealthCheckFollowPort = d.Get(prefix1 + "health_check_follow_port").(int)
		obj1.HealthCheckProtocolDisable = d.Get(prefix1 + "health_check_protocol_disable").(int)
		obj1.HealthCheck = d.Get(prefix1 + "health_check").(string)
		c.PortProto = append(c.PortProto, obj1)
	}

	c.ExternalIP = d.Get("external_ip").(string)
	c.NodeName = d.Get("node_name").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.GslbServiceIpSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj2 go_vthunder.GslbServiceIpSamplingEnable
		prefix2 := fmt.Sprintf("sampling_enable.%d.", i)
		obj2.Counters1 = d.Get(prefix2 + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj2)
	}

	c.Action = d.Get("action").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.IPAddress = d.Get("ip_address").(string)
	c.Ipv6 = d.Get("ipv6").(string)
	c.Ipv6Address = d.Get("ipv6_address").(string)
	c.HealthCheckProtocolDisable = d.Get("health_check_protocol_disable").(int)
	c.HealthCheck = d.Get("health_check").(string)

	vc.HealthCheckDisable = c
	return vc
}
