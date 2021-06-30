// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

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
	case "gcp:apigateway/api:Api":
		r = &Api{}
	case "gcp:apigateway/apiConfig:ApiConfig":
		r = &ApiConfig{}
	case "gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding":
		r = &ApiConfigIamBinding{}
	case "gcp:apigateway/apiConfigIamMember:ApiConfigIamMember":
		r = &ApiConfigIamMember{}
	case "gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy":
		r = &ApiConfigIamPolicy{}
	case "gcp:apigateway/apiIamBinding:ApiIamBinding":
		r = &ApiIamBinding{}
	case "gcp:apigateway/apiIamMember:ApiIamMember":
		r = &ApiIamMember{}
	case "gcp:apigateway/apiIamPolicy:ApiIamPolicy":
		r = &ApiIamPolicy{}
	case "gcp:apigateway/gateway:Gateway":
		r = &Gateway{}
	case "gcp:apigateway/gatewayIamBinding:GatewayIamBinding":
		r = &GatewayIamBinding{}
	case "gcp:apigateway/gatewayIamMember:GatewayIamMember":
		r = &GatewayIamMember{}
	case "gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy":
		r = &GatewayIamPolicy{}
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
		"apigateway/api",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/apiConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/apiConfigIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/apiConfigIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/apiConfigIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/apiIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/apiIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/apiIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/gateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/gatewayIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/gatewayIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"apigateway/gatewayIamPolicy",
		&module{version},
	)
}
