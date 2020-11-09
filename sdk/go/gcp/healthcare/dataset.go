// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Healthcare `Dataset` is a toplevel logical grouping of `dicomStores`, `fhirStores` and `hl7V2Stores`.
//
// To get more information about Dataset, see:
//
// * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets)
// * How-to Guides
//     * [Creating a dataset](https://cloud.google.com/healthcare/docs/how-tos/datasets)
//
// ## Example Usage
type Dataset struct {
	pulumi.CustomResourceState

	// The location for the Dataset.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for the Dataset.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The fully qualified name of this dataset
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil {
		args = &DatasetArgs{}
	}
	var resource Dataset
	err := ctx.RegisterResource("gcp:healthcare/dataset:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataset gets an existing Dataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("gcp:healthcare/dataset:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataset resources.
type datasetState struct {
	// The location for the Dataset.
	Location *string `pulumi:"location"`
	// The resource name for the Dataset.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The fully qualified name of this dataset
	SelfLink *string `pulumi:"selfLink"`
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	TimeZone *string `pulumi:"timeZone"`
}

type DatasetState struct {
	// The location for the Dataset.
	Location pulumi.StringPtrInput
	// The resource name for the Dataset.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The fully qualified name of this dataset
	SelfLink pulumi.StringPtrInput
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	TimeZone pulumi.StringPtrInput
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	// The location for the Dataset.
	Location string `pulumi:"location"`
	// The resource name for the Dataset.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	// The location for the Dataset.
	Location pulumi.StringInput
	// The resource name for the Dataset.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	TimeZone pulumi.StringPtrInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}

type DatasetInput interface {
	pulumi.Input

	ToDatasetOutput() DatasetOutput
	ToDatasetOutputWithContext(ctx context.Context) DatasetOutput
}

func (Dataset) ElementType() reflect.Type {
	return reflect.TypeOf((*Dataset)(nil)).Elem()
}

func (i Dataset) ToDatasetOutput() DatasetOutput {
	return i.ToDatasetOutputWithContext(context.Background())
}

func (i Dataset) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetOutput)
}

type DatasetOutput struct {
	*pulumi.OutputState
}

func (DatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetOutput)(nil)).Elem()
}

func (o DatasetOutput) ToDatasetOutput() DatasetOutput {
	return o
}

func (o DatasetOutput) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetOutput{})
}
