package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	Links []Link `xml:"links>sublink"`
	ID    int
}

type Link struct {
	LinkType string `xml:"linktype,attr"`
	Anchor   string `xml:"anchor"`
	Link     string `xml:"link"`
}

func LoadDocuments(path string) ([]document, error) {
	f, error := os.Open(path)
	if error != nil {
		return nil, error
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, error
	}
	defer gz.Close()
	dec := xml.NewDecoder(gz)

	dump := struct {
		Documents []document `xml:"doc"`
	}{}

	err = dec.Decode(&dump)
	if err != nil {
		return nil, error
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
