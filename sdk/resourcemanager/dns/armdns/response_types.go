//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

// RecordSetsClientCreateOrUpdateResponse contains the response from method RecordSetsClient.CreateOrUpdate.
type RecordSetsClientCreateOrUpdateResponse struct {
	// Describes a DNS record set (a collection of DNS records with the same name and type).
	RecordSet
}

// RecordSetsClientDeleteResponse contains the response from method RecordSetsClient.Delete.
type RecordSetsClientDeleteResponse struct {
	// placeholder for future response values
}

// RecordSetsClientGetResponse contains the response from method RecordSetsClient.Get.
type RecordSetsClientGetResponse struct {
	// Describes a DNS record set (a collection of DNS records with the same name and type).
	RecordSet
}

// RecordSetsClientListAllByDNSZoneResponse contains the response from method RecordSetsClient.NewListAllByDNSZonePager.
type RecordSetsClientListAllByDNSZoneResponse struct {
	// The response to a record set List operation.
	RecordSetListResult
}

// RecordSetsClientListByDNSZoneResponse contains the response from method RecordSetsClient.NewListByDNSZonePager.
type RecordSetsClientListByDNSZoneResponse struct {
	// The response to a record set List operation.
	RecordSetListResult
}

// RecordSetsClientListByTypeResponse contains the response from method RecordSetsClient.NewListByTypePager.
type RecordSetsClientListByTypeResponse struct {
	// The response to a record set List operation.
	RecordSetListResult
}

// RecordSetsClientUpdateResponse contains the response from method RecordSetsClient.Update.
type RecordSetsClientUpdateResponse struct {
	// Describes a DNS record set (a collection of DNS records with the same name and type).
	RecordSet
}

// ResourceReferenceClientGetByTargetResourcesResponse contains the response from method ResourceReferenceClient.GetByTargetResources.
type ResourceReferenceClientGetByTargetResourcesResponse struct {
	// Represents the properties of the Dns Resource Reference Result.
	ResourceReferenceResult
}

// ZonesClientCreateOrUpdateResponse contains the response from method ZonesClient.CreateOrUpdate.
type ZonesClientCreateOrUpdateResponse struct {
	// Describes a DNS zone.
	Zone
}

// ZonesClientDeleteResponse contains the response from method ZonesClient.BeginDelete.
type ZonesClientDeleteResponse struct {
	// placeholder for future response values
}

// ZonesClientGetResponse contains the response from method ZonesClient.Get.
type ZonesClientGetResponse struct {
	// Describes a DNS zone.
	Zone
}

// ZonesClientListByResourceGroupResponse contains the response from method ZonesClient.NewListByResourceGroupPager.
type ZonesClientListByResourceGroupResponse struct {
	// The response to a Zone List or ListAll operation.
	ZoneListResult
}

// ZonesClientListResponse contains the response from method ZonesClient.NewListPager.
type ZonesClientListResponse struct {
	// The response to a Zone List or ListAll operation.
	ZoneListResult
}

// ZonesClientUpdateResponse contains the response from method ZonesClient.Update.
type ZonesClientUpdateResponse struct {
	// Describes a DNS zone.
	Zone
}
