package mcclane

import (
	"github.com/ign/ipl-mcclane/configreader"
	"github.com/ign/ipl-mcclane/mcclane/models"
	"log"
)

var config models.Config

func LoadConfig() error {
	err := configreader.UnmarshalConfig("resources/config.json", &config)
	if err != nil {
		log.Println("Error loading config")
		return err
	}
	return nil
}
