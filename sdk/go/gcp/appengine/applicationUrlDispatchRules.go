// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Rules to match an HTTP request and dispatch that request to a service.
//
//
// To get more information about ApplicationUrlDispatchRules, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#UrlDispatchRule)
type ApplicationUrlDispatchRules struct {
	pulumi.CustomResourceState

	// -
	// (Required)
	// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
	DispatchRules ApplicationUrlDispatchRulesDispatchRuleArrayOutput `pulumi:"dispatchRules"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewApplicationUrlDispatchRules registers a new resource with the given unique name, arguments, and options.
func NewApplicationUrlDispatchRules(ctx *pulumi.Context,
	name string, args *ApplicationUrlDispatchRulesArgs, opts ...pulumi.ResourceOption) (*ApplicationUrlDispatchRules, error) {
	if args == nil || args.DispatchRules == nil {
		return nil, errors.New("missing required argument 'DispatchRules'")
	}
	if args == nil {
		args = &ApplicationUrlDispatchRulesArgs{}
	}
	var resource ApplicationUrlDispatchRules
	err := ctx.RegisterResource("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationUrlDispatchRules gets an existing ApplicationUrlDispatchRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationUrlDispatchRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationUrlDispatchRulesState, opts ...pulumi.ResourceOption) (*ApplicationUrlDispatchRules, error) {
	var resource ApplicationUrlDispatchRules
	err := ctx.ReadResource("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationUrlDispatchRules resources.
type applicationUrlDispatchRulesState struct {
	// -
	// (Required)
	// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
	DispatchRules []ApplicationUrlDispatchRulesDispatchRule `pulumi:"dispatchRules"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ApplicationUrlDispatchRulesState struct {
	// -
	// (Required)
	// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
	DispatchRules ApplicationUrlDispatchRulesDispatchRuleArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApplicationUrlDispatchRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationUrlDispatchRulesState)(nil)).Elem()
}

type applicationUrlDispatchRulesArgs struct {
	// -
	// (Required)
	// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
	DispatchRules []ApplicationUrlDispatchRulesDispatchRule `pulumi:"dispatchRules"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ApplicationUrlDispatchRules resource.
type ApplicationUrlDispatchRulesArgs struct {
	// -
	// (Required)
	// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
	DispatchRules ApplicationUrlDispatchRulesDispatchRuleArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApplicationUrlDispatchRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationUrlDispatchRulesArgs)(nil)).Elem()
}
