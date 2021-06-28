// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

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
	case "gcp:cloudfunctions/function:Function":
		r = &Function{}
	case "gcp:cloudfunctions/functionIamBinding:FunctionIamBinding":
		r = &FunctionIamBinding{}
	case "gcp:cloudfunctions/functionIamMember:FunctionIamMember":
		r = &FunctionIamMember{}
	case "gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy":
		r = &FunctionIamPolicy{}
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
		"cloudfunctions/function",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"cloudfunctions/functionIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"cloudfunctions/functionIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"cloudfunctions/functionIamPolicy",
		&module{version},
	)
}
