// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get information about a VM instance resource within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/instances)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
//
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_compute_instance.html.markdown.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("gcp:compute/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	// The name of the instance. One of `name` or `selfLink` must be provided.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If `selfLink` is provided, this value is ignored.  If neither `selfLink`
	// nor `project` are provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The self link of the instance. One of `name` or `selfLink` must be provided.
	SelfLink *string `pulumi:"selfLink"`
	// The zone of the instance. If `selfLink` is provided, this
	// value is ignored.  If neither `selfLink` nor `zone` are provided, the
	// provider zone is used.
	Zone *string `pulumi:"zone"`
}


// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	AllowStoppingForUpdate bool `pulumi:"allowStoppingForUpdate"`
	// List of disks attached to the instance. Structure is documented below.
	AttachedDisks []GetInstanceAttachedDisk `pulumi:"attachedDisks"`
	// The boot disk for the instance. Structure is documented below.
	BootDisks []GetInstanceBootDisk `pulumi:"bootDisks"`
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	CanIpForward bool `pulumi:"canIpForward"`
	// The CPU platform used by this instance.
	CpuPlatform string `pulumi:"cpuPlatform"`
	// Whether deletion protection is enabled on this instance.
	DeletionProtection bool `pulumi:"deletionProtection"`
	// A brief description of the resource.
	Description string `pulumi:"description"`
	DesiredStatus string `pulumi:"desiredStatus"`
	EnableDisplay bool `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance. Structure is documented below.
	GuestAccelerators []GetInstanceGuestAccelerator `pulumi:"guestAccelerators"`
	Hostname string `pulumi:"hostname"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The server-assigned unique identifier of this instance.
	InstanceId string `pulumi:"instanceId"`
	// The unique fingerprint of the labels.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// A set of key/value label pairs assigned to the instance.
	Labels map[string]string `pulumi:"labels"`
	// The machine type to create.
	MachineType string `pulumi:"machineType"`
	// Metadata key/value pairs made available within the instance.
	Metadata map[string]string `pulumi:"metadata"`
	// The unique fingerprint of the metadata.
	MetadataFingerprint string `pulumi:"metadataFingerprint"`
	MetadataStartupScript string `pulumi:"metadataStartupScript"`
	// The minimum CPU platform specified for the VM instance.
	MinCpuPlatform string `pulumi:"minCpuPlatform"`
	Name *string `pulumi:"name"`
	// The networks attached to the instance. Structure is documented below.
	NetworkInterfaces []GetInstanceNetworkInterface `pulumi:"networkInterfaces"`
	Project *string `pulumi:"project"`
	// The scheduling strategy being used by the instance.
	Schedulings []GetInstanceScheduling `pulumi:"schedulings"`
	// The scratch disks attached to the instance. Structure is documented below.
	ScratchDisks []GetInstanceScratchDisk `pulumi:"scratchDisks"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The service account to attach to the instance. Structure is documented below.
	ServiceAccounts []GetInstanceServiceAccount `pulumi:"serviceAccounts"`
	// The shielded vm config being used by the instance. Structure is documented below.
	ShieldedInstanceConfigs []GetInstanceShieldedInstanceConfig `pulumi:"shieldedInstanceConfigs"`
	// The list of tags attached to the instance.
	Tags []string `pulumi:"tags"`
	// The unique fingerprint of the tags.
	TagsFingerprint string `pulumi:"tagsFingerprint"`
	Zone *string `pulumi:"zone"`
}

