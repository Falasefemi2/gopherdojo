// 056: XML. encoding/xml uses struct tags to name elements and attributes;
// Marshal and Unmarshal convert between XML text and Go values.
//
// Scenario: a library catalogue exchanges book records as XML with partners
// that expect the same element layout.
//
// Rules:
//  1. The root element is book and the id field is its attribute.
//  2. title, author and price are child elements in that order.
//  3. encodeBook returns compact XML with no declaration and no whitespace.
//  4. decodeBook leaves absent child elements at their zero value and returns
//     an error for malformed XML.
//
// TODO: implement encodeBook and decodeBook. Keep main() as the demo harness.
package main

import (
	"encoding/xml"
	"fmt"
)

// Book is the catalogue record exchanged as XML.
type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      int      `xml:"id,attr"`
	Title   string   `xml:"title"`
	Author  string   `xml:"author"`
	Price   float64  `xml:"price"`
}

// encodeBook marshals b to compact XML.
func encodeBook(b Book) (string, error) {
	// TODO
	return "", nil
}

// decodeBook unmarshals XML into a Book.
func decodeBook(data string) (Book, error) {
	// TODO
	return Book{}, nil
}

func main() {
	data, err := encodeBook(Book{ID: 1, Title: "Go in Action", Author: "Ada", Price: 39.5})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(data)
}
