// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// An SslCertificate resource, used for HTTPS load balancing.  This resource
// represents a certificate for which the certificate secrets are created and
// managed by Google.
//
// For a resource where you provide the key, see the
// SSL Certificate resource.
//
// To get more information about ManagedSslCertificate, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
//
// > **Warning:** This resource should be used with extreme caution!  Provisioning an SSL
// certificate is complex.  Ensure that you understand the lifecycle of a
// certificate before attempting complex tasks like cert rotation automatically.
// This resource will "return" as soon as the certificate object is created,
// but post-creation the certificate object will go through a "provisioning"
// process.  The provisioning process can complete only when the domain name
// for which the certificate is created points to a target pool which, itself,
// points at the certificate.  Depending on your DNS provider, this may take
// some time, and migrating from self-managed certificates to Google-managed
// certificates may entail some downtime while the certificate provisions.
//
// In conclusion: Be extremely cautious.
type ManagedSslCertificate struct {
	pulumi.CustomResourceState

	// The unique identifier for the resource.
	CertificateId pulumi.IntOutput `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Expire time of the certificate.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// -
	// (Optional)
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).  Structure is documented below.
	Managed ManagedSslCertificateManagedPtrOutput `pulumi:"managed"`
	// -
	// (Optional)
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayOutput `pulumi:"subjectAlternativeNames"`
	// -
	// (Optional)
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewManagedSslCertificate registers a new resource with the given unique name, arguments, and options.
func NewManagedSslCertificate(ctx *pulumi.Context,
	name string, args *ManagedSslCertificateArgs, opts ...pulumi.ResourceOption) (*ManagedSslCertificate, error) {
	if args == nil {
		args = &ManagedSslCertificateArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("gcp:compute/mangedSslCertificate:MangedSslCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedSslCertificate
	err := ctx.RegisterResource("gcp:compute/managedSslCertificate:ManagedSslCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedSslCertificate gets an existing ManagedSslCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedSslCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedSslCertificateState, opts ...pulumi.ResourceOption) (*ManagedSslCertificate, error) {
	var resource ManagedSslCertificate
	err := ctx.ReadResource("gcp:compute/managedSslCertificate:ManagedSslCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedSslCertificate resources.
type managedSslCertificateState struct {
	// The unique identifier for the resource.
	CertificateId *int `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Expire time of the certificate.
	ExpireTime *string `pulumi:"expireTime"`
	// -
	// (Optional)
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).  Structure is documented below.
	Managed *ManagedSslCertificateManaged `pulumi:"managed"`
	// -
	// (Optional)
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// -
	// (Optional)
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	Type *string `pulumi:"type"`
}

type ManagedSslCertificateState struct {
	// The unique identifier for the resource.
	CertificateId pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Expire time of the certificate.
	ExpireTime pulumi.StringPtrInput
	// -
	// (Optional)
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).  Structure is documented below.
	Managed ManagedSslCertificateManagedPtrInput
	// -
	// (Optional)
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayInput
	// -
	// (Optional)
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	Type pulumi.StringPtrInput
}

func (ManagedSslCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedSslCertificateState)(nil)).Elem()
}

type managedSslCertificateArgs struct {
	// The unique identifier for the resource.
	CertificateId *int `pulumi:"certificateId"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// -
	// (Optional)
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).  Structure is documented below.
	Managed *ManagedSslCertificateManaged `pulumi:"managed"`
	// -
	// (Optional)
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Optional)
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ManagedSslCertificate resource.
type ManagedSslCertificateArgs struct {
	// The unique identifier for the resource.
	CertificateId pulumi.IntPtrInput
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// -
	// (Optional)
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).  Structure is documented below.
	Managed ManagedSslCertificateManagedPtrInput
	// -
	// (Optional)
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Optional)
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	Type pulumi.StringPtrInput
}

func (ManagedSslCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedSslCertificateArgs)(nil)).Elem()
}
