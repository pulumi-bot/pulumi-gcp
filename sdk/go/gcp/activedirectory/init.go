// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package activedirectory

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
	case "gcp:activedirectory/domain:Domain":
		r, err = NewDomain(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:activedirectory/domainTrust:DomainTrust":
		r, err = NewDomainTrust(ctx, name, nil, pulumi.URN_(urn))
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
		"activedirectory/domain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"activedirectory/domainTrust",
		&module{version},
	)
}
