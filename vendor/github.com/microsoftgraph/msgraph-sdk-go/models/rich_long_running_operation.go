package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RichLongRunningOperation 
type RichLongRunningOperation struct {
    LongRunningOperation
    // Error that caused the operation to fail.
    error PublicErrorable
    // A value between 0 and 100 that indicates the progress of the operation.
    percentageComplete *int32
    // The unique identifier for the result.
    resourceId *string
    // The type of the operation.
    type_escaped *string
}
// NewRichLongRunningOperation instantiates a new RichLongRunningOperation and sets the default values.
func NewRichLongRunningOperation()(*RichLongRunningOperation) {
    m := &RichLongRunningOperation{
        LongRunningOperation: *NewLongRunningOperation(),
    }
    odataTypeValue := "#microsoft.graph.richLongRunningOperation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRichLongRunningOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRichLongRunningOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRichLongRunningOperation(), nil
}
// GetError gets the error property value. Error that caused the operation to fail.
func (m *RichLongRunningOperation) GetError()(PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RichLongRunningOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LongRunningOperation.GetFieldDeserializers()
    res["error"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePublicErrorFromDiscriminatorValue , m.SetError)
    res["percentageComplete"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPercentageComplete)
    res["resourceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetResourceId)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetType)
    return res
}
// GetPercentageComplete gets the percentageComplete property value. A value between 0 and 100 that indicates the progress of the operation.
func (m *RichLongRunningOperation) GetPercentageComplete()(*int32) {
    return m.percentageComplete
}
// GetResourceId gets the resourceId property value. The unique identifier for the result.
func (m *RichLongRunningOperation) GetResourceId()(*string) {
    return m.resourceId
}
// GetType gets the type property value. The type of the operation.
func (m *RichLongRunningOperation) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *RichLongRunningOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LongRunningOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentageComplete", m.GetPercentageComplete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. Error that caused the operation to fail.
func (m *RichLongRunningOperation) SetError(value PublicErrorable)() {
    m.error = value
}
// SetPercentageComplete sets the percentageComplete property value. A value between 0 and 100 that indicates the progress of the operation.
func (m *RichLongRunningOperation) SetPercentageComplete(value *int32)() {
    m.percentageComplete = value
}
// SetResourceId sets the resourceId property value. The unique identifier for the result.
func (m *RichLongRunningOperation) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetType sets the type property value. The type of the operation.
func (m *RichLongRunningOperation) SetType(value *string)() {
    m.type_escaped = value
}
