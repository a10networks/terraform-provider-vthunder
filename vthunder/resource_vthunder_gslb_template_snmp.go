package vthunder

//vThunder resource GslbTemplateSnmp

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceGslbTemplateSnmp() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbTemplateSnmpCreate,
		Update: resourceGslbTemplateSnmpUpdate,
		Read:   resourceGslbTemplateSnmpRead,
		Delete: resourceGslbTemplateSnmpDelete,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"oid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"priv_proto": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"context_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"context_engine_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"security_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"community": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_proto": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"interface": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"priv_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"security_engine_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"snmp_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGslbTemplateSnmpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbTemplateSnmp (Inside resourceGslbTemplateSnmpCreate) ")
		name := d.Get("snmp_name").(string)
		data := dataToGslbTemplateSnmp(d)
		logger.Println("[INFO] received formatted data from method data to GslbTemplateSnmp --")
		d.SetId(name)
		go_vthunder.PostGslbTemplateSnmp(client.Token, data, client.Host)

		return resourceGslbTemplateSnmpRead(d, meta)

	}
	return nil
}

func resourceGslbTemplateSnmpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbTemplateSnmp (Inside resourceGslbTemplateSnmpRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbTemplateSnmp(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbTemplateSnmpUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying GslbTemplateSnmp   (Inside resourceGslbTemplateSnmpUpdate) ")
		name := d.Get("snmp_name").(string)
		data := dataToGslbTemplateSnmp(d)
		logger.Println("[INFO] received formatted data from method data to GslbTemplateSnmp ")
		d.SetId(name)
		go_vthunder.PutGslbTemplateSnmp(client.Token, name, data, client.Host)

		return resourceGslbTemplateSnmpRead(d, meta)

	}
	return nil
}

func resourceGslbTemplateSnmpDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceGslbTemplateSnmpDelete) " + name)
		err := go_vthunder.DeleteGslbTemplateSnmp(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToGslbTemplateSnmp(d *schema.ResourceData) go_vthunder.GslbTemplateSnmp {
	var vc go_vthunder.GslbTemplateSnmp
	var c go_vthunder.GslbTemplateSnmpInstance
	c.Username = d.Get("username").(string)
	c.Oid = d.Get("oid").(string)
	c.PrivProto = d.Get("priv_proto").(string)
	c.ContextName = d.Get("context_name").(string)
	c.AuthKey = d.Get("auth_key").(string)
	c.Interval = d.Get("interval").(int)
	c.ContextEngineID = d.Get("context_engine_id").(string)
	c.SecurityLevel = d.Get("security_level").(string)
	c.Community = d.Get("community").(string)
	c.AuthProto = d.Get("auth_proto").(string)
	c.Host = d.Get("host").(string)
	c.Version = d.Get("version").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.Interface = d.Get("interface").(int)
	c.PrivKey = d.Get("priv_key").(string)
	c.SecurityEngineID = d.Get("security_engine_id").(string)
	c.Port = d.Get("port").(int)
	c.SnmpName = d.Get("snmp_name").(string)

	vc.Username = c
	return vc
}
