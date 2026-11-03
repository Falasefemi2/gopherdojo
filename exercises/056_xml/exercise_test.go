package main

import "testing"

func TestEncodeBook(t *testing.T) {
	tests := []struct {
		name string
		b    Book
		want string
	}{
		{
			name: "full record",
			b:    Book{ID: 1, Title: "Go in Action", Author: "Ada", Price: 39.5},
			want: `<book id="1"><title>Go in Action</title><author>Ada</author><price>39.5</price></book>`,
		},
		{
			name: "zero value",
			b:    Book{},
			want: `<book id="0"><title></title><author></author><price>0</price></book>`,
		},
		{
			name: "escaped text",
			b:    Book{ID: 2, Title: "Go & Beyond", Author: "Bob <b>", Price: 0},
			want: `<book id="2"><title>Go &amp; Beyond</title><author>Bob &lt;b&gt;</author><price>0</price></book>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeBook(tt.b)
			if err != nil {
				t.Fatalf("encodeBook(%+v) returned error: %v", tt.b, err)
			}
			if got != tt.want {
				t.Errorf("encodeBook(%+v) = %s, want %s", tt.b, got, tt.want)
			}
		})
	}
}

func TestDecodeBook(t *testing.T) {
	tests := []struct {
		name string
		data string
		want Book
	}{
		{
			name: "full record",
			data: `<book id="1"><title>Go in Action</title><author>Ada</author><price>39.5</price></book>`,
			want: Book{ID: 1, Title: "Go in Action", Author: "Ada", Price: 39.5},
		},
		{
			name: "missing optional author",
			data: `<book id="3"><title>Untitled</title><price>1.25</price></book>`,
			want: Book{ID: 3, Title: "Untitled", Price: 1.25},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeBook(tt.data)
			if err != nil {
				t.Fatalf("decodeBook(%s) returned error: %v", tt.data, err)
			}
			if got.ID != tt.want.ID || got.Title != tt.want.Title ||
				got.Author != tt.want.Author || got.Price != tt.want.Price {
				t.Errorf("decodeBook(%s) = %+v, want %+v", tt.data, got, tt.want)
			}
		})
	}
}

func TestDecodeBookRejectsBadXML(t *testing.T) {
	tests := []struct {
		name string
		data string
	}{
		{"unclosed title", `<book id="1"><title>x</book>`},
		{"not xml", `plain text`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := decodeBook(tt.data); err == nil {
				t.Errorf("decodeBook(%s) returned nil error, want an error", tt.data)
			}
		})
	}
}

func TestBookRoundTrip(t *testing.T) {
	want := Book{ID: 7, Title: "The Go Programming Language", Author: "Donovan & Kernighan", Price: 44.99}

	data, err := encodeBook(want)
	if err != nil {
		t.Fatalf("encodeBook returned error: %v", err)
	}
	got, err := decodeBook(data)
	if err != nil {
		t.Fatalf("decodeBook returned error: %v", err)
	}
	if got.ID != want.ID || got.Title != want.Title ||
		got.Author != want.Author || got.Price != want.Price {
		t.Fatalf("round trip = %+v, want %+v", got, want)
	}
}
