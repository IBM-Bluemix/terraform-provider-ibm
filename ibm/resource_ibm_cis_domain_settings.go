package ibm

import (
	//"fmt"
	"log"
	//"strings"
	//"time"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	//"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/hashicorp/terraform/helper/validation"
)

func resourceIBMCISSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			"domain_id": {
				Type:        schema.TypeString,
				Description: "Associated CIS domain",
				Required:    true,
			},
			"waf": {
				Type:         schema.TypeString,
				Description:  "WAF setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"off", "on"}),
			},
			"waf_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl": {
				Type:         schema.TypeString,
				Description:  "SSL/TLS setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"off", "flexible", "full", "strict"}),
			},
			"certificate_status": {
				Type:        schema.TypeString,
				Description: "Certificate status",
				Computed:    true,
			},
			"ssl_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_tls_version": {
				Type:         schema.TypeString,
				Description:  "Minimum version of TLS required",
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"1.1", "1.2", "1.3", "1.4"}),
				Default:      "1.1",
			},
			"min_tls_version_modified": {
				Type:        schema.TypeString,
				Description: "Minimum version of TLS required",
				Computed:    true,
			},
			// "tls_1_3_setting": {
			// 	Type:         schema.TypeString,
			// 	Description:  "TLS_1_2 setting",
			// 	Optional:     true,
			// 	Computed:     true,
			// 	ValidateFunc: validateAllowedStringValue([]string{"off", "on"}),
			// },
			// "tls_1_3_setting_modified": {
			// 	Type:        schema.TypeString,
			// 	Description: "Date TLS_1_2 setting modified",
			// 	Computed:    true,
			// },
		},

		Create: resourceCISSettingsUpdate,
		Read:   resourceCISSettingsRead,
		Update: resourceCISSettingsUpdate,
		Delete: resourceCISSettingsDelete,
	}
}

func resourceCISSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	log.Printf("   client %v\n", cisClient)
	if err != nil {
		return err
	}

	cisId := d.Get("cis_id").(string)
	zoneId := d.Get("domain_id").(string)

	type Setting struct {
		Name  string
		Value string
	}
	var settingsArray []Setting

	settingsArray = append(settingsArray, Setting{Name: "waf", Value: d.Get("waf").(string)})
	settingsArray = append(settingsArray, Setting{Name: "ssl", Value: d.Get("ssl").(string)})
	settingsArray = append(settingsArray, Setting{Name: "min_tls_version", Value: d.Get("min_tls_version").(string)})
	//settingsArray = append(settingsArray, Setting{Name: "tls_1_3_only", Value: d.Get("tls_1_3_setting").(string)})

	for _, item := range settingsArray {
		settingsNew := v1.SettingsBody{Value: item.Value}
		_, err = cisClient.Settings().UpdateSettings(cisId, zoneId, item.Name, settingsNew)
		if err != nil {
			log.Printf("Update settings Failed on %s, %s\n", item.Name, err)
			return err
		}
	}
	// Settings are associated with a zone/domain. Use same Id
	d.SetId(zoneId)

	return resourceCISSettingsRead(d, meta)
}

func resourceCISSettingsRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}

	var settingsId string

	settingsId = d.Id()
	cisId := d.Get("cis_id").(string)
	//zoneId := d.Get("domain_id").(string)

	log.Printf("resourceCISSettingsRead - Getting Settings \n")

	settingsResults, err := cisClient.Settings().GetSettings(cisId, settingsId)
	if err != nil {
		log.Printf("resourceCISettingsRead - ListSettingss Failed\n")
		return err
	} else {

		settingsObj := *settingsResults
		d.Set("waf", settingsObj.Waf.Value)
		d.Set("waf_modified", settingsObj.Waf.ModifiedDate)
		d.Set("ssl", settingsObj.Ssl.Value)
		d.Set("certificate_status", settingsObj.Ssl.CertificateStatus)
		d.Set("ssl_modified", settingsObj.Ssl.ModifiedDate)
		d.Set("min_tls_version", settingsObj.Min_tls_version.Value)
		d.Set("min_tls_version_modified", settingsObj.Min_tls_version.ModifiedDate)
		// d.Set("tls_1_3_setting", settingsObj.Tls_1_3.Value)
		// d.Set("tls_1_3_modified", settingsObj.Tls_1_3.ModifiedDate)

	}
	return nil
}

func resourceCISSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	// Nothing to delete on CIS resource
	d.SetId("")
	return nil
}
