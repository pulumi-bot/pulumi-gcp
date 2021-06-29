module github.com/pulumi/pulumi-gcp/provider/v5

go 1.16

require (
	github.com/hashicorp/terraform-provider-google-beta v1.20.1-0.20210315160117-642085ce9b99
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.2.1
	github.com/pulumi/pulumi/sdk/v3 v3.5.2-0.20210623115523-414367963f50
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/hashicorp/terraform-provider-google-beta => github.com/pulumi/terraform-provider-google-beta v0.0.0-20210623102734-5affe8f5a9f6
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

replace github.com/pulumi/pulumi-terraform-bridge/v3 => ../../pulumi-terraform-bridge
