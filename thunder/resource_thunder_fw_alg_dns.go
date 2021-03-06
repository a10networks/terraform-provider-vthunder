package thunder

//Thunder resource FwAlgDns

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwAlgDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwAlgDnsCreate,
		Update: resourceFwAlgDnsUpdate,
		Read:   resourceFwAlgDnsRead,
		Delete: resourceFwAlgDnsDelete,
		Schema: map[string]*schema.Schema{
			"default_port_disable": {
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
	}
}

func resourceFwAlgDnsCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwAlgDns (Inside resourceFwAlgDnsCreate) ")

		data := dataToFwAlgDns(d)
		logger.Println("[INFO] received formatted data from method data to FwAlgDns --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwAlgDns(client.Token, data, client.Host)

		return resourceFwAlgDnsRead(d, meta)

	}
	return nil
}

func resourceFwAlgDnsRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwAlgDns (Inside resourceFwAlgDnsRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwAlgDns(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwAlgDnsUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgDnsRead(d, meta)
}

func resourceFwAlgDnsDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgDnsRead(d, meta)
}
func dataToFwAlgDns(d *schema.ResourceData) go_thunder.FwAlgDns {
	var vc go_thunder.FwAlgDns
	var c go_thunder.FwAlgDNSInstance
	c.DefaultPortDisable = d.Get("default_port_disable").(string)

	vc.DefaultPortDisable = c
	return vc
}
