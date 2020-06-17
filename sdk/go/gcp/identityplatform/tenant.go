// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identityplatform

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Tenant configuration in a multi-tenant project.
//
// You must enable the
// [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
// the marketplace prior to using this resource.
//
// You must [enable multi-tenancy](https://cloud.google.com/identity-platform/docs/multi-tenancy-quickstart) via
// the Cloud Console prior to creating tenants.
//
//
//
// ## Example Usage
//
// ### Identity Platform Tenant Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/identityplatform"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = identityplatform.NewTenant(ctx, "tenant", &identityplatform.TenantArgs{
// 			AllowPasswordSignup: pulumi.Bool(true),
// 			DisplayName:         pulumi.String("tenant"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Tenant struct {
	pulumi.CustomResourceState

	// Whether to allow email/password user authentication.
	AllowPasswordSignup pulumi.BoolPtrOutput `pulumi:"allowPasswordSignup"`
	// Whether authentication is disabled for the tenant. If true, the users under
	// the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
	// are not able to manage its users.
	DisableAuth pulumi.BoolPtrOutput `pulumi:"disableAuth"`
	// Human friendly display name of the tenant.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Whether to enable email link user authentication.
	EnableEmailLinkSignin pulumi.BoolPtrOutput `pulumi:"enableEmailLinkSignin"`
	// The name of the tenant that is generated by the server
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewTenant registers a new resource with the given unique name, arguments, and options.
func NewTenant(ctx *pulumi.Context,
	name string, args *TenantArgs, opts ...pulumi.ResourceOption) (*Tenant, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil {
		args = &TenantArgs{}
	}
	var resource Tenant
	err := ctx.RegisterResource("gcp:identityplatform/tenant:Tenant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenant gets an existing Tenant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantState, opts ...pulumi.ResourceOption) (*Tenant, error) {
	var resource Tenant
	err := ctx.ReadResource("gcp:identityplatform/tenant:Tenant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tenant resources.
type tenantState struct {
	// Whether to allow email/password user authentication.
	AllowPasswordSignup *bool `pulumi:"allowPasswordSignup"`
	// Whether authentication is disabled for the tenant. If true, the users under
	// the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
	// are not able to manage its users.
	DisableAuth *bool `pulumi:"disableAuth"`
	// Human friendly display name of the tenant.
	DisplayName *string `pulumi:"displayName"`
	// Whether to enable email link user authentication.
	EnableEmailLinkSignin *bool `pulumi:"enableEmailLinkSignin"`
	// The name of the tenant that is generated by the server
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type TenantState struct {
	// Whether to allow email/password user authentication.
	AllowPasswordSignup pulumi.BoolPtrInput
	// Whether authentication is disabled for the tenant. If true, the users under
	// the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
	// are not able to manage its users.
	DisableAuth pulumi.BoolPtrInput
	// Human friendly display name of the tenant.
	DisplayName pulumi.StringPtrInput
	// Whether to enable email link user authentication.
	EnableEmailLinkSignin pulumi.BoolPtrInput
	// The name of the tenant that is generated by the server
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (TenantState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantState)(nil)).Elem()
}

type tenantArgs struct {
	// Whether to allow email/password user authentication.
	AllowPasswordSignup *bool `pulumi:"allowPasswordSignup"`
	// Whether authentication is disabled for the tenant. If true, the users under
	// the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
	// are not able to manage its users.
	DisableAuth *bool `pulumi:"disableAuth"`
	// Human friendly display name of the tenant.
	DisplayName string `pulumi:"displayName"`
	// Whether to enable email link user authentication.
	EnableEmailLinkSignin *bool `pulumi:"enableEmailLinkSignin"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Tenant resource.
type TenantArgs struct {
	// Whether to allow email/password user authentication.
	AllowPasswordSignup pulumi.BoolPtrInput
	// Whether authentication is disabled for the tenant. If true, the users under
	// the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
	// are not able to manage its users.
	DisableAuth pulumi.BoolPtrInput
	// Human friendly display name of the tenant.
	DisplayName pulumi.StringInput
	// Whether to enable email link user authentication.
	EnableEmailLinkSignin pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (TenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantArgs)(nil)).Elem()
}
