# ModelsNoteFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Q** | Pointer to **string** | タイトルキーワード検索 | [optional] 
**Status** | Pointer to [**ModelsNoteStatus**](ModelsNoteStatus.md) | ステータスフィルター | [optional] 
**TemplateId** | Pointer to **string** | テンプレートIDフィルター | [optional] 
**OwnerId** | Pointer to **string** | 所有者IDフィルター | [optional] 

## Methods

### NewModelsNoteFilters

`func NewModelsNoteFilters() *ModelsNoteFilters`

NewModelsNoteFilters instantiates a new ModelsNoteFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsNoteFiltersWithDefaults

`func NewModelsNoteFiltersWithDefaults() *ModelsNoteFilters`

NewModelsNoteFiltersWithDefaults instantiates a new ModelsNoteFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQ

`func (o *ModelsNoteFilters) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ModelsNoteFilters) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ModelsNoteFilters) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ModelsNoteFilters) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsNoteFilters) GetStatus() ModelsNoteStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsNoteFilters) GetStatusOk() (*ModelsNoteStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsNoteFilters) SetStatus(v ModelsNoteStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsNoteFilters) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateId

`func (o *ModelsNoteFilters) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ModelsNoteFilters) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ModelsNoteFilters) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ModelsNoteFilters) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetOwnerId

`func (o *ModelsNoteFilters) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ModelsNoteFilters) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ModelsNoteFilters) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ModelsNoteFilters) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


