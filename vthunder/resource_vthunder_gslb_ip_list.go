package vthunder

//vThunder resource GslbIpList

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceGslbIpList() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbIpListCreate,
		Update: resourceGslbIpListUpdate,
		Read:   resourceGslbIpListRead,
		Delete: resourceGslbIpListDelete,
		Schema: map[string]*schema.Schema{
			"gslb_ip_list_obj_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"gslb_ip_list_filename": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"gslb_ip_list_addr_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ip_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"id": {
							Type:        schema.TypeInt,
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
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGslbIpListCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbIpList (Inside resourceGslbIpListCreate) ")
		name := d.Get("gslb_ip_list_obj_name").(string)
		data := dataToGslbIpList(d)
		logger.Println("[INFO] received formatted data from method data to GslbIpList --")
		d.SetId(name)
		go_vthunder.PostGslbIpList(client.Token, data, client.Host)

		return resourceGslbIpListRead(d, meta)

	}
	return nil
}

func resourceGslbIpListRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbIpList (Inside resourceGslbIpListRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbIpList(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbIpListUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying GslbIpList   (Inside resourceGslbIpListUpdate) ")
		name := d.Get("gslb_ip_list_obj_name").(string)
		data := dataToGslbIpList(d)
		logger.Println("[INFO] received formatted data from method data to GslbIpList ")
		d.SetId(name)
		go_vthunder.PutGslbIpList(client.Token, name, data, client.Host)

		return resourceGslbIpListRead(d, meta)

	}
	return nil
}

func resourceGslbIpListDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceGslbIpListDelete) " + name)
		err := go_vthunder.DeleteGslbIpList(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToGslbIpList(d *schema.ResourceData) go_vthunder.GslbIpList {
	var vc go_vthunder.GslbIpList
	var c go_vthunder.GslbIPListInstance
	c.GslbIPListObjName = d.Get("gslb_ip_list_obj_name").(string)
	c.GslbIPListFilename = d.Get("gslb_ip_list_filename").(string)

	GslbIpListAddrListCount := d.Get("gslb_ip_list_addr_list.#").(int)
	c.IP = make([]go_vthunder.GslbListGslbIPListAddrList, 0, GslbIpListAddrListCount)

	for i := 0; i < GslbIpListAddrListCount; i++ {
		var obj1 go_vthunder.GslbListGslbIPListAddrList
		prefix := fmt.Sprintf("gslb_ip_list_addr_list.%d.", i)
		obj1.IP = d.Get(prefix + "ip").(string)
		obj1.IPMask = d.Get(prefix + "ip_mask").(string)
		obj1.ID = d.Get(prefix + "id").(int)
		c.IP = append(c.IP, obj1)
	}

	c.UserTag = d.Get("user_tag").(string)

	vc.GslbIPListObjName = c
	return vc
}
