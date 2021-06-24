// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretmanager

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
	case "gcp:secretmanager/secret:Secret":
		r = &Secret{}
	case "gcp:secretmanager/secretIamBinding:SecretIamBinding":
		r = &SecretIamBinding{}
	case "gcp:secretmanager/secretIamMember:SecretIamMember":
		r = &SecretIamMember{}
	case "gcp:secretmanager/secretIamPolicy:SecretIamPolicy":
		r = &SecretIamPolicy{}
	case "gcp:secretmanager/secretVersion:SecretVersion":
		r = &SecretVersion{}
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
		"secretmanager/secret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"secretmanager/secretIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"secretmanager/secretIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"secretmanager/secretIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"secretmanager/secretVersion",
		&module{version},
	)
}
