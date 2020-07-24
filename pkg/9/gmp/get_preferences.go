package gmp

import "encoding/xml"

type GetPreferencesCommand struct {
	XMLName    xml.Name `xml:"get_preferences"`
	NVTOID     string   `xml:"nvt_oid,attr,omitempty"`
	ConfigID   string   `xml:"config_id,attr,omitempty"`
	Preference string   `xml:"preference,attr,omitempty"`
}

type GetPreferencesResponse struct {
	XMLName    xml.Name                           `xml:"get_preferences_response"`
	Status     string                             `xml:"status,attr"`
	StatusText string                             `xml:"status_text,attr"`
	Preference []GetPreferencesResponsePreference `xml:"preference"`
}

type GetPreferencesResponsePreference struct {
	NVT     GetPreferencesResponsePreferenceNVT `xml:"nvt"`
	Name    string                              `xml:"name"`
	ID      string                              `xml:"id"`
	Type    string                              `xml:"type"`
	Value   string                              `xml:"value"`
	Alt     []string                            `xml:"alt"`
	Default string                              `xml:"default"`
}

type GetPreferencesResponsePreferenceNVT struct {
	OID  string `xml:"oid,attr"`
	Name string `xml:"name"`
}
