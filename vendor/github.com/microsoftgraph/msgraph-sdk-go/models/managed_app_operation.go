package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppOperation represents an operation applied against an app registration.
type ManagedAppOperation struct {
    Entity
    // The operation name.
    displayName *string
    // The last time the app operation was modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The current state of the operation
    state *string
    // Version of the entity.
    version *string
}
// NewManagedAppOperation instantiates a new managedAppOperation and sets the default values.
func NewManagedAppOperation()(*ManagedAppOperation) {
    m := &ManagedAppOperation{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedAppOperation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagedAppOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAppOperation(), nil
}
// GetDisplayName gets the displayName property value. The operation name.
func (m *ManagedAppOperation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetState)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVersion)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The last time the app operation was modified.
func (m *ManagedAppOperation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetState gets the state property value. The current state of the operation
func (m *ManagedAppOperation) GetState()(*string) {
    return m.state
}
// GetVersion gets the version property value. Version of the entity.
func (m *ManagedAppOperation) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *ManagedAppOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The operation name.
func (m *ManagedAppOperation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The last time the app operation was modified.
func (m *ManagedAppOperation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetState sets the state property value. The current state of the operation
func (m *ManagedAppOperation) SetState(value *string)() {
    m.state = value
}
// SetVersion sets the version property value. Version of the entity.
func (m *ManagedAppOperation) SetVersion(value *string)() {
    m.version = value
}
