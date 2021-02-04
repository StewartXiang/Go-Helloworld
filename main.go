package main

import (
	"encoding/xml"
	"fmt"
)

func main()  {
	a := JsonMessage{
		CDATA{"wer<werw>"},
		1000,
		"ffd",
	}
	a2 := JsonMessage2{
		"wer<werwer>",
		1000,
		"ffd",
	}
	aByte, _ := xml.Marshal(a)
	aStr := string(aByte)
	fmt.Printf("a1: %s\n", aStr)
	a2Byte, _ := xml.Marshal(a2)
	a2Str := string(a2Byte)
	fmt.Printf("a2: %s\n", a2Str)

	bStr := `
<JsonMessage><my_name><![CDATA[wer<werw>]]></my_name><address>1000</address><phone>ffd</phone></JsonMessage>
`
	var b JsonMessage
	err := xml.Unmarshal([]byte(bStr), &b)
	fmt.Println(bStr)
	//<JsonMessage>
	//		<my_name>
	//			<![CDATA[wer<werw>]]>
	//		</my_name>
	//		<address>1000</address><phone>ffd</phone>
	//</JsonMessage>
	fmt.Println(b.Name)
	// 返回{wer<werw>}
	fmt.Println(b.Name.Text)
	// 返回wer<werw>
	fmt.Println(err)
}

type JsonMessage struct {
	Name CDATA `xml:"my_name"`
	Address int `xml:"address"`
	Phone string `xml:"phone"`
}

type CDATA struct {
	Text string `xml:",cdata"`
}


type JsonMessage2 struct {
	Name string `xml:"my_name>,cdata"`
	Address int `xml:"address"`
	Phone string `xml:"phone"`
}

