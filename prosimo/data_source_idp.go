package prosimo

import (
	"context"
	"fmt"
	"log"
	"time"

	"git.prosimo.io/prosimoio/prosimo/terraform-provider-prosimo.git/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func dataSourceIDP() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information on available identity providers.",
		ReadContext: dataSourceIDPRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Custom filters to scope specific results. Usage: filter = app_access_type==agent",
			},
			"idp_list": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idpname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idpid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"select_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"api_cred_provided": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIDPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	prosimoClient := meta.(*client.ProsimoClient)

	var diags diag.Diagnostics
	var returnedIDP []*client.IDP //append in this

	getIDP, err := prosimoClient.GetIDP(ctx)
	if err != nil {
		return diag.FromErr(err)
	}
	filter := d.Get("filter").(string)

	if filter != "" {
		for _, filteredList := range getIDP.IDPs {
			filteredMap := map[string]interface{}{}
			err := mapstructure.Decode(filteredList, &filteredMap)
			if err != nil {
				panic(err)
			}
			log.Println("FilteredMap", filteredMap)
			diags, flag := checkMainOperand(filter, filteredMap)
			if diags != nil {
				return diags
			}
			if flag {
				returnedIDP = append(returnedIDP, filteredList)

			}
		}
		if len(returnedIDP) == 0 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "No match for input attribute",
				Detail:   fmt.Sprintln("No match for input attribute"),
			})

			return diags
		}
	} else {
		for _, idpDetail := range getIDP.IDPs {
			returnedIDP = append(returnedIDP, idpDetail)
		}
	}
	d.SetId(time.Now().Format(time.RFC850))
	idpItems := flattenIDPItemsData(returnedIDP)
	d.Set("idp_list", idpItems)

	return diags
}

func flattenIDPItemsData(IdpItems []*client.IDP) []interface{} {
	if IdpItems != nil {
		ois := make([]interface{}, len(IdpItems), len(IdpItems))

		for i, IdpItem := range IdpItems {
			oi := make(map[string]interface{})

			oi["idpname"] = IdpItem.IDPName
			oi["idpid"] = IdpItem.ID
			oi["api_cred_provided"] = IdpItem.Api_Creds_Provided
			oi["auth_type"] = IdpItem.Auth_Type
			oi["select_type"] = IdpItem.Select_Type

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}
