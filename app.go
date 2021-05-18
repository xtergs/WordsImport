package main

import (
	"WordImport/api"
	"WordImport/updates"

	// "fmt"
	// "log"
	"WordImport/import/kindle"
	"flag"
	"fmt"
	"log"
	"os"
	// "strings"
	// "github.com/google/fscrypt/filesystem"
	// "github.com/google/gousb"
	// "github.com/google/gousb/usbid"
)

var CurrentVersion = "0"
var UpdateLink = "https://api.github.com/repos/xtergs/WordsImport/releases/latest"
var Host = ""

func main() {

	host := ""
	provider := flag.String("provider", "kindle", "Supported values: kindle")
	dbPath := flag.String("db", "vocab.db", "Path to vocab.db")
	userId := flag.String("u", "", "provide telegram user id")
	if Host != "" {
		host = Host
	} else {
		host = *flag.String("host", "http://localhost:8821", "")
	}
	skipUpdates := flag.Bool("skipUpdate", false, "")

	flag.Parse()

	if !*skipUpdates {
		updates.CheckNewVersion(UpdateLink, CurrentVersion)
	}

	if *provider == "kindle" {
		importFromKindle(*dbPath, *userId, host)
	}

}

func importFromKindle(dbPath string, userId string, host string) {

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Fatal(err)
	}

	words, err := kindle.ReadWords(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("importing to webserver...")
	err = api.ImportList(host, userId, words)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Will import next words:\n")
	//
	//for wordKey := range words {
	//	word := words[wordKey]
	//	fmt.Printf("%s %s %s\n", word.Word, word.Stem, word.Lang)
	//}
}
