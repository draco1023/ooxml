package ooxml

import (
	"encoding/xml"

	"github.com/google/uuid"
)

// STLang ...
type STLang string

// STHexColorRGB 3 bytes
type STHexColorRGB string

// STPanose 10 bytes
type STPanose string

// STCalendarType ...
type STCalendarType string

const (
	STCalendarType_gregorian            STCalendarType = "gregorian"
	STCalendarType_gregorianUs          STCalendarType = "gregorianUs"
	STCalendarType_gregorianMeFrench    STCalendarType = "gregorianMeFrench"
	STCalendarType_gregorianArabic      STCalendarType = "gregorianArabic"
	STCalendarType_hijri                STCalendarType = "hijri"
	STCalendarType_hebrew               STCalendarType = "hebrew"
	STCalendarType_taiwan               STCalendarType = "taiwan"
	STCalendarType_japan                STCalendarType = "japan"
	STCalendarType_thai                 STCalendarType = "thai"
	STCalendarType_korea                STCalendarType = "korea"
	STCalendarType_saka                 STCalendarType = "saka"
	STCalendarType_gregorianXlitEnglish STCalendarType = "gregorianXlitEnglish"
	STCalendarType_gregorianXlitFrench  STCalendarType = "gregorianXlitFrench"
	STCalendarType_none                 STCalendarType = "none"
)

// STAlgClass ...
type STAlgClass string

const (
	STAlgClass_hash   STAlgClass = "hash"
	STAlgClass_custom STAlgClass = "custom"
)

// STCryptProv ...
type STCryptProv string

const (
	STCryptProv_rsaAES  STCryptProv = "rsaAES"
	STCryptProv_rsaFull STCryptProv = "rsaFull"
	STCryptProv_custom  STCryptProv = "custom"
)

// STAlgType ...
type STAlgType string

const (
	STAlgType_typeAny STAlgType = "typeAny"
	STAlgType_custom  STAlgType = "custom"
)

// STColorType ...
type STColorType string

// STGuid ...
// Pattern: \{[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}\}
type STGuid string

// NewSTGuid new STGuid instance
func NewSTGuid() STGuid {
	return STGuid("{" + uuid.New().String() + "}")
}

// STOnOff ...
type STOnOff string

// STOnOff literals
const (
	STOnOff_true  STOnOff = "true"
	STOnOff_false STOnOff = "false"
	STOnOff_1     STOnOff = "1"
	STOnOff_0     STOnOff = "0"
	STOnOff_on    STOnOff = "on"
	STOnOff_off   STOnOff = "off"
)

// IsOn STOnoff as bool
func (b STOnOff) IsOn() bool {
	switch b {
	case STOnOff_true, STOnOff_1, STOnOff_on:
		return true
	}
	return false
}

// STString ...
// type STString string

// STXmlName ...
// type STXmlName string

// STTrueFalse ...
type STTrueFalse string

const (
	STTrueFalse_t     STTrueFalse = "t"
	STTrueFalse_f     STTrueFalse = "f"
	STTrueFalse_true  STTrueFalse = "true"
	STTrueFalse_false STTrueFalse = "false"
)

// STTrueFalseBlank ...
type STTrueFalseBlank string

const (
	STTrueFalseBlank_t     STTrueFalseBlank = "t"
	STTrueFalseBlank_f     STTrueFalseBlank = "f"
	STTrueFalseBlank_true  STTrueFalseBlank = "true"
	STTrueFalseBlank_false STTrueFalseBlank = "false"
	STTrueFalseBlank_      STTrueFalseBlank = ""
	STTrueFalseBlank_True  STTrueFalseBlank = "True"
	STTrueFalseBlank_False STTrueFalseBlank = "False"
)

// STUnsignedDecimalNumber ...
type STUnsignedDecimalNumber uint64

// STTwipsMeasure ...
type STTwipsMeasure struct {
	XMLName                    xml.Name `xml:"ST_TwipsMeasure"`
	STUnsignedDecimalNumber    uint64
	STPositiveUniversalMeasure *STPositiveUniversalMeasure
}

// STVerticalAlignRun ...
type STVerticalAlignRun string

const (
	STVerticalAlignRun_baseline    STVerticalAlignRun = "baseline"
	STVerticalAlignRun_superscript STVerticalAlignRun = "superscript"
	STVerticalAlignRun_subscript   STVerticalAlignRun = "subscript"
)

// STXstring ...
type STXstring string

// STXAlign ...
type STXAlign string

const (
	STXAlign_left    STXAlign = "left"
	STXAlign_center  STXAlign = "center"
	STXAlign_right   STXAlign = "right"
	STXAlign_inside  STXAlign = "inside"
	STXAlign_outside STXAlign = "outside"
)

// STYAlign ...
type STYAlign string

const (
	STYAlign_inline  STYAlign = "inline"
	STYAlign_top     STYAlign = "top"
	STYAlign_center  STYAlign = "center"
	STYAlign_bottom  STYAlign = "bottom"
	STYAlign_inside  STYAlign = "inside"
	STYAlign_outside STYAlign = "outside"
)

// STConformanceClass ...
type STConformanceClass string

const (
	STConformanceClass_strict       STConformanceClass = "strict"
	STConformanceClass_transitional STConformanceClass = "transitional"
)

// STUniversalMeasure ...
// Pattern: -?[0-9]+(\.[0-9]+)?(mm|cm|in|pt|pc|pi)
type STUniversalMeasure string

// STPositiveUniversalMeasure ...
// Pattern: [0-9]+(\.[0-9]+)?(mm|cm|in|pt|pc|pi)
type STPositiveUniversalMeasure string

// STPercentage ...
// Pattern: -?[0-9]+(\.[0-9]+)?%
type STPercentage string

// STFixedPercentage ...
// Pattern: -?((100)|([0-9][0-9]?))(\.[0-9][0-9]?)?%
type STFixedPercentage string

// STPositivePercentage ...
// Pattern: [0-9]+(\.[0-9]+)?%
type STPositivePercentage string

// STPositiveFixedPercentage ...
// Pattern: ((100)|([0-9][0-9]?))(\.[0-9][0-9]?)?%
type STPositiveFixedPercentage string
