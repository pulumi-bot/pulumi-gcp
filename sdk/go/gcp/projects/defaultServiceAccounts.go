// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows management of Google Cloud Platform project default service accounts.
//
// When certain service APIs are enabled, Google Cloud Platform automatically creates service accounts to help get started, but
// this is not recommended for production environments as per [Google's documentation](https://cloud.google.com/iam/docs/service-accounts#default).
// See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
//
// > **WARNING** Some Google Cloud products do not work if the default service accounts are deleted so it is better to `DEPRIVILEGE` as
// Google **CAN NOT** recover service accounts that have been deleted for more than 30 days.
// Also Google recommends using the `constraints/iam.automaticIamGrantsForDefaultServiceAccounts` [constraint](https://www.terraform.io/docs/providers/google/r/google_organization_policy.html)
// to disable automatic IAM Grants to default service accounts.
//
// > This resource works on a best-effort basis, as no API formally describes the default service accounts
// and it is for users who are unable to use constraints. If the default service accounts change their name
// or additional service accounts are added, this resource will need to be updated.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewDefaultServiceAccounts(ctx, "myProject", &projects.DefaultServiceAccountsArgs{
// 			Action:  pulumi.String("DELETE"),
// 			Project: pulumi.String("my-project-id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To enable the default service accounts on the resource destroy:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewDefaultServiceAccounts(ctx, "myProject", &projects.DefaultServiceAccountsArgs{
// 			Action:        pulumi.String("DISABLE"),
// 			Project:       pulumi.String("my-project-id"),
// 			RestorePolicy: pulumi.String("REVERT"),
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
// This resource does not support import
type DefaultServiceAccounts struct {
	pulumi.CustomResourceState

	// The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
	Action pulumi.StringOutput `pulumi:"action"`
	// The project ID where service accounts are created.
	Project pulumi.StringOutput `pulumi:"project"`
	// The action to be performed in the default service accounts on the resource destroy.
	// Valid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	// If set to REVERT it attempts to restore all default SAs but the DEPRIVILEGE action.
	// If set to REVERT_AND_IGNORE_FAILURE it is the same behavior as REVERT but ignores errors returned by the API.
	RestorePolicy pulumi.StringPtrOutput `pulumi:"restorePolicy"`
	// The Service Accounts changed by this resource. It is used for `REVERT` the `action` on the destroy.
	ServiceAccounts pulumi.MapOutput `pulumi:"serviceAccounts"`
}

// NewDefaultServiceAccounts registers a new resource with the given unique name, arguments, and options.
func NewDefaultServiceAccounts(ctx *pulumi.Context,
	name string, args *DefaultServiceAccountsArgs, opts ...pulumi.ResourceOption) (*DefaultServiceAccounts, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource DefaultServiceAccounts
	err := ctx.RegisterResource("gcp:projects/defaultServiceAccounts:DefaultServiceAccounts", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultServiceAccounts gets an existing DefaultServiceAccounts resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultServiceAccounts(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultServiceAccountsState, opts ...pulumi.ResourceOption) (*DefaultServiceAccounts, error) {
	var resource DefaultServiceAccounts
	err := ctx.ReadResource("gcp:projects/defaultServiceAccounts:DefaultServiceAccounts", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultServiceAccounts resources.
type defaultServiceAccountsState struct {
	// The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
	Action *string `pulumi:"action"`
	// The project ID where service accounts are created.
	Project *string `pulumi:"project"`
	// The action to be performed in the default service accounts on the resource destroy.
	// Valid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	// If set to REVERT it attempts to restore all default SAs but the DEPRIVILEGE action.
	// If set to REVERT_AND_IGNORE_FAILURE it is the same behavior as REVERT but ignores errors returned by the API.
	RestorePolicy *string `pulumi:"restorePolicy"`
	// The Service Accounts changed by this resource. It is used for `REVERT` the `action` on the destroy.
	ServiceAccounts map[string]interface{} `pulumi:"serviceAccounts"`
}

type DefaultServiceAccountsState struct {
	// The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
	Action pulumi.StringPtrInput
	// The project ID where service accounts are created.
	Project pulumi.StringPtrInput
	// The action to be performed in the default service accounts on the resource destroy.
	// Valid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	// If set to REVERT it attempts to restore all default SAs but the DEPRIVILEGE action.
	// If set to REVERT_AND_IGNORE_FAILURE it is the same behavior as REVERT but ignores errors returned by the API.
	RestorePolicy pulumi.StringPtrInput
	// The Service Accounts changed by this resource. It is used for `REVERT` the `action` on the destroy.
	ServiceAccounts pulumi.MapInput
}

func (DefaultServiceAccountsState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultServiceAccountsState)(nil)).Elem()
}

type defaultServiceAccountsArgs struct {
	// The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
	Action string `pulumi:"action"`
	// The project ID where service accounts are created.
	Project string `pulumi:"project"`
	// The action to be performed in the default service accounts on the resource destroy.
	// Valid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	// If set to REVERT it attempts to restore all default SAs but the DEPRIVILEGE action.
	// If set to REVERT_AND_IGNORE_FAILURE it is the same behavior as REVERT but ignores errors returned by the API.
	RestorePolicy *string `pulumi:"restorePolicy"`
}

// The set of arguments for constructing a DefaultServiceAccounts resource.
type DefaultServiceAccountsArgs struct {
	// The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
	Action pulumi.StringInput
	// The project ID where service accounts are created.
	Project pulumi.StringInput
	// The action to be performed in the default service accounts on the resource destroy.
	// Valid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	// If set to REVERT it attempts to restore all default SAs but the DEPRIVILEGE action.
	// If set to REVERT_AND_IGNORE_FAILURE it is the same behavior as REVERT but ignores errors returned by the API.
	RestorePolicy pulumi.StringPtrInput
}

func (DefaultServiceAccountsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultServiceAccountsArgs)(nil)).Elem()
}

type DefaultServiceAccountsInput interface {
	pulumi.Input

	ToDefaultServiceAccountsOutput() DefaultServiceAccountsOutput
	ToDefaultServiceAccountsOutputWithContext(ctx context.Context) DefaultServiceAccountsOutput
}

func (*DefaultServiceAccounts) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultServiceAccounts)(nil))
}

func (i *DefaultServiceAccounts) ToDefaultServiceAccountsOutput() DefaultServiceAccountsOutput {
	return i.ToDefaultServiceAccountsOutputWithContext(context.Background())
}

func (i *DefaultServiceAccounts) ToDefaultServiceAccountsOutputWithContext(ctx context.Context) DefaultServiceAccountsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultServiceAccountsOutput)
}

type DefaultServiceAccountsOutput struct {
	*pulumi.OutputState
}

func (DefaultServiceAccountsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultServiceAccounts)(nil))
}

func (o DefaultServiceAccountsOutput) ToDefaultServiceAccountsOutput() DefaultServiceAccountsOutput {
	return o
}

func (o DefaultServiceAccountsOutput) ToDefaultServiceAccountsOutputWithContext(ctx context.Context) DefaultServiceAccountsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DefaultServiceAccountsOutput{})
}
