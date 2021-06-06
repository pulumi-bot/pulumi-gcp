// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package certificateauthority

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
	case "gcp:certificateauthority/authority:Authority":
		r = &Authority{}
	case "gcp:certificateauthority/authorityIamBinding:AuthorityIamBinding":
		r = &AuthorityIamBinding{}
	case "gcp:certificateauthority/authorityIamMember:AuthorityIamMember":
		r = &AuthorityIamMember{}
	case "gcp:certificateauthority/authorityIamPolicy:AuthorityIamPolicy":
		r = &AuthorityIamPolicy{}
	case "gcp:certificateauthority/certificate:Certificate":
		r = &Certificate{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version := gcp.PkgVersion()
	pulumi.RegisterResourceModule(
		"gcp",
		"certificateauthority/authority",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"certificateauthority/authorityIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"certificateauthority/authorityIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"certificateauthority/authorityIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"certificateauthority/certificate",
		&module{version},
	)
}
