package vthunder

//vThunder resource GslbActiveRdt

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceGslbActiveRdt() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbActiveRdtCreate,
		Update: resourceGslbActiveRdtUpdate,
		Read:   resourceGslbActiveRdtRead,
		Delete: resourceGslbActiveRdtDelete,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"track": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sleep": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGslbActiveRdtCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbActiveRdt (Inside resourceGslbActiveRdtCreate) ")

		data := dataToGslbActiveRdt(d)
		logger.Println("[INFO] received formatted data from method data to GslbActiveRdt --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostGslbActiveRdt(client.Token, data, client.Host)

		return resourceGslbActiveRdtRead(d, meta)

	}
	return nil
}

func resourceGslbActiveRdtRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbActiveRdt (Inside resourceGslbActiveRdtRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbActiveRdt(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbActiveRdtUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceGslbActiveRdtRead(d, meta)
}

func resourceGslbActiveRdtDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceGslbActiveRdtRead(d, meta)
}
func dataToGslbActiveRdt(d *schema.ResourceData) go_vthunder.GslbActiveRdt {
	var vc go_vthunder.GslbActiveRdt
	var c go_vthunder.GslbActiveRdtInstance
	c.Domain = d.Get("domain").(string)
	c.Retry = d.Get("retry").(int)
	c.Track = d.Get("track").(int)
	c.Interval = d.Get("interval").(int)
	c.Sleep = d.Get("sleep").(int)
	c.Timeout = d.Get("timeout").(int)
	c.Icmp = d.Get("icmp").(int)
	c.Port = d.Get("port").(int)

	vc.Domain = c
	return vc
}
