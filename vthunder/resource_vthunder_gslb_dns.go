package vthunder

//vThunder resource GslbDns

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceGslbDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbDnsCreate,
		Update: resourceGslbDnsUpdate,
		Read:   resourceGslbDnsRead,
		Delete: resourceGslbDnsDelete,
		Schema: map[string]*schema.Schema{
			"action": {
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
			"logging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGslbDnsCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbDns (Inside resourceGslbDnsCreate) ")

		data := dataToGslbDns(d)
		logger.Println("[INFO] received formatted data from method data to GslbDns --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostGslbDns(client.Token, data, client.Host)

		return resourceGslbDnsRead(d, meta)

	}
	return nil
}

func resourceGslbDnsRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbDns (Inside resourceGslbDnsRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbDns(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbDnsUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceGslbDnsRead(d, meta)
}

func resourceGslbDnsDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceGslbDnsRead(d, meta)
}
func dataToGslbDns(d *schema.ResourceData) go_vthunder.GslbDns {
	var vc go_vthunder.GslbDns
	var c go_vthunder.GslbDNSInstance
	c.Action = d.Get("action").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.GslbDnsSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.GslbDnsSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.Logging = d.Get("logging").(string)
	c.Template = d.Get("template").(string)

	vc.Action = c
	return vc
}
