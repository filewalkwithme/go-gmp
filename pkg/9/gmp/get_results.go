package gmp

import (
	"encoding/xml"
	"time"
)

type GetResultsCommand struct {
	XMLName          xml.Name `xml:"get_results"`
	Filter           string   `xml:"filter,attr,omitempty"`
	FiltID           string   `xml:"filt_id,attr,omitempty"`
	TaskID           string   `xml:"task_id,attr,omitempty"`
	NotesDetails     bool     `xml:"notes_details,attr,omitempty"`
	OverridesDetails bool     `xml:"overrides_details,attr,omitempty"`
	Details          bool     `xml:"details,attr,omitempty"`
	GetCounts        bool     `xml:"get_counts,attr,omitempty"`
}

type GetResultsResponse struct {
	XMLName     xml.Name                      `xml:"get_results_response"`
	Status      string                        `xml:"status,attr"`
	StatusText  string                        `xml:"status_text,attr"`
	Result      []Result                      `xml:"result"`
	Filters     GetResultsResponseFilters     `xml:"filters"`
	Sort        GetResultsResponseSort        `xml:"sort"`
	Results     GetResultsResponseResults     `xml:"results"`
	ResultCount GetResultsResponseResultCount `xml:"result_count"`
}

type Result struct {
	ID               string          `xml:"id,attr"`
	Name             string          `xml:"name"`
	Owner            ResultOwner     `xml:"owner"`
	Comment          string          `xml:"comment"`
	CreationTime     time.Time       `xml:"creation_time"`
	ModificationTime time.Time       `xml:"modification_time"`
	UserTags         ResultUserTags  `xml:"user_tags"`
	Report           ResultReport    `xml:"result_report"`
	Task             ResultTask      `xml:"result_task"`
	Host             ResultHost      `xml:"host"`
	Port             string          `xml:"port"`
	NVT              ResultNVT       `xml:"nvt"`
	ScanNVTVersion   string          `xml:"scan_nvt_version"`
	Threat           string          `xml:"threat"`
	Severity         string          `xml:"severity"`
	QOD              ResultQOD       `xml:"qod"`
	OriginalThreat   string          `xml:"original_threat"`
	OriginalSeverity string          `xml:"original_severity"`
	Description      string          `xml:"description"`
	Delta            ResultDelta     `xml:"delta"`
	Detection        ResultDetection `xml:"detection"`
	Notes            ResultNotes     `xml:"notes"`
	Overrides        ResultOverrides `xml:"overrides"`
	Tickets          ResultTickets   `xml:"tickets"`
}

type ResultOwner struct {
	Name string `xml:"name"`
}

type ResultUserTags struct {
	Count int                 `xml:"count"`
	Tag   []ResultUserTagsTag `xml:"tag"`
}

type ResultUserTagsTag struct {
	ID      string      `xml:"id,attr"`
	Name    string      `xml:"name"`
	Value   ResultOwner `xml:"value"`
	Comment string      `xml:"comment"`
}

type ResultReport struct {
	ID string `xml:"id,attr"`
}

type ResultTask struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name"`
}

type ResultHost struct {
	Value    string          `xml:",chardata"`
	Asset    ResultHostAsset `xml:"asset"`
	Hostname string          `xml:"hostname"`
}

type ResultHostAsset struct {
	AssetID string `xml:"asset_id,attr"`
}

type ResultNVT struct {
	OID      string        `xml:"oid,attr"`
	Name     string        `xml:"name"`
	Type     string        `xml:"type"`
	Family   string        `xml:"family"`
	CVSSBase float32       `xml:"cvss_base"`
	CPE      string        `xml:"cpe"`
	Tags     string        `xml:"tags"`
	Refs     ResultNVTRefs `xml:"refs"`
}

type ResultNVTRefs struct {
	Ref []ResultNVTRefsRef `xml:"ref"`
}

type ResultNVTRefsRef struct {
	ID   string `xml:"id,attr"`
	Type string `xml:"type,attr"`
}

type ResultQOD struct {
	Value int `xml:"value"`
	Type  int `xml:"type"`
}

type ResultDelta struct {
	Value     string               `xml:",chardata"`
	Result    *Result              `xml:"result"`
	Diff      string               `xml:"diff"`
	Notes     ResultDeltaNotes     `xml:"notes"`
	Overrides ResultDeltaOverrides `xml:"overrides"`
}

type ResultDeltaNotes struct {
	Note []Note `xml:"note"`
}

type Note struct {
	ID               string          `xml:"id,attr"`
	Permissions      NotePermissions `xml:"permissions"`
	Owner            NoteOwner       `xml:"owner"`
	NVT              NoteNVT         `xml:"nvt"`
	Text             NoteText        `xml:"text"`
	CreationTime     time.Time       `xml:"creation_time"`
	ModificationTime time.Time       `xml:"modification_time"`
	Writable         bool            `xml:"writable"`
	InUse            bool            `xml:"in_use"`
	Active           bool            `xml:"active"`
	Orphan           bool            `xml:"orphan"`
	UserTags         NoteUserTags    `xml:"user_tags"`
	Hosts            string          `xml:"hosts"`
	Port             string          `xml:"port"`
	Severity         string          `xml:"severity"`
	Threat           string          `xml:"threat"`
	Task             NoteTask        `xml:"task"`
	EndTime          string          `xml:"end_time"`
	Result           NoteResult      `xml:"result"`
}

type NotePermissions struct {
	Permission []NotePermissionsPermission `xml:"permission"`
}

type NotePermissionsPermission struct {
	Name string `xml:"name"`
}

type NoteOwner struct {
	Name string `xml:"name"`
}

type NoteNVT struct {
	OID  string `xml:"oid,attr"`
	Name string `xml:"name"`
	Type string `xml:"type"`
}

type NoteText struct {
	Value   string `xml:",chardata"`
	Excerpt bool   `xml:"excerpt,attr"`
}

type NoteUserTags struct {
	Count int               `xml:"count"`
	Tag   []NoteUserTagsTag `xml:"tag"`
}

type NoteUserTagsTag struct {
	ID      int    `xml:"id,attr"`
	Name    string `xml:"name"`
	Value   string `xml:"value"`
	Comment string `xml:"comment"`
}

type NoteTask struct {
	ID int `xml:"id,attr"`
}

type NoteResult struct {
	ID          int            `xml:"id,attr"`
	Host        NoteResultHost `xml:"host"`
	Port        int            `xml:"port"`
	NVT         NoteResultNVT  `xml:"nvt"`
	Severity    string         `xml:"severity"`
	Threat      string         `xml:"threat"`
	QOD         NoteResultQOD  `xml:"qod"`
	Description string         `xml:"description"`
}

type NoteResultHost struct {
	Value string              `xml:",chardata"`
	Asset NoteResultHostAsset `xml:"asset"`
}

type NoteResultHostAsset struct {
	AssetID int `xml:"asset_id,attr"`
}

type NoteResultNVT struct {
	OID      string `xml:"oid,attr"`
	Name     string `xml:"name"`
	Type     string `xml:"type"`
	CVSSBase string `xml:"cvss_base"`
	CVE      string `xml:"cve"`
	BID      int    `xml:"bid"`
}

type NoteResultQOD struct {
	Value int `xml:"value"`
	Type  int `xml:"type"`
}

type ResultDeltaOverrides struct {
	Override []Override `xml:"override"`
}

type Override struct {
	ID               string              `xml:"id,attr"`
	Permissions      OverridePermissions `xml:"permissions"`
	Owner            OverrideOwner       `xml:"owner"`
	NVT              OverrideNVT         `xml:"nvt"`
	CreationTime     time.Time           `xml:"creation_time"`
	ModificationTime time.Time           `xml:"modification_time"`
	Writable         bool                `xml:"writable"`
	InUse            bool                `xml:"in_use"`
	Active           bool                `xml:"active"`
	Text             OverrideText        `xml:"text"`
	Threat           string              `xml:"threat"`
	Severity         string              `xml:"severity"`
	NewThreat        string              `xml:"new_threat"`
	NewSeverity      string              `xml:"new_severity"`
	Orphan           bool                `xml:"orphan"`
	UserTags         OverrideUserTags    `xml:"user_tags"`
	Hosts            string              `xml:"hosts"`
	Port             string              `xml:"port"`
	Task             OverrideTask        `xml:"task"`
	EndTime          string              `xml:"end_time"`
	Result           OverrideResult      `xml:"Result"`
}

type OverridePermissions struct {
	Permission []OverridePermissionsPermission `xml:"permission"`
}

type OverridePermissionsPermission struct {
	Name string `xml:"name"`
}

type OverrideOwner struct {
	Name string `xml:"name"`
}

type OverrideNVT struct {
	OID  string `xml:"oid,attr"`
	Name string `xml:"name"`
	Type string `xml:"type"`
}

type OverrideText struct {
	Text    string `xml:",chardata"`
	Excerpt bool   `xml:"excerpt,attr"`
}

type OverrideUserTags struct {
	Count int                   `xml:"count"`
	Tag   []OverrideUserTagsTag `xml:"tag"`
}

type OverrideUserTagsTag struct {
	ID      string `xml:"id,attr"`
	Name    string `xml:"name"`
	Value   string `xml:"value"`
	Comment string `xml:"comment"`
}

type OverrideTask struct {
	ID string `xml:"id,attr"`
}

type OverrideResult struct {
	ID          string             `xml:"id,attr"`
	Host        OverrideResultHost `xml:"host"`
	Port        string             `xml:"port"`
	NVT         OverrideResultNVT  `xml:"nvt"`
	Threat      string             `xml:"threat"`
	Severity    string             `xml:"severity"`
	QOD         OverrideResultQOD  `xml:"qod"`
	Description string             `xml:"description"`
}

type OverrideResultHost struct {
	Value string                  `xml:",chardata"`
	Asset OverrideResultHostAsset `xml:"asset"`
}

type OverrideResultHostAsset struct {
	AssetID string `xml:"asset_id,attr"`
}

type OverrideResultNVT struct {
	OID      string `xml:"oid,attr"`
	Name     string `xml:"name"`
	Type     string `xml:"type"`
	CVSSBase string `xml:"cvss_base"`
	CVE      string `xml:"cve"`
	BID      int    `xml:"bid"`
}

type OverrideResultQOD struct {
	Value int    `xml:"value"`
	Type  string `xml:"type"`
}

type ResultDetection struct {
	Result ResultDetectionResult `xml:"result"`
}

type ResultDetectionResult struct {
	ID      string                       `xml:"id,attr"`
	Details ResultDetectionResultDetails `xml:"details"`
}

type ResultDetectionResultDetails struct {
	Detail []ResultDetectionResultDetailsDetail `xml:"detail"`
}

type ResultDetectionResultDetailsDetail struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type ResultNotes struct {
	Note []Note `xml:"note"`
}

type ResultOverrides struct {
	Override []Override `xml:"override"`
}

type ResultTickets struct {
	Ticket []ResultTicketsTicket `xml:"ticket"`
}

type ResultTicketsTicket struct {
	ID string `xml:"id,attr"`
}

type GetResultsResponseFilters struct {
	ID       string                                     `xml:"id,attr"`
	Term     string                                     `xml:"term"`
	Name     string                                     `xml:"name"`
	Keywords []GetResultsResponseFiltersKeywordsKeyword `xml:"keywords"`
}

type GetResultsResponseFiltersKeywordsKeyword struct {
	Column   string `xml:"column"`
	Relation string `xml:"relation"`
	Value    string `xml:"value"`
}

type GetResultsResponseSort struct {
	Value string                      `xml:",chardata"`
	Field GetResultsResponseSortField `xml:"field"`
}

type GetResultsResponseSortField struct {
	Order string `xml:"order"`
}

type GetResultsResponseResults struct {
	Start string `xml:"start,attr"`
	Max   string `xml:"max,attr"`
}

type GetResultsResponseResultCount struct {
	Filtered int `xml:"filtered"`
	Page     int `xml:"page"`
}
