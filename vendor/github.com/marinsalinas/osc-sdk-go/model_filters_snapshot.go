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

// FiltersSnapshot One or more filters.
type FiltersSnapshot struct {
	// The account aliases of the owners of the snapshots.
	AccountAliases *[]string `json:"AccountAliases,omitempty"`
	// The account IDs of the owners of the snapshots.
	AccountIds *[]string `json:"AccountIds,omitempty"`
	// The descriptions of the snapshots.
	Descriptions *[]string `json:"Descriptions,omitempty"`
	// The account IDs of one or more users who have permissions to create volumes.
	PermissionsToCreateVolumeAccountIds *[]string `json:"PermissionsToCreateVolumeAccountIds,omitempty"`
	// If `true`, lists all public volumes. If `false`, lists all private volumes.
	PermissionsToCreateVolumeGlobalPermission *bool `json:"PermissionsToCreateVolumeGlobalPermission,omitempty"`
	// The progresses of the snapshots, as a percentage.
	Progresses *[]int64 `json:"Progresses,omitempty"`
	// The IDs of the snapshots.
	SnapshotIds *[]string `json:"SnapshotIds,omitempty"`
	// The states of the snapshots (`in-queue` \\| `completed` \\| `error`).
	States *[]string `json:"States,omitempty"`
	// The keys of the tags associated with the snapshots.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the snapshots.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the snapshots, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags *[]string `json:"Tags,omitempty"`
	// The IDs of the volumes used to create the snapshots.
	VolumeIds *[]string `json:"VolumeIds,omitempty"`
	// The sizes of the volumes used to create the snapshots, in gibibytes (GiB).
	VolumeSizes *[]int64 `json:"VolumeSizes,omitempty"`
}

// GetAccountAliases returns the AccountAliases field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetAccountAliases() []string {
	if o == nil || o.AccountAliases == nil {
		var ret []string
		return ret
	}
	return *o.AccountAliases
}

// GetAccountAliasesOk returns a tuple with the AccountAliases field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetAccountAliasesOk() ([]string, bool) {
	if o == nil || o.AccountAliases == nil {
		var ret []string
		return ret, false
	}
	return *o.AccountAliases, true
}

// HasAccountAliases returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasAccountAliases() bool {
	if o != nil && o.AccountAliases != nil {
		return true
	}

	return false
}

// SetAccountAliases gets a reference to the given []string and assigns it to the AccountAliases field.
func (o *FiltersSnapshot) SetAccountAliases(v []string) {
	o.AccountAliases = &v
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetAccountIdsOk() ([]string, bool) {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret, false
	}
	return *o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *FiltersSnapshot) SetAccountIds(v []string) {
	o.AccountIds = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetDescriptions() []string {
	if o == nil || o.Descriptions == nil {
		var ret []string
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetDescriptionsOk() ([]string, bool) {
	if o == nil || o.Descriptions == nil {
		var ret []string
		return ret, false
	}
	return *o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.
func (o *FiltersSnapshot) SetDescriptions(v []string) {
	o.Descriptions = &v
}

// GetPermissionsToCreateVolumeAccountIds returns the PermissionsToCreateVolumeAccountIds field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetPermissionsToCreateVolumeAccountIds() []string {
	if o == nil || o.PermissionsToCreateVolumeAccountIds == nil {
		var ret []string
		return ret
	}
	return *o.PermissionsToCreateVolumeAccountIds
}

// GetPermissionsToCreateVolumeAccountIdsOk returns a tuple with the PermissionsToCreateVolumeAccountIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetPermissionsToCreateVolumeAccountIdsOk() ([]string, bool) {
	if o == nil || o.PermissionsToCreateVolumeAccountIds == nil {
		var ret []string
		return ret, false
	}
	return *o.PermissionsToCreateVolumeAccountIds, true
}

// HasPermissionsToCreateVolumeAccountIds returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasPermissionsToCreateVolumeAccountIds() bool {
	if o != nil && o.PermissionsToCreateVolumeAccountIds != nil {
		return true
	}

	return false
}

// SetPermissionsToCreateVolumeAccountIds gets a reference to the given []string and assigns it to the PermissionsToCreateVolumeAccountIds field.
func (o *FiltersSnapshot) SetPermissionsToCreateVolumeAccountIds(v []string) {
	o.PermissionsToCreateVolumeAccountIds = &v
}

// GetPermissionsToCreateVolumeGlobalPermission returns the PermissionsToCreateVolumeGlobalPermission field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetPermissionsToCreateVolumeGlobalPermission() bool {
	if o == nil || o.PermissionsToCreateVolumeGlobalPermission == nil {
		var ret bool
		return ret
	}
	return *o.PermissionsToCreateVolumeGlobalPermission
}

// GetPermissionsToCreateVolumeGlobalPermissionOk returns a tuple with the PermissionsToCreateVolumeGlobalPermission field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetPermissionsToCreateVolumeGlobalPermissionOk() (bool, bool) {
	if o == nil || o.PermissionsToCreateVolumeGlobalPermission == nil {
		var ret bool
		return ret, false
	}
	return *o.PermissionsToCreateVolumeGlobalPermission, true
}

// HasPermissionsToCreateVolumeGlobalPermission returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasPermissionsToCreateVolumeGlobalPermission() bool {
	if o != nil && o.PermissionsToCreateVolumeGlobalPermission != nil {
		return true
	}

	return false
}

// SetPermissionsToCreateVolumeGlobalPermission gets a reference to the given bool and assigns it to the PermissionsToCreateVolumeGlobalPermission field.
func (o *FiltersSnapshot) SetPermissionsToCreateVolumeGlobalPermission(v bool) {
	o.PermissionsToCreateVolumeGlobalPermission = &v
}

// GetProgresses returns the Progresses field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetProgresses() []int64 {
	if o == nil || o.Progresses == nil {
		var ret []int64
		return ret
	}
	return *o.Progresses
}

// GetProgressesOk returns a tuple with the Progresses field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetProgressesOk() ([]int64, bool) {
	if o == nil || o.Progresses == nil {
		var ret []int64
		return ret, false
	}
	return *o.Progresses, true
}

// HasProgresses returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasProgresses() bool {
	if o != nil && o.Progresses != nil {
		return true
	}

	return false
}

// SetProgresses gets a reference to the given []int64 and assigns it to the Progresses field.
func (o *FiltersSnapshot) SetProgresses(v []int64) {
	o.Progresses = &v
}

// GetSnapshotIds returns the SnapshotIds field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetSnapshotIds() []string {
	if o == nil || o.SnapshotIds == nil {
		var ret []string
		return ret
	}
	return *o.SnapshotIds
}

// GetSnapshotIdsOk returns a tuple with the SnapshotIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetSnapshotIdsOk() ([]string, bool) {
	if o == nil || o.SnapshotIds == nil {
		var ret []string
		return ret, false
	}
	return *o.SnapshotIds, true
}

// HasSnapshotIds returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasSnapshotIds() bool {
	if o != nil && o.SnapshotIds != nil {
		return true
	}

	return false
}

// SetSnapshotIds gets a reference to the given []string and assigns it to the SnapshotIds field.
func (o *FiltersSnapshot) SetSnapshotIds(v []string) {
	o.SnapshotIds = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetStatesOk() ([]string, bool) {
	if o == nil || o.States == nil {
		var ret []string
		return ret, false
	}
	return *o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *FiltersSnapshot) SetStates(v []string) {
	o.States = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetTagKeysOk() ([]string, bool) {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret, false
	}
	return *o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersSnapshot) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetTagValuesOk() ([]string, bool) {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret, false
	}
	return *o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersSnapshot) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersSnapshot) SetTags(v []string) {
	o.Tags = &v
}

// GetVolumeIds returns the VolumeIds field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetVolumeIds() []string {
	if o == nil || o.VolumeIds == nil {
		var ret []string
		return ret
	}
	return *o.VolumeIds
}

// GetVolumeIdsOk returns a tuple with the VolumeIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetVolumeIdsOk() ([]string, bool) {
	if o == nil || o.VolumeIds == nil {
		var ret []string
		return ret, false
	}
	return *o.VolumeIds, true
}

// HasVolumeIds returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasVolumeIds() bool {
	if o != nil && o.VolumeIds != nil {
		return true
	}

	return false
}

// SetVolumeIds gets a reference to the given []string and assigns it to the VolumeIds field.
func (o *FiltersSnapshot) SetVolumeIds(v []string) {
	o.VolumeIds = &v
}

// GetVolumeSizes returns the VolumeSizes field value if set, zero value otherwise.
func (o *FiltersSnapshot) GetVolumeSizes() []int64 {
	if o == nil || o.VolumeSizes == nil {
		var ret []int64
		return ret
	}
	return *o.VolumeSizes
}

// GetVolumeSizesOk returns a tuple with the VolumeSizes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSnapshot) GetVolumeSizesOk() ([]int64, bool) {
	if o == nil || o.VolumeSizes == nil {
		var ret []int64
		return ret, false
	}
	return *o.VolumeSizes, true
}

// HasVolumeSizes returns a boolean if a field has been set.
func (o *FiltersSnapshot) HasVolumeSizes() bool {
	if o != nil && o.VolumeSizes != nil {
		return true
	}

	return false
}

// SetVolumeSizes gets a reference to the given []int64 and assigns it to the VolumeSizes field.
func (o *FiltersSnapshot) SetVolumeSizes(v []int64) {
	o.VolumeSizes = &v
}

type NullableFiltersSnapshot struct {
	Value        FiltersSnapshot
	ExplicitNull bool
}

func (v NullableFiltersSnapshot) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersSnapshot) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}