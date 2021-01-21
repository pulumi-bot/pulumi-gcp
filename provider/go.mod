module github.com/pulumi/pulumi-gcp/provider/v4

go 1.15

require (
	github.com/hashicorp/terraform-provider-google-beta/v3 v3.51.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.17.1-0.20201219002337-a0c25bd3e147
	github.com/pulumi/pulumi/sdk/v2 v2.15.1-0.20201202214525-260620430c4c
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/hashicorp/terraform-provider-google-beta/v3 => github.com/pulumi/terraform-provider-google-beta/v3 v3.0.0-20210112132741-2735df18e787
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

replace github.com/pulumi/pulumi/pkg/v2 => ../../pulumi/pkg

replace github.com/pulumi/pulumi/sdk/v2 => ../../pulumi/sdk
