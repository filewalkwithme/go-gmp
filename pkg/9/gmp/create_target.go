package gmp

import "encoding/xml"

type CreateTargetCommand struct {
	XMLName            xml.Name                       `xml:"create_target"`
	Name               string                         `xml:"name,omitempty"`
	Comment            string                         `xml:"comment,omitempty"`
	Copy               string                         `xml:"copy,omitempty"`
	AssetHosts         *CreateTargetAssetHosts        `xml:"asset_hosts,omitempty"`
	Hosts              string                         `xml:"hosts,omitempty"`
	ExcludeHosts       string                         `xml:"exclude_hosts,omitempty"`
	SSHCredential      *CreateTargetSSHCredential     `xml:"ssh_credential,omitempty"`
	SMBCredential      *CreateTargetSMBCredential     `xml:"smb_credential,omitempty"`
	ESXICredential     *CreateTargetESXICredential    `xml:"esxi_credential,omitempty"`
	SNMPCredential     *CreateTargetSNMPCredential    `xml:"snmp_credential,omitempty"`
	SSHLSCCredential   *CreateTargetSSHLSCCredential  `xml:"ssh_lsc_credential,omitempty"`
	SMBLSCCredential   *CreateTargetSMBLSCCredential  `xml:"smb_lsc_credential,omitempty"`
	ESXILSCCredential  *CreateTargetESXILSCCredential `xml:"esxi_lsc_credential,omitempty"`
	AliveTests         string                         `xml:"alive_tests,omitempty"`
	ReverseLookupOnly  bool                           `xml:"reverse_lookup_only,omitempty"`
	ReverseLookupUnify bool                           `xml:"reverse_lookup_unify,omitempty"`
	PortRange          string                         `xml:"port_range,omitempty"`
	PortList           *CreateTargetPortList          `xml:"port_list,omitempty"`
}

type CreateTargetAssetHosts struct {
	Filter string `xml:"filter,attr,omitempty"`
}

type CreateTargetSSHCredential struct {
	ID   string `xml:"id,attr,omitempty"`
	Port string `xml:"port,omitempty"`
}

type CreateTargetSMBCredential struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTargetESXICredential struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTargetSNMPCredential struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTargetSSHLSCCredential struct {
	ID   string `xml:"id,attr,omitempty"`
	Port string `xml:"port,omitempty"`
}

type CreateTargetSMBLSCCredential struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTargetESXILSCCredential struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTargetPortList struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTargetResponse struct {
	XMLName    xml.Name `xml:"create_target_response"`
	Status     string   `xml:"status,attr"`
	StatusText string   `xml:"status_text,attr"`
	ID         string   `xml:"id,attr"`
}
