// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a GlobalForwardingRule resource. Global forwarding rules are
// used to forward traffic to the correct load balancer for HTTP load
// balancing. Global forwarding rules can only be used for HTTP load
// balancing.
//
// For more information, see
// https://cloud.google.com/compute/docs/load-balancing/http/
//
// ## Example Usage
//
// ## Import
//
// GlobalForwardingRule can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/globalForwardingRule:GlobalForwardingRule default projects/{{project}}/global/forwardingRules/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/globalForwardingRule:GlobalForwardingRule default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/globalForwardingRule:GlobalForwardingRule default {{name}}
// ```
type GlobalForwardingRule struct {
	pulumi.CustomResourceState

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IP address that this forwarding rule serves. When a client sends
	// traffic to this IP address, the forwarding rule directs the traffic to
	// the target that you specify in the forwarding rule. The
	// loadBalancingScheme and the forwarding rule's target determine the
	// type of IP address that you can use. For detailed information, refer
	// to [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	// An address can be specified either by a literal IP address or a
	// reference to an existing Address resource. If you don't specify a
	// reserved IP address, an ephemeral IP address is assigned.
	// The value must be set to 0.0.0.0 when the target is a targetGrpcProxy
	// that has validateForProxyless field set to true.
	// For Private Service Connect forwarding rules that forward traffic to
	// Google APIs, IP address must be provided.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. When the load balancing scheme is
	// INTERNAL_SELF_MANAGED, only TCP is valid. This field must not be set if the
	// global address is configured as a purpose of PRIVATE_SERVICE_CONNECT
	// and addressType of INTERNAL
	// Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
	IpProtocol pulumi.StringOutput `pulumi:"ipProtocol"`
	// The IP Version that will be used by this global forwarding rule.
	// Possible values are `IPV4` and `IPV6`.
	IpVersion pulumi.StringPtrOutput `pulumi:"ipVersion"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// This signifies what the GlobalForwardingRule will be used for.
	// The value of INTERNAL_SELF_MANAGED means that this will be used for
	// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
	// will be used for External Global Load Balancing (HTTP(S) LB,
	// External TCP/UDP LB, SSL Proxy)
	// Note: This field must be set "" if the global address is
	// configured as a purpose of PRIVATE_SERVICE_CONNECT and addressType of INTERNAL.
	// Default value is `EXTERNAL`.
	// Possible values are `EXTERNAL` and `INTERNAL_SELF_MANAGED`.
	LoadBalancingScheme pulumi.StringPtrOutput `pulumi:"loadBalancingScheme"`
	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// Structure is documented below.
	MetadataFilters GlobalForwardingRuleMetadataFilterArrayOutput `pulumi:"metadataFilters"`
	// Name of the metadata label. The length must be between
	// 1 and 1024 characters, inclusive.
	Name pulumi.StringOutput `pulumi:"name"`
	// This field is not used for external load balancing.
	// For INTERNAL_SELF_MANAGED load balancing, this field
	// identifies the network that the load balanced IP should belong to
	// for this global forwarding rule. If this field is not specified,
	// the default network will be used.
	Network pulumi.StringOutput `pulumi:"network"`
	// This field is used along with the target field for TargetHttpProxy,
	// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
	// TargetPool, TargetInstance.
	// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target.
	// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
	// disjoint port ranges.
	// Some types of forwarding target have constraints on the acceptable
	// ports:
	// * TargetHttpProxy: 80, 8080
	// * TargetHttpsProxy: 443
	// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	//   1883, 5222
	// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	//   1883, 5222
	// * TargetVpnGateway: 500, 4500
	PortRange pulumi.StringPtrOutput `pulumi:"portRange"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The URL of the target resource to receive the matched traffic.
	// The forwarded traffic must be of a type appropriate to the target object.
	// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	// are valid.
	// For global address with a purpose of PRIVATE_SERVICE_CONNECT and
	// addressType of INTERNAL, only "all-apis" and "vpc-sc" are valid.
	Target pulumi.StringOutput `pulumi:"target"`
}

// NewGlobalForwardingRule registers a new resource with the given unique name, arguments, and options.
func NewGlobalForwardingRule(ctx *pulumi.Context,
	name string, args *GlobalForwardingRuleArgs, opts ...pulumi.ResourceOption) (*GlobalForwardingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	var resource GlobalForwardingRule
	err := ctx.RegisterResource("gcp:compute/globalForwardingRule:GlobalForwardingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalForwardingRule gets an existing GlobalForwardingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalForwardingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalForwardingRuleState, opts ...pulumi.ResourceOption) (*GlobalForwardingRule, error) {
	var resource GlobalForwardingRule
	err := ctx.ReadResource("gcp:compute/globalForwardingRule:GlobalForwardingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalForwardingRule resources.
type globalForwardingRuleState struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// The IP address that this forwarding rule serves. When a client sends
	// traffic to this IP address, the forwarding rule directs the traffic to
	// the target that you specify in the forwarding rule. The
	// loadBalancingScheme and the forwarding rule's target determine the
	// type of IP address that you can use. For detailed information, refer
	// to [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	// An address can be specified either by a literal IP address or a
	// reference to an existing Address resource. If you don't specify a
	// reserved IP address, an ephemeral IP address is assigned.
	// The value must be set to 0.0.0.0 when the target is a targetGrpcProxy
	// that has validateForProxyless field set to true.
	// For Private Service Connect forwarding rules that forward traffic to
	// Google APIs, IP address must be provided.
	IpAddress *string `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. When the load balancing scheme is
	// INTERNAL_SELF_MANAGED, only TCP is valid. This field must not be set if the
	// global address is configured as a purpose of PRIVATE_SERVICE_CONNECT
	// and addressType of INTERNAL
	// Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
	IpProtocol *string `pulumi:"ipProtocol"`
	// The IP Version that will be used by this global forwarding rule.
	// Possible values are `IPV4` and `IPV6`.
	IpVersion *string `pulumi:"ipVersion"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// This signifies what the GlobalForwardingRule will be used for.
	// The value of INTERNAL_SELF_MANAGED means that this will be used for
	// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
	// will be used for External Global Load Balancing (HTTP(S) LB,
	// External TCP/UDP LB, SSL Proxy)
	// Note: This field must be set "" if the global address is
	// configured as a purpose of PRIVATE_SERVICE_CONNECT and addressType of INTERNAL.
	// Default value is `EXTERNAL`.
	// Possible values are `EXTERNAL` and `INTERNAL_SELF_MANAGED`.
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// Structure is documented below.
	MetadataFilters []GlobalForwardingRuleMetadataFilter `pulumi:"metadataFilters"`
	// Name of the metadata label. The length must be between
	// 1 and 1024 characters, inclusive.
	Name *string `pulumi:"name"`
	// This field is not used for external load balancing.
	// For INTERNAL_SELF_MANAGED load balancing, this field
	// identifies the network that the load balanced IP should belong to
	// for this global forwarding rule. If this field is not specified,
	// the default network will be used.
	Network *string `pulumi:"network"`
	// This field is used along with the target field for TargetHttpProxy,
	// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
	// TargetPool, TargetInstance.
	// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target.
	// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
	// disjoint port ranges.
	// Some types of forwarding target have constraints on the acceptable
	// ports:
	// * TargetHttpProxy: 80, 8080
	// * TargetHttpsProxy: 443
	// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	//   1883, 5222
	// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	//   1883, 5222
	// * TargetVpnGateway: 500, 4500
	PortRange *string `pulumi:"portRange"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The URL of the target resource to receive the matched traffic.
	// The forwarded traffic must be of a type appropriate to the target object.
	// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	// are valid.
	// For global address with a purpose of PRIVATE_SERVICE_CONNECT and
	// addressType of INTERNAL, only "all-apis" and "vpc-sc" are valid.
	Target *string `pulumi:"target"`
}

type GlobalForwardingRuleState struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// The IP address that this forwarding rule serves. When a client sends
	// traffic to this IP address, the forwarding rule directs the traffic to
	// the target that you specify in the forwarding rule. The
	// loadBalancingScheme and the forwarding rule's target determine the
	// type of IP address that you can use. For detailed information, refer
	// to [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	// An address can be specified either by a literal IP address or a
	// reference to an existing Address resource. If you don't specify a
	// reserved IP address, an ephemeral IP address is assigned.
	// The value must be set to 0.0.0.0 when the target is a targetGrpcProxy
	// that has validateForProxyless field set to true.
	// For Private Service Connect forwarding rules that forward traffic to
	// Google APIs, IP address must be provided.
	IpAddress pulumi.StringPtrInput
	// The IP protocol to which this rule applies. When the load balancing scheme is
	// INTERNAL_SELF_MANAGED, only TCP is valid. This field must not be set if the
	// global address is configured as a purpose of PRIVATE_SERVICE_CONNECT
	// and addressType of INTERNAL
	// Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
	IpProtocol pulumi.StringPtrInput
	// The IP Version that will be used by this global forwarding rule.
	// Possible values are `IPV4` and `IPV6`.
	IpVersion pulumi.StringPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// This signifies what the GlobalForwardingRule will be used for.
	// The value of INTERNAL_SELF_MANAGED means that this will be used for
	// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
	// will be used for External Global Load Balancing (HTTP(S) LB,
	// External TCP/UDP LB, SSL Proxy)
	// Note: This field must be set "" if the global address is
	// configured as a purpose of PRIVATE_SERVICE_CONNECT and addressType of INTERNAL.
	// Default value is `EXTERNAL`.
	// Possible values are `EXTERNAL` and `INTERNAL_SELF_MANAGED`.
	LoadBalancingScheme pulumi.StringPtrInput
	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// Structure is documented below.
	MetadataFilters GlobalForwardingRuleMetadataFilterArrayInput
	// Name of the metadata label. The length must be between
	// 1 and 1024 characters, inclusive.
	Name pulumi.StringPtrInput
	// This field is not used for external load balancing.
	// For INTERNAL_SELF_MANAGED load balancing, this field
	// identifies the network that the load balanced IP should belong to
	// for this global forwarding rule. If this field is not specified,
	// the default network will be used.
	Network pulumi.StringPtrInput
	// This field is used along with the target field for TargetHttpProxy,
	// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
	// TargetPool, TargetInstance.
	// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target.
	// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
	// disjoint port ranges.
	// Some types of forwarding target have constraints on the acceptable
	// ports:
	// * TargetHttpProxy: 80, 8080
	// * TargetHttpsProxy: 443
	// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	//   1883, 5222
	// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	//   1883, 5222
	// * TargetVpnGateway: 500, 4500
	PortRange pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The URL of the target resource to receive the matched traffic.
	// The forwarded traffic must be of a type appropriate to the target object.
	// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	// are valid.
	// For global address with a purpose of PRIVATE_SERVICE_CONNECT and
	// addressType of INTERNAL, only "all-apis" and "vpc-sc" are valid.
	Target pulumi.StringPtrInput
}

func (GlobalForwardingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalForwardingRuleState)(nil)).Elem()
}

type globalForwardingRuleArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// The IP address that this forwarding rule serves. When a client sends
	// traffic to this IP address, the forwarding rule directs the traffic to
	// the target that you specify in the forwarding rule. The
	// loadBalancingScheme and the forwarding rule's target determine the
	// type of IP address that you can use. For detailed information, refer
	// to [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	// An address can be specified either by a literal IP address or a
	// reference to an existing Address resource. If you don't specify a
	// reserved IP address, an ephemeral IP address is assigned.
	// The value must be set to 0.0.0.0 when the target is a targetGrpcProxy
	// that has validateForProxyless field set to true.
	// For Private Service Connect forwarding rules that forward traffic to
	// Google APIs, IP address must be provided.
	IpAddress *string `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. When the load balancing scheme is
	// INTERNAL_SELF_MANAGED, only TCP is valid. This field must not be set if the
	// global address is configured as a purpose of PRIVATE_SERVICE_CONNECT
	// and addressType of INTERNAL
	// Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
	IpProtocol *string `pulumi:"ipProtocol"`
	// The IP Version that will be used by this global forwarding rule.
	// Possible values are `IPV4` and `IPV6`.
	IpVersion *string `pulumi:"ipVersion"`
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// This signifies what the GlobalForwardingRule will be used for.
	// The value of INTERNAL_SELF_MANAGED means that this will be used for
	// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
	// will be used for External Global Load Balancing (HTTP(S) LB,
	// External TCP/UDP LB, SSL Proxy)
	// Note: This field must be set "" if the global address is
	// configured as a purpose of PRIVATE_SERVICE_CONNECT and addressType of INTERNAL.
	// Default value is `EXTERNAL`.
	// Possible values are `EXTERNAL` and `INTERNAL_SELF_MANAGED`.
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// Structure is documented below.
	MetadataFilters []GlobalForwardingRuleMetadataFilter `pulumi:"metadataFilters"`
	// Name of the metadata label. The length must be between
	// 1 and 1024 characters, inclusive.
	Name *string `pulumi:"name"`
	// This field is not used for external load balancing.
	// For INTERNAL_SELF_MANAGED load balancing, this field
	// identifies the network that the load balanced IP should belong to
	// for this global forwarding rule. If this field is not specified,
	// the default network will be used.
	Network *string `pulumi:"network"`
	// This field is used along with the target field for TargetHttpProxy,
	// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
	// TargetPool, TargetInstance.
	// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target.
	// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
	// disjoint port ranges.
	// Some types of forwarding target have constraints on the acceptable
	// ports:
	// * TargetHttpProxy: 80, 8080
	// * TargetHttpsProxy: 443
	// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	//   1883, 5222
	// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	//   1883, 5222
	// * TargetVpnGateway: 500, 4500
	PortRange *string `pulumi:"portRange"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URL of the target resource to receive the matched traffic.
	// The forwarded traffic must be of a type appropriate to the target object.
	// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	// are valid.
	// For global address with a purpose of PRIVATE_SERVICE_CONNECT and
	// addressType of INTERNAL, only "all-apis" and "vpc-sc" are valid.
	Target string `pulumi:"target"`
}

// The set of arguments for constructing a GlobalForwardingRule resource.
type GlobalForwardingRuleArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// The IP address that this forwarding rule serves. When a client sends
	// traffic to this IP address, the forwarding rule directs the traffic to
	// the target that you specify in the forwarding rule. The
	// loadBalancingScheme and the forwarding rule's target determine the
	// type of IP address that you can use. For detailed information, refer
	// to [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	// An address can be specified either by a literal IP address or a
	// reference to an existing Address resource. If you don't specify a
	// reserved IP address, an ephemeral IP address is assigned.
	// The value must be set to 0.0.0.0 when the target is a targetGrpcProxy
	// that has validateForProxyless field set to true.
	// For Private Service Connect forwarding rules that forward traffic to
	// Google APIs, IP address must be provided.
	IpAddress pulumi.StringPtrInput
	// The IP protocol to which this rule applies. When the load balancing scheme is
	// INTERNAL_SELF_MANAGED, only TCP is valid. This field must not be set if the
	// global address is configured as a purpose of PRIVATE_SERVICE_CONNECT
	// and addressType of INTERNAL
	// Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
	IpProtocol pulumi.StringPtrInput
	// The IP Version that will be used by this global forwarding rule.
	// Possible values are `IPV4` and `IPV6`.
	IpVersion pulumi.StringPtrInput
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// This signifies what the GlobalForwardingRule will be used for.
	// The value of INTERNAL_SELF_MANAGED means that this will be used for
	// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
	// will be used for External Global Load Balancing (HTTP(S) LB,
	// External TCP/UDP LB, SSL Proxy)
	// Note: This field must be set "" if the global address is
	// configured as a purpose of PRIVATE_SERVICE_CONNECT and addressType of INTERNAL.
	// Default value is `EXTERNAL`.
	// Possible values are `EXTERNAL` and `INTERNAL_SELF_MANAGED`.
	LoadBalancingScheme pulumi.StringPtrInput
	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// Structure is documented below.
	MetadataFilters GlobalForwardingRuleMetadataFilterArrayInput
	// Name of the metadata label. The length must be between
	// 1 and 1024 characters, inclusive.
	Name pulumi.StringPtrInput
	// This field is not used for external load balancing.
	// For INTERNAL_SELF_MANAGED load balancing, this field
	// identifies the network that the load balanced IP should belong to
	// for this global forwarding rule. If this field is not specified,
	// the default network will be used.
	Network pulumi.StringPtrInput
	// This field is used along with the target field for TargetHttpProxy,
	// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
	// TargetPool, TargetInstance.
	// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target.
	// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
	// disjoint port ranges.
	// Some types of forwarding target have constraints on the acceptable
	// ports:
	// * TargetHttpProxy: 80, 8080
	// * TargetHttpsProxy: 443
	// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	//   1883, 5222
	// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
	//   1883, 5222
	// * TargetVpnGateway: 500, 4500
	PortRange pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URL of the target resource to receive the matched traffic.
	// The forwarded traffic must be of a type appropriate to the target object.
	// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	// are valid.
	// For global address with a purpose of PRIVATE_SERVICE_CONNECT and
	// addressType of INTERNAL, only "all-apis" and "vpc-sc" are valid.
	Target pulumi.StringInput
}

func (GlobalForwardingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalForwardingRuleArgs)(nil)).Elem()
}

type GlobalForwardingRuleInput interface {
	pulumi.Input

	ToGlobalForwardingRuleOutput() GlobalForwardingRuleOutput
	ToGlobalForwardingRuleOutputWithContext(ctx context.Context) GlobalForwardingRuleOutput
}

func (*GlobalForwardingRule) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalForwardingRule)(nil))
}

func (i *GlobalForwardingRule) ToGlobalForwardingRuleOutput() GlobalForwardingRuleOutput {
	return i.ToGlobalForwardingRuleOutputWithContext(context.Background())
}

func (i *GlobalForwardingRule) ToGlobalForwardingRuleOutputWithContext(ctx context.Context) GlobalForwardingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalForwardingRuleOutput)
}

func (i *GlobalForwardingRule) ToGlobalForwardingRulePtrOutput() GlobalForwardingRulePtrOutput {
	return i.ToGlobalForwardingRulePtrOutputWithContext(context.Background())
}

func (i *GlobalForwardingRule) ToGlobalForwardingRulePtrOutputWithContext(ctx context.Context) GlobalForwardingRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalForwardingRulePtrOutput)
}

type GlobalForwardingRulePtrInput interface {
	pulumi.Input

	ToGlobalForwardingRulePtrOutput() GlobalForwardingRulePtrOutput
	ToGlobalForwardingRulePtrOutputWithContext(ctx context.Context) GlobalForwardingRulePtrOutput
}

type globalForwardingRulePtrType GlobalForwardingRuleArgs

func (*globalForwardingRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalForwardingRule)(nil))
}

func (i *globalForwardingRulePtrType) ToGlobalForwardingRulePtrOutput() GlobalForwardingRulePtrOutput {
	return i.ToGlobalForwardingRulePtrOutputWithContext(context.Background())
}

func (i *globalForwardingRulePtrType) ToGlobalForwardingRulePtrOutputWithContext(ctx context.Context) GlobalForwardingRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalForwardingRulePtrOutput)
}

// GlobalForwardingRuleArrayInput is an input type that accepts GlobalForwardingRuleArray and GlobalForwardingRuleArrayOutput values.
// You can construct a concrete instance of `GlobalForwardingRuleArrayInput` via:
//
//          GlobalForwardingRuleArray{ GlobalForwardingRuleArgs{...} }
type GlobalForwardingRuleArrayInput interface {
	pulumi.Input

	ToGlobalForwardingRuleArrayOutput() GlobalForwardingRuleArrayOutput
	ToGlobalForwardingRuleArrayOutputWithContext(context.Context) GlobalForwardingRuleArrayOutput
}

type GlobalForwardingRuleArray []GlobalForwardingRuleInput

func (GlobalForwardingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GlobalForwardingRule)(nil))
}

func (i GlobalForwardingRuleArray) ToGlobalForwardingRuleArrayOutput() GlobalForwardingRuleArrayOutput {
	return i.ToGlobalForwardingRuleArrayOutputWithContext(context.Background())
}

func (i GlobalForwardingRuleArray) ToGlobalForwardingRuleArrayOutputWithContext(ctx context.Context) GlobalForwardingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalForwardingRuleArrayOutput)
}

// GlobalForwardingRuleMapInput is an input type that accepts GlobalForwardingRuleMap and GlobalForwardingRuleMapOutput values.
// You can construct a concrete instance of `GlobalForwardingRuleMapInput` via:
//
//          GlobalForwardingRuleMap{ "key": GlobalForwardingRuleArgs{...} }
type GlobalForwardingRuleMapInput interface {
	pulumi.Input

	ToGlobalForwardingRuleMapOutput() GlobalForwardingRuleMapOutput
	ToGlobalForwardingRuleMapOutputWithContext(context.Context) GlobalForwardingRuleMapOutput
}

type GlobalForwardingRuleMap map[string]GlobalForwardingRuleInput

func (GlobalForwardingRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GlobalForwardingRule)(nil))
}

func (i GlobalForwardingRuleMap) ToGlobalForwardingRuleMapOutput() GlobalForwardingRuleMapOutput {
	return i.ToGlobalForwardingRuleMapOutputWithContext(context.Background())
}

func (i GlobalForwardingRuleMap) ToGlobalForwardingRuleMapOutputWithContext(ctx context.Context) GlobalForwardingRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalForwardingRuleMapOutput)
}

type GlobalForwardingRuleOutput struct {
	*pulumi.OutputState
}

func (GlobalForwardingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalForwardingRule)(nil))
}

func (o GlobalForwardingRuleOutput) ToGlobalForwardingRuleOutput() GlobalForwardingRuleOutput {
	return o
}

func (o GlobalForwardingRuleOutput) ToGlobalForwardingRuleOutputWithContext(ctx context.Context) GlobalForwardingRuleOutput {
	return o
}

func (o GlobalForwardingRuleOutput) ToGlobalForwardingRulePtrOutput() GlobalForwardingRulePtrOutput {
	return o.ToGlobalForwardingRulePtrOutputWithContext(context.Background())
}

func (o GlobalForwardingRuleOutput) ToGlobalForwardingRulePtrOutputWithContext(ctx context.Context) GlobalForwardingRulePtrOutput {
	return o.ApplyT(func(v GlobalForwardingRule) *GlobalForwardingRule {
		return &v
	}).(GlobalForwardingRulePtrOutput)
}

type GlobalForwardingRulePtrOutput struct {
	*pulumi.OutputState
}

func (GlobalForwardingRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalForwardingRule)(nil))
}

func (o GlobalForwardingRulePtrOutput) ToGlobalForwardingRulePtrOutput() GlobalForwardingRulePtrOutput {
	return o
}

func (o GlobalForwardingRulePtrOutput) ToGlobalForwardingRulePtrOutputWithContext(ctx context.Context) GlobalForwardingRulePtrOutput {
	return o
}

type GlobalForwardingRuleArrayOutput struct{ *pulumi.OutputState }

func (GlobalForwardingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalForwardingRule)(nil))
}

func (o GlobalForwardingRuleArrayOutput) ToGlobalForwardingRuleArrayOutput() GlobalForwardingRuleArrayOutput {
	return o
}

func (o GlobalForwardingRuleArrayOutput) ToGlobalForwardingRuleArrayOutputWithContext(ctx context.Context) GlobalForwardingRuleArrayOutput {
	return o
}

func (o GlobalForwardingRuleArrayOutput) Index(i pulumi.IntInput) GlobalForwardingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GlobalForwardingRule {
		return vs[0].([]GlobalForwardingRule)[vs[1].(int)]
	}).(GlobalForwardingRuleOutput)
}

type GlobalForwardingRuleMapOutput struct{ *pulumi.OutputState }

func (GlobalForwardingRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GlobalForwardingRule)(nil))
}

func (o GlobalForwardingRuleMapOutput) ToGlobalForwardingRuleMapOutput() GlobalForwardingRuleMapOutput {
	return o
}

func (o GlobalForwardingRuleMapOutput) ToGlobalForwardingRuleMapOutputWithContext(ctx context.Context) GlobalForwardingRuleMapOutput {
	return o
}

func (o GlobalForwardingRuleMapOutput) MapIndex(k pulumi.StringInput) GlobalForwardingRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GlobalForwardingRule {
		return vs[0].(map[string]GlobalForwardingRule)[vs[1].(string)]
	}).(GlobalForwardingRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(GlobalForwardingRuleOutput{})
	pulumi.RegisterOutputType(GlobalForwardingRulePtrOutput{})
	pulumi.RegisterOutputType(GlobalForwardingRuleArrayOutput{})
	pulumi.RegisterOutputType(GlobalForwardingRuleMapOutput{})
}
