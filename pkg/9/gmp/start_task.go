package gmp

import "encoding/xml"

type StartTaskCommand struct {
	XMLName xml.Name `xml:"start_task"`
	TaskID  string   `xml:"task_id,attr,omitempty"`
}

type StartTaskResponse struct {
	XMLName    xml.Name `xml:"start_task_response"`
	Status     string   `xml:"status,attr"`
	StatusText string   `xml:"status_text,attr"`
	ReportID   string   `xml:"report_id"`
}
