// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package healthcare

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store_iam_policy.html.markdown.
type DicomStoreIamPolicy struct {
	pulumi.CustomResourceState

	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId pulumi.StringOutput `pulumi:"dicomStoreId"`
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewDicomStoreIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDicomStoreIamPolicy(ctx *pulumi.Context,
	name string, args *DicomStoreIamPolicyArgs, opts ...pulumi.ResourceOption) (*DicomStoreIamPolicy, error) {
	if args == nil || args.DicomStoreId == nil {
		return nil, errors.New("missing required argument 'DicomStoreId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &DicomStoreIamPolicyArgs{}
	}
	var resource DicomStoreIamPolicy
	err := ctx.RegisterResource("gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDicomStoreIamPolicy gets an existing DicomStoreIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDicomStoreIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DicomStoreIamPolicyState, opts ...pulumi.ResourceOption) (*DicomStoreIamPolicy, error) {
	var resource DicomStoreIamPolicy
	err := ctx.ReadResource("gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DicomStoreIamPolicy resources.
type dicomStoreIamPolicyState struct {
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId *string `pulumi:"dicomStoreId"`
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type DicomStoreIamPolicyState struct {
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId pulumi.StringPtrInput
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (DicomStoreIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreIamPolicyState)(nil)).Elem()
}

type dicomStoreIamPolicyArgs struct {
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId string `pulumi:"dicomStoreId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a DicomStoreIamPolicy resource.
type DicomStoreIamPolicyArgs struct {
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (DicomStoreIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreIamPolicyArgs)(nil)).Elem()
}

