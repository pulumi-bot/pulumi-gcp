// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The DefaultObjectAccessControls resources represent the Access Control
// Lists (ACLs) applied to a new object within a Google Cloud Storage bucket
// when no ACL was provided for that object. ACLs let you specify who has
// access to your bucket contents and to what extent.
//
// There are two roles that can be assigned to an entity:
//
// READERs can get an object, though the acl property will not be revealed.
// OWNERs are READERs, and they can get the acl property, update an object,
// and call all objectAccessControls methods on the object. The owner of an
// object is always an OWNER.
// For more information, see Access Control, with the caveat that this API
// uses READER and OWNER instead of READ and FULL_CONTROL.
//
// To get more information about DefaultObjectAccessControl, see:
//
// * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/defaultObjectAccessControls)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/storage/docs/access-control/create-manage-lists)
//
// ## Example Usage
// ### Storage Default Object Access Control Public
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
// 		_, err = storage.NewDefaultObjectAccessControl(ctx, "publicRule", &storage.DefaultObjectAccessControlArgs{
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
type DefaultObjectAccessControl struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The email address associated with the entity.
	Email pulumi.StringOutput `pulumi:"email"`
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	Entity pulumi.StringOutput `pulumi:"entity"`
	// The ID for the entity
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// The content generation of the object, if applied to an object.
	Generation pulumi.IntOutput `pulumi:"generation"`
	// The name of the object, if applied to an object.
	Object pulumi.StringPtrOutput `pulumi:"object"`
	// The project team associated with the entity
	ProjectTeam DefaultObjectAccessControlProjectTeamOutput `pulumi:"projectTeam"`
	// The access permission for the entity.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDefaultObjectAccessControl registers a new resource with the given unique name, arguments, and options.
func NewDefaultObjectAccessControl(ctx *pulumi.Context,
	name string, args *DefaultObjectAccessControlArgs, opts ...pulumi.ResourceOption) (*DefaultObjectAccessControl, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.Entity == nil {
		return nil, errors.New("missing required argument 'Entity'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &DefaultObjectAccessControlArgs{}
	}
	var resource DefaultObjectAccessControl
	err := ctx.RegisterResource("gcp:storage/defaultObjectAccessControl:DefaultObjectAccessControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultObjectAccessControl gets an existing DefaultObjectAccessControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultObjectAccessControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultObjectAccessControlState, opts ...pulumi.ResourceOption) (*DefaultObjectAccessControl, error) {
	var resource DefaultObjectAccessControl
	err := ctx.ReadResource("gcp:storage/defaultObjectAccessControl:DefaultObjectAccessControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultObjectAccessControl resources.
type defaultObjectAccessControlState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain *string `pulumi:"domain"`
	// The email address associated with the entity.
	Email *string `pulumi:"email"`
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	Entity *string `pulumi:"entity"`
	// The ID for the entity
	EntityId *string `pulumi:"entityId"`
	// The content generation of the object, if applied to an object.
	Generation *int `pulumi:"generation"`
	// The name of the object, if applied to an object.
	Object *string `pulumi:"object"`
	// The project team associated with the entity
	ProjectTeam *DefaultObjectAccessControlProjectTeam `pulumi:"projectTeam"`
	// The access permission for the entity.
	Role *string `pulumi:"role"`
}

type DefaultObjectAccessControlState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The domain associated with the entity.
	Domain pulumi.StringPtrInput
	// The email address associated with the entity.
	Email pulumi.StringPtrInput
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	Entity pulumi.StringPtrInput
	// The ID for the entity
	EntityId pulumi.StringPtrInput
	// The content generation of the object, if applied to an object.
	Generation pulumi.IntPtrInput
	// The name of the object, if applied to an object.
	Object pulumi.StringPtrInput
	// The project team associated with the entity
	ProjectTeam DefaultObjectAccessControlProjectTeamPtrInput
	// The access permission for the entity.
	Role pulumi.StringPtrInput
}

func (DefaultObjectAccessControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultObjectAccessControlState)(nil)).Elem()
}

type defaultObjectAccessControlArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	Entity string `pulumi:"entity"`
	// The name of the object, if applied to an object.
	Object *string `pulumi:"object"`
	// The access permission for the entity.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DefaultObjectAccessControl resource.
type DefaultObjectAccessControlArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	Entity pulumi.StringInput
	// The name of the object, if applied to an object.
	Object pulumi.StringPtrInput
	// The access permission for the entity.
	Role pulumi.StringInput
}

func (DefaultObjectAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultObjectAccessControlArgs)(nil)).Elem()
}
