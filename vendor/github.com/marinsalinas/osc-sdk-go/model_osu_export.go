/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.1
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// OsuExport Information about the OSU export.
type OsuExport struct {
	// The format of the export disk (`qcow2` \\| `vdi` \\| `vmdk`).
	DiskImageFormat string     `json:"DiskImageFormat"`
	OsuApiKey       *OsuApiKey `json:"OsuApiKey,omitempty"`
	// The name of the OSU bucket you want to export the object to.
	OsuBucket string `json:"OsuBucket"`
	// The URL of the manifest file.
	OsuManifestUrl *string `json:"OsuManifestUrl,omitempty"`
	// The prefix for the key of the OSU object. This key follows this format: `prefix + object_export_task_id + '.' + disk_image_format`.
	OsuPrefix *string `json:"OsuPrefix,omitempty"`
}

// GetDiskImageFormat returns the DiskImageFormat field value
func (o *OsuExport) GetDiskImageFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskImageFormat
}

// SetDiskImageFormat sets field value
func (o *OsuExport) SetDiskImageFormat(v string) {
	o.DiskImageFormat = v
}

// GetOsuApiKey returns the OsuApiKey field value if set, zero value otherwise.
func (o *OsuExport) GetOsuApiKey() OsuApiKey {
	if o == nil || o.OsuApiKey == nil {
		var ret OsuApiKey
		return ret
	}
	return *o.OsuApiKey
}

// GetOsuApiKeyOk returns a tuple with the OsuApiKey field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OsuExport) GetOsuApiKeyOk() (OsuApiKey, bool) {
	if o == nil || o.OsuApiKey == nil {
		var ret OsuApiKey
		return ret, false
	}
	return *o.OsuApiKey, true
}

// HasOsuApiKey returns a boolean if a field has been set.
func (o *OsuExport) HasOsuApiKey() bool {
	if o != nil && o.OsuApiKey != nil {
		return true
	}

	return false
}

// SetOsuApiKey gets a reference to the given OsuApiKey and assigns it to the OsuApiKey field.
func (o *OsuExport) SetOsuApiKey(v OsuApiKey) {
	o.OsuApiKey = &v
}

// GetOsuBucket returns the OsuBucket field value
func (o *OsuExport) GetOsuBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsuBucket
}

// SetOsuBucket sets field value
func (o *OsuExport) SetOsuBucket(v string) {
	o.OsuBucket = v
}

// GetOsuManifestUrl returns the OsuManifestUrl field value if set, zero value otherwise.
func (o *OsuExport) GetOsuManifestUrl() string {
	if o == nil || o.OsuManifestUrl == nil {
		var ret string
		return ret
	}
	return *o.OsuManifestUrl
}

// GetOsuManifestUrlOk returns a tuple with the OsuManifestUrl field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OsuExport) GetOsuManifestUrlOk() (string, bool) {
	if o == nil || o.OsuManifestUrl == nil {
		var ret string
		return ret, false
	}
	return *o.OsuManifestUrl, true
}

// HasOsuManifestUrl returns a boolean if a field has been set.
func (o *OsuExport) HasOsuManifestUrl() bool {
	if o != nil && o.OsuManifestUrl != nil {
		return true
	}

	return false
}

// SetOsuManifestUrl gets a reference to the given string and assigns it to the OsuManifestUrl field.
func (o *OsuExport) SetOsuManifestUrl(v string) {
	o.OsuManifestUrl = &v
}

// GetOsuPrefix returns the OsuPrefix field value if set, zero value otherwise.
func (o *OsuExport) GetOsuPrefix() string {
	if o == nil || o.OsuPrefix == nil {
		var ret string
		return ret
	}
	return *o.OsuPrefix
}

// GetOsuPrefixOk returns a tuple with the OsuPrefix field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OsuExport) GetOsuPrefixOk() (string, bool) {
	if o == nil || o.OsuPrefix == nil {
		var ret string
		return ret, false
	}
	return *o.OsuPrefix, true
}

// HasOsuPrefix returns a boolean if a field has been set.
func (o *OsuExport) HasOsuPrefix() bool {
	if o != nil && o.OsuPrefix != nil {
		return true
	}

	return false
}

// SetOsuPrefix gets a reference to the given string and assigns it to the OsuPrefix field.
func (o *OsuExport) SetOsuPrefix(v string) {
	o.OsuPrefix = &v
}

type NullableOsuExport struct {
	Value        OsuExport
	ExplicitNull bool
}

func (v NullableOsuExport) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOsuExport) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}