# ModelsTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | テンプレートID | 
**Name** | **string** | テンプレート名 | 
**OwnerId** | **string** | 所有者ID | 
**Owner** | [**ModelsAccountSummary**](ModelsAccountSummary.md) | 所有者情報 | 
**Fields** | [**[]ModelsField**](ModelsField.md) | フィールド一覧 | 
**UpdatedAt** | **time.Time** | 更新日時 | 
**IsUsed** | **bool** | 使用中フラグ | 

## Methods

### NewModelsTemplateResponse

`func NewModelsTemplateResponse(id string, name string, ownerId string, owner ModelsAccountSummary, fields []ModelsField, updatedAt time.Time, isUsed bool, ) *ModelsTemplateResponse`

NewModelsTemplateResponse instantiates a new ModelsTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsTemplateResponseWithDefaults

`func NewModelsTemplateResponseWithDefaults() *ModelsTemplateResponse`

NewModelsTemplateResponseWithDefaults instantiates a new ModelsTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsTemplateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ModelsTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *ModelsTemplateResponse) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ModelsTemplateResponse) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ModelsTemplateResponse) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwner

`func (o *ModelsTemplateResponse) GetOwner() ModelsAccountSummary`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelsTemplateResponse) GetOwnerOk() (*ModelsAccountSummary, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelsTemplateResponse) SetOwner(v ModelsAccountSummary)`

SetOwner sets Owner field to given value.


### GetFields

`func (o *ModelsTemplateResponse) GetFields() []ModelsField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ModelsTemplateResponse) GetFieldsOk() (*[]ModelsField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ModelsTemplateResponse) SetFields(v []ModelsField)`

SetFields sets Fields field to given value.


### GetUpdatedAt

`func (o *ModelsTemplateResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsTemplateResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsTemplateResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsUsed

`func (o *ModelsTemplateResponse) GetIsUsed() bool`

GetIsUsed returns the IsUsed field if non-nil, zero value otherwise.

### GetIsUsedOk

`func (o *ModelsTemplateResponse) GetIsUsedOk() (*bool, bool)`

GetIsUsedOk returns a tuple with the IsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsed

`func (o *ModelsTemplateResponse) SetIsUsed(v bool)`

SetIsUsed sets IsUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


