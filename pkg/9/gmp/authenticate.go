package gmp

import "encoding/xml"

type AuthenticateCommand struct {
	XMLName     xml.Name                `xml:"authenticate"`
	Credentials AuthenticateCredentials `xml:"credentials"`
}

type AuthenticateCredentials struct {
	Username string `xml:"username"`
	Password string `xml:"password"`
}

type AuthenticateResponse struct {
	XMLName    xml.Name `xml:"authenticate_response"`
	Status     string   `xml:"status,attr"`
	StatusText string   `xml:"status_text,attr"`
	Role       string   `xml:"role"`
	Timezone   string   `xml:"timezone"`
}
