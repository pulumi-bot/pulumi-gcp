module github.com/pulumi/pulumi-gcp

go 1.12

require (
	github.com/hashicorp/terraform-plugin-sdk v1.3.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.2.0
	github.com/pulumi/pulumi-terraform-bridge v1.2.1-0.20191106001426-e727be072e34
	github.com/stretchr/testify v1.4.0
	github.com/terraform-providers/terraform-provider-google-beta v0.0.0-00010101000000-000000000000
	github.com/xanzy/ssh-agent v0.2.1 // indirect
)

replace github.com/terraform-providers/terraform-provider-google-beta => github.com/pulumi/terraform-provider-google-beta v0.0.0-20191114125340-2b63c0f7957e

replace github.com/pulumi/pulumi-terraform-bridge => ../pulumi-terraform-bridge
