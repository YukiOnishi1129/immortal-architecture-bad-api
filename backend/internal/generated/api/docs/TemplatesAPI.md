# \TemplatesAPI

All URIs are relative to *https://api.mini-notion.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplatesCreateTemplate**](TemplatesAPI.md#TemplatesCreateTemplate) | **Post** /api/templates | Create template
[**TemplatesDeleteTemplate**](TemplatesAPI.md#TemplatesDeleteTemplate) | **Delete** /api/templates/{templateId} | Delete template
[**TemplatesGetTemplateById**](TemplatesAPI.md#TemplatesGetTemplateById) | **Get** /api/templates/{templateId} | Get template by ID
[**TemplatesListTemplates**](TemplatesAPI.md#TemplatesListTemplates) | **Get** /api/templates | Get templates list
[**TemplatesUpdateTemplate**](TemplatesAPI.md#TemplatesUpdateTemplate) | **Put** /api/templates/{templateId} | Update template



## TemplatesCreateTemplate

> ModelsTemplateResponse TemplatesCreateTemplate(ctx).ModelsCreateTemplateRequest(modelsCreateTemplateRequest).Execute()

Create template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mini-notion/mini-notion-api"
)

func main() {
	modelsCreateTemplateRequest := *openapiclient.NewModelsCreateTemplateRequest("Name_example", "OwnerId_example", []openapiclient.ModelsCreateFieldRequest{*openapiclient.NewModelsCreateFieldRequest("Label_example", int32(123), false)}) // ModelsCreateTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.TemplatesCreateTemplate(context.Background()).ModelsCreateTemplateRequest(modelsCreateTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.TemplatesCreateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemplatesCreateTemplate`: ModelsTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.TemplatesCreateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelsCreateTemplateRequest** | [**ModelsCreateTemplateRequest**](ModelsCreateTemplateRequest.md) |  | 

### Return type

[**ModelsTemplateResponse**](ModelsTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesDeleteTemplate

> ModelsSuccessResponse TemplatesDeleteTemplate(ctx, templateId).Execute()

Delete template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mini-notion/mini-notion-api"
)

func main() {
	templateId := "templateId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.TemplatesDeleteTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.TemplatesDeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemplatesDeleteTemplate`: ModelsSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.TemplatesDeleteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsSuccessResponse**](ModelsSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesGetTemplateById

> ModelsTemplateResponse TemplatesGetTemplateById(ctx, templateId).Execute()

Get template by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mini-notion/mini-notion-api"
)

func main() {
	templateId := "templateId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.TemplatesGetTemplateById(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.TemplatesGetTemplateById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemplatesGetTemplateById`: ModelsTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.TemplatesGetTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsTemplateResponse**](ModelsTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesListTemplates

> []ModelsTemplateResponse TemplatesListTemplates(ctx).Q(q).OwnerId(ownerId).Execute()

Get templates list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mini-notion/mini-notion-api"
)

func main() {
	q := "q_example" // string | テンプレート名のキーワード検索 (optional)
	ownerId := "ownerId_example" // string | 所有者IDフィルター (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.TemplatesListTemplates(context.Background()).Q(q).OwnerId(ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.TemplatesListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemplatesListTemplates`: []ModelsTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.TemplatesListTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesListTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | テンプレート名のキーワード検索 | 
 **ownerId** | **string** | 所有者IDフィルター | 

### Return type

[**[]ModelsTemplateResponse**](ModelsTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesUpdateTemplate

> ModelsTemplateResponse TemplatesUpdateTemplate(ctx, templateId).ModelsUpdateTemplateRequest(modelsUpdateTemplateRequest).Execute()

Update template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mini-notion/mini-notion-api"
)

func main() {
	templateId := "templateId_example" // string | 
	modelsUpdateTemplateRequest := *openapiclient.NewModelsUpdateTemplateRequest("Id_example", "Name_example", []openapiclient.ModelsUpdateFieldRequest{*openapiclient.NewModelsUpdateFieldRequest("Label_example", int32(123), false)}) // ModelsUpdateTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.TemplatesUpdateTemplate(context.Background(), templateId).ModelsUpdateTemplateRequest(modelsUpdateTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.TemplatesUpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemplatesUpdateTemplate`: ModelsTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.TemplatesUpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelsUpdateTemplateRequest** | [**ModelsUpdateTemplateRequest**](ModelsUpdateTemplateRequest.md) |  | 

### Return type

[**ModelsTemplateResponse**](ModelsTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

