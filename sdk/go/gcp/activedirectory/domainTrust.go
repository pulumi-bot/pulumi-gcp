// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package activedirectory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Adds a trust between Active Directory domains
//
// To get more information about DomainTrust, see:
//
// * [API documentation](https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains/attachTrust)
// * How-to Guides
//     * [Active Directory Trust](https://cloud.google.com/managed-microsoft-ad/docs/create-one-way-trust)
//
// > **Warning:** All arguments including `trustHandshakeSecret` will be stored in the raw
// state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// ## Example Usage
// ### Active Directory Domain Trust Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/activedirectory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := activedirectory.NewDomainTrust(ctx, "ad_domain_trust", &activedirectory.DomainTrustArgs{
// 			Domain: pulumi.String("test-managed-ad.com"),
// 			TargetDnsIpAddresses: pulumi.StringArray{
// 				pulumi.String("10.1.0.100"),
// 			},
// 			TargetDomainName:     pulumi.String("example-gcp.com"),
// 			TrustDirection:       pulumi.String("OUTBOUND"),
// 			TrustHandshakeSecret: pulumi.String("Testing1!"),
// 			TrustType:            pulumi.String("FOREST"),
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
// DomainTrust can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}
// ```
//
// ```sh
//  $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default {{project}}/{{domain}}/{{target_domain_name}}
// ```
//
// ```sh
//  $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default {{domain}}/{{target_domain_name}}
// ```
type DomainTrust struct {
	pulumi.CustomResourceState

	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
	SelectiveAuthentication pulumi.BoolPtrOutput `pulumi:"selectiveAuthentication"`
	// The target DNS server IP addresses which can resolve the remote domain involved in the trust.
	TargetDnsIpAddresses pulumi.StringArrayOutput `pulumi:"targetDnsIpAddresses"`
	// The fully qualified target domain name which will be in trust with the current domain.
	TargetDomainName pulumi.StringOutput `pulumi:"targetDomainName"`
	// The trust direction, which decides if the current domain is trusted, trusting, or both.
	// Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
	TrustDirection pulumi.StringOutput `pulumi:"trustDirection"`
	// The trust secret used for the handshake with the target domain. This will not be stored.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	TrustHandshakeSecret pulumi.StringOutput `pulumi:"trustHandshakeSecret"`
	// The type of trust represented by the trust resource.
	// Possible values are `FOREST` and `EXTERNAL`.
	TrustType pulumi.StringOutput `pulumi:"trustType"`
}

// NewDomainTrust registers a new resource with the given unique name, arguments, and options.
func NewDomainTrust(ctx *pulumi.Context,
	name string, args *DomainTrustArgs, opts ...pulumi.ResourceOption) (*DomainTrust, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.TargetDnsIpAddresses == nil {
		return nil, errors.New("invalid value for required argument 'TargetDnsIpAddresses'")
	}
	if args.TargetDomainName == nil {
		return nil, errors.New("invalid value for required argument 'TargetDomainName'")
	}
	if args.TrustDirection == nil {
		return nil, errors.New("invalid value for required argument 'TrustDirection'")
	}
	if args.TrustHandshakeSecret == nil {
		return nil, errors.New("invalid value for required argument 'TrustHandshakeSecret'")
	}
	if args.TrustType == nil {
		return nil, errors.New("invalid value for required argument 'TrustType'")
	}
	var resource DomainTrust
	err := ctx.RegisterResource("gcp:activedirectory/domainTrust:DomainTrust", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainTrust gets an existing DomainTrust resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainTrust(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainTrustState, opts ...pulumi.ResourceOption) (*DomainTrust, error) {
	var resource DomainTrust
	err := ctx.ReadResource("gcp:activedirectory/domainTrust:DomainTrust", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainTrust resources.
type domainTrustState struct {
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	Domain *string `pulumi:"domain"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
	SelectiveAuthentication *bool `pulumi:"selectiveAuthentication"`
	// The target DNS server IP addresses which can resolve the remote domain involved in the trust.
	TargetDnsIpAddresses []string `pulumi:"targetDnsIpAddresses"`
	// The fully qualified target domain name which will be in trust with the current domain.
	TargetDomainName *string `pulumi:"targetDomainName"`
	// The trust direction, which decides if the current domain is trusted, trusting, or both.
	// Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
	TrustDirection *string `pulumi:"trustDirection"`
	// The trust secret used for the handshake with the target domain. This will not be stored.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	TrustHandshakeSecret *string `pulumi:"trustHandshakeSecret"`
	// The type of trust represented by the trust resource.
	// Possible values are `FOREST` and `EXTERNAL`.
	TrustType *string `pulumi:"trustType"`
}

type DomainTrustState struct {
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	Domain pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
	SelectiveAuthentication pulumi.BoolPtrInput
	// The target DNS server IP addresses which can resolve the remote domain involved in the trust.
	TargetDnsIpAddresses pulumi.StringArrayInput
	// The fully qualified target domain name which will be in trust with the current domain.
	TargetDomainName pulumi.StringPtrInput
	// The trust direction, which decides if the current domain is trusted, trusting, or both.
	// Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
	TrustDirection pulumi.StringPtrInput
	// The trust secret used for the handshake with the target domain. This will not be stored.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	TrustHandshakeSecret pulumi.StringPtrInput
	// The type of trust represented by the trust resource.
	// Possible values are `FOREST` and `EXTERNAL`.
	TrustType pulumi.StringPtrInput
}

func (DomainTrustState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainTrustState)(nil)).Elem()
}

type domainTrustArgs struct {
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	Domain string `pulumi:"domain"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
	SelectiveAuthentication *bool `pulumi:"selectiveAuthentication"`
	// The target DNS server IP addresses which can resolve the remote domain involved in the trust.
	TargetDnsIpAddresses []string `pulumi:"targetDnsIpAddresses"`
	// The fully qualified target domain name which will be in trust with the current domain.
	TargetDomainName string `pulumi:"targetDomainName"`
	// The trust direction, which decides if the current domain is trusted, trusting, or both.
	// Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
	TrustDirection string `pulumi:"trustDirection"`
	// The trust secret used for the handshake with the target domain. This will not be stored.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	TrustHandshakeSecret string `pulumi:"trustHandshakeSecret"`
	// The type of trust represented by the trust resource.
	// Possible values are `FOREST` and `EXTERNAL`.
	TrustType string `pulumi:"trustType"`
}

// The set of arguments for constructing a DomainTrust resource.
type DomainTrustArgs struct {
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	Domain pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
	SelectiveAuthentication pulumi.BoolPtrInput
	// The target DNS server IP addresses which can resolve the remote domain involved in the trust.
	TargetDnsIpAddresses pulumi.StringArrayInput
	// The fully qualified target domain name which will be in trust with the current domain.
	TargetDomainName pulumi.StringInput
	// The trust direction, which decides if the current domain is trusted, trusting, or both.
	// Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
	TrustDirection pulumi.StringInput
	// The trust secret used for the handshake with the target domain. This will not be stored.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	TrustHandshakeSecret pulumi.StringInput
	// The type of trust represented by the trust resource.
	// Possible values are `FOREST` and `EXTERNAL`.
	TrustType pulumi.StringInput
}

func (DomainTrustArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainTrustArgs)(nil)).Elem()
}

type DomainTrustInput interface {
	pulumi.Input

	ToDomainTrustOutput() DomainTrustOutput
	ToDomainTrustOutputWithContext(ctx context.Context) DomainTrustOutput
}

func (DomainTrust) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainTrust)(nil)).Elem()
}

func (i DomainTrust) ToDomainTrustOutput() DomainTrustOutput {
	return i.ToDomainTrustOutputWithContext(context.Background())
}

func (i DomainTrust) ToDomainTrustOutputWithContext(ctx context.Context) DomainTrustOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainTrustOutput)
}

type DomainTrustOutput struct {
	*pulumi.OutputState
}

func (DomainTrustOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainTrustOutput)(nil)).Elem()
}

func (o DomainTrustOutput) ToDomainTrustOutput() DomainTrustOutput {
	return o
}

func (o DomainTrustOutput) ToDomainTrustOutputWithContext(ctx context.Context) DomainTrustOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainTrustOutput{})
}
