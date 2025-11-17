# ModelsCreateFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | ラベル | 
**Order** | **int32** | 表示順序 | 
**IsRequired** | **bool** | 必須フラグ | 

## Methods

### NewModelsCreateFieldRequest

`func NewModelsCreateFieldRequest(label string, order int32, isRequired bool, ) *ModelsCreateFieldRequest`

NewModelsCreateFieldRequest instantiates a new ModelsCreateFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCreateFieldRequestWithDefaults

`func NewModelsCreateFieldRequestWithDefaults() *ModelsCreateFieldRequest`

NewModelsCreateFieldRequestWithDefaults instantiates a new ModelsCreateFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ModelsCreateFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModelsCreateFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModelsCreateFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetOrder

`func (o *ModelsCreateFieldRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelsCreateFieldRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelsCreateFieldRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetIsRequired

`func (o *ModelsCreateFieldRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ModelsCreateFieldRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ModelsCreateFieldRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


