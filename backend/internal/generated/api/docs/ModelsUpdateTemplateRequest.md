# ModelsUpdateTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | テンプレートID | 
**Name** | **string** | テンプレート名 | 
**Fields** | [**[]ModelsUpdateFieldRequest**](ModelsUpdateFieldRequest.md) | フィールド一覧 | 

## Methods

### NewModelsUpdateTemplateRequest

`func NewModelsUpdateTemplateRequest(id string, name string, fields []ModelsUpdateFieldRequest, ) *ModelsUpdateTemplateRequest`

NewModelsUpdateTemplateRequest instantiates a new ModelsUpdateTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsUpdateTemplateRequestWithDefaults

`func NewModelsUpdateTemplateRequestWithDefaults() *ModelsUpdateTemplateRequest`

NewModelsUpdateTemplateRequestWithDefaults instantiates a new ModelsUpdateTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsUpdateTemplateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsUpdateTemplateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsUpdateTemplateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ModelsUpdateTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsUpdateTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsUpdateTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFields

`func (o *ModelsUpdateTemplateRequest) GetFields() []ModelsUpdateFieldRequest`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ModelsUpdateTemplateRequest) GetFieldsOk() (*[]ModelsUpdateFieldRequest, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ModelsUpdateTemplateRequest) SetFields(v []ModelsUpdateFieldRequest)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


