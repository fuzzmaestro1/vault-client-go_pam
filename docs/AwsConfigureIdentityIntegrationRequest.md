# AwsConfigureIdentityIntegrationRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ec2Alias** | Pointer to **string** | Configure how the AWS auth method generates entity alias when using EC2 auth. Valid values are \&quot;role_id\&quot;, \&quot;instance_id\&quot;, and \&quot;image_id\&quot;. Defaults to \&quot;role_id\&quot;. | [optional] [default to "instance_id"]
**Ec2Metadata** | Pointer to **[]string** | The metadata to include on the aliases and audit logs generated by this plugin. When set to &#x27;default&#x27;, includes: account_id, auth_type. These fields are available to add: ami_id, instance_id, region. Not editing this field means the &#x27;default&#x27; fields are included. Explicitly setting this field to empty overrides the &#x27;default&#x27; and means no metadata will be included. If not using &#x27;default&#x27;, explicit fields must be sent like: &#x27;field1,field2&#x27;. | [optional] [default to ["default"]]
**IamAlias** | Pointer to **string** | Configure how the AWS auth method generates entity aliases when using IAM auth. Valid values are \&quot;role_id\&quot;, \&quot;unique_id\&quot;, and \&quot;full_arn\&quot;. Defaults to \&quot;role_id\&quot;. | [optional] [default to "unique_id"]
**IamMetadata** | Pointer to **[]string** | The metadata to include on the aliases and audit logs generated by this plugin. When set to &#x27;default&#x27;, includes: account_id, auth_type. These fields are available to add: canonical_arn, client_arn, client_user_id, inferred_aws_region, inferred_entity_id, inferred_entity_type. Not editing this field means the &#x27;default&#x27; fields are included. Explicitly setting this field to empty overrides the &#x27;default&#x27; and means no metadata will be included. If not using &#x27;default&#x27;, explicit fields must be sent like: &#x27;field1,field2&#x27;. | [optional] [default to ["default"]]



## Methods


### NewAwsConfigureIdentityIntegrationRequest

`func NewAwsConfigureIdentityIntegrationRequest() *AwsConfigureIdentityIntegrationRequest`

NewAwsConfigureIdentityIntegrationRequest instantiates a new AwsConfigureIdentityIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConfigureIdentityIntegrationRequestWithDefaults

`func NewAwsConfigureIdentityIntegrationRequestWithDefaults() *AwsConfigureIdentityIntegrationRequest`

NewAwsConfigureIdentityIntegrationRequestWithDefaults instantiates a new AwsConfigureIdentityIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetEc2Alias

`func (o *AwsConfigureIdentityIntegrationRequest) GetEc2Alias() string`

GetEc2Alias returns the Ec2Alias field if non-nil, zero value otherwise.

### GetEc2AliasOk

`func (o *AwsConfigureIdentityIntegrationRequest) GetEc2AliasOk() (*string, bool)`

GetEc2AliasOk returns a tuple with the Ec2Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEc2Alias

`func (o *AwsConfigureIdentityIntegrationRequest) SetEc2Alias(v string)`

SetEc2Alias sets Ec2Alias field to given value.


### HasEc2Alias

`func (o *AwsConfigureIdentityIntegrationRequest) HasEc2Alias() bool`

HasEc2Alias returns a boolean if a field has been set.




### GetEc2Metadata

`func (o *AwsConfigureIdentityIntegrationRequest) GetEc2Metadata() []string`

GetEc2Metadata returns the Ec2Metadata field if non-nil, zero value otherwise.

### GetEc2MetadataOk

`func (o *AwsConfigureIdentityIntegrationRequest) GetEc2MetadataOk() (*[]string, bool)`

GetEc2MetadataOk returns a tuple with the Ec2Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEc2Metadata

`func (o *AwsConfigureIdentityIntegrationRequest) SetEc2Metadata(v []string)`

SetEc2Metadata sets Ec2Metadata field to given value.


### HasEc2Metadata

`func (o *AwsConfigureIdentityIntegrationRequest) HasEc2Metadata() bool`

HasEc2Metadata returns a boolean if a field has been set.




### GetIamAlias

`func (o *AwsConfigureIdentityIntegrationRequest) GetIamAlias() string`

GetIamAlias returns the IamAlias field if non-nil, zero value otherwise.

### GetIamAliasOk

`func (o *AwsConfigureIdentityIntegrationRequest) GetIamAliasOk() (*string, bool)`

GetIamAliasOk returns a tuple with the IamAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamAlias

`func (o *AwsConfigureIdentityIntegrationRequest) SetIamAlias(v string)`

SetIamAlias sets IamAlias field to given value.


### HasIamAlias

`func (o *AwsConfigureIdentityIntegrationRequest) HasIamAlias() bool`

HasIamAlias returns a boolean if a field has been set.




### GetIamMetadata

`func (o *AwsConfigureIdentityIntegrationRequest) GetIamMetadata() []string`

GetIamMetadata returns the IamMetadata field if non-nil, zero value otherwise.

### GetIamMetadataOk

`func (o *AwsConfigureIdentityIntegrationRequest) GetIamMetadataOk() (*[]string, bool)`

GetIamMetadataOk returns a tuple with the IamMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamMetadata

`func (o *AwsConfigureIdentityIntegrationRequest) SetIamMetadata(v []string)`

SetIamMetadata sets IamMetadata field to given value.


### HasIamMetadata

`func (o *AwsConfigureIdentityIntegrationRequest) HasIamMetadata() bool`

HasIamMetadata returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


