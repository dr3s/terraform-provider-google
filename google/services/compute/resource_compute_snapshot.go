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
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceComputeSnapshot() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeSnapshotCreate,
		Read:   resourceComputeSnapshotRead,
		Update: resourceComputeSnapshotUpdate,
		Delete: resourceComputeSnapshotDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeSnapshotImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource; provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"source_disk": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `A reference to the disk used to create this snapshot.`,
			},
			"chain_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Creates the new snapshot in the snapshot chain labeled with the
specified name. The chain name must be 1-63 characters long and
comply with RFC1035. This is an uncommon option only for advanced
service owners who needs to create separate snapshot chains, for
example, for chargeback tracking.  When you describe your snapshot
resource, this field is visible only if it has a non-empty value.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Labels to apply to this Snapshot.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"snapshot_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Encrypts the snapshot using a customer-supplied encryption key.

After you encrypt a snapshot using a customer-supplied key, you must
provide the same key if you use the snapshot later. For example, you
must provide the encryption key when you create a disk from the
encrypted snapshot in a future request.

Customer-supplied encryption keys do not protect access to metadata of
the snapshot.

If you do not provide an encryption key when creating the snapshot,
then the snapshot will be encrypted using an automatically generated
key and you do not need to provide a key to use the snapshot later.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_self_link": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The name of the encryption key that is stored in Google Cloud KMS.`,
						},
						"kms_key_service_account": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The service account used for the encryption request for the given KMS key.
If absent, the Compute Engine Service Agent service account is used.`,
						},
						"raw_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Specifies a 256-bit customer-supplied encryption key, encoded in
RFC 4648 base64 to either encrypt or decrypt this resource.`,
							Sensitive: true,
						},
						"sha256": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
encryption key that protects this resource.`,
						},
					},
				},
			},
			"source_disk_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The customer-supplied encryption key of the source snapshot. Required
if the source snapshot is protected by a customer-supplied encryption
key.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_service_account": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The service account used for the encryption request for the given KMS key.
If absent, the Compute Engine Service Agent service account is used.`,
						},
						"raw_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Specifies a 256-bit customer-supplied encryption key, encoded in
RFC 4648 base64 to either encrypt or decrypt this resource.`,
							Sensitive: true,
						},
					},
				},
			},
			"storage_locations": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Cloud Storage bucket storage location of the snapshot (regional or multi-regional).`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `A reference to the zone where the disk is hosted.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"disk_size_gb": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Size of the snapshot, specified in GB.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fingerprint used for optimistic locking of this resource. Used
internally during updates.`,
			},
			"licenses": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `A list of public visible licenses that apply to this snapshot. This
can be because the original image had licenses attached (such as a
Windows image).  snapshotEncryptionKey nested object Encrypts the
snapshot using a customer-supplied encryption key.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				},
			},
			"snapshot_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The unique identifier for the resource.`,
			},
			"storage_bytes": {
				Type:     schema.TypeInt,
				Computed: true,
				Description: `A size of the storage used by the snapshot. As snapshots share
storage, this number is expected to change with snapshot
creation/deletion.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceComputeSnapshotCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	chainNameProp, err := expandComputeSnapshotChainName(d.Get("chain_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("chain_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(chainNameProp)) && (ok || !reflect.DeepEqual(v, chainNameProp)) {
		obj["chainName"] = chainNameProp
	}
	nameProp, err := expandComputeSnapshotName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeSnapshotDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	storageLocationsProp, err := expandComputeSnapshotStorageLocations(d.Get("storage_locations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("storage_locations"); !tpgresource.IsEmptyValue(reflect.ValueOf(storageLocationsProp)) && (ok || !reflect.DeepEqual(v, storageLocationsProp)) {
		obj["storageLocations"] = storageLocationsProp
	}
	labelFingerprintProp, err := expandComputeSnapshotLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	labelsProp, err := expandComputeSnapshotEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	sourceDiskProp, err := expandComputeSnapshotSourceDisk(d.Get("source_disk"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_disk"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceDiskProp)) && (ok || !reflect.DeepEqual(v, sourceDiskProp)) {
		obj["sourceDisk"] = sourceDiskProp
	}
	zoneProp, err := expandComputeSnapshotZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}
	snapshotEncryptionKeyProp, err := expandComputeSnapshotSnapshotEncryptionKey(d.Get("snapshot_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("snapshot_encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(snapshotEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, snapshotEncryptionKeyProp)) {
		obj["snapshotEncryptionKey"] = snapshotEncryptionKeyProp
	}
	sourceDiskEncryptionKeyProp, err := expandComputeSnapshotSourceDiskEncryptionKey(d.Get("source_disk_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_disk_encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceDiskEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceDiskEncryptionKeyProp)) {
		obj["sourceDiskEncryptionKey"] = sourceDiskEncryptionKeyProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}PRE_CREATE_REPLACE_ME/createSnapshot")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Snapshot: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snapshot: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	url = regexp.MustCompile("PRE_CREATE_REPLACE_ME").ReplaceAllLiteralString(url, sourceDiskProp.(string))
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
		return fmt.Errorf("Error creating Snapshot: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating Snapshot", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Snapshot: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Snapshot %q: %#v", d.Id(), res)

	return resourceComputeSnapshotRead(d, meta)
}

func resourceComputeSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snapshot: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeSnapshot %q", d.Id()))
	}

	res, err = resourceComputeSnapshotDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeSnapshot because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeSnapshotCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("snapshot_id", flattenComputeSnapshotSnapshotId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("disk_size_gb", flattenComputeSnapshotDiskSizeGb(res["diskSizeGb"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("chain_name", flattenComputeSnapshotChainName(res["chainName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("name", flattenComputeSnapshotName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("description", flattenComputeSnapshotDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("storage_bytes", flattenComputeSnapshotStorageBytes(res["storageBytes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("storage_locations", flattenComputeSnapshotStorageLocations(res["storageLocations"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("licenses", flattenComputeSnapshotLicenses(res["licenses"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("labels", flattenComputeSnapshotLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("label_fingerprint", flattenComputeSnapshotLabelFingerprint(res["labelFingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("terraform_labels", flattenComputeSnapshotTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("effective_labels", flattenComputeSnapshotEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("source_disk", flattenComputeSnapshotSourceDisk(res["sourceDisk"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("snapshot_encryption_key", flattenComputeSnapshotSnapshotEncryptionKey(res["snapshotEncryptionKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}

	return nil
}

func resourceComputeSnapshotUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snapshot: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("label_fingerprint") || d.HasChange("effective_labels") {
		obj := make(map[string]interface{})

		labelFingerprintProp, err := expandComputeSnapshotLabelFingerprint(d.Get("label_fingerprint"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
			obj["labelFingerprint"] = labelFingerprintProp
		}
		labelsProp, err := expandComputeSnapshotEffectiveLabels(d.Get("effective_labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/snapshots/{{name}}/setLabels")
		if err != nil {
			return err
		}

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
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating Snapshot %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Snapshot %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating Snapshot", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeSnapshotRead(d, meta)
}

func resourceComputeSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snapshot: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Snapshot %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Snapshot")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting Snapshot", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Snapshot %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeSnapshotImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/global/snapshots/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeSnapshotCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotSnapshotId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeSnapshotDiskSizeGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeSnapshotChainName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotStorageBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeSnapshotStorageLocations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotLicenses(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenComputeSnapshotLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenComputeSnapshotLabelFingerprint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenComputeSnapshotEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotSourceDisk(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeSnapshotSnapshotEncryptionKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["raw_key"] =
		flattenComputeSnapshotSnapshotEncryptionKeyRawKey(original["rawKey"], d, config)
	transformed["sha256"] =
		flattenComputeSnapshotSnapshotEncryptionKeySha256(original["sha256"], d, config)
	transformed["kms_key_self_link"] =
		flattenComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(original["kmsKeyName"], d, config)
	transformed["kms_key_service_account"] =
		flattenComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(original["kmsKeyServiceAccount"], d, config)
	return []interface{}{transformed}
}
func flattenComputeSnapshotSnapshotEncryptionKeyRawKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("snapshot_encryption_key.0.raw_key")
}

func flattenComputeSnapshotSnapshotEncryptionKeySha256(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputeSnapshotChainName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotStorageLocations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotLabelFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeSnapshotSourceDisk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseZonalFieldValue("disks", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_disk: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSnapshotZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSnapshotSnapshotEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeSnapshotSnapshotEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeSnapshotSnapshotEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	transformedKmsKeySelfLink, err := expandComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	transformedKmsKeyServiceAccount, err := expandComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeSnapshotSnapshotEncryptionKeyRawKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSnapshotEncryptionKeySha256(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDiskEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeSnapshotSourceDiskEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedKmsKeyServiceAccount, err := expandComputeSnapshotSourceDiskEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeSnapshotSourceDiskEncryptionKeyRawKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDiskEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceComputeSnapshotDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := res["snapshotEncryptionKey"]; ok {
		original := v.(map[string]interface{})
		transformed := make(map[string]interface{})
		// The raw key won't be returned, so we need to use the original.
		transformed["rawKey"] = d.Get("snapshot_encryption_key.0.raw_key")
		transformed["sha256"] = original["sha256"]

		if kmsKeyName, ok := original["kmsKeyName"]; ok {
			// The response for crypto keys often includes the version of the key which needs to be removed
			// format: projects/<project>/locations/<region>/keyRings/<keyring>/cryptoKeys/<key>/cryptoKeyVersions/1
			transformed["kmsKeyName"] = strings.Split(kmsKeyName.(string), "/cryptoKeyVersions")[0]
		}

		if kmsKeyServiceAccount, ok := original["kmsKeyServiceAccount"]; ok {
			transformed["kmsKeyServiceAccount"] = kmsKeyServiceAccount
		}

		res["snapshotEncryptionKey"] = transformed
	}

	if v, ok := res["sourceDiskEncryptionKey"]; ok {
		original := v.(map[string]interface{})
		transformed := make(map[string]interface{})
		// The raw key won't be returned, so we need to use the original.
		transformed["rawKey"] = d.Get("source_disk_encryption_key.0.raw_key")
		transformed["sha256"] = original["sha256"]

		if kmsKeyName, ok := original["kmsKeyName"]; ok {
			// The response for crypto keys often includes the version of the key which needs to be removed
			// format: projects/<project>/locations/<region>/keyRings/<keyring>/cryptoKeys/<key>/cryptoKeyVersions/1
			transformed["kmsKeyName"] = strings.Split(kmsKeyName.(string), "/cryptoKeyVersions")[0]
		}

		if kmsKeyServiceAccount, ok := original["kmsKeyServiceAccount"]; ok {
			transformed["kmsKeyServiceAccount"] = kmsKeyServiceAccount
		}

		res["sourceDiskEncryptionKey"] = transformed
	}

	return res, nil
}
