// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents an InterconnectAttachment (VLAN attachment) resource. For more
// information, see Creating VLAN Attachments.
//
// ## Example Usage
//
// ## Import
//
// InterconnectAttachment can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default {{name}}
// ```
type InterconnectAttachment struct {
	pulumi.CustomResourceState

	// Whether the VLAN attachment is enabled or disabled.  When using
	// PARTNER type this will Pre-Activate the interconnect attachment
	AdminEnabled pulumi.BoolPtrOutput `pulumi:"adminEnabled"`
	// Provisioned bandwidth capacity for the interconnect attachment.
	// For attachments of type DEDICATED, the user can set the bandwidth.
	// For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	// Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	// Defaults to BPS_10G
	// Possible values are `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, and `BPS_50G`.
	Bandwidth pulumi.StringOutput `pulumi:"bandwidth"`
	// Up to 16 candidate prefixes that can be used to restrict the allocation
	// of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	// All prefixes must be within link-local address space (169.254.0.0/16)
	// and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	// an unused /29 from the supplied candidate prefix(es). The request will
	// fail if all possible /29s are in use on Google's edge. If not supplied,
	// Google will randomly select an unused /29 from all of link-local space.
	CandidateSubnets pulumi.StringArrayOutput `pulumi:"candidateSubnets"`
	// IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
	CloudRouterIpAddress pulumi.StringOutput `pulumi:"cloudRouterIpAddress"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
	CustomerRouterIpAddress pulumi.StringOutput `pulumi:"customerRouterIpAddress"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Desired availability domain for the attachment. Only available for type
	// PARTNER, at creation time. For improved reliability, customers should
	// configure a pair of attachments with one per availability domain. The
	// selected availability domain will be provided to the Partner via the
	// pairing key so that the provisioned circuit will lie in the specified
	// domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	EdgeAvailabilityDomain pulumi.StringOutput `pulumi:"edgeAvailabilityDomain"`
	// Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity
	// issues.
	GoogleReferenceId pulumi.StringOutput `pulumi:"googleReferenceId"`
	// URL of the underlying Interconnect object that this attachment's
	// traffic will traverse through. Required if type is DEDICATED, must not
	// be set if type is PARTNER.
	Interconnect pulumi.StringPtrOutput `pulumi:"interconnect"`
	// Name of the resource. Provided by the client when the resource is created. The
	// name must be 1-63 characters long, and comply with RFC1035. Specifically, the
	// name must be 1-63 characters long and match the regular expression
	// `a-z?` which means the first character must be a
	// lowercase letter, and all following characters must be a dash, lowercase
	// letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to
	// initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
	PairingKey pulumi.StringOutput `pulumi:"pairingKey"`
	// [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a
	// layer 3 Partner if they configured BGP on behalf of the customer.
	PartnerAsn pulumi.StringOutput `pulumi:"partnerAsn"`
	// Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached
	// to is of type DEDICATED.
	PrivateInterconnectInfos InterconnectAttachmentPrivateInterconnectInfoArrayOutput `pulumi:"privateInterconnectInfos"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where the regional interconnect attachment resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// URL of the cloud router to be used for dynamic routing. This router must be in
	// the same region as this InterconnectAttachment. The InterconnectAttachment will
	// automatically connect the Interconnect to the network & region within which the
	// Cloud Router is configured.
	Router pulumi.StringOutput `pulumi:"router"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// [Output Only] The current state of this attachment's functionality.
	State pulumi.StringOutput `pulumi:"state"`
	// The type of InterconnectAttachment you wish to create. Defaults to
	// DEDICATED.
	// Possible values are `DEDICATED`, `PARTNER`, and `PARTNER_PROVIDER`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
	// using PARTNER type this will be managed upstream.
	VlanTag8021q pulumi.IntOutput `pulumi:"vlanTag8021q"`
}

// NewInterconnectAttachment registers a new resource with the given unique name, arguments, and options.
func NewInterconnectAttachment(ctx *pulumi.Context,
	name string, args *InterconnectAttachmentArgs, opts ...pulumi.ResourceOption) (*InterconnectAttachment, error) {
	if args == nil || args.Router == nil {
		return nil, errors.New("missing required argument 'Router'")
	}
	if args == nil {
		args = &InterconnectAttachmentArgs{}
	}
	var resource InterconnectAttachment
	err := ctx.RegisterResource("gcp:compute/interconnectAttachment:InterconnectAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterconnectAttachment gets an existing InterconnectAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterconnectAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterconnectAttachmentState, opts ...pulumi.ResourceOption) (*InterconnectAttachment, error) {
	var resource InterconnectAttachment
	err := ctx.ReadResource("gcp:compute/interconnectAttachment:InterconnectAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InterconnectAttachment resources.
type interconnectAttachmentState struct {
	// Whether the VLAN attachment is enabled or disabled.  When using
	// PARTNER type this will Pre-Activate the interconnect attachment
	AdminEnabled *bool `pulumi:"adminEnabled"`
	// Provisioned bandwidth capacity for the interconnect attachment.
	// For attachments of type DEDICATED, the user can set the bandwidth.
	// For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	// Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	// Defaults to BPS_10G
	// Possible values are `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, and `BPS_50G`.
	Bandwidth *string `pulumi:"bandwidth"`
	// Up to 16 candidate prefixes that can be used to restrict the allocation
	// of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	// All prefixes must be within link-local address space (169.254.0.0/16)
	// and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	// an unused /29 from the supplied candidate prefix(es). The request will
	// fail if all possible /29s are in use on Google's edge. If not supplied,
	// Google will randomly select an unused /29 from all of link-local space.
	CandidateSubnets []string `pulumi:"candidateSubnets"`
	// IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
	CloudRouterIpAddress *string `pulumi:"cloudRouterIpAddress"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
	CustomerRouterIpAddress *string `pulumi:"customerRouterIpAddress"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Desired availability domain for the attachment. Only available for type
	// PARTNER, at creation time. For improved reliability, customers should
	// configure a pair of attachments with one per availability domain. The
	// selected availability domain will be provided to the Partner via the
	// pairing key so that the provisioned circuit will lie in the specified
	// domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	EdgeAvailabilityDomain *string `pulumi:"edgeAvailabilityDomain"`
	// Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity
	// issues.
	GoogleReferenceId *string `pulumi:"googleReferenceId"`
	// URL of the underlying Interconnect object that this attachment's
	// traffic will traverse through. Required if type is DEDICATED, must not
	// be set if type is PARTNER.
	Interconnect *string `pulumi:"interconnect"`
	// Name of the resource. Provided by the client when the resource is created. The
	// name must be 1-63 characters long, and comply with RFC1035. Specifically, the
	// name must be 1-63 characters long and match the regular expression
	// `a-z?` which means the first character must be a
	// lowercase letter, and all following characters must be a dash, lowercase
	// letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to
	// initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
	PairingKey *string `pulumi:"pairingKey"`
	// [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a
	// layer 3 Partner if they configured BGP on behalf of the customer.
	PartnerAsn *string `pulumi:"partnerAsn"`
	// Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached
	// to is of type DEDICATED.
	PrivateInterconnectInfos []InterconnectAttachmentPrivateInterconnectInfo `pulumi:"privateInterconnectInfos"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the regional interconnect attachment resides.
	Region *string `pulumi:"region"`
	// URL of the cloud router to be used for dynamic routing. This router must be in
	// the same region as this InterconnectAttachment. The InterconnectAttachment will
	// automatically connect the Interconnect to the network & region within which the
	// Cloud Router is configured.
	Router *string `pulumi:"router"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output Only] The current state of this attachment's functionality.
	State *string `pulumi:"state"`
	// The type of InterconnectAttachment you wish to create. Defaults to
	// DEDICATED.
	// Possible values are `DEDICATED`, `PARTNER`, and `PARTNER_PROVIDER`.
	Type *string `pulumi:"type"`
	// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
	// using PARTNER type this will be managed upstream.
	VlanTag8021q *int `pulumi:"vlanTag8021q"`
}

type InterconnectAttachmentState struct {
	// Whether the VLAN attachment is enabled or disabled.  When using
	// PARTNER type this will Pre-Activate the interconnect attachment
	AdminEnabled pulumi.BoolPtrInput
	// Provisioned bandwidth capacity for the interconnect attachment.
	// For attachments of type DEDICATED, the user can set the bandwidth.
	// For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	// Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	// Defaults to BPS_10G
	// Possible values are `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, and `BPS_50G`.
	Bandwidth pulumi.StringPtrInput
	// Up to 16 candidate prefixes that can be used to restrict the allocation
	// of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	// All prefixes must be within link-local address space (169.254.0.0/16)
	// and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	// an unused /29 from the supplied candidate prefix(es). The request will
	// fail if all possible /29s are in use on Google's edge. If not supplied,
	// Google will randomly select an unused /29 from all of link-local space.
	CandidateSubnets pulumi.StringArrayInput
	// IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
	CloudRouterIpAddress pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
	CustomerRouterIpAddress pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Desired availability domain for the attachment. Only available for type
	// PARTNER, at creation time. For improved reliability, customers should
	// configure a pair of attachments with one per availability domain. The
	// selected availability domain will be provided to the Partner via the
	// pairing key so that the provisioned circuit will lie in the specified
	// domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	EdgeAvailabilityDomain pulumi.StringPtrInput
	// Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity
	// issues.
	GoogleReferenceId pulumi.StringPtrInput
	// URL of the underlying Interconnect object that this attachment's
	// traffic will traverse through. Required if type is DEDICATED, must not
	// be set if type is PARTNER.
	Interconnect pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The
	// name must be 1-63 characters long, and comply with RFC1035. Specifically, the
	// name must be 1-63 characters long and match the regular expression
	// `a-z?` which means the first character must be a
	// lowercase letter, and all following characters must be a dash, lowercase
	// letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to
	// initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
	PairingKey pulumi.StringPtrInput
	// [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a
	// layer 3 Partner if they configured BGP on behalf of the customer.
	PartnerAsn pulumi.StringPtrInput
	// Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached
	// to is of type DEDICATED.
	PrivateInterconnectInfos InterconnectAttachmentPrivateInterconnectInfoArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the regional interconnect attachment resides.
	Region pulumi.StringPtrInput
	// URL of the cloud router to be used for dynamic routing. This router must be in
	// the same region as this InterconnectAttachment. The InterconnectAttachment will
	// automatically connect the Interconnect to the network & region within which the
	// Cloud Router is configured.
	Router pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// [Output Only] The current state of this attachment's functionality.
	State pulumi.StringPtrInput
	// The type of InterconnectAttachment you wish to create. Defaults to
	// DEDICATED.
	// Possible values are `DEDICATED`, `PARTNER`, and `PARTNER_PROVIDER`.
	Type pulumi.StringPtrInput
	// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
	// using PARTNER type this will be managed upstream.
	VlanTag8021q pulumi.IntPtrInput
}

func (InterconnectAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*interconnectAttachmentState)(nil)).Elem()
}

type interconnectAttachmentArgs struct {
	// Whether the VLAN attachment is enabled or disabled.  When using
	// PARTNER type this will Pre-Activate the interconnect attachment
	AdminEnabled *bool `pulumi:"adminEnabled"`
	// Provisioned bandwidth capacity for the interconnect attachment.
	// For attachments of type DEDICATED, the user can set the bandwidth.
	// For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	// Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	// Defaults to BPS_10G
	// Possible values are `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, and `BPS_50G`.
	Bandwidth *string `pulumi:"bandwidth"`
	// Up to 16 candidate prefixes that can be used to restrict the allocation
	// of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	// All prefixes must be within link-local address space (169.254.0.0/16)
	// and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	// an unused /29 from the supplied candidate prefix(es). The request will
	// fail if all possible /29s are in use on Google's edge. If not supplied,
	// Google will randomly select an unused /29 from all of link-local space.
	CandidateSubnets []string `pulumi:"candidateSubnets"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Desired availability domain for the attachment. Only available for type
	// PARTNER, at creation time. For improved reliability, customers should
	// configure a pair of attachments with one per availability domain. The
	// selected availability domain will be provided to the Partner via the
	// pairing key so that the provisioned circuit will lie in the specified
	// domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	EdgeAvailabilityDomain *string `pulumi:"edgeAvailabilityDomain"`
	// URL of the underlying Interconnect object that this attachment's
	// traffic will traverse through. Required if type is DEDICATED, must not
	// be set if type is PARTNER.
	Interconnect *string `pulumi:"interconnect"`
	// Name of the resource. Provided by the client when the resource is created. The
	// name must be 1-63 characters long, and comply with RFC1035. Specifically, the
	// name must be 1-63 characters long and match the regular expression
	// `a-z?` which means the first character must be a
	// lowercase letter, and all following characters must be a dash, lowercase
	// letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the regional interconnect attachment resides.
	Region *string `pulumi:"region"`
	// URL of the cloud router to be used for dynamic routing. This router must be in
	// the same region as this InterconnectAttachment. The InterconnectAttachment will
	// automatically connect the Interconnect to the network & region within which the
	// Cloud Router is configured.
	Router string `pulumi:"router"`
	// The type of InterconnectAttachment you wish to create. Defaults to
	// DEDICATED.
	// Possible values are `DEDICATED`, `PARTNER`, and `PARTNER_PROVIDER`.
	Type *string `pulumi:"type"`
	// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
	// using PARTNER type this will be managed upstream.
	VlanTag8021q *int `pulumi:"vlanTag8021q"`
}

// The set of arguments for constructing a InterconnectAttachment resource.
type InterconnectAttachmentArgs struct {
	// Whether the VLAN attachment is enabled or disabled.  When using
	// PARTNER type this will Pre-Activate the interconnect attachment
	AdminEnabled pulumi.BoolPtrInput
	// Provisioned bandwidth capacity for the interconnect attachment.
	// For attachments of type DEDICATED, the user can set the bandwidth.
	// For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	// Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	// Defaults to BPS_10G
	// Possible values are `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, and `BPS_50G`.
	Bandwidth pulumi.StringPtrInput
	// Up to 16 candidate prefixes that can be used to restrict the allocation
	// of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	// All prefixes must be within link-local address space (169.254.0.0/16)
	// and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	// an unused /29 from the supplied candidate prefix(es). The request will
	// fail if all possible /29s are in use on Google's edge. If not supplied,
	// Google will randomly select an unused /29 from all of link-local space.
	CandidateSubnets pulumi.StringArrayInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Desired availability domain for the attachment. Only available for type
	// PARTNER, at creation time. For improved reliability, customers should
	// configure a pair of attachments with one per availability domain. The
	// selected availability domain will be provided to the Partner via the
	// pairing key so that the provisioned circuit will lie in the specified
	// domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	EdgeAvailabilityDomain pulumi.StringPtrInput
	// URL of the underlying Interconnect object that this attachment's
	// traffic will traverse through. Required if type is DEDICATED, must not
	// be set if type is PARTNER.
	Interconnect pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The
	// name must be 1-63 characters long, and comply with RFC1035. Specifically, the
	// name must be 1-63 characters long and match the regular expression
	// `a-z?` which means the first character must be a
	// lowercase letter, and all following characters must be a dash, lowercase
	// letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the regional interconnect attachment resides.
	Region pulumi.StringPtrInput
	// URL of the cloud router to be used for dynamic routing. This router must be in
	// the same region as this InterconnectAttachment. The InterconnectAttachment will
	// automatically connect the Interconnect to the network & region within which the
	// Cloud Router is configured.
	Router pulumi.StringInput
	// The type of InterconnectAttachment you wish to create. Defaults to
	// DEDICATED.
	// Possible values are `DEDICATED`, `PARTNER`, and `PARTNER_PROVIDER`.
	Type pulumi.StringPtrInput
	// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
	// using PARTNER type this will be managed upstream.
	VlanTag8021q pulumi.IntPtrInput
}

func (InterconnectAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interconnectAttachmentArgs)(nil)).Elem()
}

type InterconnectAttachmentInput interface {
	pulumi.Input

	ToInterconnectAttachmentOutput() InterconnectAttachmentOutput
	ToInterconnectAttachmentOutputWithContext(ctx context.Context) InterconnectAttachmentOutput
}

func (InterconnectAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*InterconnectAttachment)(nil)).Elem()
}

func (i InterconnectAttachment) ToInterconnectAttachmentOutput() InterconnectAttachmentOutput {
	return i.ToInterconnectAttachmentOutputWithContext(context.Background())
}

func (i InterconnectAttachment) ToInterconnectAttachmentOutputWithContext(ctx context.Context) InterconnectAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterconnectAttachmentOutput)
}

type InterconnectAttachmentOutput struct {
	*pulumi.OutputState
}

func (InterconnectAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InterconnectAttachmentOutput)(nil)).Elem()
}

func (o InterconnectAttachmentOutput) ToInterconnectAttachmentOutput() InterconnectAttachmentOutput {
	return o
}

func (o InterconnectAttachmentOutput) ToInterconnectAttachmentOutputWithContext(ctx context.Context) InterconnectAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InterconnectAttachmentOutput{})
}
