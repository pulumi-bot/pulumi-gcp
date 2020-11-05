// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a billing account logging sink. For more information see
// [the official documentation](https://cloud.google.com/logging/docs/) and
// [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
//
// > **Note** You must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
// [granted on the billing account](https://cloud.google.com/billing/reference/rest/v1/billingAccounts/getIamPolicy) to
// the credentials used with this provider. [IAM roles granted on a billing account](https://cloud.google.com/billing/docs/how-to/billing-access) are separate from the
// typical IAM roles granted on a project.
type BillingAccountSink struct {
	pulumi.CustomResourceState

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions BillingAccountSinkBigqueryOptionsOutput `pulumi:"bigqueryOptions"`
	// The billing account exported to the sink.
	BillingAccount pulumi.StringOutput `pulumi:"billingAccount"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
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
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions BillingAccountSinkExclusionArrayOutput `pulumi:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// The name of the logging sink.
	Name pulumi.StringOutput `pulumi:"name"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringOutput `pulumi:"writerIdentity"`
}

// NewBillingAccountSink registers a new resource with the given unique name, arguments, and options.
func NewBillingAccountSink(ctx *pulumi.Context,
	name string, args *BillingAccountSinkArgs, opts ...pulumi.ResourceOption) (*BillingAccountSink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.BillingAccount == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccount'")
	}
	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	var resource BillingAccountSink
	err := ctx.RegisterResource("gcp:logging/billingAccountSink:BillingAccountSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingAccountSink gets an existing BillingAccountSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingAccountSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingAccountSinkState, opts ...pulumi.ResourceOption) (*BillingAccountSink, error) {
	var resource BillingAccountSink
	err := ctx.ReadResource("gcp:logging/billingAccountSink:BillingAccountSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingAccountSink resources.
type billingAccountSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *BillingAccountSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// The billing account exported to the sink.
	BillingAccount *string `pulumi:"billingAccount"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
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
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions []BillingAccountSinkExclusion `pulumi:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// The name of the logging sink.
	Name *string `pulumi:"name"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity *string `pulumi:"writerIdentity"`
}

type BillingAccountSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions BillingAccountSinkBigqueryOptionsPtrInput
	// The billing account exported to the sink.
	BillingAccount pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
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
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions BillingAccountSinkExclusionArrayInput
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// The name of the logging sink.
	Name pulumi.StringPtrInput
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringPtrInput
}

func (BillingAccountSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountSinkState)(nil)).Elem()
}

type billingAccountSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *BillingAccountSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// The billing account exported to the sink.
	BillingAccount string `pulumi:"billingAccount"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
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
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions []BillingAccountSinkExclusion `pulumi:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// The name of the logging sink.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a BillingAccountSink resource.
type BillingAccountSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions BillingAccountSinkBigqueryOptionsPtrInput
	// The billing account exported to the sink.
	BillingAccount pulumi.StringInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
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
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions BillingAccountSinkExclusionArrayInput
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// The name of the logging sink.
	Name pulumi.StringPtrInput
}

func (BillingAccountSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountSinkArgs)(nil)).Elem()
}
