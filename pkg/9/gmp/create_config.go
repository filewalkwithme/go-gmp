package gmp

import "encoding/xml"

type CreateConfigCommand struct {
	XMLName            xml.Name            `xml:"create_config"`
	Comment            string              `xml:"comment,omitempty"`
	Copy               string              `xml:"copy,omitempty"`
	GetConfigsResponse *GetConfigsResponse `xml:"get_configs_response,omitempty"`
	Scanner            string              `xml:"scanner,omitempty"`
	Name               string              `xml:"name,omitempty"`
	UsageType          string              `xml:"usage_type,omitempty"`
}

type CreateConfigResponse struct {
	XMLName    xml.Name `xml:"create_config_response"`
	Status     string   `xml:"status,attr"`
	StatusText string   `xml:"status_text,attr"`
	ID         string   `xml:"id,attr"`
}
