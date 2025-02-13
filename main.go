package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	plugin "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/spion/terraform-provider-hetznerdns/hetznerdns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return hetznerdns.Provider()
		},
	})
}
