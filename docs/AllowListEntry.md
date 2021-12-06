# AllowListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | **string** |  | 
**RuleType** | [**AllowListRules**](AllowListRules.md) |  | 
**Duration** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**AllowListStates**](AllowListStates.md) |  | [optional] 

## Methods

### NewAllowListEntry

`func NewAllowListEntry(cidrBlock string, ruleType AllowListRules, ) *AllowListEntry`

NewAllowListEntry instantiates a new AllowListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowListEntryWithDefaults

`func NewAllowListEntryWithDefaults() *AllowListEntry`

NewAllowListEntryWithDefaults instantiates a new AllowListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *AllowListEntry) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AllowListEntry) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AllowListEntry) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetRuleType

`func (o *AllowListEntry) GetRuleType() AllowListRules`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AllowListEntry) GetRuleTypeOk() (*AllowListRules, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AllowListEntry) SetRuleType(v AllowListRules)`

SetRuleType sets RuleType field to given value.


### GetDuration

`func (o *AllowListEntry) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AllowListEntry) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AllowListEntry) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AllowListEntry) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AllowListEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AllowListEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AllowListEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AllowListEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AllowListEntry) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AllowListEntry) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AllowListEntry) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AllowListEntry) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AllowListEntry) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AllowListEntry) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AllowListEntry) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AllowListEntry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetComment

`func (o *AllowListEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AllowListEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AllowListEntry) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AllowListEntry) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetState

`func (o *AllowListEntry) GetState() AllowListStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AllowListEntry) GetStateOk() (*AllowListStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AllowListEntry) SetState(v AllowListStates)`

SetState sets State field to given value.

### HasState

`func (o *AllowListEntry) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


