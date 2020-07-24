package gmp

import (
	"encoding/xml"
)

type StopTaskCommand struct {
	XMLName xml.Name `xml:"stop_task"`
	TaskID  string   `xml:"task_id,attr,omitempty"`
}

type StopTaskResponse struct {
	XMLName    xml.Name `xml:"stop_task_response"`
	Status     string   `xml:"status,attr"`
	StatusText string   `xml:"status_text,attr"`
}
