package vthunder

//vThunder resource GslbSystem

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceGslbSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbSystemCreate,
		Update: resourceGslbSystemUpdate,
		Read:   resourceGslbSystemRead,
		Delete: resourceGslbSystemDelete,
		Schema: map[string]*schema.Schema{
			"gslb_service_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gslb_site": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ip_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gslb_group": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb_device": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"hostname": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"module": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"wait": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"age_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"geo_location_iana": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb_virtual_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gslb_load_file_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location_load_filename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_name": {
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

func resourceGslbSystemCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbSystem (Inside resourceGslbSystemCreate) ")

		data := dataToGslbSystem(d)
		logger.Println("[INFO] received formatted data from method data to GslbSystem --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostGslbSystem(client.Token, data, client.Host)

		return resourceGslbSystemRead(d, meta)

	}
	return nil
}

func resourceGslbSystemRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbSystem (Inside resourceGslbSystemRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbSystem(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbSystemUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceGslbSystemRead(d, meta)
}

func resourceGslbSystemDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceGslbSystemRead(d, meta)
}
func dataToGslbSystem(d *schema.ResourceData) go_vthunder.GslbSystem {
	var vc go_vthunder.GslbSystem
	var c go_vthunder.GslbSystemInstance
	c.GslbServiceIP = d.Get("gslb_service_ip").(int)
	c.GslbSite = d.Get("gslb_site").(int)
	c.IPTTL = d.Get("ip_ttl").(int)
	c.GslbGroup = d.Get("gslb_group").(int)
	c.SlbDevice = d.Get("slb_device").(int)
	c.Hostname = d.Get("hostname").(int)
	c.Module = d.Get("module").(int)
	c.Wait = d.Get("wait").(int)
	c.AgeInterval = d.Get("age_interval").(int)
	c.GeoLocationIana = d.Get("geo_location_iana").(int)
	c.TTL = d.Get("ttl").(int)
	c.SlbServer = d.Get("slb_server").(int)
	c.SlbVirtualServer = d.Get("slb_virtual_server").(int)

	GslbLoadFileListCount := d.Get("gslb_load_file_list.#").(int)
	c.GeoLocationLoadFilename = make([]go_vthunder.GslbSystemGslbLoadFileList, 0, GslbLoadFileListCount)

	for i := 0; i < GslbLoadFileListCount; i++ {
		var obj1 go_vthunder.GslbSystemGslbLoadFileList
		prefix := fmt.Sprintf("gslb_load_file_list.%d.", i)
		obj1.GeoLocationLoadFilename = d.Get(prefix + "geo_location_load_filename").(string)
		obj1.TemplateName = d.Get(prefix + "template_name").(string)
		c.GeoLocationLoadFilename = append(c.GeoLocationLoadFilename, obj1)
	}

	vc.GslbServiceIP = c
	return vc
}
