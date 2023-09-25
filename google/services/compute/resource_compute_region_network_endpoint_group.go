// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceComputeRegionNetworkEndpointGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionNetworkEndpointGroupCreate,
		Read:   resourceComputeRegionNetworkEndpointGroupRead,
		Delete: resourceComputeRegionNetworkEndpointGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionNetworkEndpointGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateGCEName,
				Description: `Name of the resource; provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"region": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `A reference to the region where the Serverless NEGs Reside.`,
			},
			"app_engine": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Only valid when networkEndpointType is "SERVERLESS".
Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Optional serving service.
The service name must be 1-63 characters long, and comply with RFC1035.
Example value: "default", "my-service".`,
						},
						"url_mask": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `A template to parse service and version fields from a request URL.
URL mask allows for routing to multiple App Engine services without
having to create multiple Network Endpoint Groups and backend services.

For example, the request URLs "foo1-dot-appname.appspot.com/v1" and
"foo1-dot-appname.appspot.com/v2" can be backed by the same Serverless NEG with
URL mask "-dot-appname.appspot.com/". The URL mask will parse
them to { service = "foo1", version = "v1" } and { service = "foo1", version = "v2" } respectively.`,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Optional serving version.
The version must be 1-63 characters long, and comply with RFC1035.
Example value: "v1", "v2".`,
						},
					},
				},
				ConflictsWith: []string{"cloud_run", "cloud_function"},
			},
			"cloud_function": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Only valid when networkEndpointType is "SERVERLESS".
Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"function": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `A user-defined name of the Cloud Function.
The function name is case-sensitive and must be 1-63 characters long.
Example value: "func1".`,
							AtLeastOneOf: []string{"cloud_function.0.function", "cloud_function.0.url_mask"},
						},
						"url_mask": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `A template to parse function field from a request URL. URL mask allows
for routing to multiple Cloud Functions without having to create
multiple Network Endpoint Groups and backend services.

For example, request URLs "mydomain.com/function1" and "mydomain.com/function2"
can be backed by the same Serverless NEG with URL mask "/". The URL mask
will parse them to { function = "function1" } and { function = "function2" } respectively.`,
							AtLeastOneOf: []string{"cloud_function.0.function", "cloud_function.0.url_mask"},
						},
					},
				},
				ConflictsWith: []string{"cloud_run", "app_engine"},
			},
			"cloud_run": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Only valid when networkEndpointType is "SERVERLESS".
Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Cloud Run service is the main resource of Cloud Run.
The service must be 1-63 characters long, and comply with RFC1035.
Example value: "run-service".`,
							AtLeastOneOf: []string{"cloud_run.0.service", "cloud_run.0.url_mask"},
						},
						"tag": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Cloud Run tag represents the "named-revision" to provide
additional fine-grained traffic routing information.
The tag must be 1-63 characters long, and comply with RFC1035.
Example value: "revision-0010".`,
						},
						"url_mask": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `A template to parse service and tag fields from a request URL.
URL mask allows for routing to multiple Run services without having
to create multiple network endpoint groups and backend services.

For example, request URLs "foo1.domain.com/bar1" and "foo1.domain.com/bar2"
an be backed by the same Serverless Network Endpoint Group (NEG) with
URL mask ".domain.com/". The URL mask will parse them to { service="bar1", tag="foo1" }
and { service="bar2", tag="foo2" } respectively.`,
							AtLeastOneOf: []string{"cloud_run.0.service", "cloud_run.0.url_mask"},
						},
					},
				},
				ConflictsWith: []string{"cloud_function", "app_engine"},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `This field is only used for PSC.
The URL of the network to which all network endpoints in the NEG belong. Uses
"default" project network if unspecified.`,
			},
			"network_endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"SERVERLESS", "PRIVATE_SERVICE_CONNECT", ""}),
				Description:  `Type of network endpoints in this network endpoint group. Defaults to SERVERLESS Default value: "SERVERLESS" Possible values: ["SERVERLESS", "PRIVATE_SERVICE_CONNECT"]`,
				Default:      "SERVERLESS",
			},
			"psc_target_service": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The target service url used to set up private service connection to
a Google API or a PSC Producer Service Attachment.`,
			},
			"subnetwork": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `This field is only used for PSC.
Optional URL of the subnetwork to which all network endpoints in the NEG belong.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeRegionNetworkEndpointGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionNetworkEndpointGroupName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionNetworkEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkEndpointTypeProp, err := expandComputeRegionNetworkEndpointGroupNetworkEndpointType(d.Get("network_endpoint_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_endpoint_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkEndpointTypeProp)) && (ok || !reflect.DeepEqual(v, networkEndpointTypeProp)) {
		obj["networkEndpointType"] = networkEndpointTypeProp
	}
	pscTargetServiceProp, err := expandComputeRegionNetworkEndpointGroupPscTargetService(d.Get("psc_target_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("psc_target_service"); !tpgresource.IsEmptyValue(reflect.ValueOf(pscTargetServiceProp)) && (ok || !reflect.DeepEqual(v, pscTargetServiceProp)) {
		obj["pscTargetService"] = pscTargetServiceProp
	}
	networkProp, err := expandComputeRegionNetworkEndpointGroupNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	subnetworkProp, err := expandComputeRegionNetworkEndpointGroupSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subnetwork"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	cloudRunProp, err := expandComputeRegionNetworkEndpointGroupCloudRun(d.Get("cloud_run"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cloud_run"); !tpgresource.IsEmptyValue(reflect.ValueOf(cloudRunProp)) && (ok || !reflect.DeepEqual(v, cloudRunProp)) {
		obj["cloudRun"] = cloudRunProp
	}
	appEngineProp, err := expandComputeRegionNetworkEndpointGroupAppEngine(d.Get("app_engine"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine"); !tpgresource.IsEmptyValue(reflect.ValueOf(appEngineProp)) && (ok || !reflect.DeepEqual(v, appEngineProp)) {
		obj["appEngine"] = appEngineProp
	}
	cloudFunctionProp, err := expandComputeRegionNetworkEndpointGroupCloudFunction(d.Get("cloud_function"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cloud_function"); !tpgresource.IsEmptyValue(reflect.ValueOf(cloudFunctionProp)) && (ok || !reflect.DeepEqual(v, cloudFunctionProp)) {
		obj["cloudFunction"] = cloudFunctionProp
	}
	regionProp, err := expandComputeRegionNetworkEndpointGroupRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionNetworkEndpointGroup: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionNetworkEndpointGroup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating RegionNetworkEndpointGroup: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating RegionNetworkEndpointGroup", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionNetworkEndpointGroup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionNetworkEndpointGroup %q: %#v", d.Id(), res)

	return resourceComputeRegionNetworkEndpointGroupRead(d, meta)
}

func resourceComputeRegionNetworkEndpointGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionNetworkEndpointGroup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeRegionNetworkEndpointGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}

	if err := d.Set("name", flattenComputeRegionNetworkEndpointGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionNetworkEndpointGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("network_endpoint_type", flattenComputeRegionNetworkEndpointGroupNetworkEndpointType(res["networkEndpointType"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("psc_target_service", flattenComputeRegionNetworkEndpointGroupPscTargetService(res["pscTargetService"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("network", flattenComputeRegionNetworkEndpointGroupNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("subnetwork", flattenComputeRegionNetworkEndpointGroupSubnetwork(res["subnetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("cloud_run", flattenComputeRegionNetworkEndpointGroupCloudRun(res["cloudRun"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("app_engine", flattenComputeRegionNetworkEndpointGroupAppEngine(res["appEngine"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("cloud_function", flattenComputeRegionNetworkEndpointGroupCloudFunction(res["cloudFunction"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionNetworkEndpointGroupRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}

	return nil
}

func resourceComputeRegionNetworkEndpointGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionNetworkEndpointGroup: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionNetworkEndpointGroup %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "RegionNetworkEndpointGroup")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting RegionNetworkEndpointGroup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionNetworkEndpointGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionNetworkEndpointGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/networkEndpointGroups/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionNetworkEndpointGroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupNetworkEndpointType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupPscTargetService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionNetworkEndpointGroupSubnetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionNetworkEndpointGroupCloudRun(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["service"] =
		flattenComputeRegionNetworkEndpointGroupCloudRunService(original["service"], d, config)
	transformed["tag"] =
		flattenComputeRegionNetworkEndpointGroupCloudRunTag(original["tag"], d, config)
	transformed["url_mask"] =
		flattenComputeRegionNetworkEndpointGroupCloudRunUrlMask(original["urlMask"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionNetworkEndpointGroupCloudRunService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupCloudRunTag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupCloudRunUrlMask(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupAppEngine(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["service"] =
		flattenComputeRegionNetworkEndpointGroupAppEngineService(original["service"], d, config)
	transformed["version"] =
		flattenComputeRegionNetworkEndpointGroupAppEngineVersion(original["version"], d, config)
	transformed["url_mask"] =
		flattenComputeRegionNetworkEndpointGroupAppEngineUrlMask(original["urlMask"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionNetworkEndpointGroupAppEngineService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupAppEngineVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupAppEngineUrlMask(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupCloudFunction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["function"] =
		flattenComputeRegionNetworkEndpointGroupCloudFunctionFunction(original["function"], d, config)
	transformed["url_mask"] =
		flattenComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(original["urlMask"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionNetworkEndpointGroupCloudFunctionFunction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func expandComputeRegionNetworkEndpointGroupName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupNetworkEndpointType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupPscTargetService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionNetworkEndpointGroupSubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for subnetwork: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionNetworkEndpointGroupCloudRun(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandComputeRegionNetworkEndpointGroupCloudRunService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedTag, err := expandComputeRegionNetworkEndpointGroupCloudRunTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupCloudRunUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunUrlMask(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngine(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandComputeRegionNetworkEndpointGroupAppEngineService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedVersion, err := expandComputeRegionNetworkEndpointGroupAppEngineVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupAppEngineUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineUrlMask(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFunction, err := expandComputeRegionNetworkEndpointGroupCloudFunctionFunction(original["function"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFunction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["function"] = transformedFunction
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunctionFunction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
