package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnValidation 
type ColumnValidation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Default BCP 47 language tag for the description.
    defaultLanguage *string
    // Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
    descriptions []DisplayNameLocalizationable
    // The formula to validate column value. For examples, see Examples of common formulas in lists.
    formula *string
    // The OdataType property
    odataType *string
}
// NewColumnValidation instantiates a new columnValidation and sets the default values.
func NewColumnValidation()(*ColumnValidation) {
    m := &ColumnValidation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.columnValidation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateColumnValidationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateColumnValidationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnValidation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnValidation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefaultLanguage gets the defaultLanguage property value. Default BCP 47 language tag for the description.
func (m *ColumnValidation) GetDefaultLanguage()(*string) {
    return m.defaultLanguage
}
// GetDescriptions gets the descriptions property value. Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
func (m *ColumnValidation) GetDescriptions()([]DisplayNameLocalizationable) {
    return m.descriptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ColumnValidation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultLanguage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultLanguage)
    res["descriptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDisplayNameLocalizationFromDiscriminatorValue , m.SetDescriptions)
    res["formula"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFormula)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetFormula gets the formula property value. The formula to validate column value. For examples, see Examples of common formulas in lists.
func (m *ColumnValidation) GetFormula()(*string) {
    return m.formula
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ColumnValidation) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ColumnValidation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultLanguage", m.GetDefaultLanguage())
        if err != nil {
            return err
        }
    }
    if m.GetDescriptions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDescriptions())
        err := writer.WriteCollectionOfObjectValues("descriptions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("formula", m.GetFormula())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnValidation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultLanguage sets the defaultLanguage property value. Default BCP 47 language tag for the description.
func (m *ColumnValidation) SetDefaultLanguage(value *string)() {
    m.defaultLanguage = value
}
// SetDescriptions sets the descriptions property value. Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
func (m *ColumnValidation) SetDescriptions(value []DisplayNameLocalizationable)() {
    m.descriptions = value
}
// SetFormula sets the formula property value. The formula to validate column value. For examples, see Examples of common formulas in lists.
func (m *ColumnValidation) SetFormula(value *string)() {
    m.formula = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ColumnValidation) SetOdataType(value *string)() {
    m.odataType = value
}
