# ModelsCreateOrGetAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | メールアドレス | 
**Name** | **string** | 名前 | 
**Provider** | **string** | プロバイダー（例: google） | 
**ProviderAccountId** | **string** | プロバイダーのアカウントID | 
**Thumbnail** | Pointer to **string** | プロフィール画像URL | [optional] 

## Methods

### NewModelsCreateOrGetAccountRequest

`func NewModelsCreateOrGetAccountRequest(email string, name string, provider string, providerAccountId string, ) *ModelsCreateOrGetAccountRequest`

NewModelsCreateOrGetAccountRequest instantiates a new ModelsCreateOrGetAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCreateOrGetAccountRequestWithDefaults

`func NewModelsCreateOrGetAccountRequestWithDefaults() *ModelsCreateOrGetAccountRequest`

NewModelsCreateOrGetAccountRequestWithDefaults instantiates a new ModelsCreateOrGetAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ModelsCreateOrGetAccountRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsCreateOrGetAccountRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsCreateOrGetAccountRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *ModelsCreateOrGetAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsCreateOrGetAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsCreateOrGetAccountRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *ModelsCreateOrGetAccountRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelsCreateOrGetAccountRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelsCreateOrGetAccountRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderAccountId

`func (o *ModelsCreateOrGetAccountRequest) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *ModelsCreateOrGetAccountRequest) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *ModelsCreateOrGetAccountRequest) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.


### GetThumbnail

`func (o *ModelsCreateOrGetAccountRequest) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *ModelsCreateOrGetAccountRequest) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *ModelsCreateOrGetAccountRequest) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *ModelsCreateOrGetAccountRequest) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


