# ModelsUpdateNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ノートID | 
**Title** | **string** | タイトル | 
**Sections** | [**[]ModelsUpdateSectionRequest**](ModelsUpdateSectionRequest.md) | セクション | 

## Methods

### NewModelsUpdateNoteRequest

`func NewModelsUpdateNoteRequest(id string, title string, sections []ModelsUpdateSectionRequest, ) *ModelsUpdateNoteRequest`

NewModelsUpdateNoteRequest instantiates a new ModelsUpdateNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsUpdateNoteRequestWithDefaults

`func NewModelsUpdateNoteRequestWithDefaults() *ModelsUpdateNoteRequest`

NewModelsUpdateNoteRequestWithDefaults instantiates a new ModelsUpdateNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsUpdateNoteRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsUpdateNoteRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsUpdateNoteRequest) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ModelsUpdateNoteRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsUpdateNoteRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsUpdateNoteRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSections

`func (o *ModelsUpdateNoteRequest) GetSections() []ModelsUpdateSectionRequest`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ModelsUpdateNoteRequest) GetSectionsOk() (*[]ModelsUpdateSectionRequest, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ModelsUpdateNoteRequest) SetSections(v []ModelsUpdateSectionRequest)`

SetSections sets Sections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


