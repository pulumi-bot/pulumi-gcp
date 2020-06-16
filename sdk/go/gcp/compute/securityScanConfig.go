// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A ScanConfig resource contains the configurations to launch a scan.
//
// To get more information about ScanConfig, see:
//
// * [API documentation](https://cloud.google.com/security-scanner/docs/reference/rest/v1beta/projects.scanConfigs)
// * How-to Guides
//     * [Using Cloud Security Scanner](https://cloud.google.com/security-scanner/docs/scanning)
//
// > **Warning:** All arguments including `authentication.google_account.password` and `authentication.custom_account.password` will be stored in the raw
// state as plain-text.[Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets)
//
// {{% examples %}}
// ## Example Usage
// {{% example %}}
// ### Scan Config Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		scannerStaticIp, err := compute.NewAddress(ctx, "scannerStaticIp", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSecurityScanConfig(ctx, "scan-config", &compute.SecurityScanConfigArgs{
// 			DisplayName: pulumi.String("scan-config"),
// 			StartingUrls: pulumi.StringArray{
// 				scannerStaticIp.Address.ApplyT(func(address string) (string, error) {
// 					return fmt.Sprintf("%v%v", "http://", address), nil
// 				}).(pulumi.StringOutput),
// 			},
// 			TargetPlatforms: pulumi.StringArray{
// 				pulumi.String("COMPUTE"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// {{% /example %}}
// {{% /examples %}}
type SecurityScanConfig struct {
	pulumi.CustomResourceState

	// The authentication configuration.
	// If specified, service will use the authentication configuration during scanning.  Structure is documented below.
	Authentication SecurityScanConfigAuthenticationPtrOutput `pulumi:"authentication"`
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns pulumi.StringArrayOutput `pulumi:"blacklistPatterns"`
	// The user provider display name of the ScanConfig.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Controls export of scan configurations and results to Cloud Security Command Center.
	ExportToSecurityCommandCenter pulumi.StringPtrOutput `pulumi:"exportToSecurityCommandCenter"`
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	// Defaults to 15.
	MaxQps pulumi.IntPtrOutput `pulumi:"maxQps"`
	// A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The schedule of the ScanConfig  Structure is documented below.
	Schedule SecurityScanConfigSchedulePtrOutput `pulumi:"schedule"`
	// The starting URLs from which the scanner finds site pages.
	StartingUrls pulumi.StringArrayOutput `pulumi:"startingUrls"`
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms pulumi.StringArrayOutput `pulumi:"targetPlatforms"`
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
	// The authentication configuration.
	// If specified, service will use the authentication configuration during scanning.  Structure is documented below.
	Authentication *SecurityScanConfigAuthentication `pulumi:"authentication"`
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns []string `pulumi:"blacklistPatterns"`
	// The user provider display name of the ScanConfig.
	DisplayName *string `pulumi:"displayName"`
	// Controls export of scan configurations and results to Cloud Security Command Center.
	ExportToSecurityCommandCenter *string `pulumi:"exportToSecurityCommandCenter"`
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	// Defaults to 15.
	MaxQps *int `pulumi:"maxQps"`
	// A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The schedule of the ScanConfig  Structure is documented below.
	Schedule *SecurityScanConfigSchedule `pulumi:"schedule"`
	// The starting URLs from which the scanner finds site pages.
	StartingUrls []string `pulumi:"startingUrls"`
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms []string `pulumi:"targetPlatforms"`
	// Type of the user agents used for scanning
	UserAgent *string `pulumi:"userAgent"`
}

type SecurityScanConfigState struct {
	// The authentication configuration.
	// If specified, service will use the authentication configuration during scanning.  Structure is documented below.
	Authentication SecurityScanConfigAuthenticationPtrInput
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns pulumi.StringArrayInput
	// The user provider display name of the ScanConfig.
	DisplayName pulumi.StringPtrInput
	// Controls export of scan configurations and results to Cloud Security Command Center.
	ExportToSecurityCommandCenter pulumi.StringPtrInput
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	// Defaults to 15.
	MaxQps pulumi.IntPtrInput
	// A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The schedule of the ScanConfig  Structure is documented below.
	Schedule SecurityScanConfigSchedulePtrInput
	// The starting URLs from which the scanner finds site pages.
	StartingUrls pulumi.StringArrayInput
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms pulumi.StringArrayInput
	// Type of the user agents used for scanning
	UserAgent pulumi.StringPtrInput
}

func (SecurityScanConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityScanConfigState)(nil)).Elem()
}

type securityScanConfigArgs struct {
	// The authentication configuration.
	// If specified, service will use the authentication configuration during scanning.  Structure is documented below.
	Authentication *SecurityScanConfigAuthentication `pulumi:"authentication"`
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns []string `pulumi:"blacklistPatterns"`
	// The user provider display name of the ScanConfig.
	DisplayName string `pulumi:"displayName"`
	// Controls export of scan configurations and results to Cloud Security Command Center.
	ExportToSecurityCommandCenter *string `pulumi:"exportToSecurityCommandCenter"`
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	// Defaults to 15.
	MaxQps *int `pulumi:"maxQps"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The schedule of the ScanConfig  Structure is documented below.
	Schedule *SecurityScanConfigSchedule `pulumi:"schedule"`
	// The starting URLs from which the scanner finds site pages.
	StartingUrls []string `pulumi:"startingUrls"`
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms []string `pulumi:"targetPlatforms"`
	// Type of the user agents used for scanning
	UserAgent *string `pulumi:"userAgent"`
}

// The set of arguments for constructing a SecurityScanConfig resource.
type SecurityScanConfigArgs struct {
	// The authentication configuration.
	// If specified, service will use the authentication configuration during scanning.  Structure is documented below.
	Authentication SecurityScanConfigAuthenticationPtrInput
	// The blacklist URL patterns as described in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns pulumi.StringArrayInput
	// The user provider display name of the ScanConfig.
	DisplayName pulumi.StringInput
	// Controls export of scan configurations and results to Cloud Security Command Center.
	ExportToSecurityCommandCenter pulumi.StringPtrInput
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	// Defaults to 15.
	MaxQps pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The schedule of the ScanConfig  Structure is documented below.
	Schedule SecurityScanConfigSchedulePtrInput
	// The starting URLs from which the scanner finds site pages.
	StartingUrls pulumi.StringArrayInput
	// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms pulumi.StringArrayInput
	// Type of the user agents used for scanning
	UserAgent pulumi.StringPtrInput
}

func (SecurityScanConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityScanConfigArgs)(nil)).Elem()
}
