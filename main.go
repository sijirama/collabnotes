package main

import (
	"flag"
	//utils "github.com/sijirama/search-engine/utils"
	"log"
	"time"

	"github.com/sijirama/search-engine/utils"
)

var dumppath, query string

func init() {

}

func main() {

	//flag.StringVar(&dumppath, "p", "dataset.xml", "path of the dataset")
	//get our arguments
	flag.StringVar(&dumppath, "p", "comp.xml.gz", "path of the dataset")
	flag.StringVar(&query, "q", "Small wild cat", "string to search") // remove this later and make it a repl

	flag.Parse()

	log.Println("full text search in progress: ")

	// load documents
	start := time.Now()
	docs, err := utils.LoadDocuments(dumppath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	//index documents
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	//search
	start = time.Now()
	matchedIDS := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDS), time.Since(start))

	for _, id := range matchedIDS {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
