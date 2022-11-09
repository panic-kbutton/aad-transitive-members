package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParticipantLeftNotification 
type ParticipantLeftNotification struct {
    Entity
    // The call property
    call Callable
    // ID of the participant under the policy who has left the meeting.
    participantId *string
}
// NewParticipantLeftNotification instantiates a new ParticipantLeftNotification and sets the default values.
func NewParticipantLeftNotification()(*ParticipantLeftNotification) {
    m := &ParticipantLeftNotification{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.participantLeftNotification";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateParticipantLeftNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParticipantLeftNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipantLeftNotification(), nil
}
// GetCall gets the call property value. The call property
func (m *ParticipantLeftNotification) GetCall()(Callable) {
    return m.call
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParticipantLeftNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["call"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCallFromDiscriminatorValue , m.SetCall)
    res["participantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetParticipantId)
    return res
}
// GetParticipantId gets the participantId property value. ID of the participant under the policy who has left the meeting.
func (m *ParticipantLeftNotification) GetParticipantId()(*string) {
    return m.participantId
}
// Serialize serializes information the current object
func (m *ParticipantLeftNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("call", m.GetCall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("participantId", m.GetParticipantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCall sets the call property value. The call property
func (m *ParticipantLeftNotification) SetCall(value Callable)() {
    m.call = value
}
// SetParticipantId sets the participantId property value. ID of the participant under the policy who has left the meeting.
func (m *ParticipantLeftNotification) SetParticipantId(value *string)() {
    m.participantId = value
}
