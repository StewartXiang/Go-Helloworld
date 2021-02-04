package main

import (
	"encoding/xml"
	"fmt"
)

func main()  {
	a := XMLMessage{
		"sdf<sdf>",
		1000,
		"ffd",
	}
	aByte, _ := xml.Marshal(a)
	aStr := string(aByte)
	fmt.Printf("a1: %s\n", aStr)

	bStr := `
<XMLMessage><my_name><![CDATA[wer<werw>]]></my_name><address>1000</address><phone>ffd</phone></XMLMessage>
`
	var b XMLMessage
	err := xml.Unmarshal([]byte(bStr), &b)
	fmt.Println(b.Name)
	// wer<werwer>
	fmt.Println(err)
}

//这才可以
type XMLMessage struct {
	Name CDATA `xml:"my_name"`
	Address int `xml:"address"`
	Phone string `xml:"phone"`
}
type CDATA string
func (n CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct{
		S string `xml:",innerxml"`
	}{
		S: "<![CDATA[" + string(n) + "]]>",
	}, start)
}

// 这种不行
type JsonMessage2 struct {
	Name string `xml:"my_name>,cdata"`
	Address int `xml:"address"`
	Phone string `xml:"phone"`
}

