// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Authoritatively manages the default object ACLs for a Google Cloud Storage bucket
// without managing the bucket itself.
//
// > Note that for each object, its creator will have the `"OWNER"` role in addition
// to the default ACL that has been defined.
//
// For more information see
// [the official documentation](https://cloud.google.com/storage/docs/access-control/lists)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/defaultObjectAccessControls).
//
// > Want fine-grained control over default object ACLs? Use `storage.DefaultObjectAccessControl`
// to control individual role entity pairs.
//
// ## Example Usage
//
// Example creating a default object ACL on a bucket with one owner, and one reader.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storage.NewBucket(ctx, "image_store", &storage.BucketArgs{
// 			Location: pulumi.String("EU"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewDefaultObjectACL(ctx, "image_store_default_acl", &storage.DefaultObjectACLArgs{
// 			Bucket: image_store.Name,
// 			RoleEntities: pulumi.StringArray{
// 				pulumi.String("OWNER:user-my.email@gmail.com"),
// 				pulumi.String("READER:group-mygroup"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DefaultObjectACL struct {
	pulumi.CustomResourceState

	// The name of the bucket it applies to.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities pulumi.StringArrayOutput `pulumi:"roleEntities"`
}

// NewDefaultObjectACL registers a new resource with the given unique name, arguments, and options.
func NewDefaultObjectACL(ctx *pulumi.Context,
	name string, args *DefaultObjectACLArgs, opts ...pulumi.ResourceOption) (*DefaultObjectACL, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil {
		args = &DefaultObjectACLArgs{}
	}
	var resource DefaultObjectACL
	err := ctx.RegisterResource("gcp:storage/defaultObjectACL:DefaultObjectACL", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultObjectACL gets an existing DefaultObjectACL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultObjectACL(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultObjectACLState, opts ...pulumi.ResourceOption) (*DefaultObjectACL, error) {
	var resource DefaultObjectACL
	err := ctx.ReadResource("gcp:storage/defaultObjectACL:DefaultObjectACL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultObjectACL resources.
type defaultObjectACLState struct {
	// The name of the bucket it applies to.
	Bucket *string `pulumi:"bucket"`
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities []string `pulumi:"roleEntities"`
}

type DefaultObjectACLState struct {
	// The name of the bucket it applies to.
	Bucket pulumi.StringPtrInput
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities pulumi.StringArrayInput
}

func (DefaultObjectACLState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultObjectACLState)(nil)).Elem()
}

type defaultObjectACLArgs struct {
	// The name of the bucket it applies to.
	Bucket string `pulumi:"bucket"`
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities []string `pulumi:"roleEntities"`
}

// The set of arguments for constructing a DefaultObjectACL resource.
type DefaultObjectACLArgs struct {
	// The name of the bucket it applies to.
	Bucket pulumi.StringInput
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities pulumi.StringArrayInput
}

func (DefaultObjectACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultObjectACLArgs)(nil)).Elem()
}
