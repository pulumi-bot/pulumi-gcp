// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package monitoring

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
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
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/monitoring_alert_policy.html.markdown.
type AlertPolicy struct {
	pulumi.CustomResourceState

	// How to combine the results of multiple conditions to determine if an incident should be opened.
	Combiner pulumi.StringOutput `pulumi:"combiner"`
	// A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the
	// combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions.
	Conditions AlertPolicyConditionArrayOutput `pulumi:"conditions"`
	// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
	// ignored.
	CreationRecord AlertPolicyCreationRecordOutput `pulumi:"creationRecord"`
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
	// don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
	// characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
	// don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
	// characters.
	Documentation AlertPolicyDocumentationPtrOutput `pulumi:"documentation"`
	// Whether or not the policy is enabled. The default is true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The unique resource name for this policy. Its syntax is: projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
	Name pulumi.StringOutput `pulumi:"name"`
	// Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when
	// new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of
	// the NotificationChannel objects that are returned from the notificationChannels.list method. The syntax of the entries
	// in this field is 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
	NotificationChannels pulumi.StringArrayOutput `pulumi:"notificationChannels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
	// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
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
	// How to combine the results of multiple conditions to determine if an incident should be opened.
	Combiner *string `pulumi:"combiner"`
	// A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the
	// combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions.
	Conditions []AlertPolicyCondition `pulumi:"conditions"`
	// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
	// ignored.
	CreationRecord *AlertPolicyCreationRecord `pulumi:"creationRecord"`
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
	// don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
	// characters.
	DisplayName *string `pulumi:"displayName"`
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
	// don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
	// characters.
	Documentation *AlertPolicyDocumentation `pulumi:"documentation"`
	// Whether or not the policy is enabled. The default is true.
	Enabled *bool `pulumi:"enabled"`
	// The unique resource name for this policy. Its syntax is: projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
	Name *string `pulumi:"name"`
	// Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when
	// new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of
	// the NotificationChannel objects that are returned from the notificationChannels.list method. The syntax of the entries
	// in this field is 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
	NotificationChannels []string `pulumi:"notificationChannels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
	// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
}

type AlertPolicyState struct {
	// How to combine the results of multiple conditions to determine if an incident should be opened.
	Combiner pulumi.StringPtrInput
	// A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the
	// combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions.
	Conditions AlertPolicyConditionArrayInput
	// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
	// ignored.
	CreationRecord AlertPolicyCreationRecordPtrInput
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
	// don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
	// characters.
	DisplayName pulumi.StringPtrInput
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
	// don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
	// characters.
	Documentation AlertPolicyDocumentationPtrInput
	// Whether or not the policy is enabled. The default is true.
	Enabled pulumi.BoolPtrInput
	// The unique resource name for this policy. Its syntax is: projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
	Name pulumi.StringPtrInput
	// Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when
	// new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of
	// the NotificationChannel objects that are returned from the notificationChannels.list method. The syntax of the entries
	// in this field is 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
	NotificationChannels pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
	// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels pulumi.StringMapInput
}

func (AlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertPolicyState)(nil)).Elem()
}

type alertPolicyArgs struct {
	// How to combine the results of multiple conditions to determine if an incident should be opened.
	Combiner string `pulumi:"combiner"`
	// A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the
	// combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions.
	Conditions []AlertPolicyCondition `pulumi:"conditions"`
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
	// don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
	// characters.
	DisplayName string `pulumi:"displayName"`
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
	// don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
	// characters.
	Documentation *AlertPolicyDocumentation `pulumi:"documentation"`
	// Whether or not the policy is enabled. The default is true.
	Enabled *bool `pulumi:"enabled"`
	// Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when
	// new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of
	// the NotificationChannel objects that are returned from the notificationChannels.list method. The syntax of the entries
	// in this field is 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
	NotificationChannels []string `pulumi:"notificationChannels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
	// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
}

// The set of arguments for constructing a AlertPolicy resource.
type AlertPolicyArgs struct {
	// How to combine the results of multiple conditions to determine if an incident should be opened.
	Combiner pulumi.StringInput
	// A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the
	// combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions.
	Conditions AlertPolicyConditionArrayInput
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
	// don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
	// characters.
	DisplayName pulumi.StringInput
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
	// don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
	// characters.
	Documentation AlertPolicyDocumentationPtrInput
	// Whether or not the policy is enabled. The default is true.
	Enabled pulumi.BoolPtrInput
	// Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when
	// new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of
	// the NotificationChannel objects that are returned from the notificationChannels.list method. The syntax of the entries
	// in this field is 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
	NotificationChannels pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
	// entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels pulumi.StringMapInput
}

func (AlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertPolicyArgs)(nil)).Elem()
}

