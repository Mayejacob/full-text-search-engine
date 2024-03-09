package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/mayejacob/full-text-search-engine/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "small wild cat", "search query")
	flag.Parse()
	log.Println("Full text search is in progress...")
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("loaded %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("indexed %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedIds := idx.Search(query)

	log.Printf("search found %d documents in %v", len(matchedIds), time.Since(start))
	for _, id := range matchedIds {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
