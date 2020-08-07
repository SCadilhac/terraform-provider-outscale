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

// BlockDeviceMappingVmUpdate Information about the block device mapping.
type BlockDeviceMappingVmUpdate struct {
	Bsu *BsuToUpdateVm `json:"Bsu,omitempty"`
	// The name of the device.
	DeviceName *string `json:"DeviceName,omitempty"`
	// Removes the device which is included in the block device mapping of the OMI.
	NoDevice *string `json:"NoDevice,omitempty"`
	// The name of the virtual device (ephemeralN).
	VirtualDeviceName *string `json:"VirtualDeviceName,omitempty"`
}

// GetBsu returns the Bsu field value if set, zero value otherwise.
func (o *BlockDeviceMappingVmUpdate) GetBsu() BsuToUpdateVm {
	if o == nil || o.Bsu == nil {
		var ret BsuToUpdateVm
		return ret
	}
	return *o.Bsu
}

// GetBsuOk returns a tuple with the Bsu field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingVmUpdate) GetBsuOk() (BsuToUpdateVm, bool) {
	if o == nil || o.Bsu == nil {
		var ret BsuToUpdateVm
		return ret, false
	}
	return *o.Bsu, true
}

// HasBsu returns a boolean if a field has been set.
func (o *BlockDeviceMappingVmUpdate) HasBsu() bool {
	if o != nil && o.Bsu != nil {
		return true
	}

	return false
}

// SetBsu gets a reference to the given BsuToUpdateVm and assigns it to the Bsu field.
func (o *BlockDeviceMappingVmUpdate) SetBsu(v BsuToUpdateVm) {
	o.Bsu = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *BlockDeviceMappingVmUpdate) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingVmUpdate) GetDeviceNameOk() (string, bool) {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret, false
	}
	return *o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *BlockDeviceMappingVmUpdate) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *BlockDeviceMappingVmUpdate) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetNoDevice returns the NoDevice field value if set, zero value otherwise.
func (o *BlockDeviceMappingVmUpdate) GetNoDevice() string {
	if o == nil || o.NoDevice == nil {
		var ret string
		return ret
	}
	return *o.NoDevice
}

// GetNoDeviceOk returns a tuple with the NoDevice field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingVmUpdate) GetNoDeviceOk() (string, bool) {
	if o == nil || o.NoDevice == nil {
		var ret string
		return ret, false
	}
	return *o.NoDevice, true
}

// HasNoDevice returns a boolean if a field has been set.
func (o *BlockDeviceMappingVmUpdate) HasNoDevice() bool {
	if o != nil && o.NoDevice != nil {
		return true
	}

	return false
}

// SetNoDevice gets a reference to the given string and assigns it to the NoDevice field.
func (o *BlockDeviceMappingVmUpdate) SetNoDevice(v string) {
	o.NoDevice = &v
}

// GetVirtualDeviceName returns the VirtualDeviceName field value if set, zero value otherwise.
func (o *BlockDeviceMappingVmUpdate) GetVirtualDeviceName() string {
	if o == nil || o.VirtualDeviceName == nil {
		var ret string
		return ret
	}
	return *o.VirtualDeviceName
}

// GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingVmUpdate) GetVirtualDeviceNameOk() (string, bool) {
	if o == nil || o.VirtualDeviceName == nil {
		var ret string
		return ret, false
	}
	return *o.VirtualDeviceName, true
}

// HasVirtualDeviceName returns a boolean if a field has been set.
func (o *BlockDeviceMappingVmUpdate) HasVirtualDeviceName() bool {
	if o != nil && o.VirtualDeviceName != nil {
		return true
	}

	return false
}

// SetVirtualDeviceName gets a reference to the given string and assigns it to the VirtualDeviceName field.
func (o *BlockDeviceMappingVmUpdate) SetVirtualDeviceName(v string) {
	o.VirtualDeviceName = &v
}

type NullableBlockDeviceMappingVmUpdate struct {
	Value        BlockDeviceMappingVmUpdate
	ExplicitNull bool
}

func (v NullableBlockDeviceMappingVmUpdate) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBlockDeviceMappingVmUpdate) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}