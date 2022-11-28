package utils

import (
	"encoding/json"
	"github.com/extark/gateway_microservice/models"
	"io"
	"log"
	"os"
)

func ReadConf(configPath string) (configList []*models.ConfigJsonFormat, err error) {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var data []*models.ConfigJsonFormat

	err = json.Unmarshal([]byte(byteValue), &data)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return data, nil
}
