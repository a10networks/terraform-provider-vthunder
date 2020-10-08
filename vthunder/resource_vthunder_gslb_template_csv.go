package vthunder

//vThunder resource GslbTemplateCsv

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceGslbTemplateCsv() *schema.Resource {
	return &schema.Resource{
		Create: resourceGslbTemplateCsvCreate,
		Update: resourceGslbTemplateCsvUpdate,
		Read:   resourceGslbTemplateCsvRead,
		Delete: resourceGslbTemplateCsvDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"csv_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"delim_num": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"delim_char": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"multiple_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"csv_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceGslbTemplateCsvCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating GslbTemplateCsv (Inside resourceGslbTemplateCsvCreate) ")
		name := d.Get("csv_name").(string)
		data := dataToGslbTemplateCsv(d)
		logger.Println("[INFO] received formatted data from method data to GslbTemplateCsv --")
		d.SetId(name)
		go_vthunder.PostGslbTemplateCsv(client.Token, data, client.Host)

		return resourceGslbTemplateCsvRead(d, meta)

	}
	return nil
}

func resourceGslbTemplateCsvRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading GslbTemplateCsv (Inside resourceGslbTemplateCsvRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetGslbTemplateCsv(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceGslbTemplateCsvUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying GslbTemplateCsv   (Inside resourceGslbTemplateCsvUpdate) ")
		name := d.Get("csv_name").(string)
		data := dataToGslbTemplateCsv(d)
		logger.Println("[INFO] received formatted data from method data to GslbTemplateCsv ")
		d.SetId(name)
		go_vthunder.PutGslbTemplateCsv(client.Token, name, data, client.Host)

		return resourceGslbTemplateCsvRead(d, meta)

	}
	return nil
}

func resourceGslbTemplateCsvDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceGslbTemplateCsvDelete) " + name)
		err := go_vthunder.DeleteGslbTemplateCsv(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToGslbTemplateCsv(d *schema.ResourceData) go_vthunder.GslbTemplateCsv {
	var vc go_vthunder.GslbTemplateCsv
	var c go_vthunder.GslbTemplateCsvInstance
	c.CsvName = d.Get("csv_name").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.Ipv6Enable = d.Get("ipv6_enable").(int)
	c.DelimNum = d.Get("delim_num").(int)
	c.DelimChar = d.Get("delim_char").(string)

	MultipleFieldsCount := d.Get("multiple_fields.#").(int)
	c.Field = make([]go_vthunder.GslbTemplateCsvMultipleFields, 0, MultipleFieldsCount)

	for i := 0; i < MultipleFieldsCount; i++ {
		var obj1 go_vthunder.GslbTemplateCsvMultipleFields
		prefix := fmt.Sprintf("multiple_fields.%d.", i)
		obj1.Field = d.Get(prefix + "field").(int)
		obj1.CsvType = d.Get(prefix + "csv_type").(string)
		c.Field = append(c.Field, obj1)
	}

	vc.UUID = c
	return vc
}
