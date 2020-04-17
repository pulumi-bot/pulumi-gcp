// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A description of the conditions under which some aspect of your system is
// considered to be "unhealthy" and the ways to notify people or services
// about this state.
//
//
// To get more information about AlertPolicy, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/monitoring/alerts/)
type AlertPolicy struct {
	pulumi.CustomResourceState

	// -
	// (Required)
	// How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	Combiner pulumi.StringOutput `pulumi:"combiner"`
	// -
	// (Required)
	// A list of conditions for the policy. The conditions are combined by
	// AND or OR according to the combiner field. If the combined conditions
	// evaluate to true, then an incident is created. A policy can have from
	// one to six conditions.  Structure is documented below.
	Conditions AlertPolicyConditionArrayOutput `pulumi:"conditions"`
	// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
	// ignored.
	CreationRecord AlertPolicyCreationRecordOutput `pulumi:"creationRecord"`
	// -
	// (Required)
	// A short name or phrase used to identify the policy in
	// dashboards, notifications, and incidents. To avoid confusion, don't use
	// the same display name for multiple policies in the same project. The
	// name is limited to 512 Unicode characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// -
	// (Optional)
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.  Structure is documented below.
	Documentation AlertPolicyDocumentationPtrOutput `pulumi:"documentation"`
	// -
	// (Optional)
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
	// -
	// (Optional)
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
	// -
	// (Optional)
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
	if args == nil || args.Combiner == nil {
		return nil, errors.New("missing required argument 'Combiner'")
	}
	if args == nil || args.Conditions == nil {
		return nil, errors.New("missing required argument 'Conditions'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil {
		args = &AlertPolicyArgs{}
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
	// -
	// (Required)
	// How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	Combiner *string `pulumi:"combiner"`
	// -
	// (Required)
	// A list of conditions for the policy. The conditions are combined by
	// AND or OR according to the combiner field. If the combined conditions
	// evaluate to true, then an incident is created. A policy can have from
	// one to six conditions.  Structure is documented below.
	Conditions []AlertPolicyCondition `pulumi:"conditions"`
	// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
	// ignored.
	CreationRecord *AlertPolicyCreationRecord `pulumi:"creationRecord"`
	// -
	// (Required)
	// A short name or phrase used to identify the policy in
	// dashboards, notifications, and incidents. To avoid confusion, don't use
	// the same display name for multiple policies in the same project. The
	// name is limited to 512 Unicode characters.
	DisplayName *string `pulumi:"displayName"`
	// -
	// (Optional)
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.  Structure is documented below.
	Documentation *AlertPolicyDocumentation `pulumi:"documentation"`
	// -
	// (Optional)
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
	// -
	// (Optional)
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
	// -
	// (Optional)
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
}

type AlertPolicyState struct {
	// -
	// (Required)
	// How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	Combiner pulumi.StringPtrInput
	// -
	// (Required)
	// A list of conditions for the policy. The conditions are combined by
	// AND or OR according to the combiner field. If the combined conditions
	// evaluate to true, then an incident is created. A policy can have from
	// one to six conditions.  Structure is documented below.
	Conditions AlertPolicyConditionArrayInput
	// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
	// ignored.
	CreationRecord AlertPolicyCreationRecordPtrInput
	// -
	// (Required)
	// A short name or phrase used to identify the policy in
	// dashboards, notifications, and incidents. To avoid confusion, don't use
	// the same display name for multiple policies in the same project. The
	// name is limited to 512 Unicode characters.
	DisplayName pulumi.StringPtrInput
	// -
	// (Optional)
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.  Structure is documented below.
	Documentation AlertPolicyDocumentationPtrInput
	// -
	// (Optional)
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
	// -
	// (Optional)
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
	// -
	// (Optional)
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
	// -
	// (Required)
	// How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	Combiner string `pulumi:"combiner"`
	// -
	// (Required)
	// A list of conditions for the policy. The conditions are combined by
	// AND or OR according to the combiner field. If the combined conditions
	// evaluate to true, then an incident is created. A policy can have from
	// one to six conditions.  Structure is documented below.
	Conditions []AlertPolicyCondition `pulumi:"conditions"`
	// -
	// (Required)
	// A short name or phrase used to identify the policy in
	// dashboards, notifications, and incidents. To avoid confusion, don't use
	// the same display name for multiple policies in the same project. The
	// name is limited to 512 Unicode characters.
	DisplayName string `pulumi:"displayName"`
	// -
	// (Optional)
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.  Structure is documented below.
	Documentation *AlertPolicyDocumentation `pulumi:"documentation"`
	// -
	// (Optional)
	// Whether or not the policy is enabled. The default is true.
	Enabled *bool `pulumi:"enabled"`
	// -
	// (Optional)
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
	// -
	// (Optional)
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
}

// The set of arguments for constructing a AlertPolicy resource.
type AlertPolicyArgs struct {
	// -
	// (Required)
	// How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	Combiner pulumi.StringInput
	// -
	// (Required)
	// A list of conditions for the policy. The conditions are combined by
	// AND or OR according to the combiner field. If the combined conditions
	// evaluate to true, then an incident is created. A policy can have from
	// one to six conditions.  Structure is documented below.
	Conditions AlertPolicyConditionArrayInput
	// -
	// (Required)
	// A short name or phrase used to identify the policy in
	// dashboards, notifications, and incidents. To avoid confusion, don't use
	// the same display name for multiple policies in the same project. The
	// name is limited to 512 Unicode characters.
	DisplayName pulumi.StringInput
	// -
	// (Optional)
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.  Structure is documented below.
	Documentation AlertPolicyDocumentationPtrInput
	// -
	// (Optional)
	// Whether or not the policy is enabled. The default is true.
	Enabled pulumi.BoolPtrInput
	// -
	// (Optional)
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
	// -
	// (Optional)
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
