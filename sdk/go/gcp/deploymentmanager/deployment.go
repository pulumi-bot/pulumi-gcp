// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package deploymentmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A collection of resources that are deployed and managed together using
// a configuration file
//
// > **Warning:** This resource is intended only to manage a Deployment resource,
// and attempts to manage the Deployment's resources in the provider as well
// will likely result in errors or unexpected behavior as the two tools
// fight over ownership. We strongly discourage doing so unless you are an
// experienced user of both tools.
//
// In addition, due to limitations of the API, the provider will treat
// deployments in preview as recreate-only for any update operation other
// than actually deploying an in-preview deployment (i.e. `preview=true` to
// `preview=false`).
//
// ## Example Usage
//
// ## Import
//
// Deployment can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:deploymentmanager/deployment:Deployment default projects/{{project}}/deployments/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:deploymentmanager/deployment:Deployment default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:deploymentmanager/deployment:Deployment default {{name}}
// ```
type Deployment struct {
	pulumi.CustomResourceState

	// Set the policy to use for creating new resources. Only used on
	// create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
	// `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
	// the deployment will fail. Note that updating this field does not
	// actually affect the deployment, just how it is updated.
	// Default value is `CREATE_OR_ACQUIRE`.
	// Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
	CreatePolicy pulumi.StringPtrOutput `pulumi:"createPolicy"`
	// Set the policy to use for deleting new resources on update/delete.
	// Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
	// resource is deleted after removal from Deployment Manager. If
	// `ABANDON`, the resource is only removed from Deployment Manager
	// and is not actually deleted. Note that updating this field does not
	// actually change the deployment, just how it is updated.
	// Default value is `DELETE`.
	// Possible values are `ABANDON` and `DELETE`.
	DeletePolicy pulumi.StringPtrOutput `pulumi:"deletePolicy"`
	// Unique identifier for deployment. Output only.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// Optional user-provided description of deployment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Key-value pairs to apply to this labels.
	// Structure is documented below.
	Labels DeploymentLabelArrayOutput `pulumi:"labels"`
	// Output only. URL of the manifest representing the last manifest that was successfully deployed.
	Manifest pulumi.StringOutput `pulumi:"manifest"`
	// The name of the template to import, as declared in the YAML
	// configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// If set to true, a deployment is created with "shell" resources
	// that are not actually instantiated. This allows you to preview a
	// deployment. It can be updated to false to actually deploy
	// with real resources.
	// ~>**NOTE:** Deployment Manager does not allow update
	// of a deployment in preview (unless updating to preview=false). Thus,
	// the provider will force-recreate deployments if either preview is updated
	// to true or if other fields are updated while preview is true.
	Preview pulumi.BoolPtrOutput `pulumi:"preview"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. Server defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Parameters that define your deployment, including the deployment
	// configuration and relevant templates.
	// Structure is documented below.
	Target DeploymentTargetOutput `pulumi:"target"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	var resource Deployment
	err := ctx.RegisterResource("gcp:deploymentmanager/deployment:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("gcp:deploymentmanager/deployment:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
	// Set the policy to use for creating new resources. Only used on
	// create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
	// `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
	// the deployment will fail. Note that updating this field does not
	// actually affect the deployment, just how it is updated.
	// Default value is `CREATE_OR_ACQUIRE`.
	// Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
	CreatePolicy *string `pulumi:"createPolicy"`
	// Set the policy to use for deleting new resources on update/delete.
	// Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
	// resource is deleted after removal from Deployment Manager. If
	// `ABANDON`, the resource is only removed from Deployment Manager
	// and is not actually deleted. Note that updating this field does not
	// actually change the deployment, just how it is updated.
	// Default value is `DELETE`.
	// Possible values are `ABANDON` and `DELETE`.
	DeletePolicy *string `pulumi:"deletePolicy"`
	// Unique identifier for deployment. Output only.
	DeploymentId *string `pulumi:"deploymentId"`
	// Optional user-provided description of deployment.
	Description *string `pulumi:"description"`
	// Key-value pairs to apply to this labels.
	// Structure is documented below.
	Labels []DeploymentLabel `pulumi:"labels"`
	// Output only. URL of the manifest representing the last manifest that was successfully deployed.
	Manifest *string `pulumi:"manifest"`
	// The name of the template to import, as declared in the YAML
	// configuration.
	Name *string `pulumi:"name"`
	// If set to true, a deployment is created with "shell" resources
	// that are not actually instantiated. This allows you to preview a
	// deployment. It can be updated to false to actually deploy
	// with real resources.
	// ~>**NOTE:** Deployment Manager does not allow update
	// of a deployment in preview (unless updating to preview=false). Thus,
	// the provider will force-recreate deployments if either preview is updated
	// to true or if other fields are updated while preview is true.
	Preview *bool `pulumi:"preview"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Output only. Server defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// Parameters that define your deployment, including the deployment
	// configuration and relevant templates.
	// Structure is documented below.
	Target *DeploymentTarget `pulumi:"target"`
}

type DeploymentState struct {
	// Set the policy to use for creating new resources. Only used on
	// create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
	// `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
	// the deployment will fail. Note that updating this field does not
	// actually affect the deployment, just how it is updated.
	// Default value is `CREATE_OR_ACQUIRE`.
	// Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
	CreatePolicy pulumi.StringPtrInput
	// Set the policy to use for deleting new resources on update/delete.
	// Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
	// resource is deleted after removal from Deployment Manager. If
	// `ABANDON`, the resource is only removed from Deployment Manager
	// and is not actually deleted. Note that updating this field does not
	// actually change the deployment, just how it is updated.
	// Default value is `DELETE`.
	// Possible values are `ABANDON` and `DELETE`.
	DeletePolicy pulumi.StringPtrInput
	// Unique identifier for deployment. Output only.
	DeploymentId pulumi.StringPtrInput
	// Optional user-provided description of deployment.
	Description pulumi.StringPtrInput
	// Key-value pairs to apply to this labels.
	// Structure is documented below.
	Labels DeploymentLabelArrayInput
	// Output only. URL of the manifest representing the last manifest that was successfully deployed.
	Manifest pulumi.StringPtrInput
	// The name of the template to import, as declared in the YAML
	// configuration.
	Name pulumi.StringPtrInput
	// If set to true, a deployment is created with "shell" resources
	// that are not actually instantiated. This allows you to preview a
	// deployment. It can be updated to false to actually deploy
	// with real resources.
	// ~>**NOTE:** Deployment Manager does not allow update
	// of a deployment in preview (unless updating to preview=false). Thus,
	// the provider will force-recreate deployments if either preview is updated
	// to true or if other fields are updated while preview is true.
	Preview pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Output only. Server defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// Parameters that define your deployment, including the deployment
	// configuration and relevant templates.
	// Structure is documented below.
	Target DeploymentTargetPtrInput
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	// Set the policy to use for creating new resources. Only used on
	// create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
	// `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
	// the deployment will fail. Note that updating this field does not
	// actually affect the deployment, just how it is updated.
	// Default value is `CREATE_OR_ACQUIRE`.
	// Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
	CreatePolicy *string `pulumi:"createPolicy"`
	// Set the policy to use for deleting new resources on update/delete.
	// Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
	// resource is deleted after removal from Deployment Manager. If
	// `ABANDON`, the resource is only removed from Deployment Manager
	// and is not actually deleted. Note that updating this field does not
	// actually change the deployment, just how it is updated.
	// Default value is `DELETE`.
	// Possible values are `ABANDON` and `DELETE`.
	DeletePolicy *string `pulumi:"deletePolicy"`
	// Optional user-provided description of deployment.
	Description *string `pulumi:"description"`
	// Key-value pairs to apply to this labels.
	// Structure is documented below.
	Labels []DeploymentLabel `pulumi:"labels"`
	// The name of the template to import, as declared in the YAML
	// configuration.
	Name *string `pulumi:"name"`
	// If set to true, a deployment is created with "shell" resources
	// that are not actually instantiated. This allows you to preview a
	// deployment. It can be updated to false to actually deploy
	// with real resources.
	// ~>**NOTE:** Deployment Manager does not allow update
	// of a deployment in preview (unless updating to preview=false). Thus,
	// the provider will force-recreate deployments if either preview is updated
	// to true or if other fields are updated while preview is true.
	Preview *bool `pulumi:"preview"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Parameters that define your deployment, including the deployment
	// configuration and relevant templates.
	// Structure is documented below.
	Target DeploymentTarget `pulumi:"target"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	// Set the policy to use for creating new resources. Only used on
	// create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
	// `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
	// the deployment will fail. Note that updating this field does not
	// actually affect the deployment, just how it is updated.
	// Default value is `CREATE_OR_ACQUIRE`.
	// Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
	CreatePolicy pulumi.StringPtrInput
	// Set the policy to use for deleting new resources on update/delete.
	// Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
	// resource is deleted after removal from Deployment Manager. If
	// `ABANDON`, the resource is only removed from Deployment Manager
	// and is not actually deleted. Note that updating this field does not
	// actually change the deployment, just how it is updated.
	// Default value is `DELETE`.
	// Possible values are `ABANDON` and `DELETE`.
	DeletePolicy pulumi.StringPtrInput
	// Optional user-provided description of deployment.
	Description pulumi.StringPtrInput
	// Key-value pairs to apply to this labels.
	// Structure is documented below.
	Labels DeploymentLabelArrayInput
	// The name of the template to import, as declared in the YAML
	// configuration.
	Name pulumi.StringPtrInput
	// If set to true, a deployment is created with "shell" resources
	// that are not actually instantiated. This allows you to preview a
	// deployment. It can be updated to false to actually deploy
	// with real resources.
	// ~>**NOTE:** Deployment Manager does not allow update
	// of a deployment in preview (unless updating to preview=false). Thus,
	// the provider will force-recreate deployments if either preview is updated
	// to true or if other fields are updated while preview is true.
	Preview pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Parameters that define your deployment, including the deployment
	// configuration and relevant templates.
	// Structure is documented below.
	Target DeploymentTargetInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}

type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput
}

func (*Deployment) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil))
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

type DeploymentOutput struct {
	*pulumi.OutputState
}

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil))
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DeploymentOutput{})
}
