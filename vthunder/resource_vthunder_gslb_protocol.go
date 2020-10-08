package vthunder

//vThunder resource GslbProtocol

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceGslbProtocol() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbProtocolCreate,
		Update: resourceGslbProtocolUpdate,
		Read:   resourceGslbProtocolRead,
		Delete: resourceGslbProtocolDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"use_mgmt_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"msg_format_acos_2x": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ardt_response": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"conn_response": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ardt_session": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ardt_query": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"message": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"response": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ping_site": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auto_detect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_mgmt_port_for_all_partitions": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
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
			"status_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGslbProtocolCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbProtocol (Inside resourceGslbProtocolCreate) ")

		data := dataToGslbProtocol(d)
		logger.Println("[INFO] received formatted data from method data to GslbProtocol --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostGslbProtocol(client.Token, data, client.Host)

		return resourceGslbProtocolRead(d, meta)

	}
	return nil
}

func resourceGslbProtocolRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbProtocol (Inside resourceGslbProtocolRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbProtocol(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbProtocolUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceGslbProtocolRead(d, meta)
}

func resourceGslbProtocolDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceGslbProtocolRead(d, meta)
}
func dataToGslbProtocol(d *schema.ResourceData) go_vthunder.GslbProtocol {
	var vc go_vthunder.GslbProtocol
	var c go_vthunder.GslbProtocolInstance
	c.UseMgmtPort = d.Get("use_mgmt_port").(int)
	c.MsgFormatAcos2X = d.Get("msg_format_acos_2x").(int)

	var obj1 go_vthunder.GslbProtocolLimit
	prefix := "limit.0."
	obj1.ArdtResponse = d.Get(prefix + "ardt_response").(int)
	obj1.ConnResponse = d.Get(prefix + "conn_response").(int)
	obj1.ArdtSession = d.Get(prefix + "ardt_session").(int)
	obj1.ArdtQuery = d.Get(prefix + "ardt_query").(int)
	obj1.Message = d.Get(prefix + "message").(int)
	obj1.Response = d.Get(prefix + "response").(int)

	c.ArdtResponse = obj1

	c.PingSite = d.Get("ping_site").(string)
	c.AutoDetect = d.Get("auto_detect").(int)
	c.UseMgmtPortForAllPartitions = d.Get("use_mgmt_port_for_all_partitions").(int)

	EnableListCount := d.Get("enable_list.#").(int)
	c.Type = make([]go_vthunder.GslbProtocolEnableList, 0, EnableListCount)

	for i := 0; i < EnableListCount; i++ {
		var obj2 go_vthunder.GslbProtocolEnableList
		prefix2 := fmt.Sprintf("enable_list.%d.", i)
		obj2.Type = d.Get(prefix2 + "type").(string)
		c.Type = append(c.Type, obj2)
	}

	c.StatusInterval = d.Get("status_interval").(int)

	vc.UUID = c
	return vc
}
