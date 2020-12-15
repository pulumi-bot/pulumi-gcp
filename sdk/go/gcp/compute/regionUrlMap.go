// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// UrlMaps are used to route requests to a backend service based on rules
// that you define for the host and path of an incoming URL.
//
// ## Example Usage
// ### Region Url Map Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewRegionHealthCheck(ctx, "_default", &compute.RegionHealthCheckArgs{
// 			Region:           pulumi.String("us-central1"),
// 			CheckIntervalSec: pulumi.Int(1),
// 			TimeoutSec:       pulumi.Int(1),
// 			HttpHealthCheck: &compute.RegionHealthCheckHttpHealthCheckArgs{
// 				Port:        pulumi.Int(80),
// 				RequestPath: pulumi.String("/"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		login, err := compute.NewRegionBackendService(ctx, "login", &compute.RegionBackendServiceArgs{
// 			Region:     pulumi.String("us-central1"),
// 			Protocol:   pulumi.String("HTTP"),
// 			TimeoutSec: pulumi.Int(10),
// 			HealthChecks: pulumi.String(pulumi.String{
// 				_default.ID(),
// 			}),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		home, err := compute.NewRegionBackendService(ctx, "home", &compute.RegionBackendServiceArgs{
// 			Region:     pulumi.String("us-central1"),
// 			Protocol:   pulumi.String("HTTP"),
// 			TimeoutSec: pulumi.Int(10),
// 			HealthChecks: pulumi.String(pulumi.String{
// 				_default.ID(),
// 			}),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRegionUrlMap(ctx, "regionurlmap", &compute.RegionUrlMapArgs{
// 			Region:         pulumi.String("us-central1"),
// 			Description:    pulumi.String("a description"),
// 			DefaultService: home.ID(),
// 			HostRules: compute.RegionUrlMapHostRuleArray{
// 				&compute.RegionUrlMapHostRuleArgs{
// 					Hosts: pulumi.StringArray{
// 						pulumi.String("mysite.com"),
// 					},
// 					PathMatcher: pulumi.String("allpaths"),
// 				},
// 			},
// 			PathMatchers: compute.RegionUrlMapPathMatcherArray{
// 				&compute.RegionUrlMapPathMatcherArgs{
// 					Name:           pulumi.String("allpaths"),
// 					DefaultService: home.ID(),
// 					PathRules: compute.RegionUrlMapPathMatcherPathRuleArray{
// 						&compute.RegionUrlMapPathMatcherPathRuleArgs{
// 							Paths: pulumi.StringArray{
// 								pulumi.String("/home"),
// 							},
// 							Service: home.ID(),
// 						},
// 						&compute.RegionUrlMapPathMatcherPathRuleArgs{
// 							Paths: pulumi.StringArray{
// 								pulumi.String("/login"),
// 							},
// 							Service: login.ID(),
// 						},
// 					},
// 				},
// 			},
// 			Tests: compute.RegionUrlMapTestArray{
// 				&compute.RegionUrlMapTestArgs{
// 					Service: home.ID(),
// 					Host:    pulumi.String("hi.com"),
// 					Path:    pulumi.String("/home"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Region Url Map L7 Ilb Path
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewRegionHealthCheck(ctx, "_default", &compute.RegionHealthCheckArgs{
// 			HttpHealthCheck: &compute.RegionHealthCheckHttpHealthCheckArgs{
// 				Port: pulumi.Int(80),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		home, err := compute.NewRegionBackendService(ctx, "home", &compute.RegionBackendServiceArgs{
// 			Protocol:   pulumi.String("HTTP"),
// 			TimeoutSec: pulumi.Int(10),
// 			HealthChecks: pulumi.String(pulumi.String{
// 				_default.ID(),
// 			}),
// 			LoadBalancingScheme: pulumi.String("INTERNAL_MANAGED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRegionUrlMap(ctx, "regionurlmap", &compute.RegionUrlMapArgs{
// 			Description:    pulumi.String("a description"),
// 			DefaultService: home.ID(),
// 			HostRules: compute.RegionUrlMapHostRuleArray{
// 				&compute.RegionUrlMapHostRuleArgs{
// 					Hosts: pulumi.StringArray{
// 						pulumi.String("mysite.com"),
// 					},
// 					PathMatcher: pulumi.String("allpaths"),
// 				},
// 			},
// 			PathMatchers: compute.RegionUrlMapPathMatcherArray{
// 				&compute.RegionUrlMapPathMatcherArgs{
// 					Name:           pulumi.String("allpaths"),
// 					DefaultService: home.ID(),
// 					PathRules: compute.RegionUrlMapPathMatcherPathRuleArray{
// 						&compute.RegionUrlMapPathMatcherPathRuleArgs{
// 							Paths: pulumi.StringArray{
// 								pulumi.String("/home"),
// 							},
// 							RouteAction: &compute.RegionUrlMapPathMatcherPathRuleRouteActionArgs{
// 								CorsPolicy: &compute.RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyArgs{
// 									AllowCredentials: pulumi.Bool(true),
// 									AllowHeaders: pulumi.StringArray{
// 										pulumi.String("Allowed content"),
// 									},
// 									AllowMethods: pulumi.StringArray{
// 										pulumi.String("GET"),
// 									},
// 									AllowOrigins: pulumi.StringArray{
// 										pulumi.String("Allowed origin"),
// 									},
// 									ExposeHeaders: pulumi.StringArray{
// 										pulumi.String("Exposed header"),
// 									},
// 									MaxAge:   pulumi.Int(30),
// 									Disabled: pulumi.Bool(false),
// 								},
// 								FaultInjectionPolicy: &compute.RegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyArgs{
// 									Abort: &compute.RegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortArgs{
// 										HttpStatus: pulumi.Int(234),
// 										Percentage: pulumi.Float64(5.6),
// 									},
// 									Delay: &compute.RegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayArgs{
// 										FixedDelay: &compute.RegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayArgs{
// 											Seconds: pulumi.String("0"),
// 											Nanos:   pulumi.Int(50000),
// 										},
// 										Percentage: pulumi.Float64(7.8),
// 									},
// 								},
// 								RequestMirrorPolicy: &compute.RegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyArgs{
// 									BackendService: home.ID(),
// 								},
// 								RetryPolicy: &compute.RegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyArgs{
// 									NumRetries: pulumi.Int(4),
// 									PerTryTimeout: &compute.RegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutArgs{
// 										Seconds: pulumi.String("30"),
// 									},
// 									RetryConditions: pulumi.StringArray{
// 										pulumi.String("5xx"),
// 										pulumi.String("deadline-exceeded"),
// 									},
// 								},
// 								Timeout: &compute.RegionUrlMapPathMatcherPathRuleRouteActionTimeoutArgs{
// 									Seconds: pulumi.String("20"),
// 									Nanos:   pulumi.Int(750000000),
// 								},
// 								UrlRewrite: &compute.RegionUrlMapPathMatcherPathRuleRouteActionUrlRewriteArgs{
// 									HostRewrite:       pulumi.String("A replacement header"),
// 									PathPrefixRewrite: pulumi.String("A replacement path"),
// 								},
// 								WeightedBackendServices: compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArray{
// 									&compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArgs{
// 										BackendService: home.ID(),
// 										Weight:         pulumi.Int(400),
// 										HeaderAction: &compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionArgs{
// 											RequestHeadersToRemoves: pulumi.StringArray{
// 												pulumi.String("RemoveMe"),
// 											},
// 											RequestHeadersToAdds: compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionRequestHeadersToAddArray{
// 												&compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionRequestHeadersToAddArgs{
// 													HeaderName:  pulumi.String("AddMe"),
// 													HeaderValue: pulumi.String("MyValue"),
// 													Replace:     pulumi.Bool(true),
// 												},
// 											},
// 											ResponseHeadersToRemoves: pulumi.StringArray{
// 												pulumi.String("RemoveMe"),
// 											},
// 											ResponseHeadersToAdds: compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionResponseHeadersToAddArray{
// 												&compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionResponseHeadersToAddArgs{
// 													HeaderName:  pulumi.String("AddMe"),
// 													HeaderValue: pulumi.String("MyValue"),
// 													Replace:     pulumi.Bool(false),
// 												},
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Tests: compute.RegionUrlMapTestArray{
// 				&compute.RegionUrlMapTestArgs{
// 					Service: home.ID(),
// 					Host:    pulumi.String("hi.com"),
// 					Path:    pulumi.String("/home"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Region Url Map L7 Ilb Path Partial
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewRegionHealthCheck(ctx, "_default", &compute.RegionHealthCheckArgs{
// 			HttpHealthCheck: &compute.RegionHealthCheckHttpHealthCheckArgs{
// 				Port: pulumi.Int(80),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		home, err := compute.NewRegionBackendService(ctx, "home", &compute.RegionBackendServiceArgs{
// 			Protocol:   pulumi.String("HTTP"),
// 			TimeoutSec: pulumi.Int(10),
// 			HealthChecks: pulumi.String(pulumi.String{
// 				_default.ID(),
// 			}),
// 			LoadBalancingScheme: pulumi.String("INTERNAL_MANAGED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRegionUrlMap(ctx, "regionurlmap", &compute.RegionUrlMapArgs{
// 			Description:    pulumi.String("a description"),
// 			DefaultService: home.ID(),
// 			HostRules: compute.RegionUrlMapHostRuleArray{
// 				&compute.RegionUrlMapHostRuleArgs{
// 					Hosts: pulumi.StringArray{
// 						pulumi.String("mysite.com"),
// 					},
// 					PathMatcher: pulumi.String("allpaths"),
// 				},
// 			},
// 			PathMatchers: compute.RegionUrlMapPathMatcherArray{
// 				&compute.RegionUrlMapPathMatcherArgs{
// 					Name:           pulumi.String("allpaths"),
// 					DefaultService: home.ID(),
// 					PathRules: compute.RegionUrlMapPathMatcherPathRuleArray{
// 						&compute.RegionUrlMapPathMatcherPathRuleArgs{
// 							Paths: pulumi.StringArray{
// 								pulumi.String("/home"),
// 							},
// 							RouteAction: &compute.RegionUrlMapPathMatcherPathRuleRouteActionArgs{
// 								RetryPolicy: &compute.RegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyArgs{
// 									NumRetries: pulumi.Int(4),
// 									PerTryTimeout: &compute.RegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutArgs{
// 										Seconds: pulumi.String("30"),
// 									},
// 									RetryConditions: pulumi.StringArray{
// 										pulumi.String("5xx"),
// 										pulumi.String("deadline-exceeded"),
// 									},
// 								},
// 								Timeout: &compute.RegionUrlMapPathMatcherPathRuleRouteActionTimeoutArgs{
// 									Seconds: pulumi.String("20"),
// 									Nanos:   pulumi.Int(750000000),
// 								},
// 								UrlRewrite: &compute.RegionUrlMapPathMatcherPathRuleRouteActionUrlRewriteArgs{
// 									HostRewrite:       pulumi.String("A replacement header"),
// 									PathPrefixRewrite: pulumi.String("A replacement path"),
// 								},
// 								WeightedBackendServices: compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArray{
// 									&compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArgs{
// 										BackendService: home.ID(),
// 										Weight:         pulumi.Int(400),
// 										HeaderAction: &compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionArgs{
// 											ResponseHeadersToAdds: compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionResponseHeadersToAddArray{
// 												&compute.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionResponseHeadersToAddArgs{
// 													HeaderName:  pulumi.String("AddMe"),
// 													HeaderValue: pulumi.String("MyValue"),
// 													Replace:     pulumi.Bool(false),
// 												},
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Tests: compute.RegionUrlMapTestArray{
// 				&compute.RegionUrlMapTestArgs{
// 					Service: home.ID(),
// 					Host:    pulumi.String("hi.com"),
// 					Path:    pulumi.String("/home"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Region Url Map L7 Ilb Route
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewRegionHealthCheck(ctx, "_default", &compute.RegionHealthCheckArgs{
// 			HttpHealthCheck: &compute.RegionHealthCheckHttpHealthCheckArgs{
// 				Port: pulumi.Int(80),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		home, err := compute.NewRegionBackendService(ctx, "home", &compute.RegionBackendServiceArgs{
// 			Protocol:   pulumi.String("HTTP"),
// 			TimeoutSec: pulumi.Int(10),
// 			HealthChecks: pulumi.String(pulumi.String{
// 				_default.ID(),
// 			}),
// 			LoadBalancingScheme: pulumi.String("INTERNAL_MANAGED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRegionUrlMap(ctx, "regionurlmap", &compute.RegionUrlMapArgs{
// 			Description:    pulumi.String("a description"),
// 			DefaultService: home.ID(),
// 			HostRules: compute.RegionUrlMapHostRuleArray{
// 				&compute.RegionUrlMapHostRuleArgs{
// 					Hosts: pulumi.StringArray{
// 						pulumi.String("mysite.com"),
// 					},
// 					PathMatcher: pulumi.String("allpaths"),
// 				},
// 			},
// 			PathMatchers: compute.RegionUrlMapPathMatcherArray{
// 				&compute.RegionUrlMapPathMatcherArgs{
// 					Name:           pulumi.String("allpaths"),
// 					DefaultService: home.ID(),
// 					RouteRules: compute.RegionUrlMapPathMatcherRouteRuleArray{
// 						&compute.RegionUrlMapPathMatcherRouteRuleArgs{
// 							Priority: pulumi.Int(1),
// 							HeaderAction: &compute.RegionUrlMapPathMatcherRouteRuleHeaderActionArgs{
// 								RequestHeadersToRemoves: pulumi.StringArray{
// 									pulumi.String("RemoveMe2"),
// 								},
// 								RequestHeadersToAdds: compute.RegionUrlMapPathMatcherRouteRuleHeaderActionRequestHeadersToAddArray{
// 									&compute.RegionUrlMapPathMatcherRouteRuleHeaderActionRequestHeadersToAddArgs{
// 										HeaderName:  pulumi.String("AddSomethingElse"),
// 										HeaderValue: pulumi.String("MyOtherValue"),
// 										Replace:     pulumi.Bool(true),
// 									},
// 								},
// 								ResponseHeadersToRemoves: pulumi.StringArray{
// 									pulumi.String("RemoveMe3"),
// 								},
// 								ResponseHeadersToAdds: compute.RegionUrlMapPathMatcherRouteRuleHeaderActionResponseHeadersToAddArray{
// 									&compute.RegionUrlMapPathMatcherRouteRuleHeaderActionResponseHeadersToAddArgs{
// 										HeaderName:  pulumi.String("AddMe"),
// 										HeaderValue: pulumi.String("MyValue"),
// 										Replace:     pulumi.Bool(false),
// 									},
// 								},
// 							},
// 							MatchRules: compute.RegionUrlMapPathMatcherRouteRuleMatchRuleArray{
// 								&compute.RegionUrlMapPathMatcherRouteRuleMatchRuleArgs{
// 									FullPathMatch: pulumi.String("a full path"),
// 									HeaderMatches: compute.RegionUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchArray{
// 										&compute.RegionUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs{
// 											HeaderName:  pulumi.String("someheader"),
// 											ExactMatch:  pulumi.String("match this exactly"),
// 											InvertMatch: pulumi.Bool(true),
// 										},
// 									},
// 									IgnoreCase: pulumi.Bool(true),
// 									MetadataFilters: compute.RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArray{
// 										&compute.RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs{
// 											FilterMatchCriteria: pulumi.String("MATCH_ANY"),
// 											FilterLabels: compute.RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArray{
// 												&compute.RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArgs{
// 													Name:  pulumi.String("PLANET"),
// 													Value: pulumi.String("MARS"),
// 												},
// 											},
// 										},
// 									},
// 									QueryParameterMatches: compute.RegionUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArray{
// 										&compute.RegionUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArgs{
// 											Name:         pulumi.String("a query parameter"),
// 											PresentMatch: pulumi.Bool(true),
// 										},
// 									},
// 								},
// 							},
// 							UrlRedirect: &compute.RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs{
// 								HostRedirect:         pulumi.String("A host"),
// 								HttpsRedirect:        pulumi.Bool(false),
// 								PathRedirect:         pulumi.String("some/path"),
// 								RedirectResponseCode: pulumi.String("TEMPORARY_REDIRECT"),
// 								StripQuery:           pulumi.Bool(true),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Tests: compute.RegionUrlMapTestArray{
// 				&compute.RegionUrlMapTestArgs{
// 					Service: home.ID(),
// 					Host:    pulumi.String("hi.com"),
// 					Path:    pulumi.String("/home"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Region Url Map L7 Ilb Route Partial
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewRegionHealthCheck(ctx, "_default", &compute.RegionHealthCheckArgs{
// 			HttpHealthCheck: &compute.RegionHealthCheckHttpHealthCheckArgs{
// 				Port: pulumi.Int(80),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		home, err := compute.NewRegionBackendService(ctx, "home", &compute.RegionBackendServiceArgs{
// 			Protocol:   pulumi.String("HTTP"),
// 			TimeoutSec: pulumi.Int(10),
// 			HealthChecks: pulumi.String(pulumi.String{
// 				_default.ID(),
// 			}),
// 			LoadBalancingScheme: pulumi.String("INTERNAL_MANAGED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRegionUrlMap(ctx, "regionurlmap", &compute.RegionUrlMapArgs{
// 			Description:    pulumi.String("a description"),
// 			DefaultService: home.ID(),
// 			HostRules: compute.RegionUrlMapHostRuleArray{
// 				&compute.RegionUrlMapHostRuleArgs{
// 					Hosts: pulumi.StringArray{
// 						pulumi.String("mysite.com"),
// 					},
// 					PathMatcher: pulumi.String("allpaths"),
// 				},
// 			},
// 			PathMatchers: compute.RegionUrlMapPathMatcherArray{
// 				&compute.RegionUrlMapPathMatcherArgs{
// 					Name:           pulumi.String("allpaths"),
// 					DefaultService: home.ID(),
// 					RouteRules: compute.RegionUrlMapPathMatcherRouteRuleArray{
// 						&compute.RegionUrlMapPathMatcherRouteRuleArgs{
// 							Priority: pulumi.Int(1),
// 							Service:  home.ID(),
// 							HeaderAction: &compute.RegionUrlMapPathMatcherRouteRuleHeaderActionArgs{
// 								RequestHeadersToRemoves: pulumi.StringArray{
// 									pulumi.String("RemoveMe2"),
// 								},
// 							},
// 							MatchRules: compute.RegionUrlMapPathMatcherRouteRuleMatchRuleArray{
// 								&compute.RegionUrlMapPathMatcherRouteRuleMatchRuleArgs{
// 									FullPathMatch: pulumi.String("a full path"),
// 									HeaderMatches: compute.RegionUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchArray{
// 										&compute.RegionUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs{
// 											HeaderName:  pulumi.String("someheader"),
// 											ExactMatch:  pulumi.String("match this exactly"),
// 											InvertMatch: pulumi.Bool(true),
// 										},
// 									},
// 									QueryParameterMatches: compute.RegionUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArray{
// 										&compute.RegionUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArgs{
// 											Name:         pulumi.String("a query parameter"),
// 											PresentMatch: pulumi.Bool(true),
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Tests: compute.RegionUrlMapTestArray{
// 				&compute.RegionUrlMapTestArgs{
// 					Service: home.ID(),
// 					Host:    pulumi.String("hi.com"),
// 					Path:    pulumi.String("/home"),
// 				},
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
// RegionUrlMap can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default projects/{{project}}/regions/{{region}}/urlMaps/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default {{name}}
// ```
type RegionUrlMap struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// A reference to a RegionBackendService resource. This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	DefaultService pulumi.StringPtrOutput `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect RegionUrlMapDefaultUrlRedirectPtrOutput `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules RegionUrlMapHostRuleArrayOutput `pulumi:"hostRules"`
	// The unique identifier for the resource.
	MapId pulumi.IntOutput `pulumi:"mapId"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of
	// the URL if the hostRule matches the URL's host portion.
	PathMatchers RegionUrlMapPathMatcherArrayOutput `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Region in which the url map should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	// Structure is documented below.
	Tests RegionUrlMapTestArrayOutput `pulumi:"tests"`
}

// NewRegionUrlMap registers a new resource with the given unique name, arguments, and options.
func NewRegionUrlMap(ctx *pulumi.Context,
	name string, args *RegionUrlMapArgs, opts ...pulumi.ResourceOption) (*RegionUrlMap, error) {
	if args == nil {
		args = &RegionUrlMapArgs{}
	}

	var resource RegionUrlMap
	err := ctx.RegisterResource("gcp:compute/regionUrlMap:RegionUrlMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionUrlMap gets an existing RegionUrlMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionUrlMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionUrlMapState, opts ...pulumi.ResourceOption) (*RegionUrlMap, error) {
	var resource RegionUrlMap
	err := ctx.ReadResource("gcp:compute/regionUrlMap:RegionUrlMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionUrlMap resources.
type regionUrlMapState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// A reference to a RegionBackendService resource. This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect *RegionUrlMapDefaultUrlRedirect `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	Fingerprint *string `pulumi:"fingerprint"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules []RegionUrlMapHostRule `pulumi:"hostRules"`
	// The unique identifier for the resource.
	MapId *int `pulumi:"mapId"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name *string `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of
	// the URL if the hostRule matches the URL's host portion.
	PathMatchers []RegionUrlMapPathMatcher `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the url map should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	// Structure is documented below.
	Tests []RegionUrlMapTest `pulumi:"tests"`
}

type RegionUrlMapState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// A reference to a RegionBackendService resource. This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect RegionUrlMapDefaultUrlRedirectPtrInput
	// Description of this test case.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	Fingerprint pulumi.StringPtrInput
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules RegionUrlMapHostRuleArrayInput
	// The unique identifier for the resource.
	MapId pulumi.IntPtrInput
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringPtrInput
	// The name of the PathMatcher to use to match the path portion of
	// the URL if the hostRule matches the URL's host portion.
	PathMatchers RegionUrlMapPathMatcherArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the url map should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	// Structure is documented below.
	Tests RegionUrlMapTestArrayInput
}

func (RegionUrlMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionUrlMapState)(nil)).Elem()
}

type regionUrlMapArgs struct {
	// A reference to a RegionBackendService resource. This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect *RegionUrlMapDefaultUrlRedirect `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description *string `pulumi:"description"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules []RegionUrlMapHostRule `pulumi:"hostRules"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name *string `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of
	// the URL if the hostRule matches the URL's host portion.
	PathMatchers []RegionUrlMapPathMatcher `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the url map should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	// Structure is documented below.
	Tests []RegionUrlMapTest `pulumi:"tests"`
}

// The set of arguments for constructing a RegionUrlMap resource.
type RegionUrlMapArgs struct {
	// A reference to a RegionBackendService resource. This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect RegionUrlMapDefaultUrlRedirectPtrInput
	// Description of this test case.
	Description pulumi.StringPtrInput
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules RegionUrlMapHostRuleArrayInput
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringPtrInput
	// The name of the PathMatcher to use to match the path portion of
	// the URL if the hostRule matches the URL's host portion.
	PathMatchers RegionUrlMapPathMatcherArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the url map should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	// Structure is documented below.
	Tests RegionUrlMapTestArrayInput
}

func (RegionUrlMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionUrlMapArgs)(nil)).Elem()
}

type RegionUrlMapInput interface {
	pulumi.Input

	ToRegionUrlMapOutput() RegionUrlMapOutput
	ToRegionUrlMapOutputWithContext(ctx context.Context) RegionUrlMapOutput
}

func (*RegionUrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionUrlMap)(nil))
}

func (i *RegionUrlMap) ToRegionUrlMapOutput() RegionUrlMapOutput {
	return i.ToRegionUrlMapOutputWithContext(context.Background())
}

func (i *RegionUrlMap) ToRegionUrlMapOutputWithContext(ctx context.Context) RegionUrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapOutput)
}

type RegionUrlMapOutput struct {
	*pulumi.OutputState
}

func (RegionUrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionUrlMap)(nil))
}

func (o RegionUrlMapOutput) ToRegionUrlMapOutput() RegionUrlMapOutput {
	return o
}

func (o RegionUrlMapOutput) ToRegionUrlMapOutputWithContext(ctx context.Context) RegionUrlMapOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionUrlMapOutput{})
}
