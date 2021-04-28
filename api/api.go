package api

import (
	"WordImport/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func ImportList(host string, telegramId string, words map[string]models.Word) error {

	values := make([]models.Word, 0, len(words))
	for _, v := range words {
		values = append(values, v)
	}

	payload := models.WordsList{
		TelegramUserId: telegramId,
		Words:          values,
	}

	jsonValue, _ := json.Marshal(payload)

	_, err := http.Post(host+"/api/wordsImport/list", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}
