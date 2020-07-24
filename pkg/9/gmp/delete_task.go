package gmp

import (
	"encoding/xml"
)

type DeleteTaskCommand struct {
	XMLName  xml.Name `xml:"delete_task"`
	TaskID   string   `xml:"task_id,attr,omitempty"`
	Ultimate bool     `xml:"ultimate,attr,omitempty"`
}

type DeleteTaskResponse struct {
	XMLName    xml.Name `xml:"delete_task_response"`
	Status     string   `xml:"status,attr"`
	StatusText string   `xml:"status_text,attr"`
}
