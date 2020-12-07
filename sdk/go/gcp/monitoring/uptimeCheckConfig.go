// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This message configures which resources and services to monitor for availability.
//
// To get more information about UptimeCheckConfig, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.uptimeCheckConfigs)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/monitoring/uptime-checks/)
//
// > **Warning:** All arguments including `http_check.auth_info.password` will be stored in the raw
// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
//
// ## Example Usage
// ### Uptime Check Config Http
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
// 		_, err := monitoring.NewUptimeCheckConfig(ctx, "http", &monitoring.UptimeCheckConfigArgs{
// 			ContentMatchers: monitoring.UptimeCheckConfigContentMatcherArray{
// 				&monitoring.UptimeCheckConfigContentMatcherArgs{
// 					Content: pulumi.String("example"),
// 				},
// 			},
// 			DisplayName: pulumi.String("http-uptime-check"),
// 			HttpCheck: &monitoring.UptimeCheckConfigHttpCheckArgs{
// 				Body:          pulumi.String("Zm9vJTI1M0RiYXI="),
// 				ContentType:   pulumi.String("URL_ENCODED"),
// 				Path:          pulumi.String("/some-path"),
// 				Port:          pulumi.Int(8010),
// 				RequestMethod: pulumi.String("POST"),
// 			},
// 			MonitoredResource: &monitoring.UptimeCheckConfigMonitoredResourceArgs{
// 				Labels: pulumi.StringMap{
// 					"host":      pulumi.String("192.168.1.1"),
// 					"projectId": pulumi.String("my-project-name"),
// 				},
// 				Type: pulumi.String("uptime_url"),
// 			},
// 			Timeout: pulumi.String("60s"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Uptime Check Config Https
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
// 		_, err := monitoring.NewUptimeCheckConfig(ctx, "https", &monitoring.UptimeCheckConfigArgs{
// 			ContentMatchers: monitoring.UptimeCheckConfigContentMatcherArray{
// 				&monitoring.UptimeCheckConfigContentMatcherArgs{
// 					Content: pulumi.String("example"),
// 				},
// 			},
// 			DisplayName: pulumi.String("https-uptime-check"),
// 			HttpCheck: &monitoring.UptimeCheckConfigHttpCheckArgs{
// 				Path:        pulumi.String("/some-path"),
// 				Port:        pulumi.Int(443),
// 				UseSsl:      pulumi.Bool(true),
// 				ValidateSsl: pulumi.Bool(true),
// 			},
// 			MonitoredResource: &monitoring.UptimeCheckConfigMonitoredResourceArgs{
// 				Labels: pulumi.StringMap{
// 					"host":      pulumi.String("192.168.1.1"),
// 					"projectId": pulumi.String("my-project-name"),
// 				},
// 				Type: pulumi.String("uptime_url"),
// 			},
// 			Timeout: pulumi.String("60s"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Uptime Check Tcp
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
// 		check, err := monitoring.NewGroup(ctx, "check", &monitoring.GroupArgs{
// 			DisplayName: pulumi.String("uptime-check-group"),
// 			Filter:      pulumi.String("resource.metadata.name=has_substring(\"foo\")"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = monitoring.NewUptimeCheckConfig(ctx, "tcpGroup", &monitoring.UptimeCheckConfigArgs{
// 			DisplayName: pulumi.String("tcp-uptime-check"),
// 			Timeout:     pulumi.String("60s"),
// 			TcpCheck: &monitoring.UptimeCheckConfigTcpCheckArgs{
// 				Port: pulumi.Int(888),
// 			},
// 			ResourceGroup: &monitoring.UptimeCheckConfigResourceGroupArgs{
// 				ResourceType: pulumi.String("INSTANCE"),
// 				GroupId:      check.Name,
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
// UptimeCheckConfig can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig default {{name}}
// ```
type UptimeCheckConfig struct {
	pulumi.CustomResourceState

	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
	// Structure is documented below.
	ContentMatchers UptimeCheckConfigContentMatcherArrayOutput `pulumi:"contentMatchers"`
	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Contains information needed to make an HTTP or HTTPS check.
	// Structure is documented below.
	HttpCheck UptimeCheckConfigHttpCheckPtrOutput `pulumi:"httpCheck"`
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptimeUrl  gceInstance  gaeApp  awsEc2Instance  awsElbLoadBalancer
	// Structure is documented below.
	MonitoredResource UptimeCheckConfigMonitoredResourcePtrOutput `pulumi:"monitoredResource"`
	// A unique resource name for this UptimeCheckConfig. The format is
	// projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
	Name pulumi.StringOutput `pulumi:"name"`
	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	Period pulumi.StringPtrOutput `pulumi:"period"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The group resource associated with the configuration.
	// Structure is documented below.
	ResourceGroup UptimeCheckConfigResourceGroupPtrOutput `pulumi:"resourceGroup"`
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
	SelectedRegions pulumi.StringArrayOutput `pulumi:"selectedRegions"`
	// Contains information needed to make a TCP check.
	// Structure is documented below.
	TcpCheck UptimeCheckConfigTcpCheckPtrOutput `pulumi:"tcpCheck"`
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	Timeout pulumi.StringOutput `pulumi:"timeout"`
	// The id of the uptime check
	UptimeCheckId pulumi.StringOutput `pulumi:"uptimeCheckId"`
}

// NewUptimeCheckConfig registers a new resource with the given unique name, arguments, and options.
func NewUptimeCheckConfig(ctx *pulumi.Context,
	name string, args *UptimeCheckConfigArgs, opts ...pulumi.ResourceOption) (*UptimeCheckConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	var resource UptimeCheckConfig
	err := ctx.RegisterResource("gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUptimeCheckConfig gets an existing UptimeCheckConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUptimeCheckConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UptimeCheckConfigState, opts ...pulumi.ResourceOption) (*UptimeCheckConfig, error) {
	var resource UptimeCheckConfig
	err := ctx.ReadResource("gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UptimeCheckConfig resources.
type uptimeCheckConfigState struct {
	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
	// Structure is documented below.
	ContentMatchers []UptimeCheckConfigContentMatcher `pulumi:"contentMatchers"`
	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	DisplayName *string `pulumi:"displayName"`
	// Contains information needed to make an HTTP or HTTPS check.
	// Structure is documented below.
	HttpCheck *UptimeCheckConfigHttpCheck `pulumi:"httpCheck"`
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptimeUrl  gceInstance  gaeApp  awsEc2Instance  awsElbLoadBalancer
	// Structure is documented below.
	MonitoredResource *UptimeCheckConfigMonitoredResource `pulumi:"monitoredResource"`
	// A unique resource name for this UptimeCheckConfig. The format is
	// projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
	Name *string `pulumi:"name"`
	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	Period *string `pulumi:"period"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The group resource associated with the configuration.
	// Structure is documented below.
	ResourceGroup *UptimeCheckConfigResourceGroup `pulumi:"resourceGroup"`
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
	SelectedRegions []string `pulumi:"selectedRegions"`
	// Contains information needed to make a TCP check.
	// Structure is documented below.
	TcpCheck *UptimeCheckConfigTcpCheck `pulumi:"tcpCheck"`
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	Timeout *string `pulumi:"timeout"`
	// The id of the uptime check
	UptimeCheckId *string `pulumi:"uptimeCheckId"`
}

type UptimeCheckConfigState struct {
	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
	// Structure is documented below.
	ContentMatchers UptimeCheckConfigContentMatcherArrayInput
	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	DisplayName pulumi.StringPtrInput
	// Contains information needed to make an HTTP or HTTPS check.
	// Structure is documented below.
	HttpCheck UptimeCheckConfigHttpCheckPtrInput
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptimeUrl  gceInstance  gaeApp  awsEc2Instance  awsElbLoadBalancer
	// Structure is documented below.
	MonitoredResource UptimeCheckConfigMonitoredResourcePtrInput
	// A unique resource name for this UptimeCheckConfig. The format is
	// projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
	Name pulumi.StringPtrInput
	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	Period pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The group resource associated with the configuration.
	// Structure is documented below.
	ResourceGroup UptimeCheckConfigResourceGroupPtrInput
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
	SelectedRegions pulumi.StringArrayInput
	// Contains information needed to make a TCP check.
	// Structure is documented below.
	TcpCheck UptimeCheckConfigTcpCheckPtrInput
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	Timeout pulumi.StringPtrInput
	// The id of the uptime check
	UptimeCheckId pulumi.StringPtrInput
}

func (UptimeCheckConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*uptimeCheckConfigState)(nil)).Elem()
}

type uptimeCheckConfigArgs struct {
	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
	// Structure is documented below.
	ContentMatchers []UptimeCheckConfigContentMatcher `pulumi:"contentMatchers"`
	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	DisplayName string `pulumi:"displayName"`
	// Contains information needed to make an HTTP or HTTPS check.
	// Structure is documented below.
	HttpCheck *UptimeCheckConfigHttpCheck `pulumi:"httpCheck"`
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptimeUrl  gceInstance  gaeApp  awsEc2Instance  awsElbLoadBalancer
	// Structure is documented below.
	MonitoredResource *UptimeCheckConfigMonitoredResource `pulumi:"monitoredResource"`
	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	Period *string `pulumi:"period"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The group resource associated with the configuration.
	// Structure is documented below.
	ResourceGroup *UptimeCheckConfigResourceGroup `pulumi:"resourceGroup"`
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
	SelectedRegions []string `pulumi:"selectedRegions"`
	// Contains information needed to make a TCP check.
	// Structure is documented below.
	TcpCheck *UptimeCheckConfigTcpCheck `pulumi:"tcpCheck"`
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	Timeout string `pulumi:"timeout"`
}

// The set of arguments for constructing a UptimeCheckConfig resource.
type UptimeCheckConfigArgs struct {
	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
	// Structure is documented below.
	ContentMatchers UptimeCheckConfigContentMatcherArrayInput
	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	DisplayName pulumi.StringInput
	// Contains information needed to make an HTTP or HTTPS check.
	// Structure is documented below.
	HttpCheck UptimeCheckConfigHttpCheckPtrInput
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptimeUrl  gceInstance  gaeApp  awsEc2Instance  awsElbLoadBalancer
	// Structure is documented below.
	MonitoredResource UptimeCheckConfigMonitoredResourcePtrInput
	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	Period pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The group resource associated with the configuration.
	// Structure is documented below.
	ResourceGroup UptimeCheckConfigResourceGroupPtrInput
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
	SelectedRegions pulumi.StringArrayInput
	// Contains information needed to make a TCP check.
	// Structure is documented below.
	TcpCheck UptimeCheckConfigTcpCheckPtrInput
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	Timeout pulumi.StringInput
}

func (UptimeCheckConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*uptimeCheckConfigArgs)(nil)).Elem()
}

type UptimeCheckConfigInput interface {
	pulumi.Input

	ToUptimeCheckConfigOutput() UptimeCheckConfigOutput
	ToUptimeCheckConfigOutputWithContext(ctx context.Context) UptimeCheckConfigOutput
}

func (UptimeCheckConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*UptimeCheckConfig)(nil)).Elem()
}

func (i UptimeCheckConfig) ToUptimeCheckConfigOutput() UptimeCheckConfigOutput {
	return i.ToUptimeCheckConfigOutputWithContext(context.Background())
}

func (i UptimeCheckConfig) ToUptimeCheckConfigOutputWithContext(ctx context.Context) UptimeCheckConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UptimeCheckConfigOutput)
}

type UptimeCheckConfigOutput struct {
	*pulumi.OutputState
}

func (UptimeCheckConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UptimeCheckConfigOutput)(nil)).Elem()
}

func (o UptimeCheckConfigOutput) ToUptimeCheckConfigOutput() UptimeCheckConfigOutput {
	return o
}

func (o UptimeCheckConfigOutput) ToUptimeCheckConfigOutputWithContext(ctx context.Context) UptimeCheckConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UptimeCheckConfigOutput{})
}
