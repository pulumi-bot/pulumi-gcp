// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Deprecated: gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate
type MangedSslCertificate struct {
	pulumi.CustomResourceState

	// The unique identifier for the resource.
	CertificateId pulumi.IntOutput `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Expire time of the certificate.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
	// of 'MANAGED' in 'type').
	Managed MangedSslCertificateManagedPtrOutput `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
	// namespace as the managed SSL certificates.
	Name     pulumi.StringOutput `pulumi:"name"`
	Project  pulumi.StringOutput `pulumi:"project"`
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayOutput `pulumi:"subjectAlternativeNames"`
	// Enum field whose value is always 'MANAGED' - used to signal to the API which type this is. Default value: "MANAGED"
	// Possible values: ["MANAGED"]
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewMangedSslCertificate registers a new resource with the given unique name, arguments, and options.
func NewMangedSslCertificate(ctx *pulumi.Context,
	name string, args *MangedSslCertificateArgs, opts ...pulumi.ResourceOption) (*MangedSslCertificate, error) {
	if args == nil {
		args = &MangedSslCertificateArgs{}
	}
	var resource MangedSslCertificate
	err := ctx.RegisterResource("gcp:compute/mangedSslCertificate:MangedSslCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMangedSslCertificate gets an existing MangedSslCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMangedSslCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MangedSslCertificateState, opts ...pulumi.ResourceOption) (*MangedSslCertificate, error) {
	var resource MangedSslCertificate
	err := ctx.ReadResource("gcp:compute/mangedSslCertificate:MangedSslCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MangedSslCertificate resources.
type mangedSslCertificateState struct {
	// The unique identifier for the resource.
	CertificateId *int `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Expire time of the certificate.
	ExpireTime *string `pulumi:"expireTime"`
	// Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
	// of 'MANAGED' in 'type').
	Managed *MangedSslCertificateManaged `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
	// namespace as the managed SSL certificates.
	Name     *string `pulumi:"name"`
	Project  *string `pulumi:"project"`
	SelfLink *string `pulumi:"selfLink"`
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// Enum field whose value is always 'MANAGED' - used to signal to the API which type this is. Default value: "MANAGED"
	// Possible values: ["MANAGED"]
	Type *string `pulumi:"type"`
}

type MangedSslCertificateState struct {
	// The unique identifier for the resource.
	CertificateId pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Expire time of the certificate.
	ExpireTime pulumi.StringPtrInput
	// Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
	// of 'MANAGED' in 'type').
	Managed MangedSslCertificateManagedPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
	// namespace as the managed SSL certificates.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayInput
	// Enum field whose value is always 'MANAGED' - used to signal to the API which type this is. Default value: "MANAGED"
	// Possible values: ["MANAGED"]
	Type pulumi.StringPtrInput
}

func (MangedSslCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*mangedSslCertificateState)(nil)).Elem()
}

type mangedSslCertificateArgs struct {
	// The unique identifier for the resource.
	CertificateId *int `pulumi:"certificateId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
	// of 'MANAGED' in 'type').
	Managed *MangedSslCertificateManaged `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
	// namespace as the managed SSL certificates.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Enum field whose value is always 'MANAGED' - used to signal to the API which type this is. Default value: "MANAGED"
	// Possible values: ["MANAGED"]
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a MangedSslCertificate resource.
type MangedSslCertificateArgs struct {
	// The unique identifier for the resource.
	CertificateId pulumi.IntPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
	// of 'MANAGED' in 'type').
	Managed MangedSslCertificateManagedPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
	// namespace as the managed SSL certificates.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Enum field whose value is always 'MANAGED' - used to signal to the API which type this is. Default value: "MANAGED"
	// Possible values: ["MANAGED"]
	Type pulumi.StringPtrInput
}

func (MangedSslCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mangedSslCertificateArgs)(nil)).Elem()
}
