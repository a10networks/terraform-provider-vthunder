package vthunder

//vThunder resource Vrrp peer group

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVrrpPeerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVrrpPeerGroupCreate,
		Update: resourceVrrpPeerGroupUpdate,
		Read:   resourceVrrpPeerGroupRead,
		Delete: resourceVrrpPeerGroupDelete,

		Schema: map[string]*schema.Schema{
			"peer": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_peer_address_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_peer_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ip_peer_address_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_peer_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
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

func resourceVrrpPeerGroupCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating vrrp PeerGroup (Inside resourceVrrpPeerGroupCreate)")

	if client.Host != "" {
		vc := dataToVrrpPeerGroup(d)
		d.SetId("1")
		go_vthunder.PostVrrpPeerGroup(client.Token, vc, client.Host)

		return resourceVrrpPeerGroupRead(d, meta)
	}
	return nil
}

func resourceVrrpPeerGroupRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading vrrp PeerGroup (Inside resourceVrrpPeerGroupRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetVrrpPeerGroup(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No vrrp PeerGroup found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceVrrpPeerGroupUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceVrrpPeerGroupRead(d, meta)
}

func resourceVrrpPeerGroupDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceVrrpPeerGroupRead(d, meta)
}

//Utility method to instantiate Vrrp PeerGroup Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToVrrpPeerGroup(d *schema.ResourceData) go_vthunder.PeerGroup {
	var vc go_vthunder.PeerGroup

	var c go_vthunder.PeerGroupInstance

	var peer go_vthunder.Peer

	ip_count := d.Get("peer.0.ip_peer_address_cfg.#").(int)
	peer.IPPeerAddress = make([]go_vthunder.IPPeerAddressCfg, 0, ip_count)

	for i := 0; i < ip_count; i++ {
		var ip_cfg go_vthunder.IPPeerAddressCfg
		prefix := fmt.Sprintf("peer.0.ip_peer_address_cfg.%d", i)
		ip_cfg.IPPeerAddress = d.Get(prefix + ".ip_peer_address").(string)
		peer.IPPeerAddress = append(peer.IPPeerAddress, ip_cfg)
	}
	c.IPPeerAddress = peer
	vc.UUID = c

	return vc
}
