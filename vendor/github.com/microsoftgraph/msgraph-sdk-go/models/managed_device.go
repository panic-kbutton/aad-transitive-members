package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevice devices that are managed or pre-enrolled through Intune
type ManagedDevice struct {
    Entity
    // Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
    activationLockBypassCode *string
    // Android security patch level. This property is read-only.
    androidSecurityPatchLevel *string
    // The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
    azureADDeviceId *string
    // Whether the device is Azure Active Directory registered. This property is read-only.
    azureADRegistered *bool
    // The DateTime when device compliance grace period expires. This property is read-only.
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Compliance state.
    complianceState *ComplianceState
    // ConfigrMgr client enabled features. This property is read-only.
    configurationManagerClientEnabledFeatures ConfigurationManagerClientEnabledFeaturesable
    // List of ComplexType deviceActionResult objects. This property is read-only.
    deviceActionResults []DeviceActionResultable
    // Device category
    deviceCategory DeviceCategoryable
    // Device category display name. This property is read-only.
    deviceCategoryDisplayName *string
    // Device compliance policy states for this device.
    deviceCompliancePolicyStates []DeviceCompliancePolicyStateable
    // Device configuration states for this device.
    deviceConfigurationStates []DeviceConfigurationStateable
    // Possible ways of adding a mobile device to management.
    deviceEnrollmentType *DeviceEnrollmentType
    // The device health attestation state. This property is read-only.
    deviceHealthAttestationState DeviceHealthAttestationStateable
    // Name of the device. This property is read-only.
    deviceName *string
    // Device registration status.
    deviceRegistrationState *DeviceRegistrationState
    // Whether the device is Exchange ActiveSync activated. This property is read-only.
    easActivated *bool
    // Exchange ActivationSync activation time of the device. This property is read-only.
    easActivationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Exchange ActiveSync Id of the device. This property is read-only.
    easDeviceId *string
    // Email(s) for the user associated with the device. This property is read-only.
    emailAddress *string
    // Enrollment time of the device. This property is read-only.
    enrolledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Ethernet MAC. This property is read-only.
    ethernetMacAddress *string
    // Device Exchange Access State.
    exchangeAccessState *DeviceManagementExchangeAccessState
    // Device Exchange Access State Reason.
    exchangeAccessStateReason *DeviceManagementExchangeAccessStateReason
    // Last time the device contacted Exchange. This property is read-only.
    exchangeLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Free Storage in Bytes. This property is read-only.
    freeStorageSpaceInBytes *int64
    // Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
    iccid *string
    // IMEI. This property is read-only.
    imei *string
    // Device encryption status. This property is read-only.
    isEncrypted *bool
    // Device supervised status. This property is read-only.
    isSupervised *bool
    // whether the device is jail broken or rooted. This property is read-only.
    jailBroken *string
    // The date and time that the device last completed a successful sync with Intune. This property is read-only.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Automatically generated name to identify a device. Can be overwritten to a user friendly name.
    managedDeviceName *string
    // Owner type of device.
    managedDeviceOwnerType *ManagedDeviceOwnerType
    // The managementAgent property
    managementAgent *ManagementAgentType
    // Reports device management certificate expiration date. This property is read-only.
    managementCertificateExpirationDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Manufacturer of the device. This property is read-only.
    manufacturer *string
    // MEID. This property is read-only.
    meid *string
    // Model of the device. This property is read-only.
    model *string
    // Notes on the device created by IT Admin
    notes *string
    // Operating system of the device. Windows, iOS, etc. This property is read-only.
    operatingSystem *string
    // Operating system version of the device. This property is read-only.
    osVersion *string
    // Available health states for the Device Health API
    partnerReportedThreatState *ManagedDevicePartnerReportedHealthState
    // Phone number of the device. This property is read-only.
    phoneNumber *string
    // Total Memory in Bytes. This property is read-only.
    physicalMemoryInBytes *int64
    // An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
    remoteAssistanceSessionErrorDetails *string
    // Url that allows a Remote Assistance session to be established with the device. This property is read-only.
    remoteAssistanceSessionUrl *string
    // Reports if the managed iOS device is user approval enrollment. This property is read-only.
    requireUserEnrollmentApproval *bool
    // SerialNumber. This property is read-only.
    serialNumber *string
    // Subscriber Carrier. This property is read-only.
    subscriberCarrier *string
    // Total Storage in Bytes. This property is read-only.
    totalStorageSpaceInBytes *int64
    // Unique Device Identifier for iOS and macOS devices. This property is read-only.
    udid *string
    // User display name. This property is read-only.
    userDisplayName *string
    // Unique Identifier for the user associated with the device. This property is read-only.
    userId *string
    // Device user principal name. This property is read-only.
    userPrincipalName *string
    // The primary users associated with the managed device.
    users []Userable
    // Wi-Fi MAC. This property is read-only.
    wiFiMacAddress *string
}
// NewManagedDevice instantiates a new managedDevice and sets the default values.
func NewManagedDevice()(*ManagedDevice) {
    m := &ManagedDevice{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedDevice";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagedDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevice(), nil
}
// GetActivationLockBypassCode gets the activationLockBypassCode property value. Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
func (m *ManagedDevice) GetActivationLockBypassCode()(*string) {
    return m.activationLockBypassCode
}
// GetAndroidSecurityPatchLevel gets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) GetAndroidSecurityPatchLevel()(*string) {
    return m.androidSecurityPatchLevel
}
// GetAzureADDeviceId gets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureADDeviceId()(*string) {
    return m.azureADDeviceId
}
// GetAzureADRegistered gets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) GetAzureADRegistered()(*bool) {
    return m.azureADRegistered
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.complianceGracePeriodExpirationDateTime
}
// GetComplianceState gets the complianceState property value. Compliance state.
func (m *ManagedDevice) GetComplianceState()(*ComplianceState) {
    return m.complianceState
}
// GetConfigurationManagerClientEnabledFeatures gets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) GetConfigurationManagerClientEnabledFeatures()(ConfigurationManagerClientEnabledFeaturesable) {
    return m.configurationManagerClientEnabledFeatures
}
// GetDeviceActionResults gets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) GetDeviceActionResults()([]DeviceActionResultable) {
    return m.deviceActionResults
}
// GetDeviceCategory gets the deviceCategory property value. Device category
func (m *ManagedDevice) GetDeviceCategory()(DeviceCategoryable) {
    return m.deviceCategory
}
// GetDeviceCategoryDisplayName gets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) GetDeviceCategoryDisplayName()(*string) {
    return m.deviceCategoryDisplayName
}
// GetDeviceCompliancePolicyStates gets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyStateable) {
    return m.deviceCompliancePolicyStates
}
// GetDeviceConfigurationStates gets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) GetDeviceConfigurationStates()([]DeviceConfigurationStateable) {
    return m.deviceConfigurationStates
}
// GetDeviceEnrollmentType gets the deviceEnrollmentType property value. Possible ways of adding a mobile device to management.
func (m *ManagedDevice) GetDeviceEnrollmentType()(*DeviceEnrollmentType) {
    return m.deviceEnrollmentType
}
// GetDeviceHealthAttestationState gets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) GetDeviceHealthAttestationState()(DeviceHealthAttestationStateable) {
    return m.deviceHealthAttestationState
}
// GetDeviceName gets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceRegistrationState gets the deviceRegistrationState property value. Device registration status.
func (m *ManagedDevice) GetDeviceRegistrationState()(*DeviceRegistrationState) {
    return m.deviceRegistrationState
}
// GetEasActivated gets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) GetEasActivated()(*bool) {
    return m.easActivated
}
// GetEasActivationDateTime gets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.easActivationDateTime
}
// GetEasDeviceId gets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) GetEasDeviceId()(*string) {
    return m.easDeviceId
}
// GetEmailAddress gets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetEnrolledDateTime gets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.enrolledDateTime
}
// GetEthernetMacAddress gets the ethernetMacAddress property value. Ethernet MAC. This property is read-only.
func (m *ManagedDevice) GetEthernetMacAddress()(*string) {
    return m.ethernetMacAddress
}
// GetExchangeAccessState gets the exchangeAccessState property value. Device Exchange Access State.
func (m *ManagedDevice) GetExchangeAccessState()(*DeviceManagementExchangeAccessState) {
    return m.exchangeAccessState
}
// GetExchangeAccessStateReason gets the exchangeAccessStateReason property value. Device Exchange Access State Reason.
func (m *ManagedDevice) GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason) {
    return m.exchangeAccessStateReason
}
// GetExchangeLastSuccessfulSyncDateTime gets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.exchangeLastSuccessfulSyncDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationLockBypassCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActivationLockBypassCode)
    res["androidSecurityPatchLevel"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAndroidSecurityPatchLevel)
    res["azureADDeviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureADDeviceId)
    res["azureADRegistered"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAzureADRegistered)
    res["complianceGracePeriodExpirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetComplianceGracePeriodExpirationDateTime)
    res["complianceState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseComplianceState , m.SetComplianceState)
    res["configurationManagerClientEnabledFeatures"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConfigurationManagerClientEnabledFeaturesFromDiscriminatorValue , m.SetConfigurationManagerClientEnabledFeatures)
    res["deviceActionResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceActionResultFromDiscriminatorValue , m.SetDeviceActionResults)
    res["deviceCategory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceCategoryFromDiscriminatorValue , m.SetDeviceCategory)
    res["deviceCategoryDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceCategoryDisplayName)
    res["deviceCompliancePolicyStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceCompliancePolicyStateFromDiscriminatorValue , m.SetDeviceCompliancePolicyStates)
    res["deviceConfigurationStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceConfigurationStateFromDiscriminatorValue , m.SetDeviceConfigurationStates)
    res["deviceEnrollmentType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceEnrollmentType , m.SetDeviceEnrollmentType)
    res["deviceHealthAttestationState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceHealthAttestationStateFromDiscriminatorValue , m.SetDeviceHealthAttestationState)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["deviceRegistrationState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceRegistrationState , m.SetDeviceRegistrationState)
    res["easActivated"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEasActivated)
    res["easActivationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEasActivationDateTime)
    res["easDeviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEasDeviceId)
    res["emailAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailAddress)
    res["enrolledDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEnrolledDateTime)
    res["ethernetMacAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEthernetMacAddress)
    res["exchangeAccessState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementExchangeAccessState , m.SetExchangeAccessState)
    res["exchangeAccessStateReason"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementExchangeAccessStateReason , m.SetExchangeAccessStateReason)
    res["exchangeLastSuccessfulSyncDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExchangeLastSuccessfulSyncDateTime)
    res["freeStorageSpaceInBytes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetFreeStorageSpaceInBytes)
    res["iccid"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIccid)
    res["imei"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetImei)
    res["isEncrypted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEncrypted)
    res["isSupervised"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSupervised)
    res["jailBroken"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJailBroken)
    res["lastSyncDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastSyncDateTime)
    res["managedDeviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedDeviceName)
    res["managedDeviceOwnerType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedDeviceOwnerType , m.SetManagedDeviceOwnerType)
    res["managementAgent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementAgentType , m.SetManagementAgent)
    res["managementCertificateExpirationDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetManagementCertificateExpirationDate)
    res["manufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManufacturer)
    res["meid"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMeid)
    res["model"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModel)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotes)
    res["operatingSystem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOperatingSystem)
    res["osVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsVersion)
    res["partnerReportedThreatState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedDevicePartnerReportedHealthState , m.SetPartnerReportedThreatState)
    res["phoneNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPhoneNumber)
    res["physicalMemoryInBytes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetPhysicalMemoryInBytes)
    res["remoteAssistanceSessionErrorDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRemoteAssistanceSessionErrorDetails)
    res["remoteAssistanceSessionUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRemoteAssistanceSessionUrl)
    res["requireUserEnrollmentApproval"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequireUserEnrollmentApproval)
    res["serialNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSerialNumber)
    res["subscriberCarrier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubscriberCarrier)
    res["totalStorageSpaceInBytes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetTotalStorageSpaceInBytes)
    res["udid"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUdid)
    res["userDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserDisplayName)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    res["users"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserFromDiscriminatorValue , m.SetUsers)
    res["wiFiMacAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWiFiMacAddress)
    return res
}
// GetFreeStorageSpaceInBytes gets the freeStorageSpaceInBytes property value. Free Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetFreeStorageSpaceInBytes()(*int64) {
    return m.freeStorageSpaceInBytes
}
// GetIccid gets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
func (m *ManagedDevice) GetIccid()(*string) {
    return m.iccid
}
// GetImei gets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) GetImei()(*string) {
    return m.imei
}
// GetIsEncrypted gets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) GetIsEncrypted()(*bool) {
    return m.isEncrypted
}
// GetIsSupervised gets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) GetIsSupervised()(*bool) {
    return m.isSupervised
}
// GetJailBroken gets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) GetJailBroken()(*string) {
    return m.jailBroken
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetManagedDeviceName gets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetManagedDeviceOwnerType gets the managedDeviceOwnerType property value. Owner type of device.
func (m *ManagedDevice) GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType) {
    return m.managedDeviceOwnerType
}
// GetManagementAgent gets the managementAgent property value. The managementAgent property
func (m *ManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    return m.managementAgent
}
// GetManagementCertificateExpirationDate gets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) GetManagementCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.managementCertificateExpirationDate
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetMeid gets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) GetMeid()(*string) {
    return m.meid
}
// GetModel gets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) GetModel()(*string) {
    return m.model
}
// GetNotes gets the notes property value. Notes on the device created by IT Admin
func (m *ManagedDevice) GetNotes()(*string) {
    return m.notes
}
// GetOperatingSystem gets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOsVersion gets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) GetOsVersion()(*string) {
    return m.osVersion
}
// GetPartnerReportedThreatState gets the partnerReportedThreatState property value. Available health states for the Device Health API
func (m *ManagedDevice) GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState) {
    return m.partnerReportedThreatState
}
// GetPhoneNumber gets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPhysicalMemoryInBytes gets the physicalMemoryInBytes property value. Total Memory in Bytes. This property is read-only.
func (m *ManagedDevice) GetPhysicalMemoryInBytes()(*int64) {
    return m.physicalMemoryInBytes
}
// GetRemoteAssistanceSessionErrorDetails gets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionErrorDetails()(*string) {
    return m.remoteAssistanceSessionErrorDetails
}
// GetRemoteAssistanceSessionUrl gets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionUrl()(*string) {
    return m.remoteAssistanceSessionUrl
}
// GetRequireUserEnrollmentApproval gets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) GetRequireUserEnrollmentApproval()(*bool) {
    return m.requireUserEnrollmentApproval
}
// GetSerialNumber gets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetSubscriberCarrier gets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) GetSubscriberCarrier()(*string) {
    return m.subscriberCarrier
}
// GetTotalStorageSpaceInBytes gets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetTotalStorageSpaceInBytes()(*int64) {
    return m.totalStorageSpaceInBytes
}
// GetUdid gets the udid property value. Unique Device Identifier for iOS and macOS devices. This property is read-only.
func (m *ManagedDevice) GetUdid()(*string) {
    return m.udid
}
// GetUserDisplayName gets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUsers gets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) GetUsers()([]Userable) {
    return m.users
}
// GetWiFiMacAddress gets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) GetWiFiMacAddress()(*string) {
    return m.wiFiMacAddress
}
// Serialize serializes information the current object
func (m *ManagedDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetComplianceState() != nil {
        cast := (*m.GetComplianceState()).String()
        err = writer.WriteStringValue("complianceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceCategory", m.GetDeviceCategory())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicyStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceCompliancePolicyStates())
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicyStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceConfigurationStates())
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentType() != nil {
        cast := (*m.GetDeviceEnrollmentType()).String()
        err = writer.WriteStringValue("deviceEnrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceRegistrationState() != nil {
        cast := (*m.GetDeviceRegistrationState()).String()
        err = writer.WriteStringValue("deviceRegistrationState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessState() != nil {
        cast := (*m.GetExchangeAccessState()).String()
        err = writer.WriteStringValue("exchangeAccessState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessStateReason() != nil {
        cast := (*m.GetExchangeAccessStateReason()).String()
        err = writer.WriteStringValue("exchangeAccessStateReason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceOwnerType() != nil {
        cast := (*m.GetManagedDeviceOwnerType()).String()
        err = writer.WriteStringValue("managedDeviceOwnerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementAgent() != nil {
        cast := (*m.GetManagementAgent()).String()
        err = writer.WriteStringValue("managementAgent", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerReportedThreatState() != nil {
        cast := (*m.GetPartnerReportedThreatState()).String()
        err = writer.WriteStringValue("partnerReportedThreatState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUsers())
        err = writer.WriteCollectionOfObjectValues("users", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationLockBypassCode sets the activationLockBypassCode property value. Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
func (m *ManagedDevice) SetActivationLockBypassCode(value *string)() {
    m.activationLockBypassCode = value
}
// SetAndroidSecurityPatchLevel sets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) SetAndroidSecurityPatchLevel(value *string)() {
    m.androidSecurityPatchLevel = value
}
// SetAzureADDeviceId sets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureADDeviceId(value *string)() {
    m.azureADDeviceId = value
}
// SetAzureADRegistered sets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) SetAzureADRegistered(value *bool)() {
    m.azureADRegistered = value
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceGracePeriodExpirationDateTime = value
}
// SetComplianceState sets the complianceState property value. Compliance state.
func (m *ManagedDevice) SetComplianceState(value *ComplianceState)() {
    m.complianceState = value
}
// SetConfigurationManagerClientEnabledFeatures sets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(value ConfigurationManagerClientEnabledFeaturesable)() {
    m.configurationManagerClientEnabledFeatures = value
}
// SetDeviceActionResults sets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) SetDeviceActionResults(value []DeviceActionResultable)() {
    m.deviceActionResults = value
}
// SetDeviceCategory sets the deviceCategory property value. Device category
func (m *ManagedDevice) SetDeviceCategory(value DeviceCategoryable)() {
    m.deviceCategory = value
}
// SetDeviceCategoryDisplayName sets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) SetDeviceCategoryDisplayName(value *string)() {
    m.deviceCategoryDisplayName = value
}
// SetDeviceCompliancePolicyStates sets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyStateable)() {
    m.deviceCompliancePolicyStates = value
}
// SetDeviceConfigurationStates sets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) SetDeviceConfigurationStates(value []DeviceConfigurationStateable)() {
    m.deviceConfigurationStates = value
}
// SetDeviceEnrollmentType sets the deviceEnrollmentType property value. Possible ways of adding a mobile device to management.
func (m *ManagedDevice) SetDeviceEnrollmentType(value *DeviceEnrollmentType)() {
    m.deviceEnrollmentType = value
}
// SetDeviceHealthAttestationState sets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) SetDeviceHealthAttestationState(value DeviceHealthAttestationStateable)() {
    m.deviceHealthAttestationState = value
}
// SetDeviceName sets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceRegistrationState sets the deviceRegistrationState property value. Device registration status.
func (m *ManagedDevice) SetDeviceRegistrationState(value *DeviceRegistrationState)() {
    m.deviceRegistrationState = value
}
// SetEasActivated sets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) SetEasActivated(value *bool)() {
    m.easActivated = value
}
// SetEasActivationDateTime sets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.easActivationDateTime = value
}
// SetEasDeviceId sets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) SetEasDeviceId(value *string)() {
    m.easDeviceId = value
}
// SetEmailAddress sets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetEnrolledDateTime sets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrolledDateTime = value
}
// SetEthernetMacAddress sets the ethernetMacAddress property value. Ethernet MAC. This property is read-only.
func (m *ManagedDevice) SetEthernetMacAddress(value *string)() {
    m.ethernetMacAddress = value
}
// SetExchangeAccessState sets the exchangeAccessState property value. Device Exchange Access State.
func (m *ManagedDevice) SetExchangeAccessState(value *DeviceManagementExchangeAccessState)() {
    m.exchangeAccessState = value
}
// SetExchangeAccessStateReason sets the exchangeAccessStateReason property value. Device Exchange Access State Reason.
func (m *ManagedDevice) SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)() {
    m.exchangeAccessStateReason = value
}
// SetExchangeLastSuccessfulSyncDateTime sets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.exchangeLastSuccessfulSyncDateTime = value
}
// SetFreeStorageSpaceInBytes sets the freeStorageSpaceInBytes property value. Free Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetFreeStorageSpaceInBytes(value *int64)() {
    m.freeStorageSpaceInBytes = value
}
// SetIccid sets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
func (m *ManagedDevice) SetIccid(value *string)() {
    m.iccid = value
}
// SetImei sets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) SetImei(value *string)() {
    m.imei = value
}
// SetIsEncrypted sets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// SetIsSupervised sets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
// SetJailBroken sets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) SetJailBroken(value *string)() {
    m.jailBroken = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetManagedDeviceName sets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetManagedDeviceOwnerType sets the managedDeviceOwnerType property value. Owner type of device.
func (m *ManagedDevice) SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)() {
    m.managedDeviceOwnerType = value
}
// SetManagementAgent sets the managementAgent property value. The managementAgent property
func (m *ManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    m.managementAgent = value
}
// SetManagementCertificateExpirationDate sets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) SetManagementCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.managementCertificateExpirationDate = value
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetMeid sets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) SetMeid(value *string)() {
    m.meid = value
}
// SetModel sets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) SetModel(value *string)() {
    m.model = value
}
// SetNotes sets the notes property value. Notes on the device created by IT Admin
func (m *ManagedDevice) SetNotes(value *string)() {
    m.notes = value
}
// SetOperatingSystem sets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOsVersion sets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetPartnerReportedThreatState sets the partnerReportedThreatState property value. Available health states for the Device Health API
func (m *ManagedDevice) SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)() {
    m.partnerReportedThreatState = value
}
// SetPhoneNumber sets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPhysicalMemoryInBytes sets the physicalMemoryInBytes property value. Total Memory in Bytes. This property is read-only.
func (m *ManagedDevice) SetPhysicalMemoryInBytes(value *int64)() {
    m.physicalMemoryInBytes = value
}
// SetRemoteAssistanceSessionErrorDetails sets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(value *string)() {
    m.remoteAssistanceSessionErrorDetails = value
}
// SetRemoteAssistanceSessionUrl sets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionUrl(value *string)() {
    m.remoteAssistanceSessionUrl = value
}
// SetRequireUserEnrollmentApproval sets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) SetRequireUserEnrollmentApproval(value *bool)() {
    m.requireUserEnrollmentApproval = value
}
// SetSerialNumber sets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetSubscriberCarrier sets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) SetSubscriberCarrier(value *string)() {
    m.subscriberCarrier = value
}
// SetTotalStorageSpaceInBytes sets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetTotalStorageSpaceInBytes(value *int64)() {
    m.totalStorageSpaceInBytes = value
}
// SetUdid sets the udid property value. Unique Device Identifier for iOS and macOS devices. This property is read-only.
func (m *ManagedDevice) SetUdid(value *string)() {
    m.udid = value
}
// SetUserDisplayName sets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUsers sets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) SetUsers(value []Userable)() {
    m.users = value
}
// SetWiFiMacAddress sets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) SetWiFiMacAddress(value *string)() {
    m.wiFiMacAddress = value
}
