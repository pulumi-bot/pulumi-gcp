module github.com/pulumi/pulumi-gcp

go 1.12

require (
	git.apache.org/thrift.git v0.12.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/hashicorp/terraform v0.12.2
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v0.17.22-0.20190702185104-ebceea93a5da
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190709052202-629f7c54269d
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-google-beta v0.0.0-20190702191547-069d14be9e4e
	labix.org/v2/mgo v0.0.0-20140701140051-000000000287 // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)

replace github.com/pulumi/pulumi-terraform => ../pulumi-terraform/
