# AuthEnableMethodRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | Configuration for this mount, such as plugin_name. | [optional] 
**Description** | Pointer to **string** | User-friendly description for this credential backend. | [optional] 
**ExternalEntropyAccess** | Pointer to **bool** | Whether to give the mount access to Vault&#x27;s external entropy. | [optional] [default to false]
**Local** | Pointer to **bool** | Mark the mount as a local mount, which is not replicated and is unaffected by replication. | [optional] [default to false]
**Options** | Pointer to **map[string]interface{}** | The options to pass into the backend. Should be a json object with string keys and values. | [optional] 
**PluginName** | Pointer to **string** | Name of the auth plugin to use based from the name in the plugin catalog. | [optional] 
**PluginVersion** | Pointer to **string** | The semantic version of the plugin to use. | [optional] 
**SealWrap** | Pointer to **bool** | Whether to turn on seal wrapping for the mount. | [optional] [default to false]
**Type** | Pointer to **string** | The type of the backend. Example: \&quot;userpass\&quot; | [optional] 



## Methods


### NewAuthEnableMethodRequest

`func NewAuthEnableMethodRequest() *AuthEnableMethodRequest`

NewAuthEnableMethodRequest instantiates a new AuthEnableMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthEnableMethodRequestWithDefaults

`func NewAuthEnableMethodRequestWithDefaults() *AuthEnableMethodRequest`

NewAuthEnableMethodRequestWithDefaults instantiates a new AuthEnableMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetConfig

`func (o *AuthEnableMethodRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AuthEnableMethodRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AuthEnableMethodRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### HasConfig

`func (o *AuthEnableMethodRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.




### GetDescription

`func (o *AuthEnableMethodRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthEnableMethodRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthEnableMethodRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### HasDescription

`func (o *AuthEnableMethodRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.




### GetExternalEntropyAccess

`func (o *AuthEnableMethodRequest) GetExternalEntropyAccess() bool`

GetExternalEntropyAccess returns the ExternalEntropyAccess field if non-nil, zero value otherwise.

### GetExternalEntropyAccessOk

`func (o *AuthEnableMethodRequest) GetExternalEntropyAccessOk() (*bool, bool)`

GetExternalEntropyAccessOk returns a tuple with the ExternalEntropyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEntropyAccess

`func (o *AuthEnableMethodRequest) SetExternalEntropyAccess(v bool)`

SetExternalEntropyAccess sets ExternalEntropyAccess field to given value.


### HasExternalEntropyAccess

`func (o *AuthEnableMethodRequest) HasExternalEntropyAccess() bool`

HasExternalEntropyAccess returns a boolean if a field has been set.




### GetLocal

`func (o *AuthEnableMethodRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *AuthEnableMethodRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *AuthEnableMethodRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.


### HasLocal

`func (o *AuthEnableMethodRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.




### GetOptions

`func (o *AuthEnableMethodRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AuthEnableMethodRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AuthEnableMethodRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### HasOptions

`func (o *AuthEnableMethodRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.




### GetPluginName

`func (o *AuthEnableMethodRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AuthEnableMethodRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AuthEnableMethodRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### HasPluginName

`func (o *AuthEnableMethodRequest) HasPluginName() bool`

HasPluginName returns a boolean if a field has been set.




### GetPluginVersion

`func (o *AuthEnableMethodRequest) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *AuthEnableMethodRequest) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *AuthEnableMethodRequest) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.


### HasPluginVersion

`func (o *AuthEnableMethodRequest) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.




### GetSealWrap

`func (o *AuthEnableMethodRequest) GetSealWrap() bool`

GetSealWrap returns the SealWrap field if non-nil, zero value otherwise.

### GetSealWrapOk

`func (o *AuthEnableMethodRequest) GetSealWrapOk() (*bool, bool)`

GetSealWrapOk returns a tuple with the SealWrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealWrap

`func (o *AuthEnableMethodRequest) SetSealWrap(v bool)`

SetSealWrap sets SealWrap field to given value.


### HasSealWrap

`func (o *AuthEnableMethodRequest) HasSealWrap() bool`

HasSealWrap returns a boolean if a field has been set.




### GetType

`func (o *AuthEnableMethodRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthEnableMethodRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthEnableMethodRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *AuthEnableMethodRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


