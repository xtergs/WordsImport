package kindle

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type KindleWord struct {
	Id    string `db:"id"`
	Word  string `db:"word"`
	Stem  string `db:"stem"`
	Lang  string `db:"lang"`
	Usage string `db:"usage"`
}

type Usage struct {
	Usage string
}

type Word struct {
	Word   string
	Stem   string
	Lang   string
	Usages []Usage
}

func ReadWords(file string) (map[string]Word, error) {

	var connection = "file:" + file

	var db, err = sqlx.Open("sqlite3", connection)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected!\n")

	rows, err := db.Queryx(`select Words.id, word, stem, lang, usage from WORDS
	left join LOOKUPS L on WORDS.id = L.word_key
	where category != 100`, nil)
	if err != nil {
		log.Fatal(err)
	}

	results := make(map[string]Word)
	temp := KindleWord{}

	for rows.Next() {
		err := rows.StructScan(&temp)
		if err != nil {
			log.Fatal(err)
		}

		var wordKey = temp.Id
		val, ok := results[wordKey]
		if ok {
			var newUsage = Usage{Usage: temp.Usage}
			val.Usages = append(val.Usages, newUsage)
		} else {
			results[wordKey] = Word{
				Word: temp.Word,
				Stem: temp.Stem,
				Lang: temp.Lang,
				Usages: []Usage{{
					Usage: temp.Usage,
				}},
			}
		}
	}

	rows.Close()

	fmt.Println(len(results))

	return results, nil
}
