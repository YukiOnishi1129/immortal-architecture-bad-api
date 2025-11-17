# ModelsUpdateFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | フィールドID（既存フィールドの場合は必須） | [optional] 
**Label** | **string** | ラベル | 
**Order** | **int32** | 表示順序 | 
**IsRequired** | **bool** | 必須フラグ | 

## Methods

### NewModelsUpdateFieldRequest

`func NewModelsUpdateFieldRequest(label string, order int32, isRequired bool, ) *ModelsUpdateFieldRequest`

NewModelsUpdateFieldRequest instantiates a new ModelsUpdateFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsUpdateFieldRequestWithDefaults

`func NewModelsUpdateFieldRequestWithDefaults() *ModelsUpdateFieldRequest`

NewModelsUpdateFieldRequestWithDefaults instantiates a new ModelsUpdateFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsUpdateFieldRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsUpdateFieldRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsUpdateFieldRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsUpdateFieldRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ModelsUpdateFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModelsUpdateFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModelsUpdateFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetOrder

`func (o *ModelsUpdateFieldRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelsUpdateFieldRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelsUpdateFieldRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetIsRequired

`func (o *ModelsUpdateFieldRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ModelsUpdateFieldRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ModelsUpdateFieldRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


