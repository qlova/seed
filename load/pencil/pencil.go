package pencil

import (
	"encoding/xml"
	"strconv"
)

//Document is a pencil document.
type Document struct {
	XMLName    xml.Name `xml:"Document"`
	Text       string   `xml:",chardata"`
	Xmlns      string   `xml:"xmlns,attr"`
	Properties struct {
		Text     string `xml:",chardata"`
		Property []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"Property"`
	} `xml:"Properties"`
	Pages struct {
		Text string `xml:",chardata"`
		Page []struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
		} `xml:"Page"`
	} `xml:"Pages"`
}

//Page is a pencil page.
type Page struct {
	XMLName xml.Name `xml:"Page"`
	Text    string   `xml:",chardata"`
	P       string   `xml:"p,attr"`

	Properties struct {
		Text     string `xml:",chardata"`
		Property []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"Property"`
	} `xml:"Properties"`
	Content struct {
		Data []byte `xml:",innerxml"`
	} `xml:"Content"`
}

func (page Page) Width() int {
	for _, property := range page.Properties.Property {
		if property.Name == "width" {
			i, _ := strconv.Atoi(property.Text)
			return i
		}
	}
	return 0
}

func (page Page) Height() int {
	for _, property := range page.Properties.Property {
		if property.Name == "height" {
			i, _ := strconv.Atoi(property.Text)
			return i
		}
	}
	return 0
}
