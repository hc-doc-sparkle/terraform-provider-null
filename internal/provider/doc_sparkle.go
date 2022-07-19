package provider

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDocSparkle() *schema.Resource {
	return &schema.Resource{
		Description: `The ` + "`doc_sparkle`" + ` is a test resource for tfplugindocs documentation generation.`,

		Create: resourceCreate,
		Read:   resourceRead,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"triggers": {
				Description: "A map of arbitrary strings that, when changed, will force the doc_sparkle resource to be replaced, re-running any associated provisioners.",
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
			},

			"id": {
				Description: "This is set to a random value at create time.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceDocSparkleCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(fmt.Sprintf("%d", rand.Int()))
	return nil
}

func resourceDocSparkleRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDocSparkleDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
