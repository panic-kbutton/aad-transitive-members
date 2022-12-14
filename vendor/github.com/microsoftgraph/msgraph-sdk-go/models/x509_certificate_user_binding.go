package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// X509CertificateUserBinding 
type X509CertificateUserBinding struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The priority of the binding. Azure AD uses the binding with the highest priority. This value must be a non-negative integer and unique in the collection of objects in the certificateUserBindings property of an x509CertificateAuthenticationMethodConfiguration object. Required
    priority *int32
    // Defines the Azure AD user property of the user object to use for the binding. The possible values are: userPrincipalName, onPremisesUserPrincipalName, certificateUserIds. Required.
    userProperty *string
    // The field on the X.509 certificate to use for the binding. The possible values are: PrincipalName, RFC822Name, SubjectKeyIdentifier, SHA1PublicKey.
    x509CertificateField *string
}
// NewX509CertificateUserBinding instantiates a new x509CertificateUserBinding and sets the default values.
func NewX509CertificateUserBinding()(*X509CertificateUserBinding) {
    m := &X509CertificateUserBinding{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.x509CertificateUserBinding";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateX509CertificateUserBindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateX509CertificateUserBindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateUserBinding(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *X509CertificateUserBinding) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *X509CertificateUserBinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPriority)
    res["userProperty"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserProperty)
    res["x509CertificateField"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetX509CertificateField)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *X509CertificateUserBinding) GetOdataType()(*string) {
    return m.odataType
}
// GetPriority gets the priority property value. The priority of the binding. Azure AD uses the binding with the highest priority. This value must be a non-negative integer and unique in the collection of objects in the certificateUserBindings property of an x509CertificateAuthenticationMethodConfiguration object. Required
func (m *X509CertificateUserBinding) GetPriority()(*int32) {
    return m.priority
}
// GetUserProperty gets the userProperty property value. Defines the Azure AD user property of the user object to use for the binding. The possible values are: userPrincipalName, onPremisesUserPrincipalName, certificateUserIds. Required.
func (m *X509CertificateUserBinding) GetUserProperty()(*string) {
    return m.userProperty
}
// GetX509CertificateField gets the x509CertificateField property value. The field on the X.509 certificate to use for the binding. The possible values are: PrincipalName, RFC822Name, SubjectKeyIdentifier, SHA1PublicKey.
func (m *X509CertificateUserBinding) GetX509CertificateField()(*string) {
    return m.x509CertificateField
}
// Serialize serializes information the current object
func (m *X509CertificateUserBinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userProperty", m.GetUserProperty())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("x509CertificateField", m.GetX509CertificateField())
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
func (m *X509CertificateUserBinding) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *X509CertificateUserBinding) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPriority sets the priority property value. The priority of the binding. Azure AD uses the binding with the highest priority. This value must be a non-negative integer and unique in the collection of objects in the certificateUserBindings property of an x509CertificateAuthenticationMethodConfiguration object. Required
func (m *X509CertificateUserBinding) SetPriority(value *int32)() {
    m.priority = value
}
// SetUserProperty sets the userProperty property value. Defines the Azure AD user property of the user object to use for the binding. The possible values are: userPrincipalName, onPremisesUserPrincipalName, certificateUserIds. Required.
func (m *X509CertificateUserBinding) SetUserProperty(value *string)() {
    m.userProperty = value
}
// SetX509CertificateField sets the x509CertificateField property value. The field on the X.509 certificate to use for the binding. The possible values are: PrincipalName, RFC822Name, SubjectKeyIdentifier, SHA1PublicKey.
func (m *X509CertificateUserBinding) SetX509CertificateField(value *string)() {
    m.x509CertificateField = value
}
