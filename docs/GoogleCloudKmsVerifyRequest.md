# GoogleCloudKmsVerifyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | Digest to verify. This digest must use the same SHA algorithm as the underlying Cloud KMS key. The digest must be the base64-encoded binary value. This field is required. | [optional] 
**KeyVersion** | Pointer to **int32** | Integer version of the crypto key version to use for verification. This field is required. | [optional] 
**Signature** | Pointer to **string** | Base64-encoded signature to use for verification. This field is required. | [optional] 



## Methods


### NewGoogleCloudKmsVerifyRequest

`func NewGoogleCloudKmsVerifyRequest() *GoogleCloudKmsVerifyRequest`

NewGoogleCloudKmsVerifyRequest instantiates a new GoogleCloudKmsVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudKmsVerifyRequestWithDefaults

`func NewGoogleCloudKmsVerifyRequestWithDefaults() *GoogleCloudKmsVerifyRequest`

NewGoogleCloudKmsVerifyRequestWithDefaults instantiates a new GoogleCloudKmsVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDigest

`func (o *GoogleCloudKmsVerifyRequest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *GoogleCloudKmsVerifyRequest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *GoogleCloudKmsVerifyRequest) SetDigest(v string)`

SetDigest sets Digest field to given value.


### HasDigest

`func (o *GoogleCloudKmsVerifyRequest) HasDigest() bool`

HasDigest returns a boolean if a field has been set.




### GetKeyVersion

`func (o *GoogleCloudKmsVerifyRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *GoogleCloudKmsVerifyRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *GoogleCloudKmsVerifyRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.


### HasKeyVersion

`func (o *GoogleCloudKmsVerifyRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.




### GetSignature

`func (o *GoogleCloudKmsVerifyRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *GoogleCloudKmsVerifyRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *GoogleCloudKmsVerifyRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.


### HasSignature

`func (o *GoogleCloudKmsVerifyRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


