# AppendAllowListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | **string** |  | 
**RuleType** | [**AllowListRules**](AllowListRules.md) |  | 
**Comment** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **string** | The duration you would like the temporary cidr block to be active | [optional] 

## Methods

### NewAppendAllowListRequest

`func NewAppendAllowListRequest(cidrBlock string, ruleType AllowListRules, ) *AppendAllowListRequest`

NewAppendAllowListRequest instantiates a new AppendAllowListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppendAllowListRequestWithDefaults

`func NewAppendAllowListRequestWithDefaults() *AppendAllowListRequest`

NewAppendAllowListRequestWithDefaults instantiates a new AppendAllowListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *AppendAllowListRequest) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AppendAllowListRequest) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AppendAllowListRequest) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetRuleType

`func (o *AppendAllowListRequest) GetRuleType() AllowListRules`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AppendAllowListRequest) GetRuleTypeOk() (*AllowListRules, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AppendAllowListRequest) SetRuleType(v AllowListRules)`

SetRuleType sets RuleType field to given value.


### GetComment

`func (o *AppendAllowListRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AppendAllowListRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AppendAllowListRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AppendAllowListRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDuration

`func (o *AppendAllowListRequest) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AppendAllowListRequest) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AppendAllowListRequest) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AppendAllowListRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


