package hiera5

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHiera5() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHiera5Read,

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceHiera5Read(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading hiera value")

	keyName := d.Get("key").(string)
	defaultValue, defaultExists := d.GetOk("default")
	hiera := meta.(hiera5)

	v, err := hiera.value(keyName)
	log.Printf("-----------------MARKER------------------------")
	if err != nil && !defaultExists {
		log.Printf("[DEBUG] Error reading hiera value %s", err)
		log.Printf("[DEBUG] Default value is `%s`", defaultValue)
		log.Printf("[DEBUG] Default detectec as present: %t", defaultExists)
		return err
	}

	d.SetId(keyName)
	if err != nil {
		d.Set("value", defaultValue)
	} else {
		d.Set("value", v)
	}

	return nil
}
