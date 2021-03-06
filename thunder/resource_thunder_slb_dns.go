package thunder

// Thunder resource Slb DNS

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbDNS() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbDNSCreate,
		Update: resourceSlbDNSUpdate,
		Read:   resourceSlbDNSRead,
		Delete: resourceSlbDNSDelete,

		Schema: map[string]*schema.Schema{
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceSlbDNSCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	logger.Println("[INFO] Creating dns (Inside resourceSlbDNSCreate)")

	if client.Host != "" {
		vc := dataToSlbDNS(d)
		d.SetId("1")
		go_thunder.PostSlbDNS(client.Token, vc, client.Host)

		return resourceSlbDNSRead(d, meta)
	}
	return nil
}

func resourceSlbDNSRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading dns (Inside resourceSlbDNSRead)")

	client := meta.(Thunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbDNS(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No dns found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbDNSUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbDNSRead(d, meta)
}

func resourceSlbDNSDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbDNSRead(d, meta)
}

//Utility method to instantiate DNS Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbDNS(d *schema.ResourceData) go_thunder.SlbDNS {
	var vc go_thunder.SlbDNS

	var c go_thunder.SlbDNSInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable9, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable9
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
