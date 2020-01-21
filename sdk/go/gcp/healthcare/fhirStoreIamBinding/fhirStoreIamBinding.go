// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package fhirStoreIamBinding

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/healthcare/FhirStoreIamBindingCondition"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_binding.html.markdown.
type FhirStoreIamBinding struct {
	pulumi.CustomResourceState

	Condition healthcareFhirStoreIamBindingCondition.FhirStoreIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringOutput `pulumi:"fhirStoreId"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewFhirStoreIamBinding registers a new resource with the given unique name, arguments, and options.
func NewFhirStoreIamBinding(ctx *pulumi.Context,
	name string, args *FhirStoreIamBindingArgs, opts ...pulumi.ResourceOption) (*FhirStoreIamBinding, error) {
	if args == nil || args.FhirStoreId == nil {
		return nil, errors.New("missing required argument 'FhirStoreId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &FhirStoreIamBindingArgs{}
	}
	var resource FhirStoreIamBinding
	err := ctx.RegisterResource("gcp:healthcare/fhirStoreIamBinding:FhirStoreIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFhirStoreIamBinding gets an existing FhirStoreIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirStoreIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FhirStoreIamBindingState, opts ...pulumi.ResourceOption) (*FhirStoreIamBinding, error) {
	var resource FhirStoreIamBinding
	err := ctx.ReadResource("gcp:healthcare/fhirStoreIamBinding:FhirStoreIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FhirStoreIamBinding resources.
type fhirStoreIamBindingState struct {
	Condition *healthcareFhirStoreIamBindingCondition.FhirStoreIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag *string `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId *string `pulumi:"fhirStoreId"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type FhirStoreIamBindingState struct {
	Condition healthcareFhirStoreIamBindingCondition.FhirStoreIamBindingConditionPtrInput
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringPtrInput
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (FhirStoreIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreIamBindingState)(nil)).Elem()
}

type fhirStoreIamBindingArgs struct {
	Condition *healthcareFhirStoreIamBindingCondition.FhirStoreIamBindingCondition `pulumi:"condition"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId string `pulumi:"fhirStoreId"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a FhirStoreIamBinding resource.
type FhirStoreIamBindingArgs struct {
	Condition healthcareFhirStoreIamBindingCondition.FhirStoreIamBindingConditionPtrInput
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (FhirStoreIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreIamBindingArgs)(nil)).Elem()
}

