// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A description of the conditions under which some aspect of your system is
// considered to be "unhealthy" and the ways to notify people or services
// about this state.
//
// To get more information about AlertPolicy, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/monitoring/alerts/)
//
// ## Example Usage
// ### Monitoring Alert Policy Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/monitoring"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := monitoring.NewAlertPolicy(ctx, "alertPolicy", &monitoring.AlertPolicyArgs{
// 			Combiner: pulumi.String("OR"),
// 			Conditions: monitoring.AlertPolicyConditionArray{
// 				&monitoring.AlertPolicyConditionArgs{
// 					ConditionThreshold: &monitoring.AlertPolicyConditionConditionThresholdArgs{
// 						Aggregations: monitoring.AlertPolicyConditionConditionThresholdAggregationArray{
// 							&monitoring.AlertPolicyConditionConditionThresholdAggregationArgs{
// 								AlignmentPeriod:  pulumi.String("60s"),
// 								PerSeriesAligner: pulumi.String("ALIGN_RATE"),
// 							},
// 						},
// 						Comparison: pulumi.String("COMPARISON_GT"),
// 						Duration:   pulumi.String("60s"),
// 						Filter:     pulumi.String("metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\""),
// 					},
// 					DisplayName: pulumi.String("test condition"),
// 				},
// 			},
// 			DisplayName: pulumi.String("My Alert Policy"),
// 			UserLabels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
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
// AlertPolicy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:monitoring/alertPolicy:AlertPolicy default {{name}}
// ```
type AlertPolicy struct {
	pulumi.CustomResourceState

	// How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	// Possible values are `AND`, `OR`, and `AND_WITH_MATCHING_RESOURCE`.
	Combiner pulumi.StringOutput `pulumi:"combiner"`
	// A list of conditions for the policy. The conditions are combined by
	// AND or OR according to the combiner field. If the combined conditions
	// evaluate to true, then an incident is created. A policy can have from
	// one to six conditions.
	// Structure is documented below.
	Conditions AlertPolicyConditionArrayOutput `pulumi:"conditions"`
	// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
	// ignored.
	CreationRecords AlertPolicyCreationRecordArrayOutput `pulumi:"creationRecords"`
	// A short name or phrase used to identify the
	// condition in dashboards, notifications, and
	// incidents. To avoid confusion, don't use the same
	// display name for multiple conditions in the same
	// policy.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.
	// Structure is documented below.
	Documentation AlertPolicyDocumentationPtrOutput `pulumi:"documentation"`
	// Whether or not the policy is enabled. The default is true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// -
	// The unique resource name for this condition.
	// Its syntax is:
	// projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
	// [CONDITION_ID] is assigned by Stackdriver Monitoring when
	// the condition is created as part of a new or updated alerting
	// policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// Identifies the notification channels to which notifications should be
	// sent when incidents are opened or closed or when new violations occur
	// on an already opened incident. Each element of this array corresponds
	// to the name field in each of the NotificationChannel objects that are
	// returned from the notificationChannels.list method. The syntax of the
	// entries in this field is
	// `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
	NotificationChannels pulumi.StringArrayOutput `pulumi:"notificationChannels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels pulumi.StringMapOutput `pulumi:"userLabels"`
}

// NewAlertPolicy registers a new resource with the given unique name, arguments, and options.
func NewAlertPolicy(ctx *pulumi.Context,
	name string, args *AlertPolicyArgs, opts ...pulumi.ResourceOption) (*AlertPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Combiner == nil {
		return nil, errors.New("invalid value for required argument 'Combiner'")
	}
	if args.Conditions == nil {
		return nil, errors.New("invalid value for required argument 'Conditions'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource AlertPolicy
	err := ctx.RegisterResource("gcp:monitoring/alertPolicy:AlertPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertPolicy gets an existing AlertPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertPolicyState, opts ...pulumi.ResourceOption) (*AlertPolicy, error) {
	var resource AlertPolicy
	err := ctx.ReadResource("gcp:monitoring/alertPolicy:AlertPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertPolicy resources.
type alertPolicyState struct {
	// How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	// Possible values are `AND`, `OR`, and `AND_WITH_MATCHING_RESOURCE`.
	Combiner *string `pulumi:"combiner"`
	// A list of conditions for the policy. The conditions are combined by
	// AND or OR according to the combiner field. If the combined conditions
	// evaluate to true, then an incident is created. A policy can have from
	// one to six conditions.
	// Structure is documented below.
	Conditions []AlertPolicyCondition `pulumi:"conditions"`
	// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
	// ignored.
	CreationRecords []AlertPolicyCreationRecord `pulumi:"creationRecords"`
	// A short name or phrase used to identify the
	// condition in dashboards, notifications, and
	// incidents. To avoid confusion, don't use the same
	// display name for multiple conditions in the same
	// policy.
	DisplayName *string `pulumi:"displayName"`
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.
	// Structure is documented below.
	Documentation *AlertPolicyDocumentation `pulumi:"documentation"`
	// Whether or not the policy is enabled. The default is true.
	Enabled *bool `pulumi:"enabled"`
	// -
	// The unique resource name for this condition.
	// Its syntax is:
	// projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
	// [CONDITION_ID] is assigned by Stackdriver Monitoring when
	// the condition is created as part of a new or updated alerting
	// policy.
	Name *string `pulumi:"name"`
	// Identifies the notification channels to which notifications should be
	// sent when incidents are opened or closed or when new violations occur
	// on an already opened incident. Each element of this array corresponds
	// to the name field in each of the NotificationChannel objects that are
	// returned from the notificationChannels.list method. The syntax of the
	// entries in this field is
	// `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
	NotificationChannels []string `pulumi:"notificationChannels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
}

type AlertPolicyState struct {
	// How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	// Possible values are `AND`, `OR`, and `AND_WITH_MATCHING_RESOURCE`.
	Combiner pulumi.StringPtrInput
	// A list of conditions for the policy. The conditions are combined by
	// AND or OR according to the combiner field. If the combined conditions
	// evaluate to true, then an incident is created. A policy can have from
	// one to six conditions.
	// Structure is documented below.
	Conditions AlertPolicyConditionArrayInput
	// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
	// ignored.
	CreationRecords AlertPolicyCreationRecordArrayInput
	// A short name or phrase used to identify the
	// condition in dashboards, notifications, and
	// incidents. To avoid confusion, don't use the same
	// display name for multiple conditions in the same
	// policy.
	DisplayName pulumi.StringPtrInput
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.
	// Structure is documented below.
	Documentation AlertPolicyDocumentationPtrInput
	// Whether or not the policy is enabled. The default is true.
	Enabled pulumi.BoolPtrInput
	// -
	// The unique resource name for this condition.
	// Its syntax is:
	// projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
	// [CONDITION_ID] is assigned by Stackdriver Monitoring when
	// the condition is created as part of a new or updated alerting
	// policy.
	Name pulumi.StringPtrInput
	// Identifies the notification channels to which notifications should be
	// sent when incidents are opened or closed or when new violations occur
	// on an already opened incident. Each element of this array corresponds
	// to the name field in each of the NotificationChannel objects that are
	// returned from the notificationChannels.list method. The syntax of the
	// entries in this field is
	// `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
	NotificationChannels pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels pulumi.StringMapInput
}

func (AlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertPolicyState)(nil)).Elem()
}

type alertPolicyArgs struct {
	// How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	// Possible values are `AND`, `OR`, and `AND_WITH_MATCHING_RESOURCE`.
	Combiner string `pulumi:"combiner"`
	// A list of conditions for the policy. The conditions are combined by
	// AND or OR according to the combiner field. If the combined conditions
	// evaluate to true, then an incident is created. A policy can have from
	// one to six conditions.
	// Structure is documented below.
	Conditions []AlertPolicyCondition `pulumi:"conditions"`
	// A short name or phrase used to identify the
	// condition in dashboards, notifications, and
	// incidents. To avoid confusion, don't use the same
	// display name for multiple conditions in the same
	// policy.
	DisplayName string `pulumi:"displayName"`
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.
	// Structure is documented below.
	Documentation *AlertPolicyDocumentation `pulumi:"documentation"`
	// Whether or not the policy is enabled. The default is true.
	Enabled *bool `pulumi:"enabled"`
	// Identifies the notification channels to which notifications should be
	// sent when incidents are opened or closed or when new violations occur
	// on an already opened incident. Each element of this array corresponds
	// to the name field in each of the NotificationChannel objects that are
	// returned from the notificationChannels.list method. The syntax of the
	// entries in this field is
	// `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
	NotificationChannels []string `pulumi:"notificationChannels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
}

// The set of arguments for constructing a AlertPolicy resource.
type AlertPolicyArgs struct {
	// How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	// Possible values are `AND`, `OR`, and `AND_WITH_MATCHING_RESOURCE`.
	Combiner pulumi.StringInput
	// A list of conditions for the policy. The conditions are combined by
	// AND or OR according to the combiner field. If the combined conditions
	// evaluate to true, then an incident is created. A policy can have from
	// one to six conditions.
	// Structure is documented below.
	Conditions AlertPolicyConditionArrayInput
	// A short name or phrase used to identify the
	// condition in dashboards, notifications, and
	// incidents. To avoid confusion, don't use the same
	// display name for multiple conditions in the same
	// policy.
	DisplayName pulumi.StringInput
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.
	// Structure is documented below.
	Documentation AlertPolicyDocumentationPtrInput
	// Whether or not the policy is enabled. The default is true.
	Enabled pulumi.BoolPtrInput
	// Identifies the notification channels to which notifications should be
	// sent when incidents are opened or closed or when new violations occur
	// on an already opened incident. Each element of this array corresponds
	// to the name field in each of the NotificationChannel objects that are
	// returned from the notificationChannels.list method. The syntax of the
	// entries in this field is
	// `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
	NotificationChannels pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels pulumi.StringMapInput
}

func (AlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertPolicyArgs)(nil)).Elem()
}

type AlertPolicyInput interface {
	pulumi.Input

	ToAlertPolicyOutput() AlertPolicyOutput
	ToAlertPolicyOutputWithContext(ctx context.Context) AlertPolicyOutput
}

func (*AlertPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertPolicy)(nil))
}

func (i *AlertPolicy) ToAlertPolicyOutput() AlertPolicyOutput {
	return i.ToAlertPolicyOutputWithContext(context.Background())
}

func (i *AlertPolicy) ToAlertPolicyOutputWithContext(ctx context.Context) AlertPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertPolicyOutput)
}

type AlertPolicyOutput struct {
	*pulumi.OutputState
}

func (AlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertPolicy)(nil))
}

func (o AlertPolicyOutput) ToAlertPolicyOutput() AlertPolicyOutput {
	return o
}

func (o AlertPolicyOutput) ToAlertPolicyOutputWithContext(ctx context.Context) AlertPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AlertPolicyOutput{})
}
