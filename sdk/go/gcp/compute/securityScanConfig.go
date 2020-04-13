// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A ScanConfig resource contains the configurations to launch a scan.
//
// To get more information about ScanConfig, see:
//
// * [API documentation](https://cloud.google.com/security-scanner/docs/reference/rest/v1beta/projects.scanConfigs)
// * How-to Guides
//     * [Using Cloud Security Scanner](https://cloud.google.com/security-scanner/docs/scanning)
type SecurityScanConfig struct {
	pulumi.CustomResourceState

	// -
	// (Optional)
	// The authentication configuration.
	// If specified, service will use the authentication configuration during scanning.  Structure is documented below.
	Authentication SecurityScanConfigAuthenticationPtrOutput `pulumi:"authentication"`
	// -
	// (Optional)
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns pulumi.StringArrayOutput `pulumi:"blacklistPatterns"`
	// -
	// (Required)
	// The user provider display name of the ScanConfig.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// -
	// (Optional)
	// Controls export of scan configurations and results to Cloud Security Command Center.
	ExportToSecurityCommandCenter pulumi.StringPtrOutput `pulumi:"exportToSecurityCommandCenter"`
	// -
	// (Optional)
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	// Defaults to 15.
	MaxQps pulumi.IntPtrOutput `pulumi:"maxQps"`
	// A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// -
	// (Optional)
	// The schedule of the ScanConfig  Structure is documented below.
	Schedule SecurityScanConfigSchedulePtrOutput `pulumi:"schedule"`
	// -
	// (Required)
	// The starting URLs from which the scanner finds site pages.
	StartingUrls pulumi.StringArrayOutput `pulumi:"startingUrls"`
	// -
	// (Optional)
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms pulumi.StringArrayOutput `pulumi:"targetPlatforms"`
	// -
	// (Optional)
	// Type of the user agents used for scanning
	UserAgent pulumi.StringPtrOutput `pulumi:"userAgent"`
}

// NewSecurityScanConfig registers a new resource with the given unique name, arguments, and options.
func NewSecurityScanConfig(ctx *pulumi.Context,
	name string, args *SecurityScanConfigArgs, opts ...pulumi.ResourceOption) (*SecurityScanConfig, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.StartingUrls == nil {
		return nil, errors.New("missing required argument 'StartingUrls'")
	}
	if args == nil {
		args = &SecurityScanConfigArgs{}
	}
	var resource SecurityScanConfig
	err := ctx.RegisterResource("gcp:compute/securityScanConfig:SecurityScanConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityScanConfig gets an existing SecurityScanConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityScanConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityScanConfigState, opts ...pulumi.ResourceOption) (*SecurityScanConfig, error) {
	var resource SecurityScanConfig
	err := ctx.ReadResource("gcp:compute/securityScanConfig:SecurityScanConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityScanConfig resources.
type securityScanConfigState struct {
	// -
	// (Optional)
	// The authentication configuration.
	// If specified, service will use the authentication configuration during scanning.  Structure is documented below.
	Authentication *SecurityScanConfigAuthentication `pulumi:"authentication"`
	// -
	// (Optional)
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns []string `pulumi:"blacklistPatterns"`
	// -
	// (Required)
	// The user provider display name of the ScanConfig.
	DisplayName *string `pulumi:"displayName"`
	// -
	// (Optional)
	// Controls export of scan configurations and results to Cloud Security Command Center.
	ExportToSecurityCommandCenter *string `pulumi:"exportToSecurityCommandCenter"`
	// -
	// (Optional)
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	// Defaults to 15.
	MaxQps *int `pulumi:"maxQps"`
	// A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Optional)
	// The schedule of the ScanConfig  Structure is documented below.
	Schedule *SecurityScanConfigSchedule `pulumi:"schedule"`
	// -
	// (Required)
	// The starting URLs from which the scanner finds site pages.
	StartingUrls []string `pulumi:"startingUrls"`
	// -
	// (Optional)
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms []string `pulumi:"targetPlatforms"`
	// -
	// (Optional)
	// Type of the user agents used for scanning
	UserAgent *string `pulumi:"userAgent"`
}

type SecurityScanConfigState struct {
	// -
	// (Optional)
	// The authentication configuration.
	// If specified, service will use the authentication configuration during scanning.  Structure is documented below.
	Authentication SecurityScanConfigAuthenticationPtrInput
	// -
	// (Optional)
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns pulumi.StringArrayInput
	// -
	// (Required)
	// The user provider display name of the ScanConfig.
	DisplayName pulumi.StringPtrInput
	// -
	// (Optional)
	// Controls export of scan configurations and results to Cloud Security Command Center.
	ExportToSecurityCommandCenter pulumi.StringPtrInput
	// -
	// (Optional)
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	// Defaults to 15.
	MaxQps pulumi.IntPtrInput
	// A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Optional)
	// The schedule of the ScanConfig  Structure is documented below.
	Schedule SecurityScanConfigSchedulePtrInput
	// -
	// (Required)
	// The starting URLs from which the scanner finds site pages.
	StartingUrls pulumi.StringArrayInput
	// -
	// (Optional)
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms pulumi.StringArrayInput
	// -
	// (Optional)
	// Type of the user agents used for scanning
	UserAgent pulumi.StringPtrInput
}

func (SecurityScanConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityScanConfigState)(nil)).Elem()
}

type securityScanConfigArgs struct {
	// -
	// (Optional)
	// The authentication configuration.
	// If specified, service will use the authentication configuration during scanning.  Structure is documented below.
	Authentication *SecurityScanConfigAuthentication `pulumi:"authentication"`
	// -
	// (Optional)
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns []string `pulumi:"blacklistPatterns"`
	// -
	// (Required)
	// The user provider display name of the ScanConfig.
	DisplayName string `pulumi:"displayName"`
	// -
	// (Optional)
	// Controls export of scan configurations and results to Cloud Security Command Center.
	ExportToSecurityCommandCenter *string `pulumi:"exportToSecurityCommandCenter"`
	// -
	// (Optional)
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	// Defaults to 15.
	MaxQps *int `pulumi:"maxQps"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Optional)
	// The schedule of the ScanConfig  Structure is documented below.
	Schedule *SecurityScanConfigSchedule `pulumi:"schedule"`
	// -
	// (Required)
	// The starting URLs from which the scanner finds site pages.
	StartingUrls []string `pulumi:"startingUrls"`
	// -
	// (Optional)
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms []string `pulumi:"targetPlatforms"`
	// -
	// (Optional)
	// Type of the user agents used for scanning
	UserAgent *string `pulumi:"userAgent"`
}

// The set of arguments for constructing a SecurityScanConfig resource.
type SecurityScanConfigArgs struct {
	// -
	// (Optional)
	// The authentication configuration.
	// If specified, service will use the authentication configuration during scanning.  Structure is documented below.
	Authentication SecurityScanConfigAuthenticationPtrInput
	// -
	// (Optional)
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns pulumi.StringArrayInput
	// -
	// (Required)
	// The user provider display name of the ScanConfig.
	DisplayName pulumi.StringInput
	// -
	// (Optional)
	// Controls export of scan configurations and results to Cloud Security Command Center.
	ExportToSecurityCommandCenter pulumi.StringPtrInput
	// -
	// (Optional)
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	// Defaults to 15.
	MaxQps pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Optional)
	// The schedule of the ScanConfig  Structure is documented below.
	Schedule SecurityScanConfigSchedulePtrInput
	// -
	// (Required)
	// The starting URLs from which the scanner finds site pages.
	StartingUrls pulumi.StringArrayInput
	// -
	// (Optional)
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms pulumi.StringArrayInput
	// -
	// (Optional)
	// Type of the user agents used for scanning
	UserAgent pulumi.StringPtrInput
}

func (SecurityScanConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityScanConfigArgs)(nil)).Elem()
}
