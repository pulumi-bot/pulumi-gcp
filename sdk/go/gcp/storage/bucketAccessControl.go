// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Bucket ACLs can be managed authoritatively using the
// `storageBucketAcl` resource. Do not use these two resources in conjunction to manage the same bucket.
//
// The BucketAccessControls resource manages the Access Control List
// (ACLs) for a single entity/role pairing on a bucket. ACLs let you specify who
// has access to your data and to what extent.
//
// There are three roles that can be assigned to an entity:
//
// READERs can get the bucket, though no acl property will be returned, and
// list the bucket's objects.  WRITERs are READERs, and they can insert
// objects into the bucket and delete the bucket's objects.  OWNERs are
// WRITERs, and they can get the acl property of a bucket, update a bucket,
// and call all BucketAccessControls methods on the bucket.  For more
// information, see Access Control, with the caveat that this API uses
// READER, WRITER, and OWNER instead of READ, WRITE, and FULL_CONTROL.
//
//
// To get more information about BucketAccessControl, see:
//
// * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/storage/docs/access-control/lists)
//
// ## Example Usage
//
// ### Storage Bucket Access Control Public Bucket
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
// 		bucket, err := storage.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBucketAccessControl(ctx, "publicRule", &storage.BucketAccessControlArgs{
// 			Bucket: bucket.Name,
// 			Role:   pulumi.String("READER"),
// 			Entity: pulumi.String("allUsers"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type BucketAccessControl struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The email address associated with the entity.
	Email pulumi.StringOutput `pulumi:"email"`
	// The entity holding the permission, in one of the following forms:
	// user-userId
	// user-email
	// group-groupId
	// group-email
	// domain-domain
	// project-team-projectId
	// allUsers
	// allAuthenticatedUsers
	// Examples:
	// The user liz@example.com would be user-liz@example.com.
	// The group example@googlegroups.com would be
	// group-example@googlegroups.com.
	// To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	Entity pulumi.StringOutput `pulumi:"entity"`
	// The access permission for the entity.
	Role pulumi.StringPtrOutput `pulumi:"role"`
}

// NewBucketAccessControl registers a new resource with the given unique name, arguments, and options.
func NewBucketAccessControl(ctx *pulumi.Context,
	name string, args *BucketAccessControlArgs, opts ...pulumi.ResourceOption) (*BucketAccessControl, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.Entity == nil {
		return nil, errors.New("missing required argument 'Entity'")
	}
	if args == nil {
		args = &BucketAccessControlArgs{}
	}
	var resource BucketAccessControl
	err := ctx.RegisterResource("gcp:storage/bucketAccessControl:BucketAccessControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketAccessControl gets an existing BucketAccessControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketAccessControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketAccessControlState, opts ...pulumi.ResourceOption) (*BucketAccessControl, error) {
	var resource BucketAccessControl
	err := ctx.ReadResource("gcp:storage/bucketAccessControl:BucketAccessControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketAccessControl resources.
type bucketAccessControlState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain *string `pulumi:"domain"`
	// The email address associated with the entity.
	Email *string `pulumi:"email"`
	// The entity holding the permission, in one of the following forms:
	// user-userId
	// user-email
	// group-groupId
	// group-email
	// domain-domain
	// project-team-projectId
	// allUsers
	// allAuthenticatedUsers
	// Examples:
	// The user liz@example.com would be user-liz@example.com.
	// The group example@googlegroups.com would be
	// group-example@googlegroups.com.
	// To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	Entity *string `pulumi:"entity"`
	// The access permission for the entity.
	Role *string `pulumi:"role"`
}

type BucketAccessControlState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The domain associated with the entity.
	Domain pulumi.StringPtrInput
	// The email address associated with the entity.
	Email pulumi.StringPtrInput
	// The entity holding the permission, in one of the following forms:
	// user-userId
	// user-email
	// group-groupId
	// group-email
	// domain-domain
	// project-team-projectId
	// allUsers
	// allAuthenticatedUsers
	// Examples:
	// The user liz@example.com would be user-liz@example.com.
	// The group example@googlegroups.com would be
	// group-example@googlegroups.com.
	// To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	Entity pulumi.StringPtrInput
	// The access permission for the entity.
	Role pulumi.StringPtrInput
}

func (BucketAccessControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketAccessControlState)(nil)).Elem()
}

type bucketAccessControlArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The entity holding the permission, in one of the following forms:
	// user-userId
	// user-email
	// group-groupId
	// group-email
	// domain-domain
	// project-team-projectId
	// allUsers
	// allAuthenticatedUsers
	// Examples:
	// The user liz@example.com would be user-liz@example.com.
	// The group example@googlegroups.com would be
	// group-example@googlegroups.com.
	// To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	Entity string `pulumi:"entity"`
	// The access permission for the entity.
	Role *string `pulumi:"role"`
}

// The set of arguments for constructing a BucketAccessControl resource.
type BucketAccessControlArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The entity holding the permission, in one of the following forms:
	// user-userId
	// user-email
	// group-groupId
	// group-email
	// domain-domain
	// project-team-projectId
	// allUsers
	// allAuthenticatedUsers
	// Examples:
	// The user liz@example.com would be user-liz@example.com.
	// The group example@googlegroups.com would be
	// group-example@googlegroups.com.
	// To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	Entity pulumi.StringInput
	// The access permission for the entity.
	Role pulumi.StringPtrInput
}

func (BucketAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketAccessControlArgs)(nil)).Elem()
}
