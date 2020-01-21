// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package autoscalingPolicy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/AutoscalingPolicyBasicAlgorithm"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/AutoscalingPolicyBasicAlgorithmYarnConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/AutoscalingPolicySecondaryWorkerConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/AutoscalingPolicyWorkerConfig"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_autoscaling_policy.html.markdown.
type AutoscalingPolicy struct {
	pulumi.CustomResourceState

	// Basic algorithm for autoscaling.
	BasicAlgorithm dataprocAutoscalingPolicyBasicAlgorithm.AutoscalingPolicyBasicAlgorithmPtrOutput `pulumi:"basicAlgorithm"`
	// The location where the autoscaling poicy should reside. The default value is 'global'.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The "resource name" of the autoscaling policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig dataprocAutoscalingPolicySecondaryWorkerConfig.AutoscalingPolicySecondaryWorkerConfigPtrOutput `pulumi:"secondaryWorkerConfig"`
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig dataprocAutoscalingPolicyWorkerConfig.AutoscalingPolicyWorkerConfigPtrOutput `pulumi:"workerConfig"`
}

// NewAutoscalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewAutoscalingPolicy(ctx *pulumi.Context,
	name string, args *AutoscalingPolicyArgs, opts ...pulumi.ResourceOption) (*AutoscalingPolicy, error) {
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil {
		args = &AutoscalingPolicyArgs{}
	}
	var resource AutoscalingPolicy
	err := ctx.RegisterResource("gcp:dataproc/autoscalingPolicy:AutoscalingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscalingPolicy gets an existing AutoscalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscalingPolicyState, opts ...pulumi.ResourceOption) (*AutoscalingPolicy, error) {
	var resource AutoscalingPolicy
	err := ctx.ReadResource("gcp:dataproc/autoscalingPolicy:AutoscalingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoscalingPolicy resources.
type autoscalingPolicyState struct {
	// Basic algorithm for autoscaling.
	BasicAlgorithm *dataprocAutoscalingPolicyBasicAlgorithm.AutoscalingPolicyBasicAlgorithm `pulumi:"basicAlgorithm"`
	// The location where the autoscaling poicy should reside. The default value is 'global'.
	Location *string `pulumi:"location"`
	// The "resource name" of the autoscaling policy.
	Name *string `pulumi:"name"`
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId *string `pulumi:"policyId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig *dataprocAutoscalingPolicySecondaryWorkerConfig.AutoscalingPolicySecondaryWorkerConfig `pulumi:"secondaryWorkerConfig"`
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig *dataprocAutoscalingPolicyWorkerConfig.AutoscalingPolicyWorkerConfig `pulumi:"workerConfig"`
}

type AutoscalingPolicyState struct {
	// Basic algorithm for autoscaling.
	BasicAlgorithm dataprocAutoscalingPolicyBasicAlgorithm.AutoscalingPolicyBasicAlgorithmPtrInput
	// The location where the autoscaling poicy should reside. The default value is 'global'.
	Location pulumi.StringPtrInput
	// The "resource name" of the autoscaling policy.
	Name pulumi.StringPtrInput
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig dataprocAutoscalingPolicySecondaryWorkerConfig.AutoscalingPolicySecondaryWorkerConfigPtrInput
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig dataprocAutoscalingPolicyWorkerConfig.AutoscalingPolicyWorkerConfigPtrInput
}

func (AutoscalingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalingPolicyState)(nil)).Elem()
}

type autoscalingPolicyArgs struct {
	// Basic algorithm for autoscaling.
	BasicAlgorithm *dataprocAutoscalingPolicyBasicAlgorithm.AutoscalingPolicyBasicAlgorithm `pulumi:"basicAlgorithm"`
	// The location where the autoscaling poicy should reside. The default value is 'global'.
	Location *string `pulumi:"location"`
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId string `pulumi:"policyId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig *dataprocAutoscalingPolicySecondaryWorkerConfig.AutoscalingPolicySecondaryWorkerConfig `pulumi:"secondaryWorkerConfig"`
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig *dataprocAutoscalingPolicyWorkerConfig.AutoscalingPolicyWorkerConfig `pulumi:"workerConfig"`
}

// The set of arguments for constructing a AutoscalingPolicy resource.
type AutoscalingPolicyArgs struct {
	// Basic algorithm for autoscaling.
	BasicAlgorithm dataprocAutoscalingPolicyBasicAlgorithm.AutoscalingPolicyBasicAlgorithmPtrInput
	// The location where the autoscaling poicy should reside. The default value is 'global'.
	Location pulumi.StringPtrInput
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
	// begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	PolicyId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig dataprocAutoscalingPolicySecondaryWorkerConfig.AutoscalingPolicySecondaryWorkerConfigPtrInput
	// Describes how the autoscaler will operate for primary workers.
	WorkerConfig dataprocAutoscalingPolicyWorkerConfig.AutoscalingPolicyWorkerConfigPtrInput
}

func (AutoscalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalingPolicyArgs)(nil)).Elem()
}

