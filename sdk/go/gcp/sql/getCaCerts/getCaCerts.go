// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getCaCerts

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/sql/getCaCertsCert"
)

func GetCaCerts(ctx *pulumi.Context, args *GetCaCertsArgs, opts ...pulumi.InvokeOption) (*GetCaCertsResult, error) {
	var rv GetCaCertsResult
	err := ctx.Invoke("gcp:sql/getCaCerts:getCaCerts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCaCerts.
type GetCaCertsArgs struct {
	Instance string `pulumi:"instance"`
	Project *string `pulumi:"project"`
}


// A collection of values returned by getCaCerts.
type GetCaCertsResult struct {
	ActiveVersion string `pulumi:"activeVersion"`
	Certs []sqlgetCaCertsCert.GetCaCertsCert `pulumi:"certs"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Instance string `pulumi:"instance"`
	Project string `pulumi:"project"`
}

