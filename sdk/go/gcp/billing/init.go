// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp"
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
	case "gcp:billing/accountIamBinding:AccountIamBinding":
		r = &AccountIamBinding{}
	case "gcp:billing/accountIamMember:AccountIamMember":
		r = &AccountIamMember{}
	case "gcp:billing/accountIamPolicy:AccountIamPolicy":
		r = &AccountIamPolicy{}
	case "gcp:billing/budget:Budget":
		r = &Budget{}
	case "gcp:billing/subAccount:SubAccount":
		r = &SubAccount{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := gcp.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"gcp",
		"billing/accountIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"billing/accountIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"billing/accountIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"billing/budget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"billing/subAccount",
		&module{version},
	)
}
