// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuersImportBundleResponse struct for PkiIssuersImportBundleResponse
type PkiIssuersImportBundleResponse struct {
	// Net-new issuers imported as a part of this request
	ImportedIssuers []string `json:"imported_issuers"`

	// Net-new keys imported as a part of this request
	ImportedKeys []string `json:"imported_keys"`

	// A mapping of issuer_id to key_id for all issuers included in this request
	Mapping map[string]interface{} `json:"mapping"`
}

// NewPkiIssuersImportBundleResponseWithDefaults instantiates a new PkiIssuersImportBundleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuersImportBundleResponseWithDefaults() *PkiIssuersImportBundleResponse {
	var this PkiIssuersImportBundleResponse

	return &this
}

func (o PkiIssuersImportBundleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["imported_issuers"] = o.ImportedIssuers
	toSerialize["imported_keys"] = o.ImportedKeys
	toSerialize["mapping"] = o.Mapping

	return json.Marshal(toSerialize)
}