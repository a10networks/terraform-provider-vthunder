package vthunder

//vThunder resource GslbProtocolEnable

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceGslbProtocolEnable() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbProtocolEnableCreate,
		Update: resourceGslbProtocolEnableUpdate,
		Read:   resourceGslbProtocolEnableRead,
		Delete: resourceGslbProtocolEnableDelete,
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
	}
}

func resourceGslbProtocolEnableCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbProtocolEnable (Inside resourceGslbProtocolEnableCreate) ")
		name := d.Get("type").(string)
		data := dataToGslbProtocolEnable(d)
		logger.Println("[INFO] received formatted data from method data to GslbProtocolEnable --")
		d.SetId(name)
		go_vthunder.PostGslbProtocolEnable(client.Token, data, client.Host)

		return resourceGslbProtocolEnableRead(d, meta)

	}
	return nil
}

func resourceGslbProtocolEnableRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbProtocolEnable (Inside resourceGslbProtocolEnableRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbProtocolEnable(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbProtocolEnableUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying GslbProtocolEnable   (Inside resourceGslbProtocolEnableUpdate) ")
		name := d.Get("type").(string)
		data := dataToGslbProtocolEnable(d)
		logger.Println("[INFO] received formatted data from method data to GslbProtocolEnable ")
		d.SetId(name)
		go_vthunder.PutGslbProtocolEnable(client.Token, name, data, client.Host)

		return resourceGslbProtocolEnableRead(d, meta)

	}
	return nil
}

func resourceGslbProtocolEnableDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceGslbProtocolEnableDelete) " + name)
		err := go_vthunder.DeleteGslbProtocolEnable(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToGslbProtocolEnable(d *schema.ResourceData) go_vthunder.GslbProtocolEnable {
	var vc go_vthunder.GslbProtocolEnable
	var c go_vthunder.GslbProtocolEnableInstance
	c.Type = d.Get("type").(string)

	vc.Type = c
	return vc
}
