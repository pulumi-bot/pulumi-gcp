// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

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
	case "gcp:notebooks/environment:Environment":
		r, err = NewEnvironment(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:notebooks/instance:Instance":
		r, err = NewInstance(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:notebooks/instanceIamBinding:InstanceIamBinding":
		r, err = NewInstanceIamBinding(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:notebooks/instanceIamMember:InstanceIamMember":
		r, err = NewInstanceIamMember(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:notebooks/instanceIamPolicy:InstanceIamPolicy":
		r, err = NewInstanceIamPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:notebooks/location:Location":
		r, err = NewLocation(ctx, name, nil, pulumi.URN_(urn))
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
		"notebooks/environment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"notebooks/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"notebooks/instanceIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"notebooks/instanceIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"notebooks/instanceIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"notebooks/location",
		&module{version},
	)
}