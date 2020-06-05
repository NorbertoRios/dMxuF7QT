package configuration

//Field protocol description
type Field struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
	Size int    `xml:"size,attr"`
}

//Reasons codes
type Reasons struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}
