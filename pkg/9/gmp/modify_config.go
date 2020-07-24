package gmp

import "encoding/xml"

type ModifyConfigCommand struct {
	XMLName         xml.Name                     `xml:"modify_config"`
	ConfigID        string                       `xml:"config_id,attr,omitempty"`
	Name            string                       `xml:"name,omitempty"`
	Comment         string                       `xml:"comment,omitempty"`
	Scanner         string                       `xml:"scanner,omitempty"`
	Preference      *ModifyConfigPreference      `xml:"preference,omitempty"`
	FamilySelection *ModifyConfigFamilySelection `xml:"family_selection,omitempty"`
	NVTSelection    *ModifyConfigNVTSelection    `xml:"nvt_selection,omitempty"`
}

type ModifyConfigPreference struct {
	Name  string                     `xml:"name,omitempty"`
	NVT   *ModifyConfigPreferenceNVT `xml:"nvt,omitempty"`
	Value string                     `xml:"value,omitempty"`
}

type ModifyConfigPreferenceNVT struct {
	OID string `xml:"oid,attr,omitempty"`
}

type ModifyConfigFamilySelection struct {
	Growing bool                                `xml:"growing,omitempty"`
	Family  []ModifyConfigFamilySelectionFamily `xml:"family,omitempty"`
}

type ModifyConfigFamilySelectionFamily struct {
	All     bool   `xml:"all,omitempty"`
	Growing bool   `xml:"growing,omitempty"`
	Name    string `xml:"name,omitempty"`
}

type ModifyConfigNVTSelection struct {
	Family string `xml:"family,omitempty"`
	NVT    string `xml:"nvt,omitempty"`
}

type ModifyConfigResponse struct {
	XMLName    xml.Name `xml:"modify_config_response"`
	Status     string   `xml:"status,attr"`
	StatusText string   `xml:"status_text,attr"`
}
