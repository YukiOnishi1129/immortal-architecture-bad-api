# \AccountsAPI

All URIs are relative to *https://api.mini-notion.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsCreateOrGetAccount**](AccountsAPI.md#AccountsCreateOrGetAccount) | **Post** /api/accounts/auth | Create or get account via OAuth
[**AccountsGetAccountById**](AccountsAPI.md#AccountsGetAccountById) | **Get** /api/accounts/{accountId} | Get account by ID
[**AccountsGetCurrentAccount**](AccountsAPI.md#AccountsGetCurrentAccount) | **Get** /api/accounts/me | Get current account



## AccountsCreateOrGetAccount

> ModelsAccountResponse AccountsCreateOrGetAccount(ctx).ModelsCreateOrGetAccountRequest(modelsCreateOrGetAccountRequest).Execute()

Create or get account via OAuth



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
	modelsCreateOrGetAccountRequest := *openapiclient.NewModelsCreateOrGetAccountRequest("Email_example", "Name_example", "Provider_example", "ProviderAccountId_example") // ModelsCreateOrGetAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.AccountsCreateOrGetAccount(context.Background()).ModelsCreateOrGetAccountRequest(modelsCreateOrGetAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsCreateOrGetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsCreateOrGetAccount`: ModelsAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsCreateOrGetAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsCreateOrGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelsCreateOrGetAccountRequest** | [**ModelsCreateOrGetAccountRequest**](ModelsCreateOrGetAccountRequest.md) |  | 

### Return type

[**ModelsAccountResponse**](ModelsAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGetAccountById

> ModelsAccountResponse AccountsGetAccountById(ctx, accountId).Execute()

Get account by ID



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
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.AccountsGetAccountById(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsGetAccountById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsGetAccountById`: ModelsAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsGetAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetAccountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsAccountResponse**](ModelsAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGetCurrentAccount

> ModelsAccountResponse AccountsGetCurrentAccount(ctx).Execute()

Get current account



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.AccountsGetCurrentAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsGetCurrentAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsGetCurrentAccount`: ModelsAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsGetCurrentAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetCurrentAccountRequest struct via the builder pattern


### Return type

[**ModelsAccountResponse**](ModelsAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

