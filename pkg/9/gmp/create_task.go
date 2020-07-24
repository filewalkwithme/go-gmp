package gmp

import "encoding/xml"

type CreateTaskCommand struct {
	XMLName         xml.Name               `xml:"create_task"`
	Name            string                 `xml:"name,omitempty"`
	Comment         string                 `xml:"comment,omitempty"`
	Copy            string                 `xml:"copy,omitempty"`
	Alterable       bool                   `xml:"alterable,omitempty"`
	Config          *CreateTaskConfig      `xml:"config,omitempty"`
	Target          *CreateTaskTarget      `xml:"target,omitempty"`
	HostsOrdering   string                 `xml:"hosts_ordering,omitempty"`
	Scanner         *CreateTaskScanner     `xml:"scanner,omitempty"`
	Alert           []*CreateTaskAlert     `xml:"alert,omitempty"`
	Schedule        *CreateTaskSchedule    `xml:"schedule,omitempty"`
	SchedulePeriods int                    `xml:"schedule_periods,omitempty"`
	Observers       string                 `xml:"observers,omitempty"`
	Preferences     *CreateTaskPreferences `xml:"preferences,omitempty"`
}

type CreateTaskConfig struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTaskTarget struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTaskScanner struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTaskAlert struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTaskSchedule struct {
	ID string `xml:"id,attr,omitempty"`
}

type CreateTaskPreferences struct {
	Preference []*CreateTaskPreferencesPreference `xml:"preference,omitempty"`
}

type CreateTaskPreferencesPreference struct {
	ScannerName string `xml:"scanner_name,omitempty"`
	Value       string `xml:"value,omitempty"`
}

type CreateTaskResponse struct {
	XMLName    xml.Name `xml:"create_task_response"`
	Status     string   `xml:"status,attr"`
	StatusText string   `xml:"status_text,attr"`
	ID         string   `xml:"id,attr"`
}
