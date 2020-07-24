package gmp

import (
	"encoding/xml"
	"time"
)

type GetTasksCommand struct {
	XMLName          xml.Name `xml:"get_tasks"`
	TaskID           string   `xml:"task_id,attr,omitempty"`
	Filter           string   `xml:"filter,attr,omitempty"`
	FiltID           string   `xml:"filt_id,attr,omitempty"`
	Trash            bool     `xml:"trash,attr,omitempty"`
	Details          bool     `xml:"details,attr,omitempty"`
	IgnorePagination bool     `xml:"ignore_pagination,attr,omitempty"`
	SchedulesOnly    bool     `xml:"schedules_only,attr,omitempty"`
	UsageType        string   `xml:"usage_type,attr,omitempty"`
}

type GetTasksResponse struct {
	XMLName        xml.Name                  `xml:"get_tasks_response"`
	Status         string                    `xml:"status,attr"`
	StatusText     string                    `xml:"status_text,attr"`
	ApplyOverrides string                    `xml:"apply_overrides"`
	Task           []GetTasksResponseTask    `xml:"task"`
	Filters        GetTasksResponseFilters   `xml:"filters"`
	Sort           GetTasksResponseSort      `xml:"sort"`
	Tasks          GetTasksResponseTasks     `xml:"tasks"`
	TaskCount      GetTasksResponseTaskCount `xml:"task_count"`
}

type GetTasksResponseTask struct {
	ID               string                               `xml:"id,attr"`
	Owner            GetTasksResponseTaskOwner            `xml:"owner"`
	Name             string                               `xml:"name"`
	Comment          string                               `xml:"comment"`
	CreationTime     time.Time                            `xml:"creation_time"`
	ModificationTime time.Time                            `xml:"modification_time"`
	Writable         bool                                 `xml:"writable"`
	InUse            bool                                 `xml:"in_use"`
	Permissions      GetTasksResponseTaskPermissions      `xml:"permissions"`
	UserTags         GetTasksResponseTaskUserTags         `xml:"user_tags"`
	Status           string                               `xml:"status"`
	Progress         GetTasksResponseTaskProgress         `xml:"progress"`
	Alterable        bool                                 `xml:"alterable"`
	UsageType        string                               `xml:"usage_type"`
	Config           GetTasksResponseTaskConfig           `xml:"config"`
	Target           GetTasksResponseTaskTarget           `xml:"target"`
	HostsOrdering    string                               `xml:"hosts_ordering"`
	Scanner          GetTasksResponseTaskScanner          `xml:"scanner"`
	Alert            GetTasksResponseTaskAlert            `xml:"alert"`
	Observers        GetTasksResponseTaskObservers        `xml:"observers"`
	Schedule         GetTasksResponseTaskSchedule         `xml:"schedule"`
	SchedulePeriods  int                                  `xml:"schedule_periods"`
	ReportCount      GetTasksResponseTaskReportCount      `xml:"report_count"`
	Trend            string                               `xml:"trend"`
	CurrentReport    GetTasksResponseTaskCurrentReport    `xml:"current_report"`
	LastReport       GetTasksResponseTaskLastReportReport `xml:"last_report"`
	Reports          GetTasksResponseTaskReports          `xml:"reports"`
	AverageDuration  string                               `xml:"average_duration"`
	ResultCount      string                               `xml:"result_count"`
	Preferences      GetTasksResponseTaskPreferences      `xml:"preferences"`
}

type GetTasksResponseTaskOwner struct {
	Name string `xml:"name"`
}
type GetTasksResponseTaskPermissions struct {
	Permission []GetTasksResponseTaskPermissionsPermission `xml:"permission"`
}

type GetTasksResponseTaskPermissionsPermission struct {
	Name string `xml:"name"`
}

type GetTasksResponseTaskUserTags struct {
	Count int                               `xml:"count"`
	Tag   []GetTasksTesponseTaskUserTagsTag `xml:"tag"`
}

type GetTasksTesponseTaskUserTagsTag struct {
	ID      string `xml:"id,attr"`
	Name    string `xml:"name"`
	Value   string `xml:"value"`
	Comment string `xml:"comment"`
}

type GetTasksResponseTaskProgress struct {
	Value        string                                     `xml:",chardata"`
	HostProgress []GetTasksResponseTaskProgressHostProgress `xml:"host_progress"`
}

type GetTasksResponseTaskProgressHostProgress struct {
	Value string `xml:",chardata"`
	Host  string `xml:"host"`
}

type GetTasksResponseTaskConfig struct {
	ID    string `xml:"id,attr"`
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Trash bool   `xml:"trash"`
}

type GetTasksResponseTaskTarget struct {
	ID          string `xml:"id,attr"`
	Name        string `xml:"name"`
	Permissions string `xml:"permissions"`
	Trash       bool   `xml:"trash"`
}

type GetTasksResponseTaskScanner struct {
	ID          string `xml:"id,attr"`
	Name        string `xml:"name"`
	Type        int    `xml:"type"`
	Permissions string `xml:"permissions"`
}

type GetTasksResponseTaskAlert struct {
	ID          string `xml:"id,attr"`
	Name        string `xml:"name"`
	Permissions string `xml:"permissions"`
	Trash       bool   `xml:"trash"`
}

type GetTasksResponseTaskObservers struct {
	Value string                               `xml:",chardata"`
	Group []GetTasksTesponseTaskObserversGroup `xml:"group"`
	Role  []GetTasksTesponseTaskObserversRole  `xml:"role"`
}

type GetTasksTesponseTaskObserversGroup struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name"`
}

type GetTasksTesponseTaskObserversRole struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name"`
}

type GetTasksResponseTaskSchedule struct {
	ID           string    `xml:"id,attr"`
	Name         string    `xml:"name"`
	Trash        bool      `xml:"trash"`
	FirstTime    time.Time `xml:"first_time"`
	NextTime     string    `xml:"next_time"`
	Icalendar    time.Time `xml:"icalendar"`
	Period       int       `xml:"period"`
	PeriodMonths int       `xml:"period_months"`
	Duration     int       `xml:"duration"`
	Timezone     string    `xml:"timezone"`
}

type GetTasksResponseTaskReportCount struct {
	Value    string `xml:",chardata"`
	Finished int    `xml:"finished"`
}

type GetTasksResponseTaskCurrentReport struct {
	Report GetTasksResponseTaskCurrentReportReport `xml:",report"`
}

type GetTasksResponseTaskCurrentReportReport struct {
	ID        string    `xml:"id,attr"`
	Timestamp time.Time `xml:"timestamp"`
}

type GetTasksResponseTaskLastReportReport struct {
	ID          string                                          `xml:"id,attr"`
	Timestamp   time.Time                                       `xml:"timestamp"`
	ScanEnd     time.Time                                       `xml:"scan_end"`
	ResultCount GetTasksResponseTaskLastReportReportResultCount `xml:"result_count"`
	Severity    float32                                         `xml:"severity"`
}

type GetTasksResponseTaskLastReportReportResultCount struct {
	Debug         int `xml:"debug"`
	FalsePositive int `xml:"false_positive"`
	Log           int `xml:"log"`
	Info          int `xml:"info"`
	Warning       int `xml:"warning"`
	Hole          int `xml:"hole"`
}

type GetTasksResponseTaskReports struct {
	Report []GetTasksResponseTaskReportsReport `xml:"report"`
}

type GetTasksResponseTaskReportsReport struct {
	ID            string                                       `xml:"id,attr"`
	Timestamp     time.Time                                    `xml:"timestamp"`
	ScanEnd       time.Time                                    `xml:"scan_end"`
	ScanRunStatus string                                       `xml:"scan_run_status"`
	ResultCount   GetTasksResponseTaskReportsReportResultCount `xml:"result_count"`
	Severity      float32                                      `xml:"severity"`
}

type GetTasksResponseTaskReportsReportResultCount struct {
	Debug         int `xml:"debug"`
	FalsePositive int `xml:"false_positive"`
	Log           int `xml:"log"`
	Info          int `xml:"info"`
	Warning       int `xml:"warning"`
	Hole          int `xml:"hole"`
}

type GetTasksResponseTaskPreferences struct {
	Preferences []getTasksResponseTaskPreferencesPreference `xml:"preference"`
}

type getTasksResponseTaskPreferencesPreference struct {
	Name        string `xml:"name"`
	ScannerName string `xml:"scanner_name"`
	Value       string `xml:"value"`
}

type GetTasksResponseFilters struct {
	ID       string                          `xml:"id,attr"`
	Term     string                          `xml:"term"`
	Name     string                          `xml:"name"`
	Keywords GetTasksResponseFiltersKeywords `xml:"keywords"`
}

type GetTasksResponseFiltersKeywords struct {
	Keyword []GetTasksResponseFiltersKeywordsKeyword `xml:"keyword"`
}

type GetTasksResponseFiltersKeywordsKeyword struct {
	Column   string `xml:"column"`
	Relation string `xml:"relation"`
	Value    string `xml:"value"`
}

type GetTasksResponseSort struct {
	Value string                    `xml:",chardata"`
	Field GetTasksResponseSortField `xml:"field"`
}

type GetTasksResponseSortField struct {
	Order string `xml:"order"`
}

type GetTasksResponseTasks struct {
	Start int `xml:"start,attr"`
	Max   int `xml:"max,attr"`
}

type GetTasksResponseTaskCount struct {
	Filtered int `xml:"filtered"`
	Page     int `xml:"page"`
}
