package gmp

import (
	"encoding/xml"
	"time"
)

type GetConfigsCommand struct {
	XMLName     xml.Name `xml:"get_configs"`
	ConfigID    string   `xml:"config_id,attr,omitempty"`
	Filter      string   `xml:"filter,attr,omitempty"`
	FiltID      string   `xml:"filt_id,attr,omitempty"`
	Trash       bool     `xml:"trash,attr,omitempty"`
	Details     bool     `xml:"details,attr,omitempty"`
	Families    bool     `xml:"families,attr,omitempty"`
	Preferences bool     `xml:"preferences,attr,omitempty"`
	Tasks       bool     `xml:"tasks,attr,omitempty"`
}

type GetConfigsResponse struct {
	XMLName     xml.Name                      `xml:"get_configs_response"`
	Status      string                        `xml:"status,attr"`
	StatusText  string                        `xml:"status_text,attr"`
	Config      []getConfigsResponseConfig    `xml:"config"`
	Filters     getConfigsResponseFilters     `xml:"filters"`
	Sort        getConfigsResponseSort        `xml:"sort"`
	Configs     getConfigsResponseConfigs     `xml:"configs"`
	ConfigCount getConfigsResponseConfigCount `xml:"config_count"`
}

type getConfigsResponseConfig struct {
	ID               string                              `xml:"id,attr"`
	Owner            getConfigsResponseConfigOwner       `xml:"owner"`
	Name             string                              `xml:"name"`
	Comment          string                              `xml:"comment"`
	CreationTime     time.Time                           `xml:"creation_time"`
	ModificationTime time.Time                           `xml:"modification_time"`
	FamilyCount      getConfigsResponseConfigFamilyCount `xml:"family_count"`
	NVTCount         getConfigsResponseConfigNVTCount    `xml:"nvt_count"`
	Type             string                              `xml:"type"`
	UsageType        string                              `xml:"usage_type"`
	MaxNVTCount      int                                 `xml:"max_nvt_count"`
	KnownNVTCount    int                                 `xml:"known_nvt_count"`
	InUse            bool                                `xml:"in_use"`
	Writable         bool                                `xml:"writable"`
	Permissions      getConfigsResponseConfigPermissions `xml:"permissions"`
	UserTags         getConfigsResponseConfigUserTags    `xml:"user_tags"`
}

type getConfigsResponseConfigOwner struct {
	Name string `xml:"name"`
}

type getConfigsResponseConfigFamilyCount struct {
	Value   string `xml:",chardata"`
	Growing string `xml:"growing"`
}

type getConfigsResponseConfigNVTCount struct {
	Value   string `xml:",chardata"`
	Growing string `xml:"growing"`
}

type getConfigsResponseConfigPermissions struct {
	Permission []getConfigsResponseConfigPermissionsPermission `xml:"permission"`
}

type getConfigsResponseConfigPermissionsPermission struct {
	Name string `xml:"name"`
}

type getConfigsResponseConfigUserTags struct {
	Count int `xml:"count"`
}

type getConfigsResponseFilters struct {
	ID       string                            `xml:"id,attr"`
	Term     string                            `xml:"term"`
	Name     string                            `xml:"name"`
	Keywords getConfigsResponseFiltersKeywords `xml:"keywords"`
}

type getConfigsResponseFiltersKeywords struct {
	Keyword []getConfigsResponseFiltersKeywordsKeyword `xml:"keyword"`
}

type getConfigsResponseFiltersKeywordsKeyword struct {
	Column   string `xml:"column"`
	Relation string `xml:"relation"`
	Value    string `xml:"value"`
}

type getConfigsResponseSort struct {
	Value string                      `xml:",chardata"`
	Field getConfigsResponseSortField `xml:"field"`
}

type getConfigsResponseSortField struct {
	Value string `xml:",chardata"`
	Order string `xml:"order"`
}

type getConfigsResponseConfigs struct {
	Start int `xml:"start,attr"`
	Max   int `xml:"max,attr"`
}

type getConfigsResponseConfigCount struct {
	Value    string `xml:",chardata"`
	Filtered int    `xml:"filtered"`
	Page     int    `xml:"page"`
}
