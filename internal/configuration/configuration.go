package configuration

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigurationList struct {
	Keyvaults []KeyvaultConfiguration `yaml:"Keyvaults"`
}

type KeyvaultConfiguration struct {
	Name           string `yaml:"Name"`
	SubscriptionId string `yaml:"Subscription"`
}

func GetConfiguration() ConfigurationList {
	cacheVaultFile, err := os.Open("vaulty.conf")
	if err != nil {
		log.Fatal(err)
	}
	var vaults ConfigurationList
	out, _ := io.ReadAll(cacheVaultFile)

	err = yaml.Unmarshal(out, &vaults)
	if err != nil {
		log.Fatal(err)
	}

	return vaults
}
