package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingService represents a particular service offered by a booking business.
type BookingService struct {
    Entity
    // Additional information that is sent to the customer when an appointment is confirmed.
    additionalInformation *string
    // Contains the set of custom questions associated with a particular service.
    customQuestions []BookingQuestionAssignmentable
    // The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
    defaultDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The default physical location for the service.
    defaultLocation Locationable
    // The default monetary price for the service.
    defaultPrice *float64
    // Represents the type of pricing of a booking service.
    defaultPriceType *BookingPriceType
    // The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
    defaultReminders []BookingReminderable
    // A text description for the service.
    description *string
    // A service name.
    displayName *string
    // The isAnonymousJoinEnabled property
    isAnonymousJoinEnabled *bool
    // True means this service is not available to customers for booking.
    isHiddenFromCustomers *bool
    // True indicates that the appointments for the service will be held online. Default value is false.
    isLocationOnline *bool
    // The languageTag property
    languageTag *string
    // The maximum number of customers allowed in a service. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
    maximumAttendeesCount *int32
    // Additional information about this service.
    notes *string
    // The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
    postBuffer *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The time to buffer before an appointment for this service can start.
    preBuffer *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The set of policies that determine how appointments for this type of service should be created and managed.
    schedulingPolicy BookingSchedulingPolicyable
    // True indicates SMS notifications can be sent to the customers for the appointment of the service. Default value is false.
    smsNotificationsEnabled *bool
    // Represents those staff members who provide this service.
    staffMemberIds []string
    // The URL a customer uses to access the service.
    webUrl *string
}
// NewBookingService instantiates a new bookingService and sets the default values.
func NewBookingService()(*BookingService) {
    m := &BookingService{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.bookingService";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBookingServiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingServiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingService(), nil
}
// GetAdditionalInformation gets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
func (m *BookingService) GetAdditionalInformation()(*string) {
    return m.additionalInformation
}
// GetCustomQuestions gets the customQuestions property value. Contains the set of custom questions associated with a particular service.
func (m *BookingService) GetCustomQuestions()([]BookingQuestionAssignmentable) {
    return m.customQuestions
}
// GetDefaultDuration gets the defaultDuration property value. The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
func (m *BookingService) GetDefaultDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.defaultDuration
}
// GetDefaultLocation gets the defaultLocation property value. The default physical location for the service.
func (m *BookingService) GetDefaultLocation()(Locationable) {
    return m.defaultLocation
}
// GetDefaultPrice gets the defaultPrice property value. The default monetary price for the service.
func (m *BookingService) GetDefaultPrice()(*float64) {
    return m.defaultPrice
}
// GetDefaultPriceType gets the defaultPriceType property value. Represents the type of pricing of a booking service.
func (m *BookingService) GetDefaultPriceType()(*BookingPriceType) {
    return m.defaultPriceType
}
// GetDefaultReminders gets the defaultReminders property value. The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
func (m *BookingService) GetDefaultReminders()([]BookingReminderable) {
    return m.defaultReminders
}
// GetDescription gets the description property value. A text description for the service.
func (m *BookingService) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. A service name.
func (m *BookingService) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingService) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAdditionalInformation)
    res["customQuestions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBookingQuestionAssignmentFromDiscriminatorValue , m.SetCustomQuestions)
    res["defaultDuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetDefaultDuration)
    res["defaultLocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateLocationFromDiscriminatorValue , m.SetDefaultLocation)
    res["defaultPrice"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetDefaultPrice)
    res["defaultPriceType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseBookingPriceType , m.SetDefaultPriceType)
    res["defaultReminders"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBookingReminderFromDiscriminatorValue , m.SetDefaultReminders)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isAnonymousJoinEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAnonymousJoinEnabled)
    res["isHiddenFromCustomers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsHiddenFromCustomers)
    res["isLocationOnline"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsLocationOnline)
    res["languageTag"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLanguageTag)
    res["maximumAttendeesCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMaximumAttendeesCount)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotes)
    res["postBuffer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetPostBuffer)
    res["preBuffer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetPreBuffer)
    res["schedulingPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateBookingSchedulingPolicyFromDiscriminatorValue , m.SetSchedulingPolicy)
    res["smsNotificationsEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSmsNotificationsEnabled)
    res["staffMemberIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetStaffMemberIds)
    res["webUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWebUrl)
    return res
}
// GetIsAnonymousJoinEnabled gets the isAnonymousJoinEnabled property value. The isAnonymousJoinEnabled property
func (m *BookingService) GetIsAnonymousJoinEnabled()(*bool) {
    return m.isAnonymousJoinEnabled
}
// GetIsHiddenFromCustomers gets the isHiddenFromCustomers property value. True means this service is not available to customers for booking.
func (m *BookingService) GetIsHiddenFromCustomers()(*bool) {
    return m.isHiddenFromCustomers
}
// GetIsLocationOnline gets the isLocationOnline property value. True indicates that the appointments for the service will be held online. Default value is false.
func (m *BookingService) GetIsLocationOnline()(*bool) {
    return m.isLocationOnline
}
// GetLanguageTag gets the languageTag property value. The languageTag property
func (m *BookingService) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetMaximumAttendeesCount gets the maximumAttendeesCount property value. The maximum number of customers allowed in a service. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
func (m *BookingService) GetMaximumAttendeesCount()(*int32) {
    return m.maximumAttendeesCount
}
// GetNotes gets the notes property value. Additional information about this service.
func (m *BookingService) GetNotes()(*string) {
    return m.notes
}
// GetPostBuffer gets the postBuffer property value. The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
func (m *BookingService) GetPostBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.postBuffer
}
// GetPreBuffer gets the preBuffer property value. The time to buffer before an appointment for this service can start.
func (m *BookingService) GetPreBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.preBuffer
}
// GetSchedulingPolicy gets the schedulingPolicy property value. The set of policies that determine how appointments for this type of service should be created and managed.
func (m *BookingService) GetSchedulingPolicy()(BookingSchedulingPolicyable) {
    return m.schedulingPolicy
}
// GetSmsNotificationsEnabled gets the smsNotificationsEnabled property value. True indicates SMS notifications can be sent to the customers for the appointment of the service. Default value is false.
func (m *BookingService) GetSmsNotificationsEnabled()(*bool) {
    return m.smsNotificationsEnabled
}
// GetStaffMemberIds gets the staffMemberIds property value. Represents those staff members who provide this service.
func (m *BookingService) GetStaffMemberIds()([]string) {
    return m.staffMemberIds
}
// GetWebUrl gets the webUrl property value. The URL a customer uses to access the service.
func (m *BookingService) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *BookingService) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("additionalInformation", m.GetAdditionalInformation())
        if err != nil {
            return err
        }
    }
    if m.GetCustomQuestions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomQuestions())
        err = writer.WriteCollectionOfObjectValues("customQuestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("defaultDuration", m.GetDefaultDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultLocation", m.GetDefaultLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("defaultPrice", m.GetDefaultPrice())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultPriceType() != nil {
        cast := (*m.GetDefaultPriceType()).String()
        err = writer.WriteStringValue("defaultPriceType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultReminders() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDefaultReminders())
        err = writer.WriteCollectionOfObjectValues("defaultReminders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAnonymousJoinEnabled", m.GetIsAnonymousJoinEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHiddenFromCustomers", m.GetIsHiddenFromCustomers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLocationOnline", m.GetIsLocationOnline())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumAttendeesCount", m.GetMaximumAttendeesCount())
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
    {
        err = writer.WriteISODurationValue("postBuffer", m.GetPostBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("preBuffer", m.GetPreBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedulingPolicy", m.GetSchedulingPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smsNotificationsEnabled", m.GetSmsNotificationsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetStaffMemberIds() != nil {
        err = writer.WriteCollectionOfStringValues("staffMemberIds", m.GetStaffMemberIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalInformation sets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
func (m *BookingService) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
// SetCustomQuestions sets the customQuestions property value. Contains the set of custom questions associated with a particular service.
func (m *BookingService) SetCustomQuestions(value []BookingQuestionAssignmentable)() {
    m.customQuestions = value
}
// SetDefaultDuration sets the defaultDuration property value. The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
func (m *BookingService) SetDefaultDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.defaultDuration = value
}
// SetDefaultLocation sets the defaultLocation property value. The default physical location for the service.
func (m *BookingService) SetDefaultLocation(value Locationable)() {
    m.defaultLocation = value
}
// SetDefaultPrice sets the defaultPrice property value. The default monetary price for the service.
func (m *BookingService) SetDefaultPrice(value *float64)() {
    m.defaultPrice = value
}
// SetDefaultPriceType sets the defaultPriceType property value. Represents the type of pricing of a booking service.
func (m *BookingService) SetDefaultPriceType(value *BookingPriceType)() {
    m.defaultPriceType = value
}
// SetDefaultReminders sets the defaultReminders property value. The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
func (m *BookingService) SetDefaultReminders(value []BookingReminderable)() {
    m.defaultReminders = value
}
// SetDescription sets the description property value. A text description for the service.
func (m *BookingService) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. A service name.
func (m *BookingService) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsAnonymousJoinEnabled sets the isAnonymousJoinEnabled property value. The isAnonymousJoinEnabled property
func (m *BookingService) SetIsAnonymousJoinEnabled(value *bool)() {
    m.isAnonymousJoinEnabled = value
}
// SetIsHiddenFromCustomers sets the isHiddenFromCustomers property value. True means this service is not available to customers for booking.
func (m *BookingService) SetIsHiddenFromCustomers(value *bool)() {
    m.isHiddenFromCustomers = value
}
// SetIsLocationOnline sets the isLocationOnline property value. True indicates that the appointments for the service will be held online. Default value is false.
func (m *BookingService) SetIsLocationOnline(value *bool)() {
    m.isLocationOnline = value
}
// SetLanguageTag sets the languageTag property value. The languageTag property
func (m *BookingService) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetMaximumAttendeesCount sets the maximumAttendeesCount property value. The maximum number of customers allowed in a service. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
func (m *BookingService) SetMaximumAttendeesCount(value *int32)() {
    m.maximumAttendeesCount = value
}
// SetNotes sets the notes property value. Additional information about this service.
func (m *BookingService) SetNotes(value *string)() {
    m.notes = value
}
// SetPostBuffer sets the postBuffer property value. The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
func (m *BookingService) SetPostBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.postBuffer = value
}
// SetPreBuffer sets the preBuffer property value. The time to buffer before an appointment for this service can start.
func (m *BookingService) SetPreBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.preBuffer = value
}
// SetSchedulingPolicy sets the schedulingPolicy property value. The set of policies that determine how appointments for this type of service should be created and managed.
func (m *BookingService) SetSchedulingPolicy(value BookingSchedulingPolicyable)() {
    m.schedulingPolicy = value
}
// SetSmsNotificationsEnabled sets the smsNotificationsEnabled property value. True indicates SMS notifications can be sent to the customers for the appointment of the service. Default value is false.
func (m *BookingService) SetSmsNotificationsEnabled(value *bool)() {
    m.smsNotificationsEnabled = value
}
// SetStaffMemberIds sets the staffMemberIds property value. Represents those staff members who provide this service.
func (m *BookingService) SetStaffMemberIds(value []string)() {
    m.staffMemberIds = value
}
// SetWebUrl sets the webUrl property value. The URL a customer uses to access the service.
func (m *BookingService) SetWebUrl(value *string)() {
    m.webUrl = value
}
