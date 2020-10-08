package vthunder

//vThunder resource GslbGroup

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceGslbGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbGroupCreate,
		Update: resourceGslbGroupUpdate,
		Read:   resourceGslbGroupRead,
		Delete: resourceGslbGroupDelete,
		Schema: map[string]*schema.Schema{
			"config_save": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"standalone": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mgmt_interface": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dns_discover": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"config_anywhere": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"data_interface": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auto_map_primary": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"learn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"primary_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"auto_map_learn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"suffix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"config_merge": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auto_map_smart": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGslbGroupCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbGroup (Inside resourceGslbGroupCreate) ")
		name := d.Get("name").(string)
		data := dataToGslbGroup(d)
		logger.Println("[INFO] received formatted data from method data to GslbGroup --")
		d.SetId(name)
		go_vthunder.PostGslbGroup(client.Token, data, client.Host)

		return resourceGslbGroupRead(d, meta)

	}
	return nil
}

func resourceGslbGroupRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbGroup (Inside resourceGslbGroupRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbGroup(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying GslbGroup   (Inside resourceGslbGroupUpdate) ")
		name := d.Get("name").(string)
		data := dataToGslbGroup(d)
		logger.Println("[INFO] received formatted data from method data to GslbGroup ")
		d.SetId(name)
		go_vthunder.PutGslbGroup(client.Token, name, data, client.Host)

		return resourceGslbGroupRead(d, meta)

	}
	return nil
}

func resourceGslbGroupDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceGslbGroupDelete) " + name)
		err := go_vthunder.DeleteGslbGroup(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToGslbGroup(d *schema.ResourceData) go_vthunder.GslbGroup {
	var vc go_vthunder.GslbGroup
	var c go_vthunder.GslbGroupInstance
	c.ConfigSave = d.Get("config_save").(int)
	c.Enable = d.Get("enable").(int)
	c.Standalone = d.Get("standalone").(int)
	c.MgmtInterface = d.Get("mgmt_interface").(int)
	c.UserTag = d.Get("user_tag").(string)
	c.DNSDiscover = d.Get("dns_discover").(int)
	c.Priority = d.Get("priority").(int)
	c.ConfigAnywhere = d.Get("config_anywhere").(int)
	c.DataInterface = d.Get("data_interface").(int)
	c.AutoMapPrimary = d.Get("auto_map_primary").(int)
	c.Learn = d.Get("learn").(int)

	PrimaryListCount := d.Get("primary_list.#").(int)
	c.Primary = make([]go_vthunder.GslbGroupPrimaryList, 0, PrimaryListCount)

	for i := 0; i < PrimaryListCount; i++ {
		var obj1 go_vthunder.GslbGroupPrimaryList
		prefix := fmt.Sprintf("primary_list.%d.", i)
		obj1.Primary = d.Get(prefix + "primary").(string)
		c.Primary = append(c.Primary, obj1)
	}

	c.AutoMapLearn = d.Get("auto_map_learn").(int)
	c.Suffix = d.Get("suffix").(string)
	c.ConfigMerge = d.Get("config_merge").(int)
	c.AutoMapSmart = d.Get("auto_map_smart").(int)
	c.Name = d.Get("name").(string)

	vc.ConfigSave = c
	return vc
}
