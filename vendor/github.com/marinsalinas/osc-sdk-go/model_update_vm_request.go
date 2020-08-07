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

// UpdateVmRequest struct for UpdateVmRequest
type UpdateVmRequest struct {
	// One or more block device mappings of the VM.
	BlockDeviceMappings *[]BlockDeviceMappingVmUpdate `json:"BlockDeviceMappings,omitempty"`
	// If `true`, the VM is optimized for BSU I/O.
	BsuOptimized *bool `json:"BsuOptimized,omitempty"`
	// If `true`, you cannot terminate the VM using Cockpit, the CLI or the API. If `false`, you can.
	DeletionProtection *bool `json:"DeletionProtection,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// (Net only) If `true`, the source/destination check is enabled. If `false`, it is disabled. This value must be `false` for a NAT VM to perform network address translation (NAT) in a Net.
	IsSourceDestChecked *bool `json:"IsSourceDestChecked,omitempty"`
	// The name of the keypair.<br /> To complete the replacement, manually replace the old public key with the new public key in the ~/.ssh/authorized_keys file located in the VM. Restart the VM to apply the change.
	KeypairName *string `json:"KeypairName,omitempty"`
	// The performance of the VM (`standard` \\| `high` \\|  `highest`).
	Performance *string `json:"Performance,omitempty"`
	// One or more IDs of security groups for the VM.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// The Base64-encoded MIME user data.
	UserData *string `json:"UserData,omitempty"`
	// The ID of the VM.
	VmId string `json:"VmId"`
	// The VM behavior when you stop it. By default or if set to `stop`, the VM stops. If set to `restart`, the VM stops then automatically restarts. If set to `terminate`, the VM stops and is terminated.
	VmInitiatedShutdownBehavior *string `json:"VmInitiatedShutdownBehavior,omitempty"`
	// The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types).
	VmType *string `json:"VmType,omitempty"`
}

// GetBlockDeviceMappings returns the BlockDeviceMappings field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetBlockDeviceMappings() []BlockDeviceMappingVmUpdate {
	if o == nil || o.BlockDeviceMappings == nil {
		var ret []BlockDeviceMappingVmUpdate
		return ret
	}
	return *o.BlockDeviceMappings
}

// GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetBlockDeviceMappingsOk() ([]BlockDeviceMappingVmUpdate, bool) {
	if o == nil || o.BlockDeviceMappings == nil {
		var ret []BlockDeviceMappingVmUpdate
		return ret, false
	}
	return *o.BlockDeviceMappings, true
}

// HasBlockDeviceMappings returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasBlockDeviceMappings() bool {
	if o != nil && o.BlockDeviceMappings != nil {
		return true
	}

	return false
}

// SetBlockDeviceMappings gets a reference to the given []BlockDeviceMappingVmUpdate and assigns it to the BlockDeviceMappings field.
func (o *UpdateVmRequest) SetBlockDeviceMappings(v []BlockDeviceMappingVmUpdate) {
	o.BlockDeviceMappings = &v
}

// GetBsuOptimized returns the BsuOptimized field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetBsuOptimized() bool {
	if o == nil || o.BsuOptimized == nil {
		var ret bool
		return ret
	}
	return *o.BsuOptimized
}

// GetBsuOptimizedOk returns a tuple with the BsuOptimized field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetBsuOptimizedOk() (bool, bool) {
	if o == nil || o.BsuOptimized == nil {
		var ret bool
		return ret, false
	}
	return *o.BsuOptimized, true
}

// HasBsuOptimized returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasBsuOptimized() bool {
	if o != nil && o.BsuOptimized != nil {
		return true
	}

	return false
}

// SetBsuOptimized gets a reference to the given bool and assigns it to the BsuOptimized field.
func (o *UpdateVmRequest) SetBsuOptimized(v bool) {
	o.BsuOptimized = &v
}

// GetDeletionProtection returns the DeletionProtection field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetDeletionProtection() bool {
	if o == nil || o.DeletionProtection == nil {
		var ret bool
		return ret
	}
	return *o.DeletionProtection
}

// GetDeletionProtectionOk returns a tuple with the DeletionProtection field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetDeletionProtectionOk() (bool, bool) {
	if o == nil || o.DeletionProtection == nil {
		var ret bool
		return ret, false
	}
	return *o.DeletionProtection, true
}

// HasDeletionProtection returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasDeletionProtection() bool {
	if o != nil && o.DeletionProtection != nil {
		return true
	}

	return false
}

// SetDeletionProtection gets a reference to the given bool and assigns it to the DeletionProtection field.
func (o *UpdateVmRequest) SetDeletionProtection(v bool) {
	o.DeletionProtection = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateVmRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetIsSourceDestChecked returns the IsSourceDestChecked field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetIsSourceDestChecked() bool {
	if o == nil || o.IsSourceDestChecked == nil {
		var ret bool
		return ret
	}
	return *o.IsSourceDestChecked
}

// GetIsSourceDestCheckedOk returns a tuple with the IsSourceDestChecked field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetIsSourceDestCheckedOk() (bool, bool) {
	if o == nil || o.IsSourceDestChecked == nil {
		var ret bool
		return ret, false
	}
	return *o.IsSourceDestChecked, true
}

// HasIsSourceDestChecked returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasIsSourceDestChecked() bool {
	if o != nil && o.IsSourceDestChecked != nil {
		return true
	}

	return false
}

// SetIsSourceDestChecked gets a reference to the given bool and assigns it to the IsSourceDestChecked field.
func (o *UpdateVmRequest) SetIsSourceDestChecked(v bool) {
	o.IsSourceDestChecked = &v
}

// GetKeypairName returns the KeypairName field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetKeypairName() string {
	if o == nil || o.KeypairName == nil {
		var ret string
		return ret
	}
	return *o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetKeypairNameOk() (string, bool) {
	if o == nil || o.KeypairName == nil {
		var ret string
		return ret, false
	}
	return *o.KeypairName, true
}

// HasKeypairName returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasKeypairName() bool {
	if o != nil && o.KeypairName != nil {
		return true
	}

	return false
}

// SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.
func (o *UpdateVmRequest) SetKeypairName(v string) {
	o.KeypairName = &v
}

// GetPerformance returns the Performance field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetPerformance() string {
	if o == nil || o.Performance == nil {
		var ret string
		return ret
	}
	return *o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetPerformanceOk() (string, bool) {
	if o == nil || o.Performance == nil {
		var ret string
		return ret, false
	}
	return *o.Performance, true
}

// HasPerformance returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasPerformance() bool {
	if o != nil && o.Performance != nil {
		return true
	}

	return false
}

// SetPerformance gets a reference to the given string and assigns it to the Performance field.
func (o *UpdateVmRequest) SetPerformance(v string) {
	o.Performance = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetSecurityGroupIdsOk() ([]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret, false
	}
	return *o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *UpdateVmRequest) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetUserData() string {
	if o == nil || o.UserData == nil {
		var ret string
		return ret
	}
	return *o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetUserDataOk() (string, bool) {
	if o == nil || o.UserData == nil {
		var ret string
		return ret, false
	}
	return *o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given string and assigns it to the UserData field.
func (o *UpdateVmRequest) SetUserData(v string) {
	o.UserData = &v
}

// GetVmId returns the VmId field value
func (o *UpdateVmRequest) GetVmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmId
}

// SetVmId sets field value
func (o *UpdateVmRequest) SetVmId(v string) {
	o.VmId = v
}

// GetVmInitiatedShutdownBehavior returns the VmInitiatedShutdownBehavior field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetVmInitiatedShutdownBehavior() string {
	if o == nil || o.VmInitiatedShutdownBehavior == nil {
		var ret string
		return ret
	}
	return *o.VmInitiatedShutdownBehavior
}

// GetVmInitiatedShutdownBehaviorOk returns a tuple with the VmInitiatedShutdownBehavior field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetVmInitiatedShutdownBehaviorOk() (string, bool) {
	if o == nil || o.VmInitiatedShutdownBehavior == nil {
		var ret string
		return ret, false
	}
	return *o.VmInitiatedShutdownBehavior, true
}

// HasVmInitiatedShutdownBehavior returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasVmInitiatedShutdownBehavior() bool {
	if o != nil && o.VmInitiatedShutdownBehavior != nil {
		return true
	}

	return false
}

// SetVmInitiatedShutdownBehavior gets a reference to the given string and assigns it to the VmInitiatedShutdownBehavior field.
func (o *UpdateVmRequest) SetVmInitiatedShutdownBehavior(v string) {
	o.VmInitiatedShutdownBehavior = &v
}

// GetVmType returns the VmType field value if set, zero value otherwise.
func (o *UpdateVmRequest) GetVmType() string {
	if o == nil || o.VmType == nil {
		var ret string
		return ret
	}
	return *o.VmType
}

// GetVmTypeOk returns a tuple with the VmType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmRequest) GetVmTypeOk() (string, bool) {
	if o == nil || o.VmType == nil {
		var ret string
		return ret, false
	}
	return *o.VmType, true
}

// HasVmType returns a boolean if a field has been set.
func (o *UpdateVmRequest) HasVmType() bool {
	if o != nil && o.VmType != nil {
		return true
	}

	return false
}

// SetVmType gets a reference to the given string and assigns it to the VmType field.
func (o *UpdateVmRequest) SetVmType(v string) {
	o.VmType = &v
}

type NullableUpdateVmRequest struct {
	Value        UpdateVmRequest
	ExplicitNull bool
}

func (v NullableUpdateVmRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateVmRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}