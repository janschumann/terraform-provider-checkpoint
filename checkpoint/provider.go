package checkpoint

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {

	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["user"],
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["password"],
			},

			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["host"],
			},

			"session_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["session_name"],
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["insecure"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"checkpoint_host": resourceCheckpointHost(),
		},
		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"user": "User",

		"password": "Password",

		"hostname": "Hostname",

		"insecure": "Insecure",

		"session_name": "Session Name",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		User:        d.Get("user").(string),
		Password:    d.Get("password").(string),
		Hostname:    d.Get("host").(string),
		Insecure:    d.Get("insecure").(bool),
		SessionName: d.Get("session_name").(string),
	}

	return config.Client()
}
