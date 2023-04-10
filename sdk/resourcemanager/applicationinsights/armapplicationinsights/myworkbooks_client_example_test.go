//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/MyWorkbooksList.json
func ExampleMyWorkbooksClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMyWorkbooksClient().NewListByResourceGroupPager("my-resource-group", armapplicationinsights.CategoryTypeWorkbook, &armapplicationinsights.MyWorkbooksClientListByResourceGroupOptions{Tags: []string{},
		CanFetchContent: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.MyWorkbooksListResult = armapplicationinsights.MyWorkbooksListResult{
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/MyWorkbookGet.json
func ExampleMyWorkbooksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMyWorkbooksClient().Get(ctx, "my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MyWorkbook = armapplicationinsights.MyWorkbook{
	// 	Name: to.Ptr("deadb33f-8bee-4d3b-a059-9be8dac93960"),
	// 	Type: to.Ptr("microsoft.insights/myworkbooks"),
	// 	ID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourceGroups/my-resource-group/providers/microsoft.insights/myworkbooks/deadb33f-8bee-4d3b-a059-9be8dac93960"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"0": to.Ptr("TagSample01"),
	// 		"1": to.Ptr("TagSample02"),
	// 	},
	// 	Kind: to.Ptr(armapplicationinsights.SharedTypeKindUser),
	// 	Properties: &armapplicationinsights.MyWorkbookProperties{
	// 		Category: to.Ptr("workbook"),
	// 		DisplayName: to.Ptr("My New Workbook"),
	// 		SerializedData: to.Ptr("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}"),
	// 		SourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup/providers/Microsoft.Web/sites/MyTestApp-CodeLens"),
	// 		UserID: to.Ptr("userId"),
	// 		Version: to.Ptr("ME"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/MyWorkbookDelete.json
func ExampleMyWorkbooksClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMyWorkbooksClient().Delete(ctx, "my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/MyWorkbookAdd.json
func ExampleMyWorkbooksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMyWorkbooksClient().CreateOrUpdate(ctx, "my-resource-group", "deadb33f-8bee-4d3b-a059-9be8dac93960", armapplicationinsights.MyWorkbook{
		Name:     to.Ptr("deadb33f-8bee-4d3b-a059-9be8dac93960"),
		ID:       to.Ptr("c0deea5e-3344-40f2-96f8-6f8e1c3b5722"),
		Location: to.Ptr("west us"),
		Tags: map[string]*string{
			"0": to.Ptr("TagSample01"),
			"1": to.Ptr("TagSample02"),
		},
		Kind: to.Ptr(armapplicationinsights.SharedTypeKindUser),
		Properties: &armapplicationinsights.MyWorkbookProperties{
			Category:       to.Ptr("workbook"),
			DisplayName:    to.Ptr("Blah Blah Blah"),
			SerializedData: to.Ptr("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}"),
			SourceID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup/providers/Microsoft.Web/sites/MyTestApp-CodeLens"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MyWorkbook = armapplicationinsights.MyWorkbook{
	// 	Name: to.Ptr("deadb33f-8bee-4d3b-a059-9be8dac93960"),
	// 	Type: to.Ptr("microsoft.insights/myworkbooks"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/my-resource-group/providers/microsoft.insights/myworkbooks/deadb33f-8bee-4d3b-a059-9be8dac93960"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"0": to.Ptr("TagSample01"),
	// 		"1": to.Ptr("TagSample02"),
	// 	},
	// 	Kind: to.Ptr(armapplicationinsights.SharedTypeKindUser),
	// 	Properties: &armapplicationinsights.MyWorkbookProperties{
	// 		Category: to.Ptr("workbook"),
	// 		DisplayName: to.Ptr("Blah Blah Blah"),
	// 		SerializedData: to.Ptr("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}"),
	// 		SourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup/providers/Microsoft.Web/sites/MyTestApp-CodeLens"),
	// 		Version: to.Ptr("ME"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/MyWorkbookUpdate.json
func ExampleMyWorkbooksClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMyWorkbooksClient().Update(ctx, "my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", armapplicationinsights.MyWorkbook{
		Name:     to.Ptr("deadb33f-8bee-4d3b-a059-9be8dac93960"),
		Location: to.Ptr("west us"),
		Tags: map[string]*string{
			"0": to.Ptr("TagSample01"),
			"1": to.Ptr("TagSample02"),
		},
		Kind: to.Ptr(armapplicationinsights.SharedTypeKindUser),
		Properties: &armapplicationinsights.MyWorkbookProperties{
			Category:       to.Ptr("workbook"),
			DisplayName:    to.Ptr("Blah Blah Blah"),
			SerializedData: to.Ptr("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}"),
			SourceID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup/providers/Microsoft.Web/sites/MyTestApp-CodeLens"),
			Version:        to.Ptr("ME"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MyWorkbook = armapplicationinsights.MyWorkbook{
	// 	Name: to.Ptr("deadb33f-8bee-4d3b-a059-9be8dac93960"),
	// 	Type: to.Ptr("microsoft.insights/myworkbooks"),
	// 	ID: to.Ptr("c0deea5e-3344-40f2-96f8-6f8e1c3b5722"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"0": to.Ptr("TagSample01"),
	// 		"1": to.Ptr("TagSample02"),
	// 	},
	// 	Kind: to.Ptr(armapplicationinsights.SharedTypeKindUser),
	// 	Properties: &armapplicationinsights.MyWorkbookProperties{
	// 		Category: to.Ptr("workbook"),
	// 		DisplayName: to.Ptr("Blah Blah Blah"),
	// 		SerializedData: to.Ptr("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}"),
	// 		SourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup/providers/Microsoft.Web/sites/MyTestApp-CodeLens"),
	// 		Version: to.Ptr("ME"),
	// 	},
	// }
}
