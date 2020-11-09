// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gameservices

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Realm resource.
//
// To get more information about Realm, see:
//
// * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.realms)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/game-servers/docs)
//
// ## Example Usage
//
// ## Import
//
// Realm can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:gameservices/realm:Realm default projects/{{project}}/locations/{{location}}/realms/{{realm_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:gameservices/realm:Realm default {{project}}/{{location}}/{{realm_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:gameservices/realm:Realm default {{location}}/{{realm_id}}
// ```
type Realm struct {
	pulumi.CustomResourceState

	// Human readable description of the realm.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ETag of the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The labels associated with this realm. Each label is a key-value pair.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location of the Realm.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource id of the realm, of the form: 'projects/{project_id}/locations/{location}/realms/{realm_id}'. For example,
	// 'projects/my-project/locations/{location}/realms/my-realm'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// GCP region of the Realm.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// Required. Time zone where all realm-specific policies are evaluated. The value of
	// this field must be from the IANA time zone database:
	// https://www.iana.org/time-zones.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
}

// NewRealm registers a new resource with the given unique name, arguments, and options.
func NewRealm(ctx *pulumi.Context,
	name string, args *RealmArgs, opts ...pulumi.ResourceOption) (*Realm, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.TimeZone == nil {
		return nil, errors.New("missing required argument 'TimeZone'")
	}
	if args == nil {
		args = &RealmArgs{}
	}
	var resource Realm
	err := ctx.RegisterResource("gcp:gameservices/realm:Realm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealm gets an existing Realm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmState, opts ...pulumi.ResourceOption) (*Realm, error) {
	var resource Realm
	err := ctx.ReadResource("gcp:gameservices/realm:Realm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Realm resources.
type realmState struct {
	// Human readable description of the realm.
	Description *string `pulumi:"description"`
	// ETag of the resource.
	Etag *string `pulumi:"etag"`
	// The labels associated with this realm. Each label is a key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// Location of the Realm.
	Location *string `pulumi:"location"`
	// The resource id of the realm, of the form: 'projects/{project_id}/locations/{location}/realms/{realm_id}'. For example,
	// 'projects/my-project/locations/{location}/realms/my-realm'.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// GCP region of the Realm.
	RealmId *string `pulumi:"realmId"`
	// Required. Time zone where all realm-specific policies are evaluated. The value of
	// this field must be from the IANA time zone database:
	// https://www.iana.org/time-zones.
	TimeZone *string `pulumi:"timeZone"`
}

type RealmState struct {
	// Human readable description of the realm.
	Description pulumi.StringPtrInput
	// ETag of the resource.
	Etag pulumi.StringPtrInput
	// The labels associated with this realm. Each label is a key-value pair.
	Labels pulumi.StringMapInput
	// Location of the Realm.
	Location pulumi.StringPtrInput
	// The resource id of the realm, of the form: 'projects/{project_id}/locations/{location}/realms/{realm_id}'. For example,
	// 'projects/my-project/locations/{location}/realms/my-realm'.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// GCP region of the Realm.
	RealmId pulumi.StringPtrInput
	// Required. Time zone where all realm-specific policies are evaluated. The value of
	// this field must be from the IANA time zone database:
	// https://www.iana.org/time-zones.
	TimeZone pulumi.StringPtrInput
}

func (RealmState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmState)(nil)).Elem()
}

type realmArgs struct {
	// Human readable description of the realm.
	Description *string `pulumi:"description"`
	// The labels associated with this realm. Each label is a key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// Location of the Realm.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// GCP region of the Realm.
	RealmId string `pulumi:"realmId"`
	// Required. Time zone where all realm-specific policies are evaluated. The value of
	// this field must be from the IANA time zone database:
	// https://www.iana.org/time-zones.
	TimeZone string `pulumi:"timeZone"`
}

// The set of arguments for constructing a Realm resource.
type RealmArgs struct {
	// Human readable description of the realm.
	Description pulumi.StringPtrInput
	// The labels associated with this realm. Each label is a key-value pair.
	Labels pulumi.StringMapInput
	// Location of the Realm.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// GCP region of the Realm.
	RealmId pulumi.StringInput
	// Required. Time zone where all realm-specific policies are evaluated. The value of
	// this field must be from the IANA time zone database:
	// https://www.iana.org/time-zones.
	TimeZone pulumi.StringInput
}

func (RealmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmArgs)(nil)).Elem()
}
