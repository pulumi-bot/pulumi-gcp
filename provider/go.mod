module github.com/pulumi/pulumi-gcp/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.0
	github.com/pulumi/pulumi-terraform-bridge v1.8.4
	github.com/pulumi/pulumi/sdk v1.14.1-0.20200402002223-a0f615ad0938
	github.com/terraform-providers/terraform-provider-google-beta v0.0.0-20200309221941-5fc1579be217
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-google-beta => github.com/pulumi/terraform-provider-google-beta v0.0.0-20200331132738-5ed438dc5b36
	google.golang.org/api => google.golang.org/api v0.16.0
)

replace github.com/pulumi/pulumi-terraform-bridge => ../../pulumi-terraform-bridge
