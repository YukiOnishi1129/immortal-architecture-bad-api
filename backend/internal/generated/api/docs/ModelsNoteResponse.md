# ModelsNoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ノートID | 
**Title** | **string** | タイトル | 
**TemplateId** | **string** | テンプレートID | 
**TemplateName** | **string** | テンプレート名 | 
**OwnerId** | **string** | 所有者ID | 
**Owner** | [**ModelsAccountSummary**](ModelsAccountSummary.md) | 所有者情報 | 
**Status** | [**ModelsNoteStatus**](ModelsNoteStatus.md) | ステータス | 
**Sections** | [**[]ModelsSection**](ModelsSection.md) | セクション | 
**CreatedAt** | **time.Time** | 作成日時 | 
**UpdatedAt** | **time.Time** | 更新日時 | 

## Methods

### NewModelsNoteResponse

`func NewModelsNoteResponse(id string, title string, templateId string, templateName string, ownerId string, owner ModelsAccountSummary, status ModelsNoteStatus, sections []ModelsSection, createdAt time.Time, updatedAt time.Time, ) *ModelsNoteResponse`

NewModelsNoteResponse instantiates a new ModelsNoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsNoteResponseWithDefaults

`func NewModelsNoteResponseWithDefaults() *ModelsNoteResponse`

NewModelsNoteResponseWithDefaults instantiates a new ModelsNoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsNoteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsNoteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsNoteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ModelsNoteResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsNoteResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsNoteResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTemplateId

`func (o *ModelsNoteResponse) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ModelsNoteResponse) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ModelsNoteResponse) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetTemplateName

`func (o *ModelsNoteResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ModelsNoteResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ModelsNoteResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetOwnerId

`func (o *ModelsNoteResponse) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ModelsNoteResponse) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ModelsNoteResponse) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwner

`func (o *ModelsNoteResponse) GetOwner() ModelsAccountSummary`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelsNoteResponse) GetOwnerOk() (*ModelsAccountSummary, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelsNoteResponse) SetOwner(v ModelsAccountSummary)`

SetOwner sets Owner field to given value.


### GetStatus

`func (o *ModelsNoteResponse) GetStatus() ModelsNoteStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsNoteResponse) GetStatusOk() (*ModelsNoteStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsNoteResponse) SetStatus(v ModelsNoteStatus)`

SetStatus sets Status field to given value.


### GetSections

`func (o *ModelsNoteResponse) GetSections() []ModelsSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ModelsNoteResponse) GetSectionsOk() (*[]ModelsSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ModelsNoteResponse) SetSections(v []ModelsSection)`

SetSections sets Sections field to given value.


### GetCreatedAt

`func (o *ModelsNoteResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsNoteResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsNoteResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ModelsNoteResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsNoteResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsNoteResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


