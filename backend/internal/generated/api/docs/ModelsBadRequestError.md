# ModelsBadRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Message** | **string** |  | 
**Details** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewModelsBadRequestError

`func NewModelsBadRequestError(code string, message string, ) *ModelsBadRequestError`

NewModelsBadRequestError instantiates a new ModelsBadRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsBadRequestErrorWithDefaults

`func NewModelsBadRequestErrorWithDefaults() *ModelsBadRequestError`

NewModelsBadRequestErrorWithDefaults instantiates a new ModelsBadRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ModelsBadRequestError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModelsBadRequestError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModelsBadRequestError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ModelsBadRequestError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelsBadRequestError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelsBadRequestError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *ModelsBadRequestError) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ModelsBadRequestError) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ModelsBadRequestError) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ModelsBadRequestError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ModelsBadRequestError) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ModelsBadRequestError) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


