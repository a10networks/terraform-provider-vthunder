package vthunder

//vThunder resource GslbServiceGroup

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceGslbServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbServiceGroupCreate,
		Update: resourceGslbServiceGroupUpdate,
		Read:   resourceGslbServiceGroupRead,
		Delete: resourceGslbServiceGroupDelete,
		Schema: map[string]*schema.Schema{
			"service_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dependency_site": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_site_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable_site": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"persistent_mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"member": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"persistent_ipv6_mask": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"persistent_aging_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"persistent_site": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGslbServiceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbServiceGroup (Inside resourceGslbServiceGroupCreate) ")
		name := d.Get("service_group_name").(string)
		data := dataToGslbServiceGroup(d)
		logger.Println("[INFO] received formatted data from method data to GslbServiceGroup --")
		d.SetId(name)
		go_vthunder.PostGslbServiceGroup(client.Token, data, client.Host)

		return resourceGslbServiceGroupRead(d, meta)

	}
	return nil
}

func resourceGslbServiceGroupRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbServiceGroup (Inside resourceGslbServiceGroupRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbServiceGroup(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbServiceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying GslbServiceGroup   (Inside resourceGslbServiceGroupUpdate) ")
		name := d.Get("service_group_name").(string)
		data := dataToGslbServiceGroup(d)
		logger.Println("[INFO] received formatted data from method data to GslbServiceGroup ")
		d.SetId(name)
		go_vthunder.PutGslbServiceGroup(client.Token, name, data, client.Host)

		return resourceGslbServiceGroupRead(d, meta)

	}
	return nil
}

func resourceGslbServiceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceGslbServiceGroupDelete) " + name)
		err := go_vthunder.DeleteGslbServiceGroup(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToGslbServiceGroup(d *schema.ResourceData) go_vthunder.GslbServiceGroup {
	var vc go_vthunder.GslbServiceGroup
	var c go_vthunder.GslbServiceGroupInstance
	c.ServiceGroupName = d.Get("service_group_name").(string)
	c.DependencySite = d.Get("dependency_site").(int)

	DisableSiteListCount := d.Get("disable_site_list.#").(int)
	c.DisableSite = make([]go_vthunder.GslbServiceGroupDisableSiteList, 0, DisableSiteListCount)

	for i := 0; i < DisableSiteListCount; i++ {
		var obj1 go_vthunder.GslbServiceGroupDisableSiteList
		prefix1 := fmt.Sprintf("disable_site_list.%d.", i)
		obj1.DisableSite = d.Get(prefix1 + "disable_site").(string)
		c.DisableSite = append(c.DisableSite, obj1)
	}

	c.UserTag = d.Get("user_tag").(string)
	c.PersistentMask = d.Get("persistent_mask").(string)

	MemberCount := d.Get("member.#").(int)
	c.MemberName = make([]go_vthunder.GslbServiceGroupMember, 0, MemberCount)

	for i := 0; i < MemberCount; i++ {
		var obj2 go_vthunder.GslbServiceGroupMember
		prefix2 := fmt.Sprintf("member.%d.", i)
		obj2.MemberName = d.Get(prefix2 + "member_name").(string)
		c.MemberName = append(c.MemberName, obj2)
	}

	c.Disable = d.Get("disable").(int)
	c.PersistentIpv6Mask = d.Get("persistent_ipv6_mask").(int)
	c.PersistentAgingTime = d.Get("persistent_aging_time").(int)
	c.PersistentSite = d.Get("persistent_site").(int)

	vc.ServiceGroupName = c
	return vc
}
