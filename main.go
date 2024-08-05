package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/sijirama/search-engine/utils"
)

var dumppath, query string

func init() {

}

func main() {

	flag.StringVar(&dumppath, "p", "comp.xml.gz", "path of the dataset")
	//flag.StringVar(&query, "q", "Small wild cat", "string to search") // remove this later and make it a repl

	flag.Parse()

	log.Println("full text search in progress: ")

	//NOTE: load documents: uncompress, parse, assign ids to docs
	start := time.Now()
	docs, err := utils.LoadDocuments(dumppath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	//NOTE: index documents
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter exit() to exit the REPL")

	sigChan := make(chan os.Signal, 1)     // Create a channel to receive OS signals
	signal.Notify(sigChan, syscall.SIGINT) // Notify the channel on interrupt signals

	go func() {
		for {
			select {
			case <-sigChan:
				fmt.Println("\nInterrupt received, shutting down gracefully...")
				os.Exit(0)
			}
		}
	}()

	for {

		fmt.Print("\n>>> ")

		input, err := reader.ReadString('\n') // read the string
		if err != nil {
			log.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input) // Remove newline and any extra spaces

		if input == "exit" || input == "exit()" { // exit the app
			break
		}
		start = time.Now()
		matchedIDS := idx.Search(input)
		log.Printf("Search found %d documents in %v", len(matchedIDS), time.Since(start))

		for _, id := range matchedIDS {
			doc := docs[id]
			log.Printf("%d\t%s\n", id, doc.Text)
		}

	}

}
