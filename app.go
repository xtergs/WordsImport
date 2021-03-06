package main

import (
	"github.com/xtergs/WordsImport/api"
	"github.com/xtergs/WordsImport/updates"

	"flag"
	"fmt"
	"github.com/xtergs/WordsImport/import/kindle"
	"log"
	"os"
)

var CurrentVersion = "0"
var UpdateLink = "https://api.github.com/repos/xtergs/WordsImport/releases/latest"
var Host = ""

func main() {

	host := ""
	provider := flag.String("provider", "kindle", "Supported values: kindle")
	dbPath := flag.String("db", "vocab1.db", "Path to vocab.db")
	userId := flag.String("u", "", "provide telegram user id")
	if Host != "" {
		host = Host
	} else {
		host = *flag.String("host", "http://localhost:8821", "")
	}
	skipUpdates := flag.Bool("skipUpdate", false, "Skip updates check")

	flag.Parse()

	fmt.Printf("Version: %s\n\n", CurrentVersion)

	for len(*userId) <= 0 {
		fmt.Printf("Provide telegram userId:\n")
		fmt.Scanf("%s", userId)
	}

	if !*skipUpdates {
		updates.CheckNewVersion(UpdateLink, CurrentVersion)
	}

	fmt.Printf("\n")

	if *provider == "kindle" {
		importFromKindle(*dbPath, *userId, host)
	}

}

func importFromKindle(dbPath string, userId string, host string) {

	for true {
		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			fmt.Errorf(err.Error())
			fmt.Printf("Provide path to db:\n")
			fmt.Scanf("%s", &dbPath)
		} else {
			break
		}
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
