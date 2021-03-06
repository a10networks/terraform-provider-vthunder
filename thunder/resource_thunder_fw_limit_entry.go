package thunder

//Thunder resource FwLimitEntry

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwLimitEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwLimitEntryCreate,
		Update: resourceFwLimitEntryUpdate,
		Read:   resourceFwLimitEntryRead,
		Delete: resourceFwLimitEntryDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwLimitEntryCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwLimitEntry (Inside resourceFwLimitEntryCreate) ")

		data := dataToFwLimitEntry(d)
		logger.Println("[INFO] received formatted data from method data to FwLimitEntry --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwLimitEntry(client.Token, data, client.Host)

		return resourceFwLimitEntryRead(d, meta)

	}
	return nil
}

func resourceFwLimitEntryRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwLimitEntry (Inside resourceFwLimitEntryRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwLimitEntry(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwLimitEntryUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwLimitEntryRead(d, meta)
}

func resourceFwLimitEntryDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwLimitEntryRead(d, meta)
}
func dataToFwLimitEntry(d *schema.ResourceData) go_thunder.FwLimitEntry {
	var vc go_thunder.FwLimitEntry
	var c go_thunder.FwLimitEntryInstance

	vc.UUID = c
	return vc
}
