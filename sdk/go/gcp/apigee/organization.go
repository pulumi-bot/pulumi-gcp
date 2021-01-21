// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigee

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An `Organization` is the top-level container in Apigee.
//
// To get more information about Organization, see:
//
// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations)
// * How-to Guides
//     * [Creating an API organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org)
//
// ## Example Usage
// ### Apigee Organization Cloud Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/apigee"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/servicenetworking"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := organizations.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		apigeeNetwork, err := compute.NewNetwork(ctx, "apigeeNetwork", nil)
// 		if err != nil {
// 			return err
// 		}
// 		apigeeRange, err := compute.NewGlobalAddress(ctx, "apigeeRange", &compute.GlobalAddressArgs{
// 			Purpose:      pulumi.String("VPC_PEERING"),
// 			AddressType:  pulumi.String("INTERNAL"),
// 			PrefixLength: pulumi.Int(16),
// 			Network:      apigeeNetwork.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		apigeeVpcConnection, err := servicenetworking.NewConnection(ctx, "apigeeVpcConnection", &servicenetworking.ConnectionArgs{
// 			Network: apigeeNetwork.ID(),
// 			Service: pulumi.String("servicenetworking.googleapis.com"),
// 			ReservedPeeringRanges: pulumi.StringArray{
// 				apigeeRange.Name,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigee.NewOrganization(ctx, "org", &apigee.OrganizationArgs{
// 			AnalyticsRegion:   pulumi.String("us-central1"),
// 			ProjectId:         pulumi.String(current.Project),
// 			AuthorizedNetwork: apigeeNetwork.ID(),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			apigeeVpcConnection,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Apigee Organization Cloud Full
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/apigee"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/servicenetworking"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := organizations.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		apigeeNetwork, err := compute.NewNetwork(ctx, "apigeeNetwork", nil)
// 		if err != nil {
// 			return err
// 		}
// 		apigeeRange, err := compute.NewGlobalAddress(ctx, "apigeeRange", &compute.GlobalAddressArgs{
// 			Purpose:      pulumi.String("VPC_PEERING"),
// 			AddressType:  pulumi.String("INTERNAL"),
// 			PrefixLength: pulumi.Int(16),
// 			Network:      apigeeNetwork.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		apigeeVpcConnection, err := servicenetworking.NewConnection(ctx, "apigeeVpcConnection", &servicenetworking.ConnectionArgs{
// 			Network: apigeeNetwork.ID(),
// 			Service: pulumi.String("servicenetworking.googleapis.com"),
// 			ReservedPeeringRanges: pulumi.StringArray{
// 				apigeeRange.Name,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		apigeeKeyring, err := kms.NewKeyRing(ctx, "apigeeKeyring", &kms.KeyRingArgs{
// 			Location: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		apigeeKey, err := kms.NewCryptoKey(ctx, "apigeeKey", &kms.CryptoKeyArgs{
// 			KeyRing: apigeeKeyring.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		apigeeSa, err := projects.NewServiceIdentity(ctx, "apigeeSa", &projects.ServiceIdentityArgs{
// 			Project: pulumi.Any(google_project.Project.Project_id),
// 			Service: pulumi.Any(google_project_service.Apigee.Service),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		apigeeSaKeyuser, err := kms.NewCryptoKeyIAMBinding(ctx, "apigeeSaKeyuser", &kms.CryptoKeyIAMBindingArgs{
// 			CryptoKeyId: apigeeKey.ID(),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypterDecrypter"),
// 			Members: pulumi.StringArray{
// 				apigeeSa.Email.ApplyT(func(email string) (string, error) {
// 					return fmt.Sprintf("%v%v", "serviceAccount:", email), nil
// 				}).(pulumi.StringOutput),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigee.NewOrganization(ctx, "org", &apigee.OrganizationArgs{
// 			AnalyticsRegion:                  pulumi.String("us-central1"),
// 			DisplayName:                      pulumi.String("apigee-org"),
// 			Description:                      pulumi.String("Terraform-provisioned Apigee Org."),
// 			ProjectId:                        pulumi.String(current.Project),
// 			AuthorizedNetwork:                apigeeNetwork.ID(),
// 			RuntimeDatabaseEncryptionKeyName: pulumi.Any(google_kms_key.Apigee_key.Id),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			apigeeVpcConnection,
// 			apigeeSaKeyuser,
// 		}))
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
// Organization can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:apigee/organization:Organization default organizations/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:apigee/organization:Organization default {{name}}
// ```
type Organization struct {
	pulumi.CustomResourceState

	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion pulumi.StringPtrOutput `pulumi:"analyticsRegion"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork pulumi.StringPtrOutput `pulumi:"authorizedNetwork"`
	// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when 'RuntimeType'
	// is CLOUD. A base64-encoded string.
	CaCertificate pulumi.StringOutput `pulumi:"caCertificate"`
	// Description of the Apigee organization.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the Apigee organization.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Output only. Name of the Apigee organization.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project ID associated with the Apigee organization.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	RuntimeDatabaseEncryptionKeyName pulumi.StringPtrOutput `pulumi:"runtimeDatabaseEncryptionKeyName"`
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is `CLOUD`.
	// Possible values are `CLOUD` and `HYBRID`.
	RuntimeType pulumi.StringPtrOutput `pulumi:"runtimeType"`
	// Output only. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation
	// purposes only) or paid (full subscription has been purchased).
	SubscriptionType pulumi.StringOutput `pulumi:"subscriptionType"`
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource Organization
	err := ctx.RegisterResource("gcp:apigee/organization:Organization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	var resource Organization
	err := ctx.ReadResource("gcp:apigee/organization:Organization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Organization resources.
type organizationState struct {
	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion *string `pulumi:"analyticsRegion"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when 'RuntimeType'
	// is CLOUD. A base64-encoded string.
	CaCertificate *string `pulumi:"caCertificate"`
	// Description of the Apigee organization.
	Description *string `pulumi:"description"`
	// The display name of the Apigee organization.
	DisplayName *string `pulumi:"displayName"`
	// Output only. Name of the Apigee organization.
	Name *string `pulumi:"name"`
	// The project ID associated with the Apigee organization.
	ProjectId *string `pulumi:"projectId"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	RuntimeDatabaseEncryptionKeyName *string `pulumi:"runtimeDatabaseEncryptionKeyName"`
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is `CLOUD`.
	// Possible values are `CLOUD` and `HYBRID`.
	RuntimeType *string `pulumi:"runtimeType"`
	// Output only. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation
	// purposes only) or paid (full subscription has been purchased).
	SubscriptionType *string `pulumi:"subscriptionType"`
}

type OrganizationState struct {
	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion pulumi.StringPtrInput
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork pulumi.StringPtrInput
	// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when 'RuntimeType'
	// is CLOUD. A base64-encoded string.
	CaCertificate pulumi.StringPtrInput
	// Description of the Apigee organization.
	Description pulumi.StringPtrInput
	// The display name of the Apigee organization.
	DisplayName pulumi.StringPtrInput
	// Output only. Name of the Apigee organization.
	Name pulumi.StringPtrInput
	// The project ID associated with the Apigee organization.
	ProjectId pulumi.StringPtrInput
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	RuntimeDatabaseEncryptionKeyName pulumi.StringPtrInput
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is `CLOUD`.
	// Possible values are `CLOUD` and `HYBRID`.
	RuntimeType pulumi.StringPtrInput
	// Output only. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation
	// purposes only) or paid (full subscription has been purchased).
	SubscriptionType pulumi.StringPtrInput
}

func (OrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationState)(nil)).Elem()
}

type organizationArgs struct {
	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion *string `pulumi:"analyticsRegion"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// Description of the Apigee organization.
	Description *string `pulumi:"description"`
	// The display name of the Apigee organization.
	DisplayName *string `pulumi:"displayName"`
	// The project ID associated with the Apigee organization.
	ProjectId string `pulumi:"projectId"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	RuntimeDatabaseEncryptionKeyName *string `pulumi:"runtimeDatabaseEncryptionKeyName"`
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is `CLOUD`.
	// Possible values are `CLOUD` and `HYBRID`.
	RuntimeType *string `pulumi:"runtimeType"`
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion pulumi.StringPtrInput
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork pulumi.StringPtrInput
	// Description of the Apigee organization.
	Description pulumi.StringPtrInput
	// The display name of the Apigee organization.
	DisplayName pulumi.StringPtrInput
	// The project ID associated with the Apigee organization.
	ProjectId pulumi.StringInput
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	RuntimeDatabaseEncryptionKeyName pulumi.StringPtrInput
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is `CLOUD`.
	// Possible values are `CLOUD` and `HYBRID`.
	RuntimeType pulumi.StringPtrInput
}

func (OrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationArgs)(nil)).Elem()
}

type OrganizationInput interface {
	pulumi.Input

	ToOrganizationOutput() OrganizationOutput
	ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput
}

func (Organization) ElementType() reflect.Type {
	return reflect.TypeOf((*Organization)(nil)).Elem()
}

func (i Organization) ToOrganizationOutput() OrganizationOutput {
	return i.ToOrganizationOutputWithContext(context.Background())
}

func (i Organization) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationOutput)
}

type OrganizationOutput struct {
	*pulumi.OutputState
}

func (OrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationOutput)(nil)).Elem()
}

func (o OrganizationOutput) ToOrganizationOutput() OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationOutput{})
}
