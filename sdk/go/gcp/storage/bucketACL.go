// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Authoritatively manages a bucket's ACLs in Google cloud storage service (GCS). For more information see
// [the official documentation](https://cloud.google.com/storage/docs/access-control/lists)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls).
//
// Bucket ACLs can be managed non authoritatively using the `storageBucketAccessControl` resource. Do not use these two resources in conjunction to manage the same bucket.
//
// Permissions can be granted either by ACLs or Cloud IAM policies. In general, permissions granted by Cloud IAM policies do not appear in ACLs, and permissions granted by ACLs do not appear in Cloud IAM policies. The only exception is for ACLs applied directly on a bucket and certain bucket-level Cloud IAM policies, as described in [Cloud IAM relation to ACLs](https://cloud.google.com/storage/docs/access-control/iam#acls).
//
// **NOTE** This resource will not remove the `project-owners-<project_id>` entity from the `OWNER` role.
//
// ## Example Usage
//
// Example creating an ACL on a bucket with one owner, and one reader.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		_, err = storage.NewBucketACL(ctx, "image_store_acl", &storage.BucketACLArgs{
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
//
// ## Import
//
// This resource does not support import.
type BucketACL struct {
	pulumi.CustomResourceState

	// The name of the bucket it applies to.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Configure this ACL to be the default ACL.
	DefaultAcl pulumi.StringPtrOutput `pulumi:"defaultAcl"`
	// The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `roleEntity` is not.
	PredefinedAcl pulumi.StringPtrOutput `pulumi:"predefinedAcl"`
	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefinedAcl` is not.
	RoleEntities pulumi.StringArrayOutput `pulumi:"roleEntities"`
}

// NewBucketACL registers a new resource with the given unique name, arguments, and options.
func NewBucketACL(ctx *pulumi.Context,
	name string, args *BucketACLArgs, opts ...pulumi.ResourceOption) (*BucketACL, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	var resource BucketACL
	err := ctx.RegisterResource("gcp:storage/bucketACL:BucketACL", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketACL gets an existing BucketACL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketACL(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketACLState, opts ...pulumi.ResourceOption) (*BucketACL, error) {
	var resource BucketACL
	err := ctx.ReadResource("gcp:storage/bucketACL:BucketACL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketACL resources.
type bucketACLState struct {
	// The name of the bucket it applies to.
	Bucket *string `pulumi:"bucket"`
	// Configure this ACL to be the default ACL.
	DefaultAcl *string `pulumi:"defaultAcl"`
	// The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `roleEntity` is not.
	PredefinedAcl *string `pulumi:"predefinedAcl"`
	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefinedAcl` is not.
	RoleEntities []string `pulumi:"roleEntities"`
}

type BucketACLState struct {
	// The name of the bucket it applies to.
	Bucket pulumi.StringPtrInput
	// Configure this ACL to be the default ACL.
	DefaultAcl pulumi.StringPtrInput
	// The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `roleEntity` is not.
	PredefinedAcl pulumi.StringPtrInput
	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefinedAcl` is not.
	RoleEntities pulumi.StringArrayInput
}

func (BucketACLState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketACLState)(nil)).Elem()
}

type bucketACLArgs struct {
	// The name of the bucket it applies to.
	Bucket string `pulumi:"bucket"`
	// Configure this ACL to be the default ACL.
	DefaultAcl *string `pulumi:"defaultAcl"`
	// The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `roleEntity` is not.
	PredefinedAcl *string `pulumi:"predefinedAcl"`
	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefinedAcl` is not.
	RoleEntities []string `pulumi:"roleEntities"`
}

// The set of arguments for constructing a BucketACL resource.
type BucketACLArgs struct {
	// The name of the bucket it applies to.
	Bucket pulumi.StringInput
	// Configure this ACL to be the default ACL.
	DefaultAcl pulumi.StringPtrInput
	// The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `roleEntity` is not.
	PredefinedAcl pulumi.StringPtrInput
	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefinedAcl` is not.
	RoleEntities pulumi.StringArrayInput
}

func (BucketACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketACLArgs)(nil)).Elem()
}

type BucketACLInput interface {
	pulumi.Input

	ToBucketACLOutput() BucketACLOutput
	ToBucketACLOutputWithContext(ctx context.Context) BucketACLOutput
}

func (BucketACL) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketACL)(nil)).Elem()
}

func (i BucketACL) ToBucketACLOutput() BucketACLOutput {
	return i.ToBucketACLOutputWithContext(context.Background())
}

func (i BucketACL) ToBucketACLOutputWithContext(ctx context.Context) BucketACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketACLOutput)
}

type BucketACLOutput struct {
	*pulumi.OutputState
}

func (BucketACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketACLOutput)(nil)).Elem()
}

func (o BucketACLOutput) ToBucketACLOutput() BucketACLOutput {
	return o
}

func (o BucketACLOutput) ToBucketACLOutputWithContext(ctx context.Context) BucketACLOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketACLOutput{})
}
