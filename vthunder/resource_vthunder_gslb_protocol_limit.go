package vthunder

//vThunder resource GslbProtocolLimit

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceGslbProtocolLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbProtocolLimitCreate,
		Update: resourceGslbProtocolLimitUpdate,
		Read:   resourceGslbProtocolLimitRead,
		Delete: resourceGslbProtocolLimitDelete,
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
	}
}

func resourceGslbProtocolLimitCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbProtocolLimit (Inside resourceGslbProtocolLimitCreate) ")

		data := dataToGslbProtocolLimit(d)
		logger.Println("[INFO] received formatted data from method data to GslbProtocolLimit --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostGslbProtocolLimit(client.Token, data, client.Host)

		return resourceGslbProtocolLimitRead(d, meta)

	}
	return nil
}

func resourceGslbProtocolLimitRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbProtocolLimit (Inside resourceGslbProtocolLimitRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbProtocolLimit(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbProtocolLimitUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceGslbProtocolLimitRead(d, meta)
}

func resourceGslbProtocolLimitDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceGslbProtocolLimitRead(d, meta)
}
func dataToGslbProtocolLimit(d *schema.ResourceData) go_vthunder.GslbProtocolLimitDiff {
	var vc go_vthunder.GslbProtocolLimitDiff
	var c go_vthunder.GslbProtocolLimitInstance
	c.ArdtResponse = d.Get("ardt_response").(int)
	c.ConnResponse = d.Get("conn_response").(int)
	c.ArdtSession = d.Get("ardt_session").(int)
	c.ArdtQuery = d.Get("ardt_query").(int)
	c.Message = d.Get("message").(int)
	c.Response = d.Get("response").(int)

	vc.ArdtResponse = c
	return vc
}
