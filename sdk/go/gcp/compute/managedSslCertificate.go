// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
//
// ## Example Usage
// ### Managed Ssl Certificate Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultManagedSslCertificate, err := compute.NewManagedSslCertificate(ctx, "defaultManagedSslCertificate", &compute.ManagedSslCertificateArgs{
// 			Managed: &compute.ManagedSslCertificateManagedArgs{
// 				Domains: pulumi.StringArray{
// 					pulumi.String("sslcert.tf-test.club."),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultHttpHealthCheck, err := compute.NewHttpHealthCheck(ctx, "defaultHttpHealthCheck", &compute.HttpHealthCheckArgs{
// 			RequestPath:      pulumi.String("/"),
// 			CheckIntervalSec: pulumi.Int(1),
// 			TimeoutSec:       pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultBackendService, err := compute.NewBackendService(ctx, "defaultBackendService", &compute.BackendServiceArgs{
// 			PortName:   pulumi.String("http"),
// 			Protocol:   pulumi.String("HTTP"),
// 			TimeoutSec: pulumi.Int(10),
// 			HealthChecks: pulumi.String(pulumi.String{
// 				defaultHttpHealthCheck.ID(),
// 			}),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultURLMap, err := compute.NewURLMap(ctx, "defaultURLMap", &compute.URLMapArgs{
// 			Description:    pulumi.String("a description"),
// 			DefaultService: defaultBackendService.ID(),
// 			HostRules: compute.URLMapHostRuleArray{
// 				&compute.URLMapHostRuleArgs{
// 					Hosts: pulumi.StringArray{
// 						pulumi.String("sslcert.tf-test.club"),
// 					},
// 					PathMatcher: pulumi.String("allpaths"),
// 				},
// 			},
// 			PathMatchers: compute.URLMapPathMatcherArray{
// 				&compute.URLMapPathMatcherArgs{
// 					Name:           pulumi.String("allpaths"),
// 					DefaultService: defaultBackendService.ID(),
// 					PathRules: compute.URLMapPathMatcherPathRuleArray{
// 						&compute.URLMapPathMatcherPathRuleArgs{
// 							Paths: pulumi.StringArray{
// 								pulumi.String("/*"),
// 							},
// 							Service: defaultBackendService.ID(),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultTargetHttpsProxy, err := compute.NewTargetHttpsProxy(ctx, "defaultTargetHttpsProxy", &compute.TargetHttpsProxyArgs{
// 			UrlMap: defaultURLMap.ID(),
// 			SslCertificates: pulumi.StringArray{
// 				defaultManagedSslCertificate.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		zone, err := dns.NewManagedZone(ctx, "zone", &dns.ManagedZoneArgs{
// 			DnsName: pulumi.String("sslcert.tf-test.club."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultGlobalForwardingRule, err := compute.NewGlobalForwardingRule(ctx, "defaultGlobalForwardingRule", &compute.GlobalForwardingRuleArgs{
// 			Target:    defaultTargetHttpsProxy.ID(),
// 			PortRange: pulumi.String("443"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dns.NewRecordSet(ctx, "set", &dns.RecordSetArgs{
// 			Type:        pulumi.String("A"),
// 			Ttl:         pulumi.Int(3600),
// 			ManagedZone: zone.Name,
// 			Rrdatas: pulumi.StringArray{
// 				defaultGlobalForwardingRule.IpAddress,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ManagedSslCertificate can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/managedSslCertificate:ManagedSslCertificate default projects/{{project}}/global/sslCertificates/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/managedSslCertificate:ManagedSslCertificate default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/managedSslCertificate:ManagedSslCertificate default {{name}}
// ```
type ManagedSslCertificate struct {
	pulumi.CustomResourceState

	// The unique identifier for the resource.
	CertificateId pulumi.IntOutput `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Expire time of the certificate.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).
	// Structure is documented below.
	Managed ManagedSslCertificateManagedPtrOutput `pulumi:"managed"`
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
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	// Default value is `MANAGED`.
	// Possible values are `MANAGED`.
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
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Expire time of the certificate.
	ExpireTime *string `pulumi:"expireTime"`
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).
	// Structure is documented below.
	Managed *ManagedSslCertificateManaged `pulumi:"managed"`
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
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	// Default value is `MANAGED`.
	// Possible values are `MANAGED`.
	Type *string `pulumi:"type"`
}

type ManagedSslCertificateState struct {
	// The unique identifier for the resource.
	CertificateId pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Expire time of the certificate.
	ExpireTime pulumi.StringPtrInput
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).
	// Structure is documented below.
	Managed ManagedSslCertificateManagedPtrInput
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
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	// Default value is `MANAGED`.
	// Possible values are `MANAGED`.
	Type pulumi.StringPtrInput
}

func (ManagedSslCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedSslCertificateState)(nil)).Elem()
}

type managedSslCertificateArgs struct {
	// The unique identifier for the resource.
	CertificateId *int `pulumi:"certificateId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).
	// Structure is documented below.
	Managed *ManagedSslCertificateManaged `pulumi:"managed"`
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
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	// Default value is `MANAGED`.
	// Possible values are `MANAGED`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ManagedSslCertificate resource.
type ManagedSslCertificateArgs struct {
	// The unique identifier for the resource.
	CertificateId pulumi.IntPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of `MANAGED` in `type`).
	// Structure is documented below.
	Managed ManagedSslCertificateManagedPtrInput
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
	// Enum field whose value is always `MANAGED` - used to signal to the API
	// which type this is.
	// Default value is `MANAGED`.
	// Possible values are `MANAGED`.
	Type pulumi.StringPtrInput
}

func (ManagedSslCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedSslCertificateArgs)(nil)).Elem()
}

type ManagedSslCertificateInput interface {
	pulumi.Input

	ToManagedSslCertificateOutput() ManagedSslCertificateOutput
	ToManagedSslCertificateOutputWithContext(ctx context.Context) ManagedSslCertificateOutput
}

func (*ManagedSslCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedSslCertificate)(nil))
}

func (i *ManagedSslCertificate) ToManagedSslCertificateOutput() ManagedSslCertificateOutput {
	return i.ToManagedSslCertificateOutputWithContext(context.Background())
}

func (i *ManagedSslCertificate) ToManagedSslCertificateOutputWithContext(ctx context.Context) ManagedSslCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedSslCertificateOutput)
}

type ManagedSslCertificateOutput struct {
	*pulumi.OutputState
}

func (ManagedSslCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedSslCertificate)(nil))
}

func (o ManagedSslCertificateOutput) ToManagedSslCertificateOutput() ManagedSslCertificateOutput {
	return o
}

func (o ManagedSslCertificateOutput) ToManagedSslCertificateOutputWithContext(ctx context.Context) ManagedSslCertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedSslCertificateOutput{})
}
