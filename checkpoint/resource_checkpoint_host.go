package checkpoint

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/janschumann/checkpoint-go-sdk/service/host"
)

func resourceCheckpointHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceCheckpointHostCreate,
		Update: resourceCheckpointHostUpdate,
		Delete: resourceCheckpointHostDelete,
		Read:   resourceCheckpointHostRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			"ip4_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIP4Address,
			},

			"ip6_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIP6Address,
			},

			"uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCheckpointHostCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*CPClient).host
	input := &host.AddHostInput{}

	if v, ok := d.GetOk("name"); ok {
		input.Name = v.(string)
	}
	if v, ok := d.GetOk("ip4_address"); ok {
		input.Ip4Address = v.(string)
	}
	if v, ok := d.GetOk("ip6_address"); ok {
		input.Ip6Address = v.(string)
	}

	out, err := client.AddHost(input)
	if err != nil {
		return err
	}
	_, err = client.Publish()
	if err != nil {
		return err
	}

	d.SetId(out.Uid)

	return resourceCheckpointHostRead(d, meta)
}

func resourceCheckpointHostUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*CPClient).host

	input := &host.SetHostInput{
		Uid: d.Id(),
	}

	if d.HasChange("ip4_address") {
		input.Ip4Address = d.Get("ip4_address").(string)
	}
	if d.HasChange("ip6_address") {
		input.Ip6Address = d.Get("ip6_address").(string)
	}

	_, err := client.SetHost(input)
	if err != nil {
		return err
	}
	_, err = client.Publish()
	if err != nil {
		return err
	}

	return resourceCheckpointHostRead(d, meta)
}

func resourceCheckpointHostDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*CPClient).host
	_, err := client.DeleteHost(&host.ShowHostInput{
		Uid: d.Id(),
	})
	if err != nil {
		return err
	}
	_, err = client.Publish()
	if err != nil {
		return err
	}

	return nil
}

func resourceCheckpointHostRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*CPClient).host
	out, err := client.ShowHost(&host.ShowHostInput{
		Uid: d.Id(),
	})
	if err != nil {
		return err
	}

	_ = d.Set("uid", out.Uid)
	_ = d.Set("name", out.Name)
	_ = d.Set("ip4_address", out.Ip4Address)
	_ = d.Set("ip6_address", out.Ip6Address)

	return nil
}
