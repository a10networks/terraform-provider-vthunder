package vthunder

//vThunder resource GslbGeoLocation

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceGslbGeoLocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbGeoLocationCreate,
		Update: resourceGslbGeoLocationUpdate,
		Read:   resourceGslbGeoLocationRead,
		Delete: resourceGslbGeoLocationDelete,
		Schema: map[string]*schema.Schema{
			"geo_locn_obj_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"geo_locn_multiple_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"first_ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"first_ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geol_ipv4_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ip_addr2": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv6_addr2": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geol_ipv6_mask": {
							Type:        schema.TypeInt,
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGslbGeoLocationCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbGeoLocation (Inside resourceGslbGeoLocationCreate) ")
		name := d.Get("geo_locn_obj_name").(string)
		data := dataToGslbGeoLocation(d)
		logger.Println("[INFO] received formatted data from method data to GslbGeoLocation --")
		d.SetId(name)
		go_vthunder.PostGslbGeoLocation(client.Token, data, client.Host)

		return resourceGslbGeoLocationRead(d, meta)

	}
	return nil
}

func resourceGslbGeoLocationRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbGeoLocation (Inside resourceGslbGeoLocationRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbGeoLocation(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbGeoLocationUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying GslbGeoLocation   (Inside resourceGslbGeoLocationUpdate) ")
		name := d.Get("geo_locn_obj_name").(string)
		data := dataToGslbGeoLocation(d)
		logger.Println("[INFO] received formatted data from method data to GslbGeoLocation ")
		d.SetId(name)
		go_vthunder.PutGslbGeoLocation(client.Token, name, data, client.Host)

		return resourceGslbGeoLocationRead(d, meta)

	}
	return nil
}

func resourceGslbGeoLocationDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceGslbGeoLocationDelete) " + name)
		err := go_vthunder.DeleteGslbGeoLocation(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToGslbGeoLocation(d *schema.ResourceData) go_vthunder.GslbGeoLocation {
	var vc go_vthunder.GslbGeoLocation
	var c go_vthunder.GslbGeoLocationInstance
	c.GeoLocnObjName = d.Get("geo_locn_obj_name").(string)

	GeoLocnMultipleAddressesCount := d.Get("geo_locn_multiple_addresses.#").(int)
	c.FirstIPAddress = make([]go_vthunder.GslbGeoLocationLocnMultipleAddresses, 0, GeoLocnMultipleAddressesCount)

	for i := 0; i < GeoLocnMultipleAddressesCount; i++ {
		var obj1 go_vthunder.GslbGeoLocationLocnMultipleAddresses
		prefix := fmt.Sprintf("geo_locn_multiple_addresses.%d.", i)
		obj1.FirstIPAddress = d.Get(prefix + "first_ip_address").(string)
		obj1.FirstIpv6Address = d.Get(prefix + "first_ipv6_address").(string)
		obj1.GeolIpv4Mask = d.Get(prefix + "geol_ipv4_mask").(string)
		obj1.IPAddr2 = d.Get(prefix + "ip_addr2").(string)
		obj1.Ipv6Addr2 = d.Get(prefix + "ipv6_addr2").(string)
		obj1.GeolIpv6Mask = d.Get(prefix + "geol_ipv6_mask").(int)
		c.FirstIPAddress = append(c.FirstIPAddress, obj1)
	}

	c.UserTag = d.Get("user_tag").(string)

	vc.GeoLocnObjName = c
	return vc
}
