package gmp

import (
	"encoding/xml"
	"time"
)

type GetScannersCommand struct {
	XMLName   xml.Name `xml:"get_scanners"`
	ScannerID string   `xml:"scanner_id,attr,omitempty"`
	Filter    string   `xml:"filter,attr,omitempty"`
	FiltID    string   `xml:"filt_id,attr,omitempty"`
	Trash     bool     `xml:"trash,attr,omitempty"`
	Details   bool     `xml:"details,attr,omitempty"`
}

type GetScannersResponse struct {
	XMLName    xml.Name                        `xml:"get_scanners_response"`
	Status     string                          `xml:"status,attr"`
	StatusText string                          `xml:"status_text,attr"`
	Scanner    []GetScannersResponseScanner    `xml:"scanner"`
	Filters    GetScannersResponseFilters      `xml:"filters"`
	Sort       GetScannersResponseSort         `xml:"sort"`
	Scanners   GetScannersResponseScanners     `xml:"scanners"`
	Count      GetScannersResponseScannerCount `xml:"scanner_count"`
}

type GetScannersResponseScanner struct {
	ID               string                                `xml:"id,attr"`
	Owner            GetScannersResponseScannerOwner       `xml:"owner"`
	Name             string                                `xml:"name"`
	Comment          string                                `xml:"comment"`
	Copy             string                                `xml:"copy"`
	CreationTime     time.Time                             `xml:"creation_time"`
	ModificationTime time.Time                             `xml:"modification_time"`
	Writable         bool                                  `xml:"writable"`
	InUse            bool                                  `xml:"in_use"`
	Permissions      GetScannersResponseScannerPermissions `xml:"permissions"`
	UserTags         GetScannersResponseScannerUserTags    `xml:"user_tags"`
	CAPubInfo        CertificateInfo                       `xml:"ca_pub_info"`
	CertificateInfo  CertificateInfo                       `xml:"certificate_info"`
	Host             string                                `xml:"host"`
	Port             string                                `xml:"port"`
	Type             string                                `xml:"type"`
	CAPub            string                                `xml:"ca_pub"`
	Credential       GetScannersResponseScannerCredential  `xml:"credential"`
	Configs          GetScannersResponseScannerConfigs     `xml:"configs"`
	Tasks            GetScannersResponseScannerTasks       `xml:"tasks"`
}

type GetScannersResponseScannerOwner struct {
	Name string `xml:"name"`
}

type GetScannersResponseScannerPermissions struct {
	Permission []GetScannersResponseScannerPermissionsPermission `xml:"permission"`
}

type GetScannersResponseScannerPermissionsPermission struct {
	Name string `xml:"name"`
}

type GetScannersResponseScannerUserTags struct {
	Count int                                     `xml:"count"`
	Tag   []GetScannersResponseScannerUserTagsTag `xml:"tag"`
}

type GetScannersResponseScannerUserTagsTag struct {
	Name    string `xml:"name"`
	Value   string `xml:"value"`
	Comment string `xml:"comment"`
}

type GetScannersResponseScannerCredential struct {
	ID    string `xml:"id,attr"`
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Trash bool   `xml:"trash"`
}

type GetScannersResponseScannerConfigs struct {
	Config []GetScannersResponseScannerConfigsConfig `xml:"config"`
}

type GetScannersResponseScannerConfigsConfig struct {
	ID          string `xml:"id,attr"`
	Name        string `xml:"name"`
	Permissions string `xml:"permissions"`
}

type GetScannersResponseScannerTasks struct {
	Task []GetScannersResponseScannerTasksTask `xml:"task"`
}

type GetScannersResponseScannerTasksTask struct {
	ID          string `xml:"id,attr"`
	Name        string `xml:"name"`
	Permissions string `xml:"permissions"`
}

type GetScannersResponseFilters struct {
	ID       string                             `xml:"id,attr"`
	Term     string                             `xml:"term"`
	Name     string                             `xml:"name"`
	Keywords GetScannersResponseFiltersKeywords `xml:"keywords"`
}

type GetScannersResponseFiltersKeywords struct {
	Keyword []GetScannersResponseFiltersKeywordsKeyword `xml:"keyword"`
}

type GetScannersResponseFiltersKeywordsKeyword struct {
	Column   string `xml:"column"`
	Relation string `xml:"relation"`
	Value    string `xml:"value"`
}

type GetScannersResponseSort struct {
	Value string                       `xml:",chardata"`
	Field GetScannersResponseSortField `xml:"field"`
}

type GetScannersResponseSortField struct {
	Value string `xml:",chardata"`
	Order string `xml:"order"`
}

type GetScannersResponseScanners struct {
	Start int `xml:"start,attr"`
	Max   int `xml:"max,attr"`
}

type GetScannersResponseScannerCount struct {
	Value    int `xml:",chardata"`
	Filtered int `xml:"filtered"`
	Page     int `xml:"page"`
}
