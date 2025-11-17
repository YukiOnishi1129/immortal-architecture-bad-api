# \NotesAPI

All URIs are relative to *https://api.mini-notion.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotesCreateNote**](NotesAPI.md#NotesCreateNote) | **Post** /api/notes | Create note
[**NotesDeleteNote**](NotesAPI.md#NotesDeleteNote) | **Delete** /api/notes/{noteId} | Delete note
[**NotesGetNoteById**](NotesAPI.md#NotesGetNoteById) | **Get** /api/notes/{noteId} | Get note by ID
[**NotesListNotes**](NotesAPI.md#NotesListNotes) | **Get** /api/notes | Get notes list
[**NotesPublishNote**](NotesAPI.md#NotesPublishNote) | **Post** /api/notes/{noteId}/publish | Publish note
[**NotesUnpublishNote**](NotesAPI.md#NotesUnpublishNote) | **Post** /api/notes/{noteId}/unpublish | Unpublish note
[**NotesUpdateNote**](NotesAPI.md#NotesUpdateNote) | **Put** /api/notes/{noteId} | Update note



## NotesCreateNote

> ModelsNoteResponse NotesCreateNote(ctx).ModelsCreateNoteRequest(modelsCreateNoteRequest).Execute()

Create note



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
	modelsCreateNoteRequest := *openapiclient.NewModelsCreateNoteRequest("Title_example", "TemplateId_example") // ModelsCreateNoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesCreateNote(context.Background()).ModelsCreateNoteRequest(modelsCreateNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesCreateNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesCreateNote`: ModelsNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesCreateNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotesCreateNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelsCreateNoteRequest** | [**ModelsCreateNoteRequest**](ModelsCreateNoteRequest.md) |  | 

### Return type

[**ModelsNoteResponse**](ModelsNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesDeleteNote

> ModelsSuccessResponse NotesDeleteNote(ctx, noteId).Execute()

Delete note



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
	noteId := "noteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesDeleteNote(context.Background(), noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesDeleteNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesDeleteNote`: ModelsSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesDeleteNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesDeleteNoteRequest struct via the builder pattern


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


## NotesGetNoteById

> ModelsNoteResponse NotesGetNoteById(ctx, noteId).Execute()

Get note by ID



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
	noteId := "noteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesGetNoteById(context.Background(), noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesGetNoteById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesGetNoteById`: ModelsNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesGetNoteById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesGetNoteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsNoteResponse**](ModelsNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesListNotes

> []ModelsNoteResponse NotesListNotes(ctx).Q(q).Status(status).TemplateId(templateId).OwnerId(ownerId).Execute()

Get notes list



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
	q := "q_example" // string | タイトルキーワード検索 (optional)
	status := openapiclient.Models.NoteStatus("Draft") // ModelsNoteStatus | ステータスフィルター (optional)
	templateId := "templateId_example" // string | テンプレートIDフィルター (optional)
	ownerId := "ownerId_example" // string | 所有者IDフィルター (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesListNotes(context.Background()).Q(q).Status(status).TemplateId(templateId).OwnerId(ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesListNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesListNotes`: []ModelsNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesListNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotesListNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | タイトルキーワード検索 | 
 **status** | [**ModelsNoteStatus**](ModelsNoteStatus.md) | ステータスフィルター | 
 **templateId** | **string** | テンプレートIDフィルター | 
 **ownerId** | **string** | 所有者IDフィルター | 

### Return type

[**[]ModelsNoteResponse**](ModelsNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesPublishNote

> ModelsNoteResponse NotesPublishNote(ctx, noteId).Execute()

Publish note



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
	noteId := "noteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesPublishNote(context.Background(), noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesPublishNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesPublishNote`: ModelsNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesPublishNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesPublishNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsNoteResponse**](ModelsNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesUnpublishNote

> ModelsNoteResponse NotesUnpublishNote(ctx, noteId).Execute()

Unpublish note



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
	noteId := "noteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesUnpublishNote(context.Background(), noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesUnpublishNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesUnpublishNote`: ModelsNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesUnpublishNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesUnpublishNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsNoteResponse**](ModelsNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesUpdateNote

> ModelsNoteResponse NotesUpdateNote(ctx, noteId).ModelsUpdateNoteRequest(modelsUpdateNoteRequest).Execute()

Update note



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
	noteId := "noteId_example" // string | 
	modelsUpdateNoteRequest := *openapiclient.NewModelsUpdateNoteRequest("Id_example", "Title_example", []openapiclient.ModelsUpdateSectionRequest{*openapiclient.NewModelsUpdateSectionRequest("Id_example", "Content_example")}) // ModelsUpdateNoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesUpdateNote(context.Background(), noteId).ModelsUpdateNoteRequest(modelsUpdateNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesUpdateNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesUpdateNote`: ModelsNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesUpdateNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesUpdateNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelsUpdateNoteRequest** | [**ModelsUpdateNoteRequest**](ModelsUpdateNoteRequest.md) |  | 

### Return type

[**ModelsNoteResponse**](ModelsNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

