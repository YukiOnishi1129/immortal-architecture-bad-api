# ModelsErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | エラーコード | 
**Message** | **string** | エラーメッセージ | 
**Details** | Pointer to **interface{}** | 詳細情報（オプション） | [optional] 

## Methods

### NewModelsErrorResponse

`func NewModelsErrorResponse(code string, message string, ) *ModelsErrorResponse`

NewModelsErrorResponse instantiates a new ModelsErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsErrorResponseWithDefaults

`func NewModelsErrorResponseWithDefaults() *ModelsErrorResponse`

NewModelsErrorResponseWithDefaults instantiates a new ModelsErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ModelsErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModelsErrorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModelsErrorResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ModelsErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelsErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelsErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *ModelsErrorResponse) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ModelsErrorResponse) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ModelsErrorResponse) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ModelsErrorResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ModelsErrorResponse) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ModelsErrorResponse) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


