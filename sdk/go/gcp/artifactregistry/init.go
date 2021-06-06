// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactregistry

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
	case "gcp:artifactregistry/repository:Repository":
		r = &Repository{}
	case "gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding":
		r = &RepositoryIamBinding{}
	case "gcp:artifactregistry/repositoryIamMember:RepositoryIamMember":
		r = &RepositoryIamMember{}
	case "gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy":
		r = &RepositoryIamPolicy{}
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
		"artifactregistry/repository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"artifactregistry/repositoryIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"artifactregistry/repositoryIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"artifactregistry/repositoryIamPolicy",
		&module{version},
	)
}
