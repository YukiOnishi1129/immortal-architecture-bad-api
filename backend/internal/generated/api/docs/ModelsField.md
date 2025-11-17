# ModelsField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | フィールドID | 
**Label** | **string** | ラベル | 
**Order** | **int32** | 表示順序 | 
**IsRequired** | **bool** | 必須フラグ | 

## Methods

### NewModelsField

`func NewModelsField(id string, label string, order int32, isRequired bool, ) *ModelsField`

NewModelsField instantiates a new ModelsField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsFieldWithDefaults

`func NewModelsFieldWithDefaults() *ModelsField`

NewModelsFieldWithDefaults instantiates a new ModelsField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsField) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *ModelsField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModelsField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModelsField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetOrder

`func (o *ModelsField) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelsField) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelsField) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetIsRequired

`func (o *ModelsField) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ModelsField) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ModelsField) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


