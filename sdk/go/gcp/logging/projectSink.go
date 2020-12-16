// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a project-level logging sink. For more information see
// [the official documentation](https://cloud.google.com/logging/docs/),
// [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs)
// and
// [API](https://cloud.google.com/logging/docs/reference/v2/rest/).
//
// > **Note:** You must have [granted the "Logs Configuration Writer"](https://cloud.google.com/logging/docs/access-control) IAM role (`roles/logging.configWriter`) to the credentials used with this provider.
//
// > **Note** You must [enable the Cloud Resource Manager API](https://console.cloud.google.com/apis/library/cloudresourcemanager.googleapis.com)
//
// ## Import
//
// Project-level logging sinks can be imported using their URI, e.g.
//
// ```sh
//  $ pulumi import gcp:logging/projectSink:ProjectSink my_sink projects/my-project/sinks/my-sink
// ```
type ProjectSink struct {
	pulumi.CustomResourceState

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions ProjectSinkBigqueryOptionsOutput `pulumi:"bigqueryOptions"`
	// A description of this exclusion.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions ProjectSinkExclusionArrayOutput `pulumi:"exclusions"`
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity pulumi.BoolPtrOutput `pulumi:"uniqueWriterIdentity"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringOutput `pulumi:"writerIdentity"`
}

// NewProjectSink registers a new resource with the given unique name, arguments, and options.
func NewProjectSink(ctx *pulumi.Context,
	name string, args *ProjectSinkArgs, opts ...pulumi.ResourceOption) (*ProjectSink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	var resource ProjectSink
	err := ctx.RegisterResource("gcp:logging/projectSink:ProjectSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectSink gets an existing ProjectSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectSinkState, opts ...pulumi.ResourceOption) (*ProjectSink, error) {
	var resource ProjectSink
	err := ctx.ReadResource("gcp:logging/projectSink:ProjectSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectSink resources.
type projectSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *ProjectSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// A description of this exclusion.
	Description *string `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination *string `pulumi:"destination"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []ProjectSinkExclusion `pulumi:"exclusions"`
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `pulumi:"name"`
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project *string `pulumi:"project"`
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity *bool `pulumi:"uniqueWriterIdentity"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity *string `pulumi:"writerIdentity"`
}

type ProjectSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions ProjectSinkBigqueryOptionsPtrInput
	// A description of this exclusion.
	Description pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringPtrInput
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled pulumi.BoolPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions ProjectSinkExclusionArrayInput
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringPtrInput
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringPtrInput
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity pulumi.BoolPtrInput
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringPtrInput
}

func (ProjectSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectSinkState)(nil)).Elem()
}

type projectSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *ProjectSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// A description of this exclusion.
	Description *string `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination string `pulumi:"destination"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []ProjectSinkExclusion `pulumi:"exclusions"`
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `pulumi:"name"`
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project *string `pulumi:"project"`
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity *bool `pulumi:"uniqueWriterIdentity"`
}

// The set of arguments for constructing a ProjectSink resource.
type ProjectSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions ProjectSinkBigqueryOptionsPtrInput
	// A description of this exclusion.
	Description pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringInput
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled pulumi.BoolPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions ProjectSinkExclusionArrayInput
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringPtrInput
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringPtrInput
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity pulumi.BoolPtrInput
}

func (ProjectSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectSinkArgs)(nil)).Elem()
}

type ProjectSinkInput interface {
	pulumi.Input

	ToProjectSinkOutput() ProjectSinkOutput
	ToProjectSinkOutputWithContext(ctx context.Context) ProjectSinkOutput
}

func (*ProjectSink) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectSink)(nil))
}

func (i *ProjectSink) ToProjectSinkOutput() ProjectSinkOutput {
	return i.ToProjectSinkOutputWithContext(context.Background())
}

func (i *ProjectSink) ToProjectSinkOutputWithContext(ctx context.Context) ProjectSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectSinkOutput)
}

type ProjectSinkOutput struct {
	*pulumi.OutputState
}

func (ProjectSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectSink)(nil))
}

func (o ProjectSinkOutput) ToProjectSinkOutput() ProjectSinkOutput {
	return o
}

func (o ProjectSinkOutput) ToProjectSinkOutputWithContext(ctx context.Context) ProjectSinkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProjectSinkOutput{})
}
