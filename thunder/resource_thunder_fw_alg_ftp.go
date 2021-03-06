package thunder

//Thunder resource FwAlgFtp

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwAlgFtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwAlgFtpCreate,
		Update: resourceFwAlgFtpUpdate,
		Read:   resourceFwAlgFtpRead,
		Delete: resourceFwAlgFtpDelete,
		Schema: map[string]*schema.Schema{
			"default_port_disable": {
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwAlgFtpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwAlgFtp (Inside resourceFwAlgFtpCreate) ")

		data := dataToFwAlgFtp(d)
		logger.Println("[INFO] received formatted data from method data to FwAlgFtp --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwAlgFtp(client.Token, data, client.Host)

		return resourceFwAlgFtpRead(d, meta)

	}
	return nil
}

func resourceFwAlgFtpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwAlgFtp (Inside resourceFwAlgFtpRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwAlgFtp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwAlgFtpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgFtpRead(d, meta)
}

func resourceFwAlgFtpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgFtpRead(d, meta)
}
func dataToFwAlgFtp(d *schema.ResourceData) go_thunder.FwAlgFtp {
	var vc go_thunder.FwAlgFtp
	var c go_thunder.FwAlgFtpInstance
	c.DefaultPortDisable = d.Get("default_port_disable").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.FwAlgFtpSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.FwAlgFtpSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.DefaultPortDisable = c
	return vc
}
