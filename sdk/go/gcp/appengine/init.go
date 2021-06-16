// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
		r = &Application{}
	case "gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules":
		r = &ApplicationUrlDispatchRules{}
	case "gcp:appengine/domainMapping:DomainMapping":
		r = &DomainMapping{}
	case "gcp:appengine/engineSplitTraffic:EngineSplitTraffic":
		r = &EngineSplitTraffic{}
	case "gcp:appengine/firewallRule:FirewallRule":
		r = &FirewallRule{}
	case "gcp:appengine/flexibleAppVersion:FlexibleAppVersion":
		r = &FlexibleAppVersion{}
	case "gcp:appengine/standardAppVersion:StandardAppVersion":
		r = &StandardAppVersion{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := gcp.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
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
