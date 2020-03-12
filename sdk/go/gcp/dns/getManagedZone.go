// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package dns

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides access to a zone's attributes within Google Cloud DNS.
// For more information see
// [the official documentation](https://cloud.google.com/dns/zones/)
// and
// [API](https://cloud.google.com/dns/api/v1/managedZones).
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/dns_managed_zone.html.markdown.
func LookupManagedZone(ctx *pulumi.Context, args *LookupManagedZoneArgs, opts ...pulumi.InvokeOption) (*LookupManagedZoneResult, error) {
	var rv LookupManagedZoneResult
	err := ctx.Invoke("gcp:dns/getManagedZone:getManagedZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getManagedZone.
type LookupManagedZoneArgs struct {
	// A unique name for the resource.
	Name string `pulumi:"name"`
	// The ID of the project for the Google Cloud DNS zone.
	Project *string `pulumi:"project"`
}


// A collection of values returned by getManagedZone.
type LookupManagedZoneResult struct {
	// A textual description field.
	Description string `pulumi:"description"`
	// The fully qualified DNS name of this zone, e.g. `example.com.`.
	DnsName string `pulumi:"dnsName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The list of nameservers that will be authoritative for this
	// domain. Use NS records to redirect from your DNS provider to these names,
	// thus making Google Cloud DNS authoritative for this zone.
	NameServers []string `pulumi:"nameServers"`
	Project *string `pulumi:"project"`
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources.
	Visibility string `pulumi:"visibility"`
}

