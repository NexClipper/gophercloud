package nodes

import (
	"github.com/nexclipper/gophercloud"
	"github.com/nexclipper/gophercloud/pagination"
)

type nodeResult struct {
	gophercloud.Result
}

// Extract interprets any nodeResult as a Node, if possible.
func (r nodeResult) Extract() (*Node, error) {
	var s Node
	err := r.ExtractInto(&s)
	return &s, err
}

// Extract interprets a BootDeviceResult as BootDeviceOpts, if possible.
func (r BootDeviceResult) Extract() (*BootDeviceOpts, error) {
	var s BootDeviceOpts
	err := r.ExtractInto(&s)
	return &s, err
}

// Extract interprets a SupportedBootDeviceResult as an array of supported boot devices, if possible.
func (r SupportedBootDeviceResult) Extract() ([]string, error) {
	var s struct {
		Devices []string `json:"supported_boot_devices"`
	}

	err := r.ExtractInto(&s)
	return s.Devices, err
}

// Extract interprets a ValidateResult as NodeValidation, if possible.
func (r ValidateResult) Extract() (*NodeValidation, error) {
	var s NodeValidation
	err := r.ExtractInto(&s)
	return &s, err
}

func (r nodeResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

func ExtractNodesInto(r pagination.Page, v interface{}) error {
	return r.(NodePage).Result.ExtractIntoSlicePtr(v, "nodes")
}

// Extract interprets a BIOSSettingsResult as an array of BIOSSetting structs, if possible.
func (r ListBIOSSettingsResult) Extract() ([]BIOSSetting, error) {
	var s struct {
		Settings []BIOSSetting `json:"bios"`
	}

	err := r.ExtractInto(&s)
	return s.Settings, err
}

// Extract interprets a SingleBIOSSettingResult as a BIOSSetting struct, if possible.
func (r GetBIOSSettingResult) Extract() (*BIOSSetting, error) {
	var s SingleBIOSSetting
	err := r.ExtractInto(&s)
	return &s.Setting, err
}

// Extract interprets a VendorPassthruMethod as
func (r VendorPassthruMethodsResult) Extract() (*VendorPassthruMethods, error) {
	var s VendorPassthruMethods
	err := r.ExtractInto(&s)
	return &s, err
}

func (r GetAllSubscriptionsVendorPassthruResult) Extract() (*GetAllSubscriptionsVendorPassthru, error) {
	var s GetAllSubscriptionsVendorPassthru
	err := r.ExtractInto(&s)
	return &s, err
}

func (r SubscriptionVendorPassthruResult) Extract() (*SubscriptionVendorPassthru, error) {
	var s SubscriptionVendorPassthru
	err := r.ExtractInto(&s)
	return &s, err
}

// Node represents a node in the OpenStack Bare Metal API.
type Node struct {
	// Whether automated cleaning is enabled or disabled on this node.
	// Requires microversion 1.47 or later.
	AutomatedClean *bool `json:"automated_clean"`

	// UUID for the resource.
	UUID string `json:"uuid"`

	// Identifier for the Node resource. May be undefined. Certain words are reserved.
	Name string `json:"name"`

	// Current power state of this Node. Usually, “power on” or “power off”, but may be “None”
	// if Ironic is unable to determine the power state (eg, due to hardware failure).
	PowerState string `json:"power_state"`

	// A power state transition has been requested, this field represents the requested (ie, “target”)
	// state either “power on”, “power off”, “rebooting”, “soft power off” or “soft rebooting”.
	TargetPowerState string `json:"target_power_state"`

	// Current provisioning state of this Node.
	ProvisionState string `json:"provision_state"`

	// A provisioning action has been requested, this field represents the requested (ie, “target”) state. Note
	// that a Node may go through several states during its transition to this target state. For instance, when
	// requesting an instance be deployed to an AVAILABLE Node, the Node may go through the following state
	// change progression: AVAILABLE -> DEPLOYING -> DEPLOYWAIT -> DEPLOYING -> ACTIVE
	TargetProvisionState string `json:"target_provision_state"`

	// Whether or not this Node is currently in “maintenance mode”. Setting a Node into maintenance mode removes it
	// from the available resource pool and halts some internal automation. This can happen manually (eg, via an API
	// request) or automatically when Ironic detects a hardware fault that prevents communication with the machine.
	Maintenance bool `json:"maintenance"`

	// Description of the reason why this Node was placed into maintenance mode
	MaintenanceReason string `json:"maintenance_reason"`

	// Fault indicates the active fault detected by ironic, typically the Node is in “maintenance mode”. None means no
	// fault has been detected by ironic. “power failure” indicates ironic failed to retrieve power state from this
	// node. There are other possible types, e.g., “clean failure” and “rescue abort failure”.
	Fault string `json:"fault"`

	// Error from the most recent (last) transaction that started but failed to finish.
	LastError string `json:"last_error"`

	// Name of an Ironic Conductor host which is holding a lock on this node, if a lock is held. Usually “null”,
	// but this field can be useful for debugging.
	Reservation string `json:"reservation"`

	// Name of the driver.
	Driver string `json:"driver"`

	// The metadata required by the driver to manage this Node. List of fields varies between drivers, and can be
	// retrieved from the /v1/drivers/<DRIVER_NAME>/properties resource.
	DriverInfo map[string]interface{} `json:"driver_info"`

	// Metadata set and stored by the Node’s driver. This field is read-only.
	DriverInternalInfo map[string]interface{} `json:"driver_internal_info"`

	// Characteristics of this Node. Populated by ironic-inspector during inspection. May be edited via the REST
	// API at any time.
	Properties map[string]interface{} `json:"properties"`

	// Used to customize the deployed image. May include root partition size, a base 64 encoded config drive, and other
	// metadata. Note that this field is erased automatically when the instance is deleted (this is done by requesting
	// the Node provision state be changed to DELETED).
	InstanceInfo map[string]interface{} `json:"instance_info"`

	// ID of the Nova instance associated with this Node.
	InstanceUUID string `json:"instance_uuid"`

	// ID of the chassis associated with this Node. May be empty or None.
	ChassisUUID string `json:"chassis_uuid"`

	// Set of one or more arbitrary metadata key and value pairs.
	Extra map[string]interface{} `json:"extra"`

	// Whether console access is enabled or disabled on this node.
	ConsoleEnabled bool `json:"console_enabled"`

	// The current RAID configuration of the node. Introduced with the cleaning feature.
	RAIDConfig map[string]interface{} `json:"raid_config"`

	// The requested RAID configuration of the node, which will be applied when the Node next transitions
	// through the CLEANING state. Introduced with the cleaning feature.
	TargetRAIDConfig map[string]interface{} `json:"target_raid_config"`

	// Current clean step. Introduced with the cleaning feature.
	CleanStep map[string]interface{} `json:"clean_step"`

	// Current deploy step.
	DeployStep map[string]interface{} `json:"deploy_step"`

	// String which can be used by external schedulers to identify this Node as a unit of a specific type of resource.
	// For more details, see: https://docs.openstack.org/ironic/latest/install/configure-nova-flavors.html
	ResourceClass string `json:"resource_class"`

	// BIOS interface for a Node, e.g. “redfish”.
	BIOSInterface string `json:"bios_interface"`

	// Boot interface for a Node, e.g. “pxe”.
	BootInterface string `json:"boot_interface"`

	// Console interface for a node, e.g. “no-console”.
	ConsoleInterface string `json:"console_interface"`

	// Deploy interface for a node, e.g. “iscsi”.
	DeployInterface string `json:"deploy_interface"`

	// Interface used for node inspection, e.g. “no-inspect”.
	InspectInterface string `json:"inspect_interface"`

	// For out-of-band node management, e.g. “ipmitool”.
	ManagementInterface string `json:"management_interface"`

	// Network Interface provider to use when plumbing the network connections for this Node.
	NetworkInterface string `json:"network_interface"`

	// used for performing power actions on the node, e.g. “ipmitool”.
	PowerInterface string `json:"power_interface"`

	// Used for configuring RAID on this node, e.g. “no-raid”.
	RAIDInterface string `json:"raid_interface"`

	// Interface used for node rescue, e.g. “no-rescue”.
	RescueInterface string `json:"rescue_interface"`

	// Used for attaching and detaching volumes on this node, e.g. “cinder”.
	StorageInterface string `json:"storage_interface"`

	// Array of traits for this node.
	Traits []string `json:"traits"`

	// For vendor-specific functionality on this node, e.g. “no-vendor”.
	VendorInterface string `json:"vendor_interface"`

	// Conductor group for a node. Case-insensitive string up to 255 characters, containing a-z, 0-9, _, -, and ..
	ConductorGroup string `json:"conductor_group"`

	// The node is protected from undeploying, rebuilding and deletion.
	Protected bool `json:"protected"`

	// Reason the node is marked as protected.
	ProtectedReason string `json:"protected_reason"`

	// A string or UUID of the tenant who owns the baremetal node.
	Owner string `json:"owner"`

	// Static network configuration to use during deployment and cleaning.
	NetworkData map[string]interface{} `json:"network_data"`
}

// NodePage abstracts the raw results of making a List() request against
// the API. As OpenStack extensions may freely alter the response bodies of
// structures returned to the client, you may only safely access the data
// provided through the ExtractNodes call.
type NodePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a page contains no Node results.
func (r NodePage) IsEmpty() (bool, error) {
	s, err := ExtractNodes(r)
	return len(s) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (r NodePage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"nodes_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// ExtractNodes interprets the results of a single page from a List() call,
// producing a slice of Node entities.
func ExtractNodes(r pagination.Page) ([]Node, error) {
	var s []Node
	err := ExtractNodesInto(r, &s)
	return s, err
}

// GetResult is the response from a Get operation. Call its Extract
// method to interpret it as a Node.
type GetResult struct {
	nodeResult
}

// CreateResult is the response from a Create operation.
type CreateResult struct {
	nodeResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as a Node.
type UpdateResult struct {
	nodeResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ValidateResult is the response from a Validate operation. Call its Extract
// method to interpret it as a NodeValidation struct.
type ValidateResult struct {
	gophercloud.Result
}

// InjectNMIResult is the response from an InjectNMI operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type InjectNMIResult struct {
	gophercloud.ErrResult
}

// BootDeviceResult is the response from a GetBootDevice operation. Call its Extract
// method to interpret it as a BootDeviceOpts struct.
type BootDeviceResult struct {
	gophercloud.Result
}

// SetBootDeviceResult is the response from a SetBootDevice operation. Call its Extract
// method to interpret it as a BootDeviceOpts struct.
type SetBootDeviceResult struct {
	gophercloud.ErrResult
}

// SupportedBootDeviceResult is the response from a GetSupportedBootDevices operation. Call its Extract
// method to interpret it as an array of supported boot device values.
type SupportedBootDeviceResult struct {
	gophercloud.Result
}

// ChangePowerStateResult is the response from a ChangePowerState operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type ChangePowerStateResult struct {
	gophercloud.ErrResult
}

// ListBIOSSettingsResult is the response from a ListBIOSSettings operation. Call its Extract
// method to interpret it as an array of BIOSSetting structs.
type ListBIOSSettingsResult struct {
	gophercloud.Result
}

// GetBIOSSettingResult is the response from a GetBIOSSetting operation. Call its Extract
// method to interpret it as a BIOSSetting struct.
type GetBIOSSettingResult struct {
	gophercloud.Result
}

// VendorPassthruMethodsResult is the response from a GetVendorPassthruMethods operation. Call its Extract
// method to interpret it as an array of allowed vendor methods.
type VendorPassthruMethodsResult struct {
	gophercloud.Result
}

// GetAllSubscriptionsVendorPassthruResult is the response from GetAllSubscriptions operation. Call its
// Extract method to interpret it as a GetAllSubscriptionsVendorPassthru struct.
type GetAllSubscriptionsVendorPassthruResult struct {
	gophercloud.Result
}

// SubscriptionVendorPassthruResult is the response from GetSubscription and CreateSubscription operation. Call its Extract
// method to interpret it as a SubscriptionVendorPassthru struct.
type SubscriptionVendorPassthruResult struct {
	gophercloud.Result
}

// DeleteSubscriptionVendorPassthruResult is the response from DeleteSubscription operation. Call its
// ExtractErr method to determine if the call succeeded of failed.
type DeleteSubscriptionVendorPassthruResult struct {
	gophercloud.ErrResult
}

// Each element in the response will contain a “result” variable, which will have a value of “true” or “false”, and
// also potentially a reason. A value of nil indicates that the Node’s driver does not support that interface.
type DriverValidation struct {
	Result bool   `json:"result"`
	Reason string `json:"reason"`
}

//  Ironic validates whether the Node’s driver has enough information to manage the Node. This polls each interface on
//  the driver, and returns the status of that interface as an DriverValidation struct.
type NodeValidation struct {
	BIOS       DriverValidation `json:"bios"`
	Boot       DriverValidation `json:"boot"`
	Console    DriverValidation `json:"console"`
	Deploy     DriverValidation `json:"deploy"`
	Inspect    DriverValidation `json:"inspect"`
	Management DriverValidation `json:"management"`
	Network    DriverValidation `json:"network"`
	Power      DriverValidation `json:"power"`
	RAID       DriverValidation `json:"raid"`
	Rescue     DriverValidation `json:"rescue"`
	Storage    DriverValidation `json:"storage"`
}

// A particular BIOS setting for a node in the OpenStack Bare Metal API.
type BIOSSetting struct {

	// Identifier for the BIOS setting.
	Name string `json:"name"`

	// Value of the BIOS setting.
	Value string `json:"value"`

	// The following fields are returned in microversion 1.74 or later
	// when using the `details` option

	// The type of setting - Enumeration, String, Integer, or Boolean.
	AttributeType string `json:"attribute_type"`

	// The allowable value for an Enumeration type setting.
	AllowableValues []string `json:"allowable_values"`

	// The lowest value for an Integer type setting.
	LowerBound *int `json:"lower_bound"`

	// The highest value for an Integer type setting.
	UpperBound *int `json:"upper_bound"`

	// Minimum length for a String type setting.
	MinLength *int `json:"min_length"`

	// Maximum length for a String type setting.
	MaxLength *int `json:"max_length"`

	// Whether or not this setting is read only.
	ReadOnly *bool `json:"read_only"`

	// Whether or not a reset is required after changing this setting.
	ResetRequired *bool `json:"reset_required"`

	// Whether or not this setting's value is unique to this node, e.g.
	// a serial number.
	Unique *bool `json:"unique"`
}

type SingleBIOSSetting struct {
	Setting BIOSSetting
}

// ChangeStateResult is the response from any state change operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type ChangeStateResult struct {
	gophercloud.ErrResult
}

type VendorPassthruMethods struct {
	CreateSubscription  CreateSubscriptionMethod  `json:"create_subscription,omitempty"`
	DeleteSubscription  DeleteSubscriptionMethod  `json:"delete_subscription,omitempty"`
	GetSubscription     GetSubscriptionMethod     `json:"get_subscription,omitempty"`
	GetAllSubscriptions GetAllSubscriptionsMethod `json:"get_all_subscriptions,omitempty"`
}

// Below you can find all vendor passthru methods structs

type CreateSubscriptionMethod struct {
	HTTPMethods          []string `json:"http_methods"`
	Async                bool     `json:"async"`
	Description          string   `json:"description"`
	Attach               bool     `json:"attach"`
	RequireExclusiveLock bool     `json:"require_exclusive_lock"`
}

type DeleteSubscriptionMethod struct {
	HTTPMethods          []string `json:"http_methods"`
	Async                bool     `json:"async"`
	Description          string   `json:"description"`
	Attach               bool     `json:"attach"`
	RequireExclusiveLock bool     `json:"require_exclusive_lock"`
}

type GetSubscriptionMethod struct {
	HTTPMethods          []string `json:"http_methods"`
	Async                bool     `json:"async"`
	Description          string   `json:"description"`
	Attach               bool     `json:"attach"`
	RequireExclusiveLock bool     `json:"require_exclusive_lock"`
}

type GetAllSubscriptionsMethod struct {
	HTTPMethods          []string `json:"http_methods"`
	Async                bool     `json:"async"`
	Description          string   `json:"description"`
	Attach               bool     `json:"attach"`
	RequireExclusiveLock bool     `json:"require_exclusive_lock"`
}

// A List of subscriptions from a node in the OpenStack Bare Metal API.
type GetAllSubscriptionsVendorPassthru struct {
	Context      string              `json:"@odata.context"`
	Etag         string              `json:"@odata.etag"`
	Id           string              `json:"@odata.id"`
	Type         string              `json:"@odata.type"`
	Description  string              `json:"Description"`
	Name         string              `json:"Name"`
	Members      []map[string]string `json:"Members"`
	MembersCount int                 `json:"Members@odata.count"`
}

// A Subscription from a node in the OpenStack Bare Metal API.
type SubscriptionVendorPassthru struct {
	Id          string   `json:"Id"`
	Context     string   `json:"Context"`
	Destination string   `json:"Destination"`
	EventTypes  []string `json:"EventTypes"`
	Protocol    string   `json:"Protocol"`
}
