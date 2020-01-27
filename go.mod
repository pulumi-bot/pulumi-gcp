module github.com/pulumi/pulumi-gcp

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.9.1-0.20200123184238-4d8027f05757
	github.com/pulumi/pulumi-terraform-bridge v1.5.2
	github.com/stretchr/testify v1.4.1-0.20191106224347-f1bd0923b832
	github.com/terraform-providers/terraform-provider-google-beta v0.0.0-00010101000000-000000000000
	github.com/xanzy/ssh-agent v0.2.1 // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-google-beta => github.com/pulumi/terraform-provider-google-beta v0.0.0-20200124105704-9d10a44a3d71
)

replace github.com/pulumi/pulumi-terraform-bridge => ../pulumi-terraform-bridge
