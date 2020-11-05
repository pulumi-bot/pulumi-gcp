// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package activedirectory

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a Microsoft AD domain
//
// To get more information about Domain, see:
//
// * [API documentation](https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains)
// * How-to Guides
//     * [Managed Microsoft Active Directory Quickstart](https://cloud.google.com/managed-microsoft-ad/docs/quickstarts)
//
// ## Example Usage
type Domain struct {
	pulumi.CustomResourceState

	// The name of delegated administrator account used to perform Active Directory operations.
	// If not specified, setupadmin will be used.
	Admin pulumi.StringPtrOutput `pulumi:"admin"`
	// The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
	// If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks pulumi.StringArrayOutput `pulumi:"authorizedNetworks"`
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would
	// be chosen for an Active Directory set up on an internal network.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Resource labels that can contain user-provided metadata
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations pulumi.StringArrayOutput `pulumi:"locations"`
	// The unique name of the domain using the format: 'projects/{project}/locations/global/domains/{domainName}'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
	// Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	ReservedIpRange pulumi.StringOutput `pulumi:"reservedIpRange"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.ReservedIpRange == nil {
		return nil, errors.New("invalid value for required argument 'ReservedIpRange'")
	}
	var resource Domain
	err := ctx.RegisterResource("gcp:activedirectory/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("gcp:activedirectory/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// The name of delegated administrator account used to perform Active Directory operations.
	// If not specified, setupadmin will be used.
	Admin *string `pulumi:"admin"`
	// The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
	// If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks []string `pulumi:"authorizedNetworks"`
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	DomainName *string `pulumi:"domainName"`
	// The fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would
	// be chosen for an Active Directory set up on an internal network.
	Fqdn *string `pulumi:"fqdn"`
	// Resource labels that can contain user-provided metadata
	Labels map[string]string `pulumi:"labels"`
	// Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations []string `pulumi:"locations"`
	// The unique name of the domain using the format: 'projects/{project}/locations/global/domains/{domainName}'.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
	// Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	ReservedIpRange *string `pulumi:"reservedIpRange"`
}

type DomainState struct {
	// The name of delegated administrator account used to perform Active Directory operations.
	// If not specified, setupadmin will be used.
	Admin pulumi.StringPtrInput
	// The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
	// If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks pulumi.StringArrayInput
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	DomainName pulumi.StringPtrInput
	// The fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would
	// be chosen for an Active Directory set up on an internal network.
	Fqdn pulumi.StringPtrInput
	// Resource labels that can contain user-provided metadata
	Labels pulumi.StringMapInput
	// Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations pulumi.StringArrayInput
	// The unique name of the domain using the format: 'projects/{project}/locations/global/domains/{domainName}'.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
	// Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	ReservedIpRange pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// The name of delegated administrator account used to perform Active Directory operations.
	// If not specified, setupadmin will be used.
	Admin *string `pulumi:"admin"`
	// The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
	// If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks []string `pulumi:"authorizedNetworks"`
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	DomainName string `pulumi:"domainName"`
	// Resource labels that can contain user-provided metadata
	Labels map[string]string `pulumi:"labels"`
	// Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations []string `pulumi:"locations"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
	// Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	ReservedIpRange string `pulumi:"reservedIpRange"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The name of delegated administrator account used to perform Active Directory operations.
	// If not specified, setupadmin will be used.
	Admin pulumi.StringPtrInput
	// The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
	// If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks pulumi.StringArrayInput
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	DomainName pulumi.StringInput
	// Resource labels that can contain user-provided metadata
	Labels pulumi.StringMapInput
	// Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
	// Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	ReservedIpRange pulumi.StringInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}
