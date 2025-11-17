# ModelsCreateNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | タイトル | 
**TemplateId** | **string** | テンプレートID | 
**Sections** | Pointer to [**[]ModelsCreateSectionRequest**](ModelsCreateSectionRequest.md) | セクション（オプション） | [optional] 

## Methods

### NewModelsCreateNoteRequest

`func NewModelsCreateNoteRequest(title string, templateId string, ) *ModelsCreateNoteRequest`

NewModelsCreateNoteRequest instantiates a new ModelsCreateNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCreateNoteRequestWithDefaults

`func NewModelsCreateNoteRequestWithDefaults() *ModelsCreateNoteRequest`

NewModelsCreateNoteRequestWithDefaults instantiates a new ModelsCreateNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ModelsCreateNoteRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsCreateNoteRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsCreateNoteRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTemplateId

`func (o *ModelsCreateNoteRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ModelsCreateNoteRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ModelsCreateNoteRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetSections

`func (o *ModelsCreateNoteRequest) GetSections() []ModelsCreateSectionRequest`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ModelsCreateNoteRequest) GetSectionsOk() (*[]ModelsCreateSectionRequest, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ModelsCreateNoteRequest) SetSections(v []ModelsCreateSectionRequest)`

SetSections sets Sections field to given value.

### HasSections

`func (o *ModelsCreateNoteRequest) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


