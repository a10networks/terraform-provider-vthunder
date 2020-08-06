package vthunder

//vThunder resource AcosCloudIntegrationEcosystemOracle

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAcosCloudIntegrationEcosystemOracle() *schema.Resource {
	return &schema.Resource{
		Create: resourceAcosCloudIntegrationEcosystemOracleCreate,
		Update: resourceAcosCloudIntegrationEcosystemOracleUpdate,
		Read:   resourceAcosCloudIntegrationEcosystemOracleRead,
		Delete: resourceAcosCloudIntegrationEcosystemOracleDelete,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"private_key_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"tenancy_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"service_label": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"compartment_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fingerprint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceAcosCloudIntegrationEcosystemOracleCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating AcosCloudIntegrationEcosystemOracle (Inside resourceAcosCloudIntegrationEcosystemOracleCreate) ")
		data := dataToAcosCloudIntegrationEcosystemOracle(d)
		logger.Println("[INFO] received formatted data from method data to AcosCloudIntegrationEcosystemOracle --")
		d.SetId("1")
		go_vthunder.PostAcosCloudIntegrationEcosystemOracle(client.Token, data, client.Host)

		return resourceAcosCloudIntegrationEcosystemOracleRead(d, meta)

	}
	return nil
}

func resourceAcosCloudIntegrationEcosystemOracleRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading AcosCloudIntegrationEcosystemOracle (Inside resourceAcosCloudIntegrationEcosystemOracleRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetAcosCloudIntegrationEcosystemOracle(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceAcosCloudIntegrationEcosystemOracleUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceAcosCloudIntegrationEcosystemOracleRead(d, meta)
}

func resourceAcosCloudIntegrationEcosystemOracleDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceAcosCloudIntegrationEcosystemOracleRead(d, meta)
}
func dataToAcosCloudIntegrationEcosystemOracle(d *schema.ResourceData) go_vthunder.EcosystemOracle {
	var vc go_vthunder.EcosystemOracle
	var c go_vthunder.EcosystemOracleInstance

	c.ServiceLabel = d.Get("service_label").(string)

	c.Ipv4Address = d.Get("ipv4_address").(string)

	c.Ipv6Address = d.Get("ipv6_address").(string)

	c.HostName = d.Get("host_name").(string)

	c.Port = d.Get("port").(int)

	c.HealthCheckInterval = d.Get("health_check_interval").(string)

	c.Action = d.Get("action").(string)

	c.CompartmentID = d.Get("compartment_id").(string)

	c.TenancyID = d.Get("tenancy_id").(string)

	c.UserID = d.Get("user_id").(string)

	c.Fingerprint = d.Get("fingerprint").(string)

	c.PrivateKeyPath = d.Get("private_key_path").(string)

	vc.UUID = c
	return vc
}
