// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a folder-level logging sink. For more information see
// [the official documentation](https://cloud.google.com/logging/docs/) and
// [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
//
// Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
// granted to the credentials used with this provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/logging"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storage.NewBucket(ctx, "log-bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = organizations.NewFolder(ctx, "my-folder", &organizations.FolderArgs{
// 			DisplayName: pulumi.String("My folder"),
// 			Parent:      pulumi.String("organizations/123456"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logging.NewFolderSink(ctx, "my-sink", &logging.FolderSinkArgs{
// 			Description: pulumi.String("some explaination on what this is"),
// 			Folder:      my_folder.Name,
// 			Destination: log_bucket.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "storage.googleapis.com/", name), nil
// 			}).(pulumi.StringOutput),
// 			Filter: pulumi.String("resource.type = gce_instance AND severity >= WARNING"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMBinding(ctx, "log-writer", &projects.IAMBindingArgs{
// 			Role: pulumi.String("roles/storage.objectCreator"),
// 			Members: pulumi.StringArray{
// 				my_sink.WriterIdentity,
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
// Folder-level logging sinks can be imported using this format
//
// ```sh
//  $ pulumi import gcp:logging/folderSink:FolderSink my_sink folders/{{folder_id}}/sinks/{{name}}
// ```
type FolderSink struct {
	pulumi.CustomResourceState

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions FolderSinkBigqueryOptionsOutput `pulumi:"bigqueryOptions"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
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
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions FolderSinkExclusionArrayOutput `pulumi:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren pulumi.BoolPtrOutput `pulumi:"includeChildren"`
	// The name of the logging sink.
	Name pulumi.StringOutput `pulumi:"name"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringOutput `pulumi:"writerIdentity"`
}

// NewFolderSink registers a new resource with the given unique name, arguments, and options.
func NewFolderSink(ctx *pulumi.Context,
	name string, args *FolderSinkArgs, opts ...pulumi.ResourceOption) (*FolderSink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	var resource FolderSink
	err := ctx.RegisterResource("gcp:logging/folderSink:FolderSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderSink gets an existing FolderSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderSinkState, opts ...pulumi.ResourceOption) (*FolderSink, error) {
	var resource FolderSink
	err := ctx.ReadResource("gcp:logging/folderSink:FolderSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderSink resources.
type folderSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *FolderSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description *string `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
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
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions []FolderSinkExclusion `pulumi:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder *string `pulumi:"folder"`
	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren *bool `pulumi:"includeChildren"`
	// The name of the logging sink.
	Name *string `pulumi:"name"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity *string `pulumi:"writerIdentity"`
}

type FolderSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions FolderSinkBigqueryOptionsPtrInput
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
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
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled pulumi.BoolPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions FolderSinkExclusionArrayInput
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder pulumi.StringPtrInput
	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren pulumi.BoolPtrInput
	// The name of the logging sink.
	Name pulumi.StringPtrInput
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringPtrInput
}

func (FolderSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderSinkState)(nil)).Elem()
}

type folderSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *FolderSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description *string `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
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
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions []FolderSinkExclusion `pulumi:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder string `pulumi:"folder"`
	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren *bool `pulumi:"includeChildren"`
	// The name of the logging sink.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a FolderSink resource.
type FolderSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions FolderSinkBigqueryOptionsPtrInput
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
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
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled pulumi.BoolPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions FolderSinkExclusionArrayInput
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder pulumi.StringInput
	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren pulumi.BoolPtrInput
	// The name of the logging sink.
	Name pulumi.StringPtrInput
}

func (FolderSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderSinkArgs)(nil)).Elem()
}

type FolderSinkInput interface {
	pulumi.Input

	ToFolderSinkOutput() FolderSinkOutput
	ToFolderSinkOutputWithContext(ctx context.Context) FolderSinkOutput
}

func (*FolderSink) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderSink)(nil))
}

func (i *FolderSink) ToFolderSinkOutput() FolderSinkOutput {
	return i.ToFolderSinkOutputWithContext(context.Background())
}

func (i *FolderSink) ToFolderSinkOutputWithContext(ctx context.Context) FolderSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderSinkOutput)
}

type FolderSinkOutput struct {
	*pulumi.OutputState
}

func (FolderSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderSink)(nil))
}

func (o FolderSinkOutput) ToFolderSinkOutput() FolderSinkOutput {
	return o
}

func (o FolderSinkOutput) ToFolderSinkOutputWithContext(ctx context.Context) FolderSinkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FolderSinkOutput{})
}
