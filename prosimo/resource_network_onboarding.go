package prosimo

import (
	"context"
	"log"
	"time"

	"git.prosimo.io/prosimoio/prosimo/terraform-provider-prosimo.git/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetworkOnboarding() *schema.Resource {
	return &schema.Resource{
		Description:   "Use this resource to onboard networks.",
		CreateContext: resourceNetworkOnboardingCreate,
		ReadContext:   resourceNetworkOnboardingRead,
		DeleteContext: resourceNetworkOnboardingDelete,
		UpdateContext: resourceNetworkOnboardingUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "name for the application",
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pam_cname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deployed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_cloud": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloud_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(client.GetCloudTypeOptions(), false),
							Description:  "public or private cloud",
						},
						"connection_option": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(client.GetCloudConnectionOptions(), false),
							Description:  "public or private cloud",
						},
						"cloud_creds_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "cloud application account name.",
						},
						"region_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of cloud region",
						},
						"cloud_networks": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "VPC ID",
									},
									"vnet": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "VNET ID",
									},
									"connector_placement": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(client.GetConnectorPlacementOptions(), false),
										Description:  "Infra VPC, Workload VPC or none.",
									},
									"hub_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "(Required if transit-gateway is selected) tgw-id",
									},
									"connectivity_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "transit-gateway, vpc-peering",
									},
									"subnets": {
										Type:        schema.TypeList,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "subnet cider list",
									},
									"connector_settings": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bandwidth": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "EX: small, medium, large",
												},
												"bandwidth_name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "EX: <1 Gbps, >1 Gbps",
												},
												"instance_type": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "EX: t3.medium, t3.large",
												},
											},
										},
									},
								},
							},
						},
						"connect_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "connector",
						},
					},
				},
			},
			"policies": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Select policy name.e.g: ALLOW-ALL-NETWORKS, DENY-ALL-NETWORKS or Custom Policies",
			},
			"onboard_app": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Set this to true if you would like the network to be onboarded to fabric",
			},
			"decommission_app": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Set this to true if you would like the network  to be offboarded from fabric",
			},
			"wait_for_rollout": {
				Type:        schema.TypeBool,
				Description: "Wait for the rollout of the task to complete. Defaults to true.",
				Default:     true,
				Optional:    true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
	}
}

func resourceNetworkOnboardingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	prosimoClient := meta.(*client.ProsimoClient)
	appOffboardFlag := d.Get("decommission_app").(bool)
	if appOffboardFlag {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid  decommission_app flag.",
			Detail:   "decommission_app can't be set to true while creating network onboarding resource.",
		})
		return diags
	}

	// CloudCredName := d.Get("app_name").(string)
	networkOnboardops := &client.NetworkOnboardoptns{
		Name: d.Get("name").(string),
	}

	onboardresponse, err := prosimoClient.NetworkOnboard(ctx, networkOnboardops)
	if err != nil {
		return diag.FromErr(err)
	}
	networkOnboardops.ID = onboardresponse.NetworkOnboardResponseData.ID

	diags = resourceNetworkOnboardingSettings(ctx, d, meta, networkOnboardops)
	d.SetId(networkOnboardops.ID)
	if d.Get("onboard_app").(bool) {
		res, err2 := prosimoClient.OnboardNetworkDeployment(ctx, networkOnboardops.ID)
		if err2 != nil {
			return diag.FromErr(err)
		}
		if d.Get("wait_for_rollout").(bool) {
			log.Printf("[DEBUG] Waiting for task id %s to complete", res.NetworkDeploymentResops.TaskID)
			err = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate),
				retryUntilTaskCompleteNetwork(ctx, d, meta, res.NetworkDeploymentResops.TaskID, networkOnboardops))
			if err != nil {
				return diag.FromErr(err)
			}
			log.Printf("[INFO] task %s is successful", res.NetworkDeploymentResops.TaskID)
		}
	}

	resourceNetworkOnboardingRead(ctx, d, meta)

	return diags
}

func resourceNetworkOnboardingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	prosimoClient := meta.(*client.ProsimoClient)
	appOnboardFlag := d.Get("onboard_app").(bool)
	appOffboardFlag := d.Get("decommission_app").(bool)
	if appOnboardFlag && appOffboardFlag {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid onboard_app and decommission_app flag combination.",
			Detail:   "Both onboard_app and decommission_app have been set to true.",
		})
		return diags
	}

	updateReq := false
	if d.HasChange("name") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Can't Modify App Name",
			Detail:   "App Name can't be modified",
		})
		return diags
	}

	if d.HasChange("public_cloud") {
		updateReq = true
	}
	if d.HasChange("policies") {
		updateReq = true
	}
	if d.HasChange("onboard_app") && !d.IsNewResource() {
		updateReq = true
	}
	if d.HasChange("decommission_app") && !d.IsNewResource() {
		updateReq = true
	}

	if updateReq {
		networkOnboardops := &client.NetworkOnboardoptns{
			Name: d.Get("name").(string),
			ID:   d.Id(),
		}
		_, networkOnboardops = resourceNetworkOnboardingSettingsUpdate(ctx, d, meta, networkOnboardops)
		if d.Get("decommission_app").(bool) {
			onboardresponse, err := prosimoClient.OffboardNetworkDeployment(ctx, networkOnboardops.ID)
			if err != nil {
				return diag.FromErr(err)
			}
			if d.Get("wait_for_rollout").(bool) {
				log.Printf("[DEBUG] Waiting for task id %s to complete", onboardresponse.NetworkDeploymentResops.TaskID)
				err = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate),
					retryUntilTaskComplete(ctx, d, meta, onboardresponse.NetworkDeploymentResops.TaskID))
				if err != nil {
					return diag.FromErr(err)
				}
				log.Printf("[DEBUG] task %s is successful", onboardresponse.NetworkDeploymentResops.TaskID)
			}
		} else if d.Get("onboard_app").(bool) {
			networkOnboardSettingsDbObj, err := prosimoClient.GetNetworkSettings(ctx, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			if networkOnboardSettingsDbObj.Deployed {
				res, err := prosimoClient.OnboardNetworkDeploymentPost(ctx, networkOnboardops)
				if err != nil {
					return diag.FromErr(err)
				}

				if d.Get("wait_for_rollout").(bool) {
					log.Printf("[DEBUG] Waiting for task id %s to complete", res.NetworkDeploymentResops.TaskID)
					err = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate),
						retryUntilTaskCompleteNetwork(ctx, d, meta, res.NetworkDeploymentResops.TaskID, networkOnboardops))
					if err != nil {
						return diag.FromErr(err)
					}
					log.Printf("[INFO] task %s is successful", res.NetworkDeploymentResops.TaskID)
				}
			} else {
				res, err2 := prosimoClient.OnboardNetworkDeployment(ctx, networkOnboardops.ID)
				if err2 != nil {
					return diag.FromErr(err)
				}
				if d.Get("wait_for_rollout").(bool) {
					log.Printf("[DEBUG] Waiting for task id %s to complete", res.NetworkDeploymentResops.TaskID)
					err = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate),
						retryUntilTaskCompleteNetwork(ctx, d, meta, res.NetworkDeploymentResops.TaskID, networkOnboardops))
					if err != nil {
						return diag.FromErr(err)
					}
					log.Printf("[INFO] task %s is successful", res.NetworkDeploymentResops.TaskID)
				}
			}
		}

	}
	resourceNetworkOnboardingRead(ctx, d, meta)

	return diags
}
func resourceNetworkOnboardingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	prosimoClient := meta.(*client.ProsimoClient)
	// log.Printf("resourceAppOnboardingRead %s", d.Id())
	networkOnboardSettingsDbObj, err := prosimoClient.GetNetworkSettings(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(networkOnboardSettingsDbObj.ID)
	d.Set("name", networkOnboardSettingsDbObj.Name)
	d.Set("pam_cname", networkOnboardSettingsDbObj.PamCname)
	d.Set("deployed", networkOnboardSettingsDbObj.Deployed)
	d.Set("status", networkOnboardSettingsDbObj.Status)
	d.Set("policies", d.Get("policies").([]interface{}))
	d.Set("onboard_app", networkOnboardSettingsDbObj.Deployed)
	d.Set("decommission_app", d.Get("decommission_app").(bool))

	return diags

}

func resourceNetworkOnboardingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	prosimoClient := meta.(*client.ProsimoClient)

	networkOffBoardSettingsID := d.Id()

	appSummary, err := prosimoClient.GetNetworkSettings(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if appSummary.Status != "DEPLOYED" {
		err := prosimoClient.DeleteNetworkDeployment(ctx, networkOffBoardSettingsID)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Can't Delete an Onboarded App",
			Detail:   "App is Onboarded. First decommission the App and then try deleting.",
		})
		return diags
	}
	return diags

}

func resourceNetworkOnboardingSettings(ctx context.Context, d *schema.ResourceData, meta interface{}, networkOnboardops *client.NetworkOnboardoptns) diag.Diagnostics {
	prosimoClient := meta.(*client.ProsimoClient)
	var diags diag.Diagnostics

	// Cloud configuration.
	if v, ok := d.GetOk("public_cloud"); ok {
		publiccloudOptsConfig := v.(*schema.Set).List()[0].(map[string]interface{})
		cloudCreds, err := prosimoClient.GetCloudCredsByName(ctx, publiccloudOptsConfig["cloud_creds_name"].(string))
		if err != nil {
			return diag.FromErr(err)
		}
		cloudNetworkInputList := []client.CloudNetworkops{}
		if v, ok := publiccloudOptsConfig["cloud_networks"]; ok {
			cloudNetworkListConfig := v.(*schema.Set).List()
			for _, cloudNetwork := range cloudNetworkListConfig {
				cloudNetworkConfig := cloudNetwork.(map[string]interface{})
				cloudNetworkInput := &client.CloudNetworkops{
					ConnectivityType: cloudNetworkConfig["connectivity_type"].(string),
					HubID:            cloudNetworkConfig["hub_id"].(string),
					Subnets:          expandStringList(cloudNetworkConfig["subnets"].([]interface{})),
				}
				connectorPlacement := cloudNetworkConfig["connector_placement"].(string)
				if connectorPlacement == client.WorkloadVpcConnectorPlacementOptions {
					cloudNetworkInput.ConnectorPlacement = client.AppConnectorPlacementOptions
				} else if connectorPlacement == client.InfraVPCConnectorPlacementOptions {
					cloudNetworkInput.ConnectorPlacement = client.InfraConnectorPlacementOptions
				} else {
					cloudNetworkInput.ConnectorPlacement = connectorPlacement
				}
				if cloudCreds.CloudType == client.AzureCloudType {
					cloudNetworkInput.CloudNetworkID = cloudNetworkConfig["vnet"].(string)
				} else {
					cloudNetworkInput.CloudNetworkID = cloudNetworkConfig["vpc"].(string)
				}
				if cloudNetworkInput.ConnectorPlacement == client.AppConnectorPlacementOptions || cloudNetworkInput.ConnectorPlacement == client.InfraConnectorPlacementOptions {
					if v, ok := cloudNetworkConfig["connector_settings"]; ok {
						connectorsettingConfig := v.(*schema.Set).List()[0].(map[string]interface{})
						connectorsettingInput := &client.ConnectorSettings{
							Bandwidth:     connectorsettingConfig["bandwidth"].(string),
							BandwidthName: connectorsettingConfig["bandwidth_name"].(string),
							InstanceType:  connectorsettingConfig["instance_type"].(string),
						}
						cloudNetworkInput.Connectorsettings = connectorsettingInput
					}
				}
				cloudNetworkInputList = append(cloudNetworkInputList, *cloudNetworkInput)
			}
		}
		publicCloudoptn := &client.PublicCloud{
			CloudType:        publiccloudOptsConfig["cloud_type"].(string),
			ConnectionOption: publiccloudOptsConfig["connection_option"].(string),
			CloudKeyID:       cloudCreds.ID,
			CloudRegion:      publiccloudOptsConfig["region_name"].(string),
			CloudNetworks:    cloudNetworkInputList,
			ConnectType:      publiccloudOptsConfig["connect_type"].(string),
		}
		networkOnboardops.PublicCloud = publicCloudoptn

	}
	err1 := prosimoClient.NetworkOnboardCloud(ctx, networkOnboardops)
	if err1 != nil {
		return diag.FromErr(err1)
	}
	// Securirty policy configuration.
	networkPolicyList := []client.Policyops{}
	if v, ok := d.GetOk("policies"); ok {
		inputPolicies := expandStringList(v.([]interface{}))
		for _, inputpolicy := range inputPolicies {
			networkPolicy := client.Policyops{}
			policyDbObj, err := prosimoClient.GetPolicyByName(ctx, inputpolicy)
			if err != nil {
				return diag.FromErr(err)
			}
			networkPolicy.ID = policyDbObj.ID
			networkPolicyList = append(networkPolicyList, networkPolicy)
		}
	}
	policyList := &client.Security{
		Policies: networkPolicyList,
	}
	securityInput := &client.NetworkSecurityInput{
		Security: policyList,
	}
	err2 := prosimoClient.NetworkOnboardSecurity(ctx, securityInput, networkOnboardops.ID)
	if err2 != nil {
		return diag.FromErr(err2)
	}
	return diags
}

func resourceNetworkOnboardingSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}, networkOnboardops *client.NetworkOnboardoptns) (diag.Diagnostics, *client.NetworkOnboardoptns) {
	prosimoClient := meta.(*client.ProsimoClient)
	var diags diag.Diagnostics

	// Cloud configuration.
	if v, ok := d.GetOk("public_cloud"); ok {
		publiccloudOptsConfig := v.(*schema.Set).List()[0].(map[string]interface{})
		cloudCreds, err := prosimoClient.GetCloudCredsByName(ctx, publiccloudOptsConfig["cloud_creds_name"].(string))
		if err != nil {
			return diag.FromErr(err), nil
		}
		cloudNetworkInputList := []client.CloudNetworkops{}
		if v, ok := publiccloudOptsConfig["cloud_networks"]; ok {
			cloudNetworkListConfig := v.(*schema.Set).List()
			for _, cloudNetwork := range cloudNetworkListConfig {
				cloudNetworkConfig := cloudNetwork.(map[string]interface{})
				cloudNetworkInput := &client.CloudNetworkops{
					ConnectivityType:   cloudNetworkConfig["connectivity_type"].(string),
					HubID:              cloudNetworkConfig["hub_id"].(string),
					ConnectorPlacement: cloudNetworkConfig["connector_placement"].(string),
					Subnets:            expandStringList(cloudNetworkConfig["subnets"].([]interface{})),
				}
				connectorPlacement := cloudNetworkConfig["connector_placement"].(string)
				if connectorPlacement == client.WorkloadVpcConnectorPlacementOptions {
					cloudNetworkInput.ConnectorPlacement = client.AppConnectorPlacementOptions
				} else if connectorPlacement == client.InfraVPCConnectorPlacementOptions {
					cloudNetworkInput.ConnectorPlacement = client.InfraConnectorPlacementOptions
				} else {
					cloudNetworkInput.ConnectorPlacement = connectorPlacement
				}
				if cloudCreds.CloudType == client.AzureCloudType {
					cloudNetworkInput.CloudNetworkID = cloudNetworkConfig["vnet"].(string)
				} else {
					cloudNetworkInput.CloudNetworkID = cloudNetworkConfig["vpc"].(string)
				}
				if cloudNetworkInput.ConnectorPlacement == client.AppConnectorPlacementOptions || cloudNetworkInput.ConnectorPlacement == client.InfraConnectorPlacementOptions {
					if v, ok := cloudNetworkConfig["connector_settings"]; ok {
						connectorsettingConfig := v.(*schema.Set).List()[0].(map[string]interface{})
						connectorsettingInput := &client.ConnectorSettings{
							Bandwidth:     connectorsettingConfig["bandwidth"].(string),
							BandwidthName: connectorsettingConfig["bandwidth_name"].(string),
							InstanceType:  connectorsettingConfig["instance_type"].(string),
						}
						cloudNetworkInput.Connectorsettings = connectorsettingInput
					}
				}
				cloudNetworkInputList = append(cloudNetworkInputList, *cloudNetworkInput)
			}
		}

		publicCloudoptn := &client.PublicCloud{
			CloudType:        publiccloudOptsConfig["cloud_type"].(string),
			ConnectionOption: publiccloudOptsConfig["connection_option"].(string),
			CloudKeyID:       cloudCreds.ID,
			CloudRegion:      publiccloudOptsConfig["region_name"].(string),
			CloudNetworks:    cloudNetworkInputList,
			ConnectType:      publiccloudOptsConfig["connect_type"].(string),
		}
		networkOnboardops.PublicCloud = publicCloudoptn

	}

	// Securirty policy configuration.
	networkPolicyList := []client.Policyops{}
	if v, ok := d.GetOk("policies"); ok {
		inputPolicies := expandStringList(v.([]interface{}))
		for _, inputpolicy := range inputPolicies {
			networkPolicy := client.Policyops{}
			policyDbObj, err := prosimoClient.GetPolicyByName(ctx, inputpolicy)
			if err != nil {
				return diag.FromErr(err), nil
			}
			networkPolicy.ID = policyDbObj.ID
			networkPolicyList = append(networkPolicyList, networkPolicy)
		}
	}
	policyList := &client.Security{
		Policies: networkPolicyList,
	}

	networkOnboardops.Security = policyList

	return diags, networkOnboardops
}
