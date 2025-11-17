# ModelsAccountSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | アカウントID | 
**FirstName** | **string** | 名前 | 
**LastName** | **string** | 苗字 | 
**Thumbnail** | Pointer to **string** | プロフィール画像URL | [optional] 

## Methods

### NewModelsAccountSummary

`func NewModelsAccountSummary(id string, firstName string, lastName string, ) *ModelsAccountSummary`

NewModelsAccountSummary instantiates a new ModelsAccountSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAccountSummaryWithDefaults

`func NewModelsAccountSummaryWithDefaults() *ModelsAccountSummary`

NewModelsAccountSummaryWithDefaults instantiates a new ModelsAccountSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsAccountSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsAccountSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsAccountSummary) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *ModelsAccountSummary) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ModelsAccountSummary) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ModelsAccountSummary) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ModelsAccountSummary) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ModelsAccountSummary) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ModelsAccountSummary) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetThumbnail

`func (o *ModelsAccountSummary) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *ModelsAccountSummary) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *ModelsAccountSummary) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *ModelsAccountSummary) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


