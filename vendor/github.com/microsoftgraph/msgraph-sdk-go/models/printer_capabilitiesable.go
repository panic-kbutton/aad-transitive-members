package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterCapabilitiesable 
type PrinterCapabilitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBottomMargins()([]int32)
    GetCollation()(*bool)
    GetColorModes()([]PrintColorMode)
    GetContentTypes()([]string)
    GetCopiesPerJob()(IntegerRangeable)
    GetDpis()([]int32)
    GetDuplexModes()([]PrintDuplexMode)
    GetFeedOrientations()([]PrinterFeedOrientation)
    GetFinishings()([]PrintFinishing)
    GetInputBins()([]string)
    GetIsColorPrintingSupported()(*bool)
    GetIsPageRangeSupported()(*bool)
    GetLeftMargins()([]int32)
    GetMediaColors()([]string)
    GetMediaSizes()([]string)
    GetMediaTypes()([]string)
    GetMultipageLayouts()([]PrintMultipageLayout)
    GetOdataType()(*string)
    GetOrientations()([]PrintOrientation)
    GetOutputBins()([]string)
    GetPagesPerSheet()([]int32)
    GetQualities()([]PrintQuality)
    GetRightMargins()([]int32)
    GetScalings()([]PrintScaling)
    GetSupportsFitPdfToPage()(*bool)
    GetTopMargins()([]int32)
    SetBottomMargins(value []int32)()
    SetCollation(value *bool)()
    SetColorModes(value []PrintColorMode)()
    SetContentTypes(value []string)()
    SetCopiesPerJob(value IntegerRangeable)()
    SetDpis(value []int32)()
    SetDuplexModes(value []PrintDuplexMode)()
    SetFeedOrientations(value []PrinterFeedOrientation)()
    SetFinishings(value []PrintFinishing)()
    SetInputBins(value []string)()
    SetIsColorPrintingSupported(value *bool)()
    SetIsPageRangeSupported(value *bool)()
    SetLeftMargins(value []int32)()
    SetMediaColors(value []string)()
    SetMediaSizes(value []string)()
    SetMediaTypes(value []string)()
    SetMultipageLayouts(value []PrintMultipageLayout)()
    SetOdataType(value *string)()
    SetOrientations(value []PrintOrientation)()
    SetOutputBins(value []string)()
    SetPagesPerSheet(value []int32)()
    SetQualities(value []PrintQuality)()
    SetRightMargins(value []int32)()
    SetScalings(value []PrintScaling)()
    SetSupportsFitPdfToPage(value *bool)()
    SetTopMargins(value []int32)()
}
