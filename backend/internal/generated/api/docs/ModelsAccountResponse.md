# ModelsAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | アカウントID | 
**Email** | **string** | メールアドレス | 
**FirstName** | **string** | 名前 | 
**LastName** | **string** | 苗字 | 
**FullName** | **string** | フルネーム | 
**Thumbnail** | Pointer to **string** | プロフィール画像URL | [optional] 
**LastLoginAt** | **time.Time** | 最終ログイン日時 | 
**CreatedAt** | **time.Time** | 作成日時 | 
**UpdatedAt** | **time.Time** | 更新日時 | 

## Methods

### NewModelsAccountResponse

`func NewModelsAccountResponse(id string, email string, firstName string, lastName string, fullName string, lastLoginAt time.Time, createdAt time.Time, updatedAt time.Time, ) *ModelsAccountResponse`

NewModelsAccountResponse instantiates a new ModelsAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAccountResponseWithDefaults

`func NewModelsAccountResponseWithDefaults() *ModelsAccountResponse`

NewModelsAccountResponseWithDefaults instantiates a new ModelsAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsAccountResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *ModelsAccountResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsAccountResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsAccountResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *ModelsAccountResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ModelsAccountResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ModelsAccountResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ModelsAccountResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ModelsAccountResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ModelsAccountResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetFullName

`func (o *ModelsAccountResponse) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ModelsAccountResponse) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ModelsAccountResponse) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetThumbnail

`func (o *ModelsAccountResponse) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *ModelsAccountResponse) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *ModelsAccountResponse) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *ModelsAccountResponse) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *ModelsAccountResponse) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *ModelsAccountResponse) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *ModelsAccountResponse) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.


### GetCreatedAt

`func (o *ModelsAccountResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsAccountResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsAccountResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ModelsAccountResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsAccountResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsAccountResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


