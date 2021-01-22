// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicedirectory

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
	case "gcp:servicedirectory/endpoint:Endpoint":
		r, err = NewEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:servicedirectory/namespace:Namespace":
		r, err = NewNamespace(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding":
		r, err = NewNamespaceIamBinding(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:servicedirectory/namespaceIamMember:NamespaceIamMember":
		r, err = NewNamespaceIamMember(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:servicedirectory/namespaceIamPolicy:NamespaceIamPolicy":
		r, err = NewNamespaceIamPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:servicedirectory/service:Service":
		r, err = NewService(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:servicedirectory/serviceIamBinding:ServiceIamBinding":
		r, err = NewServiceIamBinding(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:servicedirectory/serviceIamMember:ServiceIamMember":
		r, err = NewServiceIamMember(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:servicedirectory/serviceIamPolicy:ServiceIamPolicy":
		r, err = NewServiceIamPolicy(ctx, name, nil, pulumi.URN_(urn))
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
		"servicedirectory/endpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"servicedirectory/namespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"servicedirectory/namespaceIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"servicedirectory/namespaceIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"servicedirectory/namespaceIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"servicedirectory/service",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"servicedirectory/serviceIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"servicedirectory/serviceIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"servicedirectory/serviceIamPolicy",
		&module{version},
	)
}
