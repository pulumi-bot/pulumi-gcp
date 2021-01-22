// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package folder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows management of Organization policies for a Google Folder. For more information see
// [the official documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
// [API](https://cloud.google.com/resource-manager/reference/rest/v1/folders/setOrgPolicy).
//
// ## Example Usage
//
// To set policy with a [boolean constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-boolean-constraints):
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/folder"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := folder.NewOrganizationPolicy(ctx, "serialPortPolicy", &folder.OrganizationPolicyArgs{
// 			BooleanPolicy: &folder.OrganizationPolicyBooleanPolicyArgs{
// 				Enforced: pulumi.Bool(true),
// 			},
// 			Constraint: pulumi.String("compute.disableSerialPortAccess"),
// 			Folder:     pulumi.String("folders/123456789"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To set a policy with a [list constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-list-constraints):
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/folder"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := folder.NewOrganizationPolicy(ctx, "servicesPolicy", &folder.OrganizationPolicyArgs{
// 			Constraint: pulumi.String("serviceuser.services"),
// 			Folder:     pulumi.String("folders/123456789"),
// 			ListPolicy: &folder.OrganizationPolicyListPolicyArgs{
// 				Allow: &folder.OrganizationPolicyListPolicyAllowArgs{
// 					All: pulumi.Bool(true),
// 				},
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
// Or to deny some services, use the following instead:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/folder"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := folder.NewOrganizationPolicy(ctx, "servicesPolicy", &folder.OrganizationPolicyArgs{
// 			Constraint: pulumi.String("serviceuser.services"),
// 			Folder:     pulumi.String("folders/123456789"),
// 			ListPolicy: &folder.OrganizationPolicyListPolicyArgs{
// 				Deny: &folder.OrganizationPolicyListPolicyDenyArgs{
// 					Values: pulumi.StringArray{
// 						pulumi.String("cloudresourcemanager.googleapis.com"),
// 					},
// 				},
// 				SuggestedValue: pulumi.String("compute.googleapis.com"),
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
// To restore the default folder organization policy, use the following instead:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/folder"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := folder.NewOrganizationPolicy(ctx, "servicesPolicy", &folder.OrganizationPolicyArgs{
// 			Constraint: pulumi.String("serviceuser.services"),
// 			Folder:     pulumi.String("folders/123456789"),
// 			RestorePolicy: &folder.OrganizationPolicyRestorePolicyArgs{
// 				Default: pulumi.Bool(true),
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
// Folder organization policies can be imported using any of the follow formats
//
// ```sh
//  $ pulumi import gcp:folder/organizationPolicy:OrganizationPolicy policy folders/folder-1234/constraints/serviceuser.services
// ```
//
// ```sh
//  $ pulumi import gcp:folder/organizationPolicy:OrganizationPolicy policy folder-1234/serviceuser.services
// ```
type OrganizationPolicy struct {
	pulumi.CustomResourceState

	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy OrganizationPolicyBooleanPolicyPtrOutput `pulumi:"booleanPolicy"`
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint pulumi.StringOutput `pulumi:"constraint"`
	// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It
	// can also be used to allow or deny all values. Structure is documented below.
	ListPolicy OrganizationPolicyListPolicyPtrOutput `pulumi:"listPolicy"`
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy OrganizationPolicyRestorePolicyPtrOutput `pulumi:"restorePolicy"`
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Version of the Policy. Default version is 0.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewOrganizationPolicy registers a new resource with the given unique name, arguments, and options.
func NewOrganizationPolicy(ctx *pulumi.Context,
	name string, args *OrganizationPolicyArgs, opts ...pulumi.ResourceOption) (*OrganizationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Constraint == nil {
		return nil, errors.New("invalid value for required argument 'Constraint'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	var resource OrganizationPolicy
	err := ctx.RegisterResource("gcp:folder/organizationPolicy:OrganizationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationPolicy gets an existing OrganizationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationPolicyState, opts ...pulumi.ResourceOption) (*OrganizationPolicy, error) {
	var resource OrganizationPolicy
	err := ctx.ReadResource("gcp:folder/organizationPolicy:OrganizationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationPolicy resources.
type organizationPolicyState struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy *OrganizationPolicyBooleanPolicy `pulumi:"booleanPolicy"`
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint *string `pulumi:"constraint"`
	// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag *string `pulumi:"etag"`
	// The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
	Folder *string `pulumi:"folder"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It
	// can also be used to allow or deny all values. Structure is documented below.
	ListPolicy *OrganizationPolicyListPolicy `pulumi:"listPolicy"`
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy *OrganizationPolicyRestorePolicy `pulumi:"restorePolicy"`
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime *string `pulumi:"updateTime"`
	// Version of the Policy. Default version is 0.
	Version *int `pulumi:"version"`
}

type OrganizationPolicyState struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy OrganizationPolicyBooleanPolicyPtrInput
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint pulumi.StringPtrInput
	// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag pulumi.StringPtrInput
	// The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
	Folder pulumi.StringPtrInput
	// A policy that can define specific values that are allowed or denied for the given constraint. It
	// can also be used to allow or deny all values. Structure is documented below.
	ListPolicy OrganizationPolicyListPolicyPtrInput
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy OrganizationPolicyRestorePolicyPtrInput
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime pulumi.StringPtrInput
	// Version of the Policy. Default version is 0.
	Version pulumi.IntPtrInput
}

func (OrganizationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationPolicyState)(nil)).Elem()
}

type organizationPolicyArgs struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy *OrganizationPolicyBooleanPolicy `pulumi:"booleanPolicy"`
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint string `pulumi:"constraint"`
	// The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
	Folder string `pulumi:"folder"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It
	// can also be used to allow or deny all values. Structure is documented below.
	ListPolicy *OrganizationPolicyListPolicy `pulumi:"listPolicy"`
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy *OrganizationPolicyRestorePolicy `pulumi:"restorePolicy"`
	// Version of the Policy. Default version is 0.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a OrganizationPolicy resource.
type OrganizationPolicyArgs struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy OrganizationPolicyBooleanPolicyPtrInput
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint pulumi.StringInput
	// The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
	Folder pulumi.StringInput
	// A policy that can define specific values that are allowed or denied for the given constraint. It
	// can also be used to allow or deny all values. Structure is documented below.
	ListPolicy OrganizationPolicyListPolicyPtrInput
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy OrganizationPolicyRestorePolicyPtrInput
	// Version of the Policy. Default version is 0.
	Version pulumi.IntPtrInput
}

func (OrganizationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationPolicyArgs)(nil)).Elem()
}

type OrganizationPolicyInput interface {
	pulumi.Input

	ToOrganizationPolicyOutput() OrganizationPolicyOutput
	ToOrganizationPolicyOutputWithContext(ctx context.Context) OrganizationPolicyOutput
}

func (*OrganizationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationPolicy)(nil))
}

func (i *OrganizationPolicy) ToOrganizationPolicyOutput() OrganizationPolicyOutput {
	return i.ToOrganizationPolicyOutputWithContext(context.Background())
}

func (i *OrganizationPolicy) ToOrganizationPolicyOutputWithContext(ctx context.Context) OrganizationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationPolicyOutput)
}

type OrganizationPolicyOutput struct {
	*pulumi.OutputState
}

func (OrganizationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationPolicy)(nil))
}

func (o OrganizationPolicyOutput) ToOrganizationPolicyOutput() OrganizationPolicyOutput {
	return o
}

func (o OrganizationPolicyOutput) ToOrganizationPolicyOutputWithContext(ctx context.Context) OrganizationPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationPolicyOutput{})
}
