module github.com/pulumi/pulumi-gcp/provider/v3

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.11.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.5.4
	github.com/pulumi/pulumi/sdk/v2 v2.5.1-0.20200701223250-45d2fa95d60b
	github.com/terraform-providers/terraform-provider-google-beta v0.0.0-20200309221941-5fc1579be217
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-google-beta => github.com/pulumi/terraform-provider-google-beta v0.0.0-20200804214836-39302eaa6cb2
	google.golang.org/api => google.golang.org/api v0.26.0
)

replace github.com/pulumi/pulumi/pkg/v2 => ../../pulumi/pkg

replace github.com/pulumi/pulumi/sdk/v2 => ../../pulumi/sdk
