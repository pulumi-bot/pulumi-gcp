// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp/sdk/v2/go/gcp"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "gcp:appengine/application:Application":
		r, err = NewApplication(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules":
		r, err = NewApplicationUrlDispatchRules(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:appengine/domainMapping:DomainMapping":
		r, err = NewDomainMapping(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:appengine/engineSplitTraffic:EngineSplitTraffic":
		r, err = NewEngineSplitTraffic(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:appengine/firewallRule:FirewallRule":
		r, err = NewFirewallRule(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:appengine/flexibleAppVersion:FlexibleAppVersion":
		r, err = NewFlexibleAppVersion(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:appengine/standardAppVersion:StandardAppVersion":
		r, err = NewStandardAppVersion(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := gcp.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"gcp",
		"appengine/application",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"appengine/applicationUrlDispatchRules",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"appengine/domainMapping",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"appengine/engineSplitTraffic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"appengine/firewallRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"appengine/flexibleAppVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"appengine/standardAppVersion",
		&module{version},
	)
}
