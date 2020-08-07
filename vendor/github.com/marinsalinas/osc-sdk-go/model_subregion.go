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

// Subregion Information about the Subregion.
type Subregion struct {
	// The name of the Region containing the Subregion.
	RegionName *string `json:"RegionName,omitempty"`
	// The state of the Subregion (`available` \\| `information` \\| `impaired` \\| `unavailable`).
	State *string `json:"State,omitempty"`
	// The name of the Subregion.
	SubregionName *string `json:"SubregionName,omitempty"`
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *Subregion) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subregion) GetRegionNameOk() (string, bool) {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret, false
	}
	return *o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *Subregion) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *Subregion) SetRegionName(v string) {
	o.RegionName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Subregion) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subregion) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Subregion) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Subregion) SetState(v string) {
	o.State = &v
}

// GetSubregionName returns the SubregionName field value if set, zero value otherwise.
func (o *Subregion) GetSubregionName() string {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret
	}
	return *o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subregion) GetSubregionNameOk() (string, bool) {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret, false
	}
	return *o.SubregionName, true
}

// HasSubregionName returns a boolean if a field has been set.
func (o *Subregion) HasSubregionName() bool {
	if o != nil && o.SubregionName != nil {
		return true
	}

	return false
}

// SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.
func (o *Subregion) SetSubregionName(v string) {
	o.SubregionName = &v
}

type NullableSubregion struct {
	Value        Subregion
	ExplicitNull bool
}

func (v NullableSubregion) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSubregion) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}