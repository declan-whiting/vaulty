package configuration

import (
	"io"
	"log"
	"os"

	"github.com/declan-whiting/vaulty/internal/models"
	"gopkg.in/yaml.v3"
)

type ConfigrationService struct{}

func NewConfigurationService() *ConfigrationService {
	return new(ConfigrationService)
}

func (cs *ConfigrationService) GetConfiguration() models.ConfigurationList {
	cacheVaultFile, err := os.Open("vaulty.conf")
	if err != nil {
		log.Fatal(err)
	}
	var vaults models.ConfigurationList
	out, _ := io.ReadAll(cacheVaultFile)

	err = yaml.Unmarshal(out, &vaults)
	if err != nil {
		log.Fatal(err)
	}

	return vaults
}
