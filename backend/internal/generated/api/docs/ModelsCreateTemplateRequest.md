# ModelsCreateTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | テンプレート名 | 
**OwnerId** | **string** | 所有者ID | 
**Fields** | [**[]ModelsCreateFieldRequest**](ModelsCreateFieldRequest.md) | フィールド一覧 | 

## Methods

### NewModelsCreateTemplateRequest

`func NewModelsCreateTemplateRequest(name string, ownerId string, fields []ModelsCreateFieldRequest, ) *ModelsCreateTemplateRequest`

NewModelsCreateTemplateRequest instantiates a new ModelsCreateTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCreateTemplateRequestWithDefaults

`func NewModelsCreateTemplateRequestWithDefaults() *ModelsCreateTemplateRequest`

NewModelsCreateTemplateRequestWithDefaults instantiates a new ModelsCreateTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelsCreateTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsCreateTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsCreateTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *ModelsCreateTemplateRequest) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ModelsCreateTemplateRequest) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ModelsCreateTemplateRequest) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetFields

`func (o *ModelsCreateTemplateRequest) GetFields() []ModelsCreateFieldRequest`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ModelsCreateTemplateRequest) GetFieldsOk() (*[]ModelsCreateFieldRequest, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ModelsCreateTemplateRequest) SetFields(v []ModelsCreateFieldRequest)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


